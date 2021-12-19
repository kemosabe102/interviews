package grep

// SearchInput is the input needed to parse a file
type SearchInput struct {
	Pattern  string
	FilePath string
}

// LinesFound is the response returned with each line where the searched for value exists
type LinesFound struct {
	LinesWithPattern []string
}
