package main

import (
	"fmt"

	"grep-go/grep"
)

func main() {
	config := LoadConfig()

	fmt.Println(grep.SearchFile(&grep.SearchInput{
		FilePath: config.FilePath,
		Pattern:  config.Pattern,
	}))
}
