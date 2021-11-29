package createVerifyCode

import (
	"math/rand"
	"time"
)

const (
	OnlyNumber = 0
	NumberAndLetter = 1
)

/*
	產生驗證碼
*/
func GenerateVerifyCode(codeType int, number int) string{

	choose1 := [10]string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	choose2 := [62]string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z",
	}

	rand.Seed(time.Now().UnixNano())

	code := ""

	// 產生驗證碼部分
	switch codeType {
	case OnlyNumber:
		for i := 0; i < number; i++ {
			randInt := rand.Intn(len(choose1))
			code += choose1[randInt]
		}
		break
	case NumberAndLetter:
		for i := 0; i < number; i++ {
			randInt := rand.Intn(len(choose2))
			code += choose2[randInt]
		}
		break
	}

	return code
}

/*
	產生邀請碼
*/
func GenerateInviteCode() string{
	var code 	string

	chars := [36]string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z",
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 6; i++ {
		randInt := rand.Intn(len(chars))
		code += chars[randInt]
	}

	return code
}