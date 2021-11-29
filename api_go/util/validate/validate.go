package validate

import (
	"math/rand"
	"regexp"
)

// TODO 格式有問題
func CheckEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@+[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

//func CheckPassword(password string) bool {
//	Re := regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`)
//	return Re.MatchString(password)
//}
func CheckPassword(password string) bool {
	numberRe := regexp.MustCompile(`[0-9]`).MatchString(password)
	englishRe := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	if numberRe && englishRe {
		Re := regexp.MustCompile(`^[a-zA-Z0-9]{6,12}$`)
		return Re.MatchString(password)
	} else {
		return false
	}
}

func BuildInvitationCode() (invitationCode string) {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}



func CheckPayPassword(password string) bool {
	Re := regexp.MustCompile(`^[0-9]{5,6}$`)
	return Re.MatchString(password)
}

func CheckAccount(account string) bool {
	Re := regexp.MustCompile(`^[0-9]`)
	return Re.MatchString(account)
}

func CheckPhone(phone string) (string, bool) {
	// 判斷是否大於10碼
	Re := regexp.MustCompile(`^[0-9]{10,}$`)
	exist := Re.MatchString(phone)
	if exist {
		return "", true
	} else {
		return "電話號碼格式錯誤", false
	}

}

func CheckVerifyCode(code string) (string, bool) {
	// 判斷是否為4碼
	Re := regexp.MustCompile(`^[0-9]{4,}$`)
	exist := Re.MatchString(code)
	if exist {
		return "", true
	} else {
		return "驗證碼格式錯誤", false
	}

}
