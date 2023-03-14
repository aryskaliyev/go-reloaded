package functions

import (
	"regexp"
	"strconv"
)

func Bin2Dec(word string) string {
	if regexp.MustCompile(`[01]+`).MatchString(word) {
		tmp, _ := strconv.ParseInt(word, 2, 64)
		word = strconv.FormatInt(tmp, 10)
	}
	return word
}
