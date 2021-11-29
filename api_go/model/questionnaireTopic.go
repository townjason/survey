package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
)

type QuestionnaireTopic struct {
	Id              int64  `table:"id"`
	QuestionnaireId int64  `table:"questionnaire_id"`
	Title           string `table:"title"`
	Type            int64  `table:"type"`
	InputType       string `table:"input_type"`
	IsUpdate        bool   `table:"is_update"`
	IsRequired      bool   `table:"is_required"`
}

type ReveiveQuestionnaireTopicData struct {
	Id                             int64                            `json:"id"`
	Title                          string                           `json:"title"`
	Type                           int64                            `json:"type"`
	Answer                         string                           `json:"answer"`
	ReveiveQuestionnaireAnswerData []ReveiveQuestionnaireAnswerData `json:"questionnaireAnswer"`
}

func (model *QuestionnaireTopic) SetId(id int64) *QuestionnaireTopic {
	model.Id = id
	return model
}

func (model *QuestionnaireTopic) SetQuestionnaireId(questionnaireId int64) *QuestionnaireTopic {
	model.QuestionnaireId = questionnaireId
	return model
}

func (model *QuestionnaireTopic) GetListByQuestionnaireId() []map[string]interface{} {
	var questionnaireAnswer QuestionnaireAnswer

	var dataList = make([]map[string]interface{}, 0)
	Model(model).Select([]string{"id", "title", "type", "input_type", "is_required"}).
		Where("questionnaire_id", "=", model.QuestionnaireId).
		Get(func(rows *sql.Rows) (isBreak bool) {
			log.Error(rows.Scan(&model.Id, &model.Title, &model.Type, &model.InputType, &model.IsRequired))
			data := map[string]interface{}{
				"id":                  model.Id,
				"title":               model.Title,
				"type":                model.Type,
				"inputType":           model.InputType,
				"isRequired":          model.IsRequired,
				"answer":              "",
				"questionnaireAnswer": questionnaireAnswer.SetQuestionnaireTopicId(model.Id).GetListByQuestionnaireTopicId(),
			}
			if model.Type == 3 {
				data["answer"] = 0
			}
			dataList = append(dataList, data)
			return
		})
	return dataList
}
