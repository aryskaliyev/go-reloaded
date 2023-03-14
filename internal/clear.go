package internal

func Clear(arr, types []string) ([]string, []string) {
	for i := 0; i < len(types); i++ {
		if types[i] == "hex" || types[i] == "low" || types[i] == "cap" || types[i] == "up" || types[i] == "bin" {
			if i+1 < len(types) {
				types = append(types[:i], types[i+1:]...)
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr, types
}
