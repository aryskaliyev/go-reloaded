package main

import (
	"fmt"
	"go-reloaded/internal"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 3 {
		internal.Manage(args[1], args[2])
	} else {
		fmt.Println("[!] Error: [Provide 2 arguments.]")
	}
}
