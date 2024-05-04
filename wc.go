package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	// Check if args has any flags

	if len(os.Args) < 3 {
		printUsage()
	}

	var responses [4]int

	//countLines := flag.Bool("l", false, "Count lines")
	// countWords := flag.Bool("w", false, "Count words")
	countBytes := flag.Bool("c", false, "Count bytes")
	// countChars := flag.Bool("m", false, "Count characters")
	flag.Parse() // Parse the flags

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

	if countBytes != nil && *countBytes {
		characters, err := countCharacters(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		responses[2] = characters
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

func printUsage() {
	fmt.Println("Usage: [flags] [filename]")
	fmt.Println("Flags:")
	fmt.Println("  -c  Count bytes")
	fmt.Println("  -l  Count lines")
	fmt.Println("  -m  Count characters")
	fmt.Println("  -w  Count words")

	os.Exit(0)
}
