package crypto

import (
	"math/rand"
	"time"
)

func generateCode(lenght int, chars []string) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < lenght; i++ {
		randInt := rand.Intn(len(chars))
		code += chars[randInt]
	}
	return code
}

// 產生隨機碼驗證(大小寫英文+數字)
func GenerateRandomCode(lenght int) string {
	return generateCode(lenght, []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z",
	})
}

// 產生隨機碼驗證(數字)
func GenerateRandomNumber(lenght int) string {
	return generateCode(lenght, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",})
}
