//+build test

package main

import (
	"api/config"
	"api/database/mysql"
	"api/util/log"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.InitConfig()
	mysql.DatabaseOpen()
	log.Clean()
}