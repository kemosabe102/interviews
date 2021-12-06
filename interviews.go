package interviews

import "strings"

// Examples for testing
// parse_accept_language(
// "en-US, fr-CA, fr-FR",  # the client's Accept-Language header, a string
// ["fr-FR", "en-US"]      # the server's supported languages, a set of strings
// )
// returns: ["en-US", "fr-FR"]

func parseAcceptLanguage(languageHeader string, serverSupportedLanguages []string) []string {
	// TODO: if only zero or one accepted languages are in header, handle case
	if len(languageHeader) == 0 {
		return nil
	}

	supportedLanguageTagsToReturn := make([]string, 0, len(serverSupportedLanguages))

	serverLanguageMap := arrayToMap(serverSupportedLanguages)
	languageHeaderArray := stringToArray(languageHeader, ", ")

	for _, lh := range languageHeaderArray {
		if _, ok := serverLanguageMap[lh]; ok {
			supportedLanguageTagsToReturn = append(supportedLanguageTagsToReturn, lh)
		}
	}

	return supportedLanguageTagsToReturn
}

func stringToArray(str string, delimiter string) []string {
	return strings.Split(str, delimiter)
}

func arrayToMap(strArray []string) map[string]int  {
	// TODO: if only zero or one array values, handle case
	mappedArray := make(map[string]int)
	for i, s := range strArray {
		mappedArray[s] = i
	}
	return mappedArray
}