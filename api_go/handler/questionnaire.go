package handler

import (
	"api/config"
	"api/content"
	"api/model"
	"api/util"
	"api/util/crypto"
	. "api/util/fcm"
	. "api/util/hit"
	"api/util/image"
	"api/util/kafka"
	"api/util/log"
	"api/util/validate"
	"encoding/json"
	"os/exec"
	"strconv"
	"time"
)

type QuestionnaireHandler content.Handler

type dataReceivedSurveyHandler struct {
	QuestionnaireRecodeId         string                                `json:"questionnaireRecodeId"`
	QrCodeDate                    string                                `json:"qrCodeDate"`
	Name                          string                                `json:"name"`
	Phone                         string                                `json:"phone"`
	UserId                        int64                                 `json:"userId"`
	WriteType                     string                                `json:"writeType"`
	Gender                        int                                   `json:"gender"`
	Age                           int                                   `json:"age"`
	ReveiveQuestionnaireTopicData []model.ReveiveQuestionnaireTopicData `json:"questionnaireTopic"`
}

type qrCodeDate struct {
	QuestionnaireId             int64  `json:"questionnaireId"`
	No                          string `json:"no"`
	TableNo                     string `json:"tableNo"`
	People                      int    `json:"people"`
	StoreId                     int64  `json:"StoreId"`
	AdminId                     int64  `json:"adminId"`
	MealTimeId                  int64  `json:"mealTimeId"`
	Code                        string `json:"code"`
	QuestionnaireQrcodeRecordId int64  `json:"questionnaireQrcodeRecordId"`
}

func (handler *QuestionnaireHandler) GetQuestionnaireInfo() interface{} {
	var data dataReceivedSurveyHandler
	var qrCodeDate qrCodeDate
	var questionnaire model.Questionnaire
	var questionnaireTopic model.QuestionnaireTopic
	var questionnaireRecord model.QuestionnaireRecord
	var brand model.Brand
	var store model.Store
	var users model.Users

	parameter, err := crypto.KeyDecrypt(handler.Parameter)
	log.Error(err)

	if err := json.Unmarshal([]byte(parameter), &data); err != nil {
		return util.RS{Message: "", Status: false}
	} else if data.WriteType == "" {
		return util.RS{Message: "", Status: false}
	} else if err := json.Unmarshal([]byte(data.QrCodeDate), &qrCodeDate); err != nil {
		return util.RS{Message: "", Status: false}
	} else if qrCodeDate.QuestionnaireId == 0 {
		return util.RS{Message: "問卷不存在", Status: false}
	} else if qrCodeDate.Code == "" {
		return util.RS{Message: "問卷編碼不存在", Status: false}
	} else if !questionnaire.SetId(qrCodeDate.QuestionnaireId).QueryOne().IsOpen {
		return util.RS{Message: "此問卷已關閉", Status: false}
	} else if !questionnaire.IsRepeatWrite && data.Phone != "" && questionnaireRecord.SetQuestionnaireId(qrCodeDate.QuestionnaireId).SetPhone(data.Phone).QueryOne().Id != 0 {
		return util.RS{Message: "已填寫過此問卷", Status: false}
	} else if userId := users.SetPhone(data.Phone).SetCountryCode("+886").GetUsersIdByPhone().Id; userId <= 0 {
		return util.RS{Message: "查無會員", Status: false}
	} else {
		//判斷是否有設時限
		if questionnaire.IsLimitTime {
			questionnaireRecord.SetEndedAt(time.Now().Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(questionnaire.LimitTime) + time.Second*time.Duration(0)))
		}

		//新增問卷紀錄資料
		if questionnaireRecodeId, err :=
			questionnaireRecord.
				SetId(0).
				SetQuestionnaireId(qrCodeDate.QuestionnaireId).
				SetCode(qrCodeDate.Code).
				SetTableNo(qrCodeDate.TableNo).
				SetNo(qrCodeDate.No).
				SetPeople(qrCodeDate.People).
				SetMealTimeId(qrCodeDate.MealTimeId).
				SetWriteType(data.WriteType).
				SetUserId(data.UserId).
				SetStoreId(questionnaire.StoreId).
				SetUseStoreId(qrCodeDate.StoreId).
				SetName(data.Name).
				SetAdminId(qrCodeDate.AdminId).
				SetMealTimeId(qrCodeDate.MealTimeId).
				SetQrcodeRecordId(qrCodeDate.QuestionnaireQrcodeRecordId).
				SetPhone("").
				Insert(); err != nil { //新增問卷
			return util.RS{Message: "問卷發生異常", Status: true}
		} else {
			encryptQuestionnaireRecodeId, err := crypto.KeyEncrypt(strconv.Itoa(int(questionnaireRecodeId)))
			log.Error(err)

			dataList := map[string]interface{}{
				"name":    questionnaire.Name,
				"content": questionnaire.Content,
				"imagePath": image.ReturnPhotoPath(If(questionnaire.ImagePath == "", brand.SetId(store.SetId(questionnaire.StoreId).QueryOne().BrandId).QueryOne().Image,
					questionnaire.ImagePath).(string)),
				"congratulationText":    questionnaire.CongratulationText,
				"questionnaireTopic":    questionnaireTopic.SetQuestionnaireId(qrCodeDate.QuestionnaireId).GetListByQuestionnaireId(),
				"questionnaireRecodeId": encryptQuestionnaireRecodeId,
				"buildInItem":           questionnaire.BuildInItem,
				"brandName":             brand.Name,
				"storeName":             store.Name,
			}

			return util.RS{Message: "", Status: true, Data: dataList}
		}
	}
}

