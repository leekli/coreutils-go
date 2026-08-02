package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
	"wc/utils"
)

// Global variables
var isStdIn bool = false
var filePathGiven string
var fileContents []byte

func main() {
	// Retrieve any args given by the user
	argsGiven := os.Args

	// First check if a standard input (stdin) pipe has been used
	isStdIn = utils.CheckIsStdIn()

	// If there is something in stdin, then read it and save to file contents
	if (isStdIn) {
		stdInput, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Println("Failed to read from stdin")
		}

		fileContents = stdInput
	}

	// If no stdin, then check and then read the file name/file path given by user and save to file contents
	if (!isStdIn) {
		filePathGiven = argsGiven[len(argsGiven)-1]

		if !utils.CheckFileExists(filePathGiven) {
			panic("File does not exist, or file is not valid.")
		}

		fileContents = utils.ReadFile(filePathGiven)
	}

	// If no flag has been provided, then return the total bytes, lines and words count to user
	if (len(argsGiven) == 1 && isStdIn) || (len(argsGiven) == 2 && !isStdIn) {
		formattedResult := CountWithNoFlags()

		fmt.Println(formattedResult)
	} else if len(argsGiven) >= 2 {
	// Check which flag the user has provided and return the result
		switch argsGiven[1] {
		case "-c":
			// Count Total Bytes
			totalByteCount := CountBytes(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalByteCount, filePathGiven)

			fmt.Println(formattedResult)
		case "-l":
			// Count Total Lines
			totalLineCount := CountLines(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalLineCount, filePathGiven)

			fmt.Println(formattedResult)
		case "-w":
			// Count Total Words
			totalWordCount := CountWords(fileContents)

			formattedResult := fmt.Sprintf("   %d %s", totalWordCount, filePathGiven)

			fmt.Println(formattedResult)
		case "-m":
			// Count Total Characters
			totalCharCount := CountChars(fileContents)

			formattedResult := fmt.Sprintf("    %d %s", totalCharCount, filePathGiven)

			fmt.Println(formattedResult)
		default:
			fmt.Println("Invalid argument flag given.")
		}
	}

}

func CountBytes(fileContents []byte) int {
	totalBytes := len(fileContents)

	return totalBytes
}

func CountLines(fileContents []byte) int {
	lineCount := 0

	for _, byte := range fileContents {
		if byte == '\n' {
			lineCount++
		}
	}

	return lineCount
}

func CountWords(fileContents []byte) int {
	fileContentsToString := string(fileContents)

	fileContentsSplitByWords := strings.Fields(fileContentsToString)

	wordCount := len(fileContentsSplitByWords)

	return wordCount
}

func CountChars(fileContents []byte) int {
	return utf8.RuneCountInString(string(fileContents))
}

func CountWithNoFlags() string {
	totalByteCount := CountBytes(fileContents)
	totalLineCount := CountLines(fileContents)
	totalWordCount := CountWords(fileContents)

	totalCountString := fmt.Sprintf("    %d   %d   %d %s", totalLineCount, totalWordCount, totalByteCount, filePathGiven)

	return totalCountString
}