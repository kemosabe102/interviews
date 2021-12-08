package languageHeaderParser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Examples for testing
// parse_accept_language(
// "en-US, fr-CA, fr-FR",  # the client's Accept-Language header, a string
// ["fr-FR", "en-US"]      # the server's supported languages, a set of strings
// )
// returns: ["en-US", "fr-FR"]
//
// parse_accept_language("fr-CA, fr-FR", ["en-US", "fr-FR"])
// returns: ["fr-FR"]
//
// parse_accept_language("en-US", ["en-US", "fr-CA"])
// returns: ["en-US"]

func TestParseAcceptLanguage_MultipleAcceptHeaderLanguages(t *testing.T) {
	// arrange
	acceptedLanguageHeaders := "fr-CA, fr-FR"
	serverLanguagesSupported := []string{"en-US", "fr-FR"}
	expectedReturnValue := []string{"fr-FR"}

	// act
	actualLanguagesAcceptedFromServer := parseAcceptLanguage(acceptedLanguageHeaders, serverLanguagesSupported)

	// assert
	assert.Equalf(t, expectedReturnValue, actualLanguagesAcceptedFromServer, "server returned %v but expected %v", actualLanguagesAcceptedFromServer, expectedReturnValue)
}

func TestParseAcceptLanguage_NoAcceptHeaderLanguagePassed(t *testing.T) {
	// arrange
	acceptedLanguageHeaders := ""
	serverLanguagesSupported := []string{"en-US", "fr-FR"}
	expectedReturnValue := []string{""}

	// act
	actualLanguagesAcceptedFromServer := parseAcceptLanguage(acceptedLanguageHeaders, serverLanguagesSupported)

	// assert
	assert.Equalf(t, expectedReturnValue, actualLanguagesAcceptedFromServer, "server returned %v but expected %v", actualLanguagesAcceptedFromServer, expectedReturnValue)
}

