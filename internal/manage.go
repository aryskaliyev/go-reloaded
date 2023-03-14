package internal

import (
	"fmt"
	"strings"
)

func Manage(input, output string) {
	data := OpenFile(input, output)

	newlines := strings.Split(data, "\n")

	res := []string{}

	for _, str := range newlines {
		arr, types := Split(str)
		arr = Functify(arr, types)
		arr, types = Clear(arr, types)
		arr = Articlify(arr)
		spaces := Spacify(types)
		for i := range arr {
			fmt.Println(arr[i], types[i])
		}
		ans := Stringify(arr, spaces)
		res = append(res, ans)
	}
	SaveFile(output, strings.Join(res, "\n"))
}
