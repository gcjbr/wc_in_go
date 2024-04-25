package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Check if args were passed
	if len(os.Args) < 2 {
		fmt.Println("You must specify a file after the command")
		os.Exit(1)
	}

	filePath := os.Args[len(os.Args)-1]

	// Check if file exists
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(filePath + ": open: No such file or directory")
	}

	// Ensure the file is closed after the function returns
	defer file.Close()

	// Parse flags
	c := flag.Bool("c", false, "Count characters in the file")

	flag.Parse()

	if c != nil && *c {

		characters, err := countCharacters(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(characters)
	}

}

func countCharacters(file *os.File) (int, error) {

	var count int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil

}
