//+build test

package main

import (
	"api/config"
	"api/database/mysql"
	"github.com/gin-gonic/gin"
	"api/util/log"
)

func init() {
	gin.SetMode(gin.TestMode)
	config.InitConfig()
	mysql.DatabaseOpen()
	log.Clean()
}