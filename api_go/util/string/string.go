package string

import (
	"strings"
)

func TrimQuotes(s string) string {
	return strings.Replace(strings.Replace(s, "\"", "\\", -1), "'", "\\'", -1)
}

