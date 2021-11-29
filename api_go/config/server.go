package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configServer struct {
	Port         string
	Host         string
	FileHost     string
	FcmServerKey string
	CrunFcmServerKey string
	OinAppFcmKey string
}

var ServerInfo configServer

func initConfigServerInfo() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		ServerInfo.Port = "9100"
		ServerInfo.Host = "https://file.oinapp.com/"
		ServerInfo.FileHost = "https://file.oinapp.com/"
		ServerInfo.FcmServerKey = "AAAABYBLxsM:APA91bGtTC7L7AU3-YQO3XP5z-qcM7mY6dt0AYWI8hmrJgPawZDSBm6SdIdZFBfIRGUdkIEF0MiW_F9PgYHdo4b-GYB2_SlAR-39SesaLDTgAShgc3wHVYTchTtbq5TFgW28brCfONwn"
		ServerInfo.CrunFcmServerKey = "AAAA-bBV7Os:APA91bEBGZOpakYLKLlSS07N1SHunuBluxcE0U4F4j6AcRBZdB3YneDPOB2JXIcSdyB6gto-UItRaFd85_Q7av_tpQIkWB-v8SdrCXxA0TxsYElx0NKJDBXfOYOS-bWv_j46XHv127Ej"
		ServerInfo.OinAppFcmKey = "AAAA8jRoxL4:APA91bHTwZdShwBuWVhdsLnE192FAbgXVJ5Kb7JQcTixha-A7PPAhxuB0JZy-qcm-wRjw8JxLkbfJEX02NQIOU05K574cDbdbl7q3SkEbz9lSvBVR7mFG7nh4IxIhS10kW_g0KXNccQV"
	case gin.DebugMode:
		ServerInfo.Port = os.Getenv("PORT")
		ServerInfo.Host = os.Getenv("HOST")
		ServerInfo.FileHost = os.Getenv("FILE_HOST")
		ServerInfo.FcmServerKey = "AAAAO_7fpGk:APA91bF8l3Gwnfni4b8ly2V6UncmuU-dUPRH1eeyWx4Oyk8rZDXxf8Vn4RLsskVDGPspTXoNpDRxXyjBNEMhNSE8lyJcnSpkT8DFSM7o9qsNFEPx1NaRNf6jrzpC-Loplk1G1CPagdyA"
		ServerInfo.CrunFcmServerKey = "AAAA-bBV7Os:APA91bEBGZOpakYLKLlSS07N1SHunuBluxcE0U4F4j6AcRBZdB3YneDPOB2JXIcSdyB6gto-UItRaFd85_Q7av_tpQIkWB-v8SdrCXxA0TxsYElx0NKJDBXfOYOS-bWv_j46XHv127Ej"
		ServerInfo.OinAppFcmKey = "AAAA8jRoxL4:APA91bHTwZdShwBuWVhdsLnE192FAbgXVJ5Kb7JQcTixha-A7PPAhxuB0JZy-qcm-wRjw8JxLkbfJEX02NQIOU05K574cDbdbl7q3SkEbz9lSvBVR7mFG7nh4IxIhS10kW_g0KXNccQV"
	case gin.TestMode:
		ServerInfo.Port = "9100"
		ServerInfo.Host = "http://dev.oinapp.com/"
		ServerInfo.FileHost = "http://file-test.oinapp.com/"
		ServerInfo.FcmServerKey = "AAAAO_7fpGk:APA91bF8l3Gwnfni4b8ly2V6UncmuU-dUPRH1eeyWx4Oyk8rZDXxf8Vn4RLsskVDGPspTXoNpDRxXyjBNEMhNSE8lyJcnSpkT8DFSM7o9qsNFEPx1NaRNf6jrzpC-Loplk1G1CPagdyA"
		ServerInfo.CrunFcmServerKey = "AAAA-bBV7Os:APA91bEBGZOpakYLKLlSS07N1SHunuBluxcE0U4F4j6AcRBZdB3YneDPOB2JXIcSdyB6gto-UItRaFd85_Q7av_tpQIkWB-v8SdrCXxA0TxsYElx0NKJDBXfOYOS-bWv_j46XHv127Ej"
		ServerInfo.OinAppFcmKey = "AAAA8jRoxL4:APA91bHTwZdShwBuWVhdsLnE192FAbgXVJ5Kb7JQcTixha-A7PPAhxuB0JZy-qcm-wRjw8JxLkbfJEX02NQIOU05K574cDbdbl7q3SkEbz9lSvBVR7mFG7nh4IxIhS10kW_g0KXNccQV"
	}
}
