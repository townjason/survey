package factory

import (
	"api/content"
	"api/handler"
	"api/util"
	"api/util/log"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func buildHandlerValue(tempHandlerType reflect.Type) reflect.Value {
	switch tempHandlerType.String() {
	case "handler.QuestionnaireHandler":
        newHandler := reflect.New(tempHandlerType).Elem().Interface().(handler.QuestionnaireHandler)
		return reflect.ValueOf(&newHandler)
	default:
		return reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())
	}
}

func LaunchHandler(tempHandler interface{}, context content.Context, c *gin.Context) []reflect.Value {
	newHandlerValue := buildHandlerValue(reflect.TypeOf(tempHandler).Elem())

	if newHandlerValue.IsNil() {
		log.Error(errors.New("handler not found"))
		c.JSON(http.StatusOK, util.RS{Message: "handler not found", Status: false})
		return nil
	}
	
	newHandlerValueElem := newHandlerValue.Elem()

	if newHandlerValueElem.Kind() == reflect.Struct {
		userId := c.GetInt64("user_id")

		f1 := newHandlerValueElem.FieldByName("Parameter")
		f1.SetString(context.Parameters)

		f2 := newHandlerValueElem.FieldByName("Context")
		f2.Set(reflect.ValueOf(c))

		if userId != 0 {
			f3 := newHandlerValueElem.FieldByName("UserId")
			f3.SetInt(userId)
		}else {
			f3 := newHandlerValueElem.FieldByName("UserId")
			f3.SetInt(0)
		}
	} else {
		c.JSON(http.StatusOK, util.RS{Message: "call api error", Status: false})
		return nil
	}

	return newHandlerValue.MethodByName(context.Action).Call([]reflect.Value{})
}