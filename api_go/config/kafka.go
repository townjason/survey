package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configKafka struct {
	Ip string
	Branch string
}

var Kafka configKafka

func initConfigKafka() {
	if len(os.Args) == 2 {
		Kafka.Branch = os.Args[1]
	}

	switch gin.Mode() {
	case gin.ReleaseMode:
		if len(os.Args) == 2 {
			if os.Args[1] == "master_app_com" {
				Kafka.Ip = "192.168.50.111:9092"
			}
		} else {
			Kafka.Ip = "10.140.0.8:9092"
		}		
	case gin.DebugMode:
		Kafka.Ip = os.Getenv("KAFKA_IP")
		Kafka.Branch = os.Getenv("BRANCH")
	case gin.TestMode:
		Kafka.Ip = "192.168.50.111:9092"
	}
}