package strings_filter

import "net/url"

func ValidateHttpsURL(urlString string) bool {
	// Parse the URL
	urlParsed, err := url.ParseRequestURI(urlString)
	if err != nil {
		return false
	}
	// Check that it's an HTTPS URL
	if urlParsed.Scheme != "https" {
		return false
	}
	return true
}
