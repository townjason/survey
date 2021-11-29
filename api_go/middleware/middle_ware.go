package middleware

import (
	"api/util"
	"api/util/crypto"
	"api/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const CsrfToken = "c20d19eb0fa930"

// 判斷是否 使用者登入的 middleware
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int64
		status := true
		token := c.GetHeader("Auth-Token")
		/** check token is empty or null or undefined */
		if c.Request.RequestURI == "/api/auth" && (token == "" || token == "null") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token != ""{
			// 解析 Token
			id, status = ParseToken(token)
		}

		if status {
			c.Set("user_id", id)
			c.Next()
			return
		} else {
			// 未授權的身份
			c.AbortWithStatusJSON(http.StatusOK, util.RS{Message: "logout", Status: false})
			return
		}
	}
}

func verifyCsrfToken(token string) bool {
	if token == "" {
		return false
	}

	tokenInfo, err := crypto.KeyDecrypt(token)
	if err != nil {
		log.Error(err)
		return false
	}

	return tokenInfo == CsrfToken
}

// go run main.go csrf:create
func GenerateCsrfToken() {
	token, err := crypto.KeyEncrypt(CsrfToken)
	log.Error(err)
	fmt.Print(token)
}

func GenerateToken(userId int64) (string, error) {
	temp := strconv.FormatInt(userId, 10) + ";" + time.Now().Format("20060102150405")
	token, err := crypto.KeyEncrypt(temp)
	if err != nil {
		return "", err
	}

	token, err = crypto.KeyEncrypt(token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (int64,  bool) {
	if token == "" {
		return 0, false
	}


	tokenInfo, err := crypto.KeyDecrypt(token)
	if err != nil {
		log.Error(err)
		return 0, false
	}

	spiltStr := strings.Split(tokenInfo, ";")

	if len(spiltStr) == 2 {
		userId := spiltStr[0]

		userId64, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			log.Error(err)
			return 0, false
		}

		return userId64, true
	} else {
		return 0, false
	}
}
