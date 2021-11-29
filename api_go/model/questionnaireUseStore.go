package model

type QuestionnaireUseStore struct {
	Id              int64 `table:"id"`
	QuestionnaireId int64 `table:"questionnaire_id"`
	StoreId         int64 `table:"store_id"`
}

func (model *QuestionnaireUseStore) SetQuestionnaireId(questionnaireId int64) *QuestionnaireUseStore {
	model.QuestionnaireId = questionnaireId
	return model
}