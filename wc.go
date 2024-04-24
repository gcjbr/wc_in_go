package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if args were passed
	if len(os.Args) < 2 {
		fmt.Println("You must specify a file after the command")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Check if file exists
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(filePath + ": open: No such file or directory")
	}

	// Ensure the file is closed after the function returns
	defer file.Close()

}
