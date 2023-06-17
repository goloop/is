package is

import "unicode"

// Digit checks whether a string consists only of numbers.
//
// This method returns True if all characters in the string are numbers and
// the string is not empty. This includes digits (0-9), numeric characters
// that have a specific meaning in non-positional number systems (such as
// base 2, 8, or 16 number systems), and Unicode digit characters.
//
// Example usage:
//
//	is.Digit("1234")     // Output: true
//	is.Digit("Ⅳ")       // Output: true
//	is.Digit("1234abc")  // Output: false
func Digit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Numeric is similar to Digit, but also recognizes numeric characters
// used in other number systems.
//
// Returns true if all characters in the string are any numeric and
// false otherwise. It recognizes numbers (including numbers with a
// decimal point), Roman numerals, Chinese numerals, and more.
//
// Example usage:
//
//	is.Numeric("1234")     // Output: true
//	is.Numeric("Ⅳ")        // Output: true
//	is.Numeric("1234abc")  // Output: false
func Numeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

// Decimal returns true if all characters in the string are decimal digits
// and the string is not empty. Decimal digits are only the characters with
// the numbers 0 through 9. It does not recognize any other numeric characters
// such as Roman numerals or digits from non-positional number systems.
//
// Example usage:
//
//	is.Decimal("1234")     // Output: true
//	is.Decimal("Ⅳ")       // Output: false
//	is.Decimal("1234abc")  // Output: false
//	is.Decimal("1234.56")  // Output: false
func Decimal(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

// Alpha checks whether a string consists only of alphabetic
// characters (letters).
//
// It includes both lowercase and uppercase English letters (from 'a' to
// 'z' and from 'A' to 'Z'). It does not recognize digits, special characters,
// or letters from non-Latin alphabets.
//
// Example usage:
//
//	is.Alpha("abc")   // Output: true
//	is.Alpha("abc1")  // Output: false, contains a digit
//	is.Alpha("abc!")  // Output: false, contains a special character
//	is.Alpha("abcΔ")  // Output: false, contains a non-Latin letter
func Alpha(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Alnum checks whether a string consists only of alphabetic
// characters (letters) and numbers.
//
// It includes both lowercase and uppercase English letters (from 'a' to
// 'z' and from 'A' to 'Z') and digits (from '0' to '9'). It doesn't recognize
// special characters or letters/numbers from non-Latin alpha./number systems.
//
// Example usage:
//
//	is.Alnum("abc123")      // Output: true
//	is.Alnum("abc")         // Output: true
//	is.Alnum("123")         // Output: true
//	is.Alnum("abc!")        // Output: false
//	is.Alnum("abcΔ")        // Output: false
func Alnum(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

// Lower checks whether a string consists only of lowercase
// alphabetic characters.
//
// Example usage:
//
//	is.Lower("abc")       // Output: true
//	is.Lower("Abc")       // Output: false, contains an uppercase letter
//	is.Lower("abc123")    // Output: false, contains a number
//	is.Lower("abc!")      // Output: false, contains a special character
func Lower(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

// Upper checks whether a string consists only of uppercase
// alphabetic characters.
//
// Example usage:
//
//	is.Upper("ABC")      // Output: true
//	is.Upper("AbC")      // Output: false, contains a lowercase letter
//	is.Upper("ABC123")   // Output: false, contains a number
//	is.Upper("ABC!")     // Output: false, contains a special character
func Upper(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// Title checks whether a string is a titlecased string.
//
// In a titlecased string, upper- and title-case characters may only
// follow uncased characters and lowercase characters only cased ones.
//
// Example usage:
//
//	is.Title("Hello World")   // Output: true
//	is.Title("Hello world")   // Output: false, 'world' starts with a lowercase
//	is.Title("HELLO WORLD")   // Output: false, all letters are uppercase
//	is.Title("hELLO wORLD")   // Output: false, words start with a lowercase
func Title(s string) bool {
	if len(s) == 0 {
		return false
	}

	prevCased := true

	for _, r := range s {
		if unicode.IsUpper(r) || unicode.IsTitle(r) {
			if prevCased {
				return false
			}
			prevCased = true
		} else if unicode.IsLower(r) {
			if !prevCased {
				return false
			}
		} else {
			prevCased = false
		}
	}

	return true
}

// Space checks whether a string consists only of whitespace characters.
// Whitespace characters includes spaces, tabs, newlines, and other
// Unicode whitespace characters.
//
// This function is analogous to Python's str.isspace() method.
//
// Example usage:
//
//	is.Space(" \t\n")    // Output: true
//	is.Space("Hello")    // Output: false
//	is.Space(" ")        // Output: true
//	is.Space("\n\t ")    // Output: true
//	is.Space("")         // Output: false, an empty string has no characters
func Space(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}
