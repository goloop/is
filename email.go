package is

import (
	"regexp"
	"strings"
)

var (
	// Regular expression pattern for a valid email slug (local/domain part).
	slugPattern = `[\p{L}\p{N}]{1,}([\p{L}\p{N}.-]{1,}[\p{L}\p{N}]{1,})?`

	// Compiled regex for the local part of the email.
	localRegex = regexp.MustCompile(`^` + slugPattern + `$`)

	// Compiled regex for the domain part of the email.
	domainRegex = regexp.MustCompile(`^` + slugPattern + `\.[\p{L}\p{N}]{2,}$`)
)

// Email validates an email address and returns a boolean result.
// The function checks for email length, splits the email into local and
// domain parts, validates the length of the local part, and finally uses
// regular expressions to check the format of each part.
//
// Example usage:
//
//	is.Email("test@example.com")  // Output: true
//	is.Email("TEST@EXAMPLE.COM")  // Output: true
//	is.Email(" test@example.com") // Output: false
//	is.Email("test@example.c")    // Output: false
//
// The function performs only a format check, so it does not clean spaces
// at the beginning and end of a line, does not remove tab characters and
// carriage returns to a new line. You can use the g.Wedd and g.Trim
// functions for it:
//
//	is.Email(g.Trim(" test@example.com\n"))   // Output: true
//	is.Email(g.Weed("test\t@example.com\n"))  // Output: true
func Email(email string) bool {
	// Convert to lowercase to handle case sensitivity.
	email = strings.ToLower(email)

	// Check the overall length.
	if len(email) > 254 || len(email) < 6 {
		return false
	}

	// Split the email into local and domain parts.
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	local, domain := parts[0], parts[1]

	// Check the length of the local part.
	if len(local) > 64 || len(local) < 1 {
		return false
	}

	// Check if both parts are valid.
	if !localRegex.MatchString(local) || !domainRegex.MatchString(domain) {
		return false
	}

	return true
}
