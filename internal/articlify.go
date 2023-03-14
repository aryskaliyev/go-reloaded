package internal

func Articlify(arr []string) []string {
	for i := range arr {
		if i+1 < len(arr) {
			if (arr[i] == "A" || arr[i] == "a") && (arr[i+1][0] == 'a' || arr[i+1][0] == 'e' || arr[i+1][0] == 'i' || arr[i+1][0] == 'o' || arr[i+1][0] == 'u' || arr[i+1][0] == 'h') {
				arr[i] = arr[i] + "n"
			}
		}
	}
	return arr
}
