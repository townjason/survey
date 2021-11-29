package main

import (
	"api/config"
	"api/database/mysql"
	"api/content"
	"api/factory"
	"api/util"
	"api/util/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func init() {
	_ = godotenv.Load(".env")
	gin.SetMode(gin.DebugMode)
	config.InitConfig()
	mysql.DatabaseOpen()
	log.Clean()
}
   
func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET","POST"},
		AllowHeaders:    []string{"Content-Type", "Access-Control-Allow-Origin"},
	}))

	apiRouter := router.Group("/api")

	//apiRouter.Use(middle_ware.AuthToken())
	//{
	//	apiRouter.POST("/auth", api)
	//	apiRouter.POST("", api)
	//}

	apiRouter.POST("", api)

	log.Error(router.Run(":" + config.ServerInfo.Port))
}

func api(c *gin.Context) {
	var context content.Context
	userId := c.GetInt64("user_id")
	if err := c.ShouldBindJSON(&context); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, util.RS{Message: "should bind JSON error", Status: false})
		return
	} else if a, ok := factory.ActionFactoryAuth[context.Action]; ok && userId <= 0 {
		c.JSON(http.StatusOK, util.RS{Message: "api auth failure", Status: false})
		return
	} else {
		if !ok {
			a = factory.ActionFactory[context.Action]
		}

		// at := reflect.TypeOf(a).Elem()
		// a = reflect.New(at).Interface()
		// s := reflect.ValueOf(a).Elem()
		
		// if s.Kind() == reflect.Struct {
			// f1 := s.FieldByName("Parameter")
			// f1.SetString(context.Parameters)

			// f2 := s.FieldByName("Context")
			// f2.Set(reflect.ValueOf(c))

			// if userId != 0 {
			// 	f3 := s.FieldByName("UserId")
			// 	f3.SetInt(userId)
			// }else {
			// 	f3 := s.FieldByName("UserId")
			// 	f3.SetInt(0)
			// }
		// } else {
		// 	c.JSON(http.StatusOK, util.RS{Message: "call api error", Status: false})
		// 	return
		// }

		result := factory.LaunchHandler(a, context, c)

		if len(result) > 0 {
			c.SecureJSON(http.StatusOK, result[0].Interface())
			return
		}
	}
}
