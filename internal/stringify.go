package internal

func Stringify(arr []string, spaces []bool) string {
	ans := ""

	for i := range arr {
		ans += arr[i]
		if spaces[i] {
			ans += " "
		}
	}
	return ans
}
