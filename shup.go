package main

import (
	"os"
	"fmt"

	"github.com/taikedz/shell-up/shup"
)

func main() {
	if len(os.Args) > 1 {
		shup.Collate(os.Args[1])
	} else {
		fmt.Printf("Please provide a file\n")
		os.Exit(1)
	}
}