func (handler *QuestionnaireHandler) InsertQuestionnaireRecord() interface{} {
	type coupon struct {
		Id        int64  `json:"id"`
		Amount    int    `json:"amount"`
		StartTime string `json:"startTime"`
	}

	var (
		data                           dataReceivedSurveyHandler
		questionnaireRecord            model.QuestionnaireRecord
		questionnaireRecordCheck       model.QuestionnaireRecord
		questionnaireAnswerAdmin       model.QuestionnaireAnswerAdmin
		questionnaireAnswer            model.QuestionnaireAnswer
		questionnaireQrcodeRecordAdmin model.QuestionnaireQrcodeRecordAdmin
		questionnaire                  model.Questionnaire
		users                          model.Users
		fcm                            FCM //推播功能宣告
		couponDataList                 []coupon
	)


	if err := json.Unmarshal([]byte(handler.Parameter), &data); err != nil {
		return util.RS{Message: "", Status: false}
	} else if questionnaireRecodeIdString, err := crypto.KeyDecrypt(data.QuestionnaireRecodeId); err != nil {
		return util.RS{Message: "", Status: false}
	} else if questionnaireRecodeId, err := strconv.Atoi(questionnaireRecodeIdString); err != nil {
		return util.RS{Message: "", Status: false}
	} else if !questionnaire.SetId(questionnaireRecord.SetId(int64(questionnaireRecodeId)).QueryOne().QuestionnaireId).QueryOne().IsOpen {
		return util.RS{Message: "此問卷已關閉", Status: false}
	} else if data.Name = TrimZero(data.Name); data.Name == "" && questionnaire.GetRequire("name", questionnaire.BuildInItem) {
		return util.RS{Message: "請填寫姓名", Status: false}
	} else if data.Phone = TrimZero(data.Phone); data.Phone == "" && questionnaire.GetRequire("phone", questionnaire.BuildInItem) {
		return util.RS{Message: "請填寫電話", Status: false}
	} else if data.Gender <= 0 && questionnaire.GetRequire("gender", questionnaire.BuildInItem) {
		return util.RS{Message: "請選擇性別", Status: false}
	} else if data.Age <= 0 && questionnaire.GetRequire("old", questionnaire.BuildInItem) {
		return util.RS{Message: "年齡錯誤", Status: false}
	} else if _, status := validate.CheckPhone(data.Phone); questionnaire.GetRequire("phone", questionnaire.BuildInItem) && !status {
		return util.RS{Message: "電話格式錯誤", Status: false}
	} else if !questionnaire.IsRepeatWrite && questionnaireRecordCheck.SetIsWrite(true).SetId(0).SetQuestionnaireId(questionnaireRecord.QuestionnaireId).SetPhone(data.Phone).QueryOne().Id != 0 {
		return util.RS{Message: "您已填寫過此問卷", Status: false}
	} else if questionnaire.IsLimitTime && questionnaireRecord.EndedAt.Before(time.Now()) {
		return util.RS{Message: "問卷已過期", Status: false}
	} else if userId := users.SetPhone(data.Phone).SetCountryCode("+886").GetUsersIdByPhone().Id; userId <= 0 {
		return util.RS{Message: "查無會員", Status: false}
	} else if err, questionnaireAnswerIdList :=
		questionnaireRecord.
			SetPhone(data.Phone).
			SetName(data.Name).
			SetAge(data.Age).
			SetGender(data.Gender).
			SetUserId(userId).
			Write(data.ReveiveQuestionnaireTopicData); err != nil { //填寫問卷
		return util.RS{Message: "填寫失敗", Status: true}
	} else {
		log.Error(json.Unmarshal([]byte(questionnaire.Coupon), &couponDataList))

		//判斷是否要發送優惠卷與電話是否已註冊並且判斷優惠卷剩餘張數是否足夠
		if len(couponDataList) > 0 && users.Id != 0 {
			count := 0
			//產生uuid
			out, _ := exec.Command("uuidgen").Output()
			uuid := string(out)
			//建立要傳送的資料(記得傳入uuid
			jsonMap := map[string]interface{}{
				"uuid":            uuid,
				"adminId":         questionnaire.AdminId,
				"userId":          users.Id,
				"questionnaireId": questionnaire.Id,
			}

			_, err := kafka.Push("QuestionnaireSendCoupon", jsonMap)
			log.Error(err)
			dataList := make(map[string]interface{}, 0)
			kafka.Listen("QuestionnaireApiSendCoupon", func(value []byte) (isBreak bool) {

				var kafKaRS util.KafKaRS
				if err := json.Unmarshal(value, &kafKaRS); err != nil {
					log.Error(err)
					dataList["status"] = false
					dataList["message"] = "error"
					return true
				} else if uuid == kafKaRS.Uuid {
					dataList["status"] = kafKaRS.Status
					dataList["message"] = kafKaRS.Message
					return true
				} else if count > 10 {
					dataList["status"] = false
					dataList["message"] = "time out"
					return true
				}

				count++
				return false
			})

			//kafka執行成功就推播
			if dataList["status"] == true {
				//傳送推播給app用戶///////////////////////
				fcm.SetBody(questionnaire.CongratulationText).SetData(map[string]interface{}{
					"action":       "GoCoupon",
					"id":       	0,
					"type":         "coupon",
					"title":        "非常感謝您抽空填寫問卷",
					"body":         questionnaire.CongratulationText,
				}).SetUserId(users.Id).
					SetStoreId(questionnaire.StoreId).
					SetTitle("非常感謝您抽空填寫問卷").
					SendByStoreId()

				fcm.SetBody(questionnaire.CongratulationText).SetData(map[string]interface{}{
					"action":       "GoCoupon",
					"id":       	0,
					"type":         "coupon",
					"title":        "非常感謝您抽空填寫問卷",
					"body":         questionnaire.CongratulationText,
				}).SetUserId(users.Id).
					SetTitle("非常感謝您抽空填寫問卷").
					SetAppId(1).
					SetFcmServerKey(config.ServerInfo.OinAppFcmKey).
					Send()


				////////////////////////////////////////
			}
		}

		adminIds := questionnaireAnswerAdmin.GetAdminIdList(questionnaireAnswerIdList)

		//判斷是否需要通知問卷發送者
		if questionnaireAnswer.HaveNoticeSender(questionnaireAnswerIdList) {
			adminIds = append(adminIds, questionnaireRecord.AdminId)
		}

		//取得問卷QRCode紀錄警示者
		if questionnaireRecord.QrcodeRecordId != 0 {
			questionnaireQrcodeRecordAdmin.SetQrcodeRecordId(questionnaireRecord.QrcodeRecordId).QueryAll(func(rs *model.QuestionnaireQrcodeRecordAdmin) {
				if rs.AdminId != 0 {
					adminIds = append(adminIds, rs.AdminId)
				}
			})
		}

		//針對選項取得後台通知人員id並推播到智慧雲App
		fcm.SetBody("有警示問卷，請前往問卷查看處理。").
			SetAdminIdList(adminIds).
			SetTitle("警示問卷").
			SetFcmServerKey(config.ServerInfo.FcmServerKey).
			SetData(map[string]interface{}{
				"action":           "questionnaire",
				"questionRecordId": questionnaireRecodeId,
				"questionName":     questionnaire.Name,
			}).
			SendByAdminIdList()

		//針對選項取得後台通知人員id並推播到碳佐小幫手App
		fcm.SetBody("有警示問卷，請前往問卷查看處理。").
			SetAdminIdList(adminIds).
			SetTitle("碳佐麻里-警示問卷").
			SetFcmServerKey(config.ServerInfo.CrunFcmServerKey).
			SetData(map[string]interface{}{
				"action":           "questionnaire",
				"questionRecordId": questionnaireRecodeId,
				"questionName":     questionnaire.Name,
			}).
			SendByAdminIdList()

		return util.RS{Message: "填寫成功", Status: true}
	}
}
