package strings_filter

import "net/url"

func AddHexPrefix(hex string) string {
	if len(hex) >= 2 && hex[0:2] == "0x" {
		return hex
	}
	return "0x" + hex
}

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
