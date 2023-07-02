package is

import (
	"fmt"
	"regexp"

	"github.com/goloop/g"
)

var (
	// The base64StrictModeRegex is a regex pattern used to validate
	// standard Base64 encoded strings.
	// It matches strings that:
	//  - consist of sequences of 4 valid Base64 characters (A-Za-z0-9+/),
	//    repeated 0 or more times.
	//  - optionally, end with a sequence of 2 valid Base64 characters
	//    followed by '==', or 3 valid Base64 characters followed by '=',
	//    or 4 valid Base64 characters.
	base64StrictModeRegex = regexp.MustCompile(
		`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|` +
			`[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`,
	)

	// The base64SafeModeRegex is similar to base64Regex, but for URL-safe
	// Base64 encoding. The '+' and '/' characters are replaced with '-'
	// and '_' respectively. And can be missing paddings (== or =).
	base64SafeModeRegex = regexp.MustCompile(
		`^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}|` +
			`[A-Za-z0-9-_]{3}|[A-Za-z0-9-_]{4})$`,
	)
)

// Base64 validates whether a given string is encoded in Base64.
// By default, the function checks for strict compliance with
// the base64 format. The second argument to the function can be
// as false to check that the base64 URL is safe.
//
// Example usage:
//
//	// Strict mode.
//	is.Base64("SGVsbG8sIHdvcmxkIQ==")       // Returns: true
//	is.Base64("SGVsbG8sIHdvcmxkIQ==", true) // Returns: true
//	is.Base64("SGVsbG8sIHdvcmxkIQ")         // Returns: false
//	is.Base64("SGVsbG8sIHdvcmxkIQ", true)   // Returns: false
//
//	// Safe mode.
//	is.Base64("SGVsbG8sIHdvcmxkIQ", false)   // Returns: true
//	is.Base64("SGVsbG8sIHdvcmxkIQ==", false) // Returns: true
func Base64(v string, strict ...bool) bool {
	// Default mode is a strict.
	if l := len(strict); l == 0 || l != 0 && g.All(strict...) {
		return base64StrictModeRegex.MatchString(v)
	}

	fmt.Println("Safe")
	return base64SafeModeRegex.MatchString(v)
}
