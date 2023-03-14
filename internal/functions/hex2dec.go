package functions

import (
	"regexp"
	"strconv"
)

func Hex2Dec(word string) string {
	if regexp.MustCompile(`[0-9a-fA-F]+`).MatchString(word) {
		tmp, _ := strconv.ParseInt(word, 16, 64)
		word = strconv.FormatInt(tmp, 10)
	}
	return word
}
