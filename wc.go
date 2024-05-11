package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type response struct {
	lines int
	words int
	bytes int
	chars int
}

type flags struct {
	countLines bool
	countWords bool
	countBytes bool
	countChars bool
}

func main() {

	// Check if args has any flags

	if len(os.Args) < 3 {
		printUsage()
	}

	expandFlags()

	var responses response
	var flags flags

	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countBytes := flag.Bool("c", false, "Count bytes")
	countChars := flag.Bool("m", false, "Count characters")

	flag.Parse() // Parse the flags

	flags.countLines = *countLines
	flags.countWords = *countWords
	flags.countBytes = *countBytes
	flags.countChars = *countChars

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

	// Check flags

	if flags.countLines {
		lines, err := countLinesFromFile(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		responses.lines = lines
	}

	if flags.countWords {
		words, err := wordsCount(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		responses.words = words
	}

	if flags.countBytes {
		bytes, err := countCharBytes(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		responses.bytes = bytes
	}

	if flags.countChars {
		characters, err := countCharacters(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		responses.chars = characters
	}

	//printResponses(responses, flags)

}

func wordsCount(file *os.File) (int, error) {
	wordCount := 0
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

func countCharacters(file *os.File) (int, error) {
	file.Seek(0, 0)
	charCount := 0

	reader := bufio.NewReader(file)

	for {
		_, _, err := reader.ReadRune()

		charCount++
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}

	}

	return charCount, nil

}

func countLinesFromFile(file *os.File) (int, error) {
	file.Seek(0, 0)
	count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil

}

func countCharBytes(file *os.File) (int, error) {
	file.Seek(0, 0)
	count := 0

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

func expandFlags() {
	originalArgs := os.Args[1:]
	expandedArgs := []string{}

	for _, arg := range originalArgs {
		if strings.HasPrefix(arg, "-") && len(arg) > 2 {
			for _, ch := range arg[1:] {
				expandedArgs = append(expandedArgs, "-"+string(ch))
			}
		} else {
			expandedArgs = append(expandedArgs, arg)
		}
	}

	os.Args = append([]string{os.Args[0]}, expandedArgs...)
}
