package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
)

type QuestionnaireAnswer struct {
	Id                   int64  `table:"id"`
	QuestionnaireTopicId int64  `table:"questionnaire_topic_id"`
	Title                string `table:"title"`
	Sort                 int64  `table:"sort"`
	IsNoticeSender       bool   `table:"is_notice_sender"`
	IsShowInput          bool   `table:"is_show_input"`
	IsUpdate             bool   `table:"is_update"`
}

type ReveiveQuestionnaireAnswerData struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	IsSelect  bool   `json:"isSelect"`
	InputText string `json:"inputText"`
}

func (model *QuestionnaireAnswer) SetId(id int64) *QuestionnaireAnswer {
	model.Id = id
	return model
}

func (model *QuestionnaireAnswer) SetQuestionnaireTopicId(questionnaireTopicId int64) *QuestionnaireAnswer {
	model.QuestionnaireTopicId = questionnaireTopicId
	return model
}

func (model *QuestionnaireAnswer) QueryOne() *QuestionnaireAnswer {
	log.Error(Model(model).Select([]string{"title"}).
		Where("id", "=", model.Id).
		Find().
		Scan(&model.Title))
	return model
}

func (model *QuestionnaireAnswer) GetListByQuestionnaireTopicId() []map[string]interface{} {
	var dataList = make([]map[string]interface{}, 0)
	Model(model).Select([]string{"id", "title", "sort", "is_show_input"}).
		Where("questionnaire_topic_id", "=", model.QuestionnaireTopicId).
		OrderBy([]string{"sort"}, []string{"asc"}).
		Get(func(rows *sql.Rows) (isBreak bool) {
			log.Error(rows.Scan(&model.Id, &model.Title, &model.Sort, &model.IsShowInput))

			var data = map[string]interface{}{
				"id":          model.Id,
				"title":       model.Title,
				"isSelect":    false,
				"inputText":   "",
				"isShowInput": model.IsShowInput,
			}
			dataList = append(dataList, data)
			return
		})
	return dataList
}

func (model *QuestionnaireAnswer) HaveNoticeSender(questionnaireAnswerIdList []interface{}) (isNoticeSender bool) {
	if len(questionnaireAnswerIdList) > 0 {
		Model(model).Select([]string{"is_notice_sender"}).
			WhereIn("and", "questionnaire_answer.id", questionnaireAnswerIdList).
			Get(func(rows *sql.Rows) (isBreak bool) {
				log.Error(rows.Scan(&model.IsNoticeSender))

				if model.IsNoticeSender{
					isNoticeSender = true
				}

				return
			})
	}
	return isNoticeSender
}
