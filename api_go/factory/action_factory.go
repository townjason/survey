package factory

import (
	"api/handler"
)

var (
	questionnaireHandler handler.QuestionnaireHandler
)

var ActionFactoryAuth = map[string]interface{}{

}

var ActionFactory = map[string]interface{}{
	//======================================================================
	//				Questionnaire Api
	//======================================================================
	"GetQuestionnaireInfo":      &questionnaireHandler,
	"InsertQuestionnaireRecord": &questionnaireHandler,
}
