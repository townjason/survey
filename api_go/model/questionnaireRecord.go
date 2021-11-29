package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
	"strconv"
	"time"
)

type QuestionnaireRecord struct {
	Id              int64     `table:"id"`
	QuestionnaireId int64     `table:"questionnaire_id"`
	StoreId         int64     `table:"store_id"`
	UserId          int64     `table:"user_id"`
	MealTimeId      int64     `table:"meal_time_id"`
	Code            string    `table:"code"`
	Name            string    `table:"name"`
	Phone           string    `table:"phone"`
	Age             int       `table:"age"`
	Gender          int       `table:"gender"`
	TableNo         string    `table:"table_no"`
	No              string    `table:"no"`
	People          int       `table:"people"`
	Status          int       `table:"status"`
	PreviousStatus  int       `table:"previous_status"`
	IsWrite         bool      `table:"is_write"`
	WriteType       string    `table:"write_type"`
	CreatedAt       time.Time `table:"created_at"`
	UpdatedAt       time.Time `table:"updated_at"`
	EndedAt         time.Time `table:"ended_at"`
	AdminId         int64     `table:"admin_id"`
	QrcodeRecordId  int64     `table:"qrcode_record_id"`
	UseStoreId  	int64     `table:"use_store_id"`
}

func (model *QuestionnaireRecord) SetId(id int64) *QuestionnaireRecord {
	model.Id = id
	return model
}

func (model *QuestionnaireRecord) SetTableNo(tableNo string) *QuestionnaireRecord {
	model.TableNo = tableNo
	return model
}

func (model *QuestionnaireRecord) SetNo(no string) *QuestionnaireRecord {
	model.No = no
	return model
}

func (model *QuestionnaireRecord) SetPeople(people int) *QuestionnaireRecord {
	model.People = people
	return model
}

func (model *QuestionnaireRecord) SetMealTimeId(mealTimeId int64) *QuestionnaireRecord {
	model.MealTimeId = mealTimeId
	return model
}

func (model *QuestionnaireRecord) SetCode(code string) *QuestionnaireRecord {
	model.Code = code
	return model
}

func (model *QuestionnaireRecord) SetAdminId(adminId int64) *QuestionnaireRecord {
	model.AdminId = adminId
	return model
}

func (model *QuestionnaireRecord) SetWriteType(writeType string) *QuestionnaireRecord {
	model.WriteType = writeType
	return model
}

func (model *QuestionnaireRecord) SetUserId(userId int64) *QuestionnaireRecord {
	model.UserId = userId
	return model
}

func (model *QuestionnaireRecord) SetStatus(status int) *QuestionnaireRecord {
	model.Status = status
	return model
}

func (model *QuestionnaireRecord) SetPreviousStatus(previousStatus int) *QuestionnaireRecord {
	model.PreviousStatus = previousStatus
	return model
}

func (model *QuestionnaireRecord) SetStoreId(storeId int64) *QuestionnaireRecord {
	model.StoreId = storeId
	return model
}

func (model *QuestionnaireRecord) SetQuestionnaireId(questionnaireId int64) *QuestionnaireRecord {
	model.QuestionnaireId = questionnaireId
	return model
}

func (model *QuestionnaireRecord) SetQrcodeRecordId(qrcodeRecordId int64) *QuestionnaireRecord {
	model.QrcodeRecordId = qrcodeRecordId
	return model
}

func (model *QuestionnaireRecord) SetName(name string) *QuestionnaireRecord {
	model.Name = name
	return model
}

func (model *QuestionnaireRecord) SetPhone(phone string) *QuestionnaireRecord {
	model.Phone = phone
	return model
}

func (model *QuestionnaireRecord) SetAge(age int) *QuestionnaireRecord {
	model.Age = age
	return model
}

func (model *QuestionnaireRecord) SetGender(gender int) *QuestionnaireRecord {
	model.Gender = gender
	return model
}

func (model *QuestionnaireRecord) SetEndedAt(endedAt time.Time) *QuestionnaireRecord {
	model.EndedAt = endedAt
	return model
}

func (model *QuestionnaireRecord) SetIsWrite(isWrite bool) *QuestionnaireRecord {
	model.IsWrite = isWrite
	return model
}

func (model *QuestionnaireRecord) SetUpdatedAt(updatedAt time.Time) *QuestionnaireRecord {
	model.UpdatedAt = updatedAt
	return model
}

func (model *QuestionnaireRecord) SetUseStoreId(useStoreId int64) *QuestionnaireRecord {
	model.UseStoreId = useStoreId
	return model
}

func (model *QuestionnaireRecord) QueryOne() *QuestionnaireRecord {
	table := Model(model)

	if model.Id != 0 {
		table.Where("id", "=", model.Id)
	}

	if model.QuestionnaireId != 0 {
		table.Where("questionnaire_id", "=", model.QuestionnaireId)
	}

	if model.Phone != "" {
		table.Where("phone", "=", model.Phone)
	}

	if model.IsWrite {
		table.Where("is_write", "=", model.IsWrite)
	}

	if model.UserId > 0 {
		table.Where("user_id", "=", model.UserId)
	}

	log.Error(table.Select([]string{"id", "questionnaire_id", "ended_at", "admin_id", "qrcode_record_id"}).
		Find().
		Scan(&model.Id, &model.QuestionnaireId, &model.EndedAt, &model.AdminId, &model.QrcodeRecordId))
	return model
}

