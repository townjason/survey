package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
)

type QuestionnaireAnswerAdmin struct {
	Id                    int64 `table:"id"`
	QuestionnaireAnswerId int64 `table:"questionnaire_answer_id"`
	AdminId               int64 `table:"admin_id"`
}

func (model *QuestionnaireAnswerAdmin) SetQuestionnaireAnswerStoreId(questionnaireAnswerId int64) *QuestionnaireAnswerAdmin {
	model.QuestionnaireAnswerId = questionnaireAnswerId
	return model
}

func (model *QuestionnaireAnswerAdmin) GetAdminIdList(questionnaireAnswerIdList []interface{}) []int64 {
	var adminIdList []int64

	if len(questionnaireAnswerIdList) > 0{
		Model(model).Select([]string{"admin_id"}).
			WhereIn("and", "questionnaire_answer_admin.questionnaire_answer_id", questionnaireAnswerIdList).
			Get(func(rows *sql.Rows) (isBreak bool) {
				log.Error(rows.Scan(&model.AdminId))
				adminIdList = append(adminIdList, model.AdminId)
				return
			})
	}
	return adminIdList
}