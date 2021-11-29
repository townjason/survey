package model

import (
	. "api/database/mysql"
)

type QuestionnaireRecordAnswer struct {
	Id                       int64  `table:"id"`
	QuestionnaireRecordId    int64  `table:"questionnaire_record_id"`
	QuestionnaireTopicId     int64  `table:"questionnaire_topic_id"`
	QuestionnaireTopicType   int64  `table:"questionnaire_topic_type"`
	QuestionnaireTopicTitle  string `table:"questionnaire_topic_title"`
	QuestionnaireAnswerId    int64  `table:"questionnaire_answer_id"`
	QuestionnaireAnswerTitle string `table:"questionnaire_answer_title"`
}

func (model *QuestionnaireRecordAnswer) SetId(id int64) *QuestionnaireRecordAnswer {
	model.Id = id
	return model
}

func (model *QuestionnaireRecordAnswer) SetQuestionnaireRecordId(questionnaireRecordId int64) *QuestionnaireRecordAnswer {
	model.QuestionnaireRecordId = questionnaireRecordId
	return model
}

func (model *QuestionnaireRecordAnswer) SetQuestionnaireTopicId(questionnaireTopicId int64) *QuestionnaireRecordAnswer {
	model.QuestionnaireTopicId = questionnaireTopicId
	return model
}

func (model *QuestionnaireRecordAnswer) SetQuestionnaireTopicType(questionnaireTopicType int64) *QuestionnaireRecordAnswer {
	model.QuestionnaireTopicType = questionnaireTopicType
	return model
}

func (model *QuestionnaireRecordAnswer) SetQuestionnaireTopicTitle(questionnaireTopicTitle string) *QuestionnaireRecordAnswer {
	model.QuestionnaireTopicTitle = questionnaireTopicTitle
	return model
}

func (model *QuestionnaireRecordAnswer) SetQuestionnaireAnswerId(questionnaireAnswerId int64) *QuestionnaireRecordAnswer {
	model.QuestionnaireAnswerId = questionnaireAnswerId
	return model
}

func (model *QuestionnaireRecordAnswer) SetQuestionnaireAnswerTitle(questionnaireAnswerTitle string) *QuestionnaireRecordAnswer {
	model.QuestionnaireAnswerTitle = questionnaireAnswerTitle
	return model
}

func (model *QuestionnaireRecordAnswer) Insert(database Database) (int64, error) {
	return database.Model(model).Insert()
}
