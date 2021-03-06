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
		return util.RS{Message: "???????????????", Status: false}
	} else if qrCodeDate.Code == "" {
		return util.RS{Message: "?????????????????????", Status: false}
	} else if !questionnaire.SetId(qrCodeDate.QuestionnaireId).QueryOne().IsOpen {
		return util.RS{Message: "??????????????????", Status: false}
	} else if !questionnaire.IsRepeatWrite && data.Phone != "" && questionnaireRecord.SetQuestionnaireId(qrCodeDate.QuestionnaireId).SetPhone(data.Phone).QueryOne().Id != 0 {
		return util.RS{Message: "?????????????????????", Status: false}
	} else if userId := users.SetPhone(data.Phone).SetCountryCode("+886").GetUsersIdByPhone().Id; userId <= 0 {
		return util.RS{Message: "????????????", Status: false}
	} else {
		//????????????????????????
		if questionnaire.IsLimitTime {
			questionnaireRecord.SetEndedAt(time.Now().Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(questionnaire.LimitTime) + time.Second*time.Duration(0)))
		}

		//????????????????????????
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
				Insert(); err != nil { //????????????
			return util.RS{Message: "??????????????????", Status: true}
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
		fcm                            FCM //??????????????????
		couponDataList                 []coupon
	)


	if err := json.Unmarshal([]byte(handler.Parameter), &data); err != nil {
		return util.RS{Message: "", Status: false}
	} else if questionnaireRecodeIdString, err := crypto.KeyDecrypt(data.QuestionnaireRecodeId); err != nil {
		return util.RS{Message: "", Status: false}
	} else if questionnaireRecodeId, err := strconv.Atoi(questionnaireRecodeIdString); err != nil {
		return util.RS{Message: "", Status: false}
	} else if !questionnaire.SetId(questionnaireRecord.SetId(int64(questionnaireRecodeId)).QueryOne().QuestionnaireId).QueryOne().IsOpen {
		return util.RS{Message: "??????????????????", Status: false}
	} else if data.Name = TrimZero(data.Name); data.Name == "" && questionnaire.GetRequire("name", questionnaire.BuildInItem) {
		return util.RS{Message: "???????????????", Status: false}
	} else if data.Phone = TrimZero(data.Phone); data.Phone == "" && questionnaire.GetRequire("phone", questionnaire.BuildInItem) {
		return util.RS{Message: "???????????????", Status: false}
	} else if data.Gender <= 0 && questionnaire.GetRequire("gender", questionnaire.BuildInItem) {
		return util.RS{Message: "???????????????", Status: false}
	} else if data.Age <= 0 && questionnaire.GetRequire("old", questionnaire.BuildInItem) {
		return util.RS{Message: "????????????", Status: false}
	} else if _, status := validate.CheckPhone(data.Phone); questionnaire.GetRequire("phone", questionnaire.BuildInItem) && !status {
		return util.RS{Message: "??????????????????", Status: false}
	} else if !questionnaire.IsRepeatWrite && questionnaireRecordCheck.SetIsWrite(true).SetId(0).SetQuestionnaireId(questionnaireRecord.QuestionnaireId).SetPhone(data.Phone).QueryOne().Id != 0 {
		return util.RS{Message: "????????????????????????", Status: false}
	} else if questionnaire.IsLimitTime && questionnaireRecord.EndedAt.Before(time.Now()) {
		return util.RS{Message: "???????????????", Status: false}
	} else if userId := users.SetPhone(data.Phone).SetCountryCode("+886").GetUsersIdByPhone().Id; userId <= 0 {
		return util.RS{Message: "????????????", Status: false}
	} else if err, questionnaireAnswerIdList :=
		questionnaireRecord.
			SetPhone(data.Phone).
			SetName(data.Name).
			SetAge(data.Age).
			SetGender(data.Gender).
			SetUserId(userId).
			Write(data.ReveiveQuestionnaireTopicData); err != nil { //????????????
		return util.RS{Message: "????????????", Status: true}
	} else {
		log.Error(json.Unmarshal([]byte(questionnaire.Coupon), &couponDataList))

		//???????????????????????????????????????????????????????????????????????????????????????????????????
		if len(couponDataList) > 0 && users.Id != 0 {
			count := 0
			//??????uuid
			out, _ := exec.Command("uuidgen").Output()
			uuid := string(out)
			//????????????????????????(????????????uuid
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

			//kafka?????????????????????
			if dataList["status"] == true {
				//???????????????app??????///////////////////////
				fcm.SetBody(questionnaire.CongratulationText).SetData(map[string]interface{}{
					"action":       "GoCoupon",
					"id":       	0,
					"type":         "coupon",
					"title":        "?????????????????????????????????",
					"body":         questionnaire.CongratulationText,
				}).SetUserId(users.Id).
					SetStoreId(questionnaire.StoreId).
					SetTitle("?????????????????????????????????").
					SendByStoreId()

				fcm.SetBody(questionnaire.CongratulationText).SetData(map[string]interface{}{
					"action":       "GoCoupon",
					"id":       	0,
					"type":         "coupon",
					"title":        "?????????????????????????????????",
					"body":         questionnaire.CongratulationText,
				}).SetUserId(users.Id).
					SetTitle("?????????????????????????????????").
					SetAppId(1).
					SetFcmServerKey(config.ServerInfo.OinAppFcmKey).
					Send()


				////////////////////////////////////////
			}
		}

		adminIds := questionnaireAnswerAdmin.GetAdminIdList(questionnaireAnswerIdList)

		//???????????????????????????????????????
		if questionnaireAnswer.HaveNoticeSender(questionnaireAnswerIdList) {
			adminIds = append(adminIds, questionnaireRecord.AdminId)
		}

		//????????????QRCode???????????????
		if questionnaireRecord.QrcodeRecordId != 0 {
			questionnaireQrcodeRecordAdmin.SetQrcodeRecordId(questionnaireRecord.QrcodeRecordId).QueryAll(func(rs *model.QuestionnaireQrcodeRecordAdmin) {
				if rs.AdminId != 0 {
					adminIds = append(adminIds, rs.AdminId)
				}
			})
		}

		//????????????????????????????????????id?????????????????????App
		fcm.SetBody("????????????????????????????????????????????????").
			SetAdminIdList(adminIds).
			SetTitle("????????????").
			SetFcmServerKey(config.ServerInfo.FcmServerKey).
			SetData(map[string]interface{}{
				"action":           "questionnaire",
				"questionRecordId": questionnaireRecodeId,
				"questionName":     questionnaire.Name,
			}).
			SendByAdminIdList()

		//????????????????????????????????????id???????????????????????????App
		fcm.SetBody("????????????????????????????????????????????????").
			SetAdminIdList(adminIds).
			SetTitle("????????????-????????????").
			SetFcmServerKey(config.ServerInfo.CrunFcmServerKey).
			SetData(map[string]interface{}{
				"action":           "questionnaire",
				"questionRecordId": questionnaireRecodeId,
				"questionName":     questionnaire.Name,
			}).
			SendByAdminIdList()

		return util.RS{Message: "????????????", Status: true}
	}
}
