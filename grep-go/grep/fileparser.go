package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// SearchFile scans each line of a file and returns the lines that include the searched for pattern
func SearchFile(searchInput *SearchInput) (LinesFound, error) {
	file, err := os.Open(searchInput.FilePath)
	if err != nil {
		log.Fatalf("SearchFile failed with: %v", err)
	}
	defer file.Close()

	var linesFound LinesFound

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchInput.Pattern) {
			linesFound.LinesWithPattern = append(linesFound.LinesWithPattern, scanner.Text())
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return linesFound, nil
}
