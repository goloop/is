package is

import "regexp"

// The nicknameRegex is any word character (a-z, A-Z, 0-9, _)
// or an underscore, between 1 and 15 times.
var nicknameRegex = regexp.MustCompile(`^[\w_]{1,15}$`)

// Nickname checks if the provided string is a valid nickname. It allows
// alphanumeric characters and underscores, and the length must be within
// the range of 1 to 15 characters. This is useful when validating user
// input for account creation or profile updates.
//
// Example usage:
//
// nickname := "User123"
//
//	if is.Nickname(nickname) {
//	    fmt.Println("Nickname is valid!")
//	} else {
//	    fmt.Println("Invalid nickname!")
//	}
func Nickname(nickname string) bool {
	return nicknameRegex.MatchString(nickname)
}
