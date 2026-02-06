package validator

import (
	"errors"
	"net/url"
	"regexp"
)

// urlRegex is used for basic structural validation of URLs.
// This avoids unnecessary parsing of clearly invalid inputs.
var urlRegex = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)

// ValidateURL validates a raw URL string provided by the user.
// It ensures the URL is well formed, uses a supported scheme,
// and can be safely processed by the application.
func ValidateURL(rawURL string) error {
	// Validation 1: Check the empty input
	if rawURL == "" {
		return errors.New("url cannot be empty!")
	}

	// Validation 2: Perform regex based validation
	if !urlRegex.MatchString(rawURL) {
		return errors.New("url format is invalid!")
	}

	// Validation 3: Parse the URL using Go's standard library
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.New("url parsing failed!")
	}

	// Validation 4: Ensure the http/https is supported
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("only http and https schemes are supported")
	}

	return nil
}
