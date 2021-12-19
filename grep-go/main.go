package main

import (
	"grep-go/grep"
)

func main() {
	config := LoadConfig()

	_, _ = grep.SearchFile(&grep.SearchInput{
		FilePath: config.FilePath,
		Pattern:  config.Pattern,
	})
}
