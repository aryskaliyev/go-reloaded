package internal

import (
	"go-reloaded/internal/functions"
)

func Functify(arr, types []string) []string {
	for i, v := range types {
		if v == "bin" {
			cnt := functions.Quantify(arr[i])
			arr[i] = ""

			for j := 1; j <= cnt; j++ {
				if i-j >= 0 {
					if types[i-j] == "word" {
						arr[i-j] = functions.Bin2Dec(arr[i-j])
					} else {
						cnt++
					}
				} else {
					break
				}
			}
		} else if v == "hex" {
			cnt := functions.Quantify(arr[i])
			arr[i] = ""

			for j := 1; j <= cnt; j++ {
				if i-j >= 0 {
					if types[i-j] == "word" {
						arr[i-j] = functions.Hex2Dec(arr[i-j])
					} else {
						cnt++
					}
				} else {
					break
				}
			}
		} else if v == "up" {
			cnt := functions.Quantify(arr[i])
			arr[i] = ""

			for j := 1; j <= cnt; j++ {
				if i-j >= 0 {
					if types[i-j] == "word" {
						arr[i-j] = functions.Up(arr[i-j])
					} else {
						cnt++
					}
				} else {
					break
				}
			}
		} else if v == "low" {
			cnt := functions.Quantify(arr[i])
			arr[i] = ""

			for j := 1; j <= cnt; j++ {
				if i-j >= 0 {
					if types[i-j] == "word" {
						arr[i-j] = functions.Low(arr[i-j])
					} else {
						cnt++
					}
				} else {
					break
				}
			}
		} else if v == "cap" {
			cnt := functions.Quantify(arr[i])
			arr[i] = ""

			for j := 1; j <= cnt; j++ {
				if i-j >= 0 {
					if types[i-j] == "word" {
						arr[i-j] = functions.Cap(arr[i-j])
					} else {
						cnt++
					}
				} else {
					break
				}
			}
		}
	}

	return arr
}
