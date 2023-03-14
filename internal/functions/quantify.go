package functions

import (
	"strconv"
)

func Quantify(num string) int {
	cnt, _ := strconv.Atoi(num)
	return cnt
}
