package internal

import (
	"fmt"
	"os"
)

func SaveFile(output string, ans string) {
	err := os.WriteFile(output, []byte(ans), 0666)
	if err != nil {
		fmt.Printf("[!] Error: [%s]\n\n", err.Error())
		fmt.Println("Program exited.")
		os.Exit(0)
	} else {
		fmt.Println("[+] File saved successfully.")
	}
}
