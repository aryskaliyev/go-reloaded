package internal

import (
	"fmt"
	"os"
	"regexp"
)

func OpenFile(input, output string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9_]+(?P<ext>.txt)$`)

	if re.MatchString(input) {
		fmt.Printf("[+] [%s]: correct name and extension.\n", input)
	} else {
		fmt.Printf("[!] Error: [[%s] incorrect file name (0-9a-zA-Z_) or extension (.txt).]\n\n", input)
		fmt.Println("Program exited.")
		os.Exit(0)
	}

	if re.MatchString(output) {
		fmt.Printf("[+] [%s]: correct name and extension.\n", output)
	} else {
		fmt.Printf("[!] Error: [[%s] incorrect file name (0-9a-zA-Z_) or extension (.txt).]\n\n", output)
		fmt.Println("Program exited.")
		os.Exit(0)
	}

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("[!] Error: [%s]\n\n", err.Error())
		fmt.Println("Program exited.")
		os.Exit(0)
	} else {
		fmt.Println("[+] Data read successfully.")
	}

	return string(data)
}
