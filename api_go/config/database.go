package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configDatabase struct {
	Host string
	Port string
	Name string
	Username string
	Password string
}

var Database configDatabase

func initConfigDatabase() {

	switch gin.Mode() {
	case gin.ReleaseMode:
		Database.Host = "34.80.238.245"
		Database.Port = "3306"
		Database.Name = "oin"
		Database.Username = "oin"
		Database.Password = "mys20105602qL"
	case gin.DebugMode:
		Database.Host = os.Getenv("DB_HOST")
		Database.Port = os.Getenv("DB_PORT")
		Database.Name = os.Getenv("DB_NAME")
		Database.Username = os.Getenv("DB_USER")
		Database.Password = os.Getenv("DB_PASSWORD")
	case gin.TestMode:
		Database.Host = "192.168.50.111"
		Database.Port = "3306"
		Database.Name = "oin"
		Database.Username = "oin"
		Database.Password = "oin"
	}
}