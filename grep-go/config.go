package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"grep-go/grep"
)

// LoadConfig returns the search input from environment variables
func LoadConfig() grep.SearchInput {
	var (
		filePath = os.Getenv("FILEPATH")
		pattern  = os.Getenv("PATTERN")
	)

	if filePath == "" {
		filePath = getFilePathFromUserInput()
	}
	checkFilePathExists(os.Getenv("FILEPATH"))

	if pattern == "" {
		pattern = getPatternFromUserInput()
	}

	return grep.SearchInput{
		FilePath: filePath,
		Pattern:  pattern,
	}
}

func getFilePathFromUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file path:")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		switch runtime.GOOS {
		// convert CRLF to LF
		case "windows":
			return strings.Replace(text, "\r\n", "", -1)
		default:
			return strings.Replace(text, "\n", "", -1)
		}
	}
}

func getPatternFromUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter pattern to search for:")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		switch runtime.GOOS {
		// convert CRLF to LF
		case "windows":
			return strings.Replace(text, "\r\n", "", -1)
		default:
			return strings.Replace(text, "\n", "", -1)
		}
	}
}

// checkFilePathExists asserts that the chosen value exists on the local file system by panicking if it doesn't
func checkFilePathExists(fp string) {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		panic(fmt.Errorf("chosen option %s does not exist", fp))
	}
}
