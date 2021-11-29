package sms

import (
	"api/util/log"
	"fmt"
	"net/http"
	"strings"
)

const (
	// 內文長度
	LongContent   = 0
	NormalContent = 1
)

/*
	寄送簡訊
*/
func SendSMS(phone string, sms string, contentType int) (isOK bool, statusCode int) {

	var phoneurl string
	switch contentType {
	case LongContent:
		phoneurl = "http://smexpress.mitake.com.tw:7002/SpLmGet"
	case NormalContent:
		phoneurl = "http://smexpress.mitake.com.tw:9600/SmSendGet.asp"
	}
	userName := "59318165"
	password := "usj394018"
	dstaddr := phone
	encoding := "UTF8"
	smbody := sms
	response := "http://192.168.1.200/smreply.asp"
	//message :=
	var builder strings.Builder

	fmt.Fprintf(&builder, "%s?username=%s&password=%s&dstaddr=%s&encoding=%s&smbody=%s&response=%s&CharsetURL=utf-8", phoneurl, userName, password, dstaddr, encoding, smbody, response)

	fmt.Println(builder.String())

	resp, err := http.Get(builder.String())
	defer resp.Body.Close()

	if err != nil {
		log.Error(err)
		statusCode = 404
	} else {
		if resp.StatusCode == 200 {
			isOK = true
		}
		statusCode = resp.StatusCode
	}

	return
}