package is

import (
	"regexp"
	"unicode/utf8"
)

// Regular expressions for nickname validation
var (
	// Unicode regex: Allows Unicode letters (\p{L}), numbers (\p{N}),
	// and underscores (_), 1 to 15 characters.
	unicodeNicknameRegex = regexp.MustCompile(`^[\p{L}\p{N}_]{1,15}$`)

	// ASCII regex: Allows ASCII letters (A-Za-z), numbers (0-9),
	// and underscores (_), 1 to 15 characters.
	asciiNicknameRegex = regexp.MustCompile(`^[A-Za-z0-9_]{1,15}$`)
)

// Nickname checks if the provided string is a valid nickname.
//
// In non-strict mode (default), it allows Unicode letters, numbers,
// and underscores, and the length must be within 1 to 15 characters.
// This is useful when validating user input for account creation or
// profile updates in international contexts.
//
// In strict mode, activated by passing `true` as the optional parameter,
// it only allows ASCII letters, numbers, and underscores, within the same
// length constraints.
//
// Consistent with the package's no-cleaning principle, the input is validated
// as-is: surrounding whitespace is NOT trimmed, so " user" is invalid. Trim
// the input beforehand (e.g. with the companion 'g' package) if needed.
//
// Example usage:
//
//	nickname := "User123"
//
//	if is.Nickname(nickname) {
//	    fmt.Println("Nickname is valid (non-strict mode)!")
//	} else {
//	    fmt.Println("Invalid nickname!")
//	}
//
//	if is.Nickname(nickname, true) {
//	    fmt.Println("Nickname is valid (strict mode)!")
//	} else {
//	    fmt.Println("Invalid nickname in strict mode!")
//	}
func Nickname(nickname string, strict ...bool) bool {
	// Check the length in characters (runes). The input is validated as-is,
	// without trimming, per the package's no-cleaning principle.
	length := utf8.RuneCountInString(nickname)
	if length < 1 || length > 15 {
		return false
	}

	// Determine if strict mode is enabled.
	isStrict := len(strict) > 0 && strict[0]

	// Select the appropriate regex based on the mode.
	var regex *regexp.Regexp
	if isStrict {
		regex = asciiNicknameRegex
	} else {
		regex = unicodeNicknameRegex
	}

	return regex.MatchString(nickname)
}
