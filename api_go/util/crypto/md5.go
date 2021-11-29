package crypto

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	if str == "" {
		return ""
	}
	data := []byte(str + key)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