func (model *QuestionnaireRecord) QueryAll() ([]map[string]interface{}, int) {
	var dataList = make([]map[string]interface{}, 0)

	table := Model(model)
	tableCount := Model(model)

	if model.StoreId != 0 {
		table.Where("questionnaire_record.store_id", "=", model.StoreId)
		tableCount.Where("questionnaire_record.store_id", "=", model.StoreId)
	}

	table.Select([]string{"id"}).
		Get(func(rows *sql.Rows) (isBreak bool) {
			log.Error(rows.Scan(&model.Id))
			data := map[string]interface{}{
				"id": model.Id,
			}
			dataList = append(dataList, data)
			return
		})

	count := tableCount.Count() //取得目前資料筆數
	return dataList, count
}

func (model *QuestionnaireRecord) Write(data []ReveiveQuestionnaireTopicData) (error, []interface{}) {
	var (
		questionnaireRecordAnswer QuestionnaireRecordAnswer
		questionnaireAnswerAdmin  QuestionnaireAnswerAdmin
		questionnaireAnswer       QuestionnaireAnswer
		dbError                   error
	)

	questionnaireAnswerIdList := make([]interface{}, 0)

	NewDB().Transaction(func(db Database) {
		local, err := time.LoadLocation("Asia/Taipei") //修改成台北時間
		log.Error(err)
		timeNow, err := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), local)
		log.Error(err)

		//更新問卷資料並判斷是否有異常
		if err := model.SetIsWrite(true).SetUpdatedAt(timeNow).Update([]string{"user_id", "name", "phone", "age", "gender", "is_write", "updated_at"}, db); err != nil {
			dbError = err
			log.Error(err)
			log.Error(db.Rollback())
			return
		} else {
			//判斷新增問卷使用店家是否有異常
			for _, value := range data {
				answer, _ := strconv.Atoi(value.Answer)
				if value.Type == 6 { //若為複選題,則取得所有答案陣列
					for _, answerData := range value.ReveiveQuestionnaireAnswerData {
						if answerData.IsSelect { //判斷此答案是否被勾選
							_, err := questionnaireRecordAnswer.
								SetQuestionnaireAnswerId(answerData.Id).
								SetQuestionnaireAnswerTitle(answerData.InputText). //將勾選的答案標題與輸入框內容串在一起
								SetQuestionnaireRecordId(model.Id).
								SetQuestionnaireTopicId(value.Id).
								SetQuestionnaireTopicTitle(value.Title).
								SetQuestionnaireTopicType(value.Type).
								Insert(db)
							if err != nil {
								dbError = err
								log.Error(err)
								log.Error(db.Rollback())
								return
							}

							//將選項id用array記起來
							questionnaireAnswerIdList = append(questionnaireAnswerIdList, answerData.Id)
						}
					}
				} else {
					if value.Type == 3 && answer != 0 { //若為星級選項,則找出回答選項所對應的題目選項id
						questionnaireRecordAnswer.SetQuestionnaireAnswerId(value.ReveiveQuestionnaireAnswerData[answer-1].Id)
						questionnaireRecordAnswer.SetQuestionnaireAnswerTitle(value.Answer)
					} else if value.Type == 5 { //若為自填選項,則預設第一個題目選項id
						questionnaireRecordAnswer.SetQuestionnaireAnswerId(value.ReveiveQuestionnaireAnswerData[0].Id)
						questionnaireRecordAnswer.SetQuestionnaireAnswerTitle(value.Answer)
					} else {
						var inputText string

						//取得輸入框的內容
						for _, answerData := range value.ReveiveQuestionnaireAnswerData {
							if answerData.Id == int64(answer) { //判斷此答案是否被選擇
								inputText = answerData.InputText
							}
						}

						//找出選擇的答案所對應的文字
						questionnaireRecordAnswer.SetQuestionnaireAnswerTitle(inputText)

						questionnaireRecordAnswer.SetQuestionnaireAnswerId(int64(answer))
					}

					//若為自填題目並且未填寫答案時,則不新增資料
					if !(value.Type == 5 && value.Answer == "") {
						_, err := questionnaireRecordAnswer.
							SetQuestionnaireRecordId(model.Id).
							SetQuestionnaireTopicId(value.Id).
							SetQuestionnaireTopicTitle(value.Title).
							SetQuestionnaireTopicType(value.Type).
							Insert(db)
						if err != nil {
							dbError = err
							log.Error(err)
							log.Error(db.Rollback())
							return
						}

						//將選項id用array記起來
						questionnaireAnswerIdList = append(questionnaireAnswerIdList, questionnaireRecordAnswer.QuestionnaireAnswerId)
					}
				}
			}
			//判斷答案是否有通知人員,有的話則將問卷紀錄狀態更新至未處理
			if len(questionnaireAnswerIdList) != 0 && (Model(&questionnaireAnswerAdmin).WhereIn("and", "questionnaire_answer_id", questionnaireAnswerIdList).Count() > 0 || questionnaireAnswer.HaveNoticeSender(questionnaireAnswerIdList)) {
				//更新問卷紀錄狀態
				err = model.SetStatus(1).SetPreviousStatus(1).Update([]string{"status", "previous_status"}, db)
				if err != nil {
					dbError = err
					log.Error(err)
					log.Error(db.Rollback())
					return
				}
			}
		}
		log.Error(db.Commit())
	})

	return dbError, questionnaireAnswerIdList
}

func (model *QuestionnaireRecord) Update(columns []string, database Database) error {
	table := database.Model(model)

	if model.Id != 0 {
		table.Where("id", "=", model.Id)
	}

	return table.Update(columns)
}

func (model *QuestionnaireRecord) Insert() (int64, error) {
	model.CreatedAt = time.Now()
	model.Status = 0

	return Model(model).Insert()
}
