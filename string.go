package is

import (
	"strings"
	"unicode"
)

// decimalSeparators holds the runes accepted by Numeric as the single
// decimal/grouping mark between digits. Kept at package level so it is
// allocated once, not on every Numeric call.
var decimalSeparators = map[rune]struct{}{
	'.': {}, // Dot
	',': {}, // Comma
	'·': {}, // Middle dot
	'・': {}, // Japanese separator
	'٫': {}, // Arabic decimal point
	'،': {}, // Arabic comma
	'۔': {}, // Urdu full stop
}

// Digit checks whether a string consists only of numbers.
//
// This method returns true if all characters in the string are numbers and
// the string is not empty. This includes digits (0-9), numeric characters
// that have a specific meaning in non-positional number systems (such as
// base 2, 8, or 16 number systems), and Unicode digit characters.
//
// Example usage:
//
//	is.Digit("1234")     // Output: true
//	is.Digit("Ⅳ")       // Output: false
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

// Numeric checks whether a string represents a number: an optional leading
// '+' or '-' sign, decimal digits, and at most one decimal/grouping
// separator. The digits may come from any writing system whose decimal
// digits are classified as Unicode "Number, decimal digit" (category Nd) —
// e.g. ASCII, Arabic-Indic, Devanagari, Thai — so localized input validates
// out of the box.
//
// What it deliberately rejects:
//   - letter-based numeral systems (Hebrew, Greek, Armenian, Coptic, …),
//     whose glyphs are letters, not digits, and would otherwise let ordinary
//     words pass as "numbers";
//   - ideographic and alphabetic numerals (CJK 一二三, Roman Ⅳ), which are
//     not decimal digits;
//   - a string with no digit at all (a lone separator or sign such as "."
//     or "-").
//
// The function returns true only when every character is valid AND at least
// one decimal digit is present.
//
// Example usage:
//
//	is.Numeric("1234")     // Output: true
//	is.Numeric("3.14")     // Output: true
//	is.Numeric("-456,789") // Output: true (comma as a separator)
//	is.Numeric("٣.١٤")     // Output: true (Arabic-Indic digits)
//	is.Numeric("Ⅳ")       // Output: false (Roman numeral, not a digit)
//	is.Numeric("一二三")    // Output: false (CJK numerals, not digits)
//	is.Numeric("1234abc")  // Output: false
//	is.Numeric("1.2.3")    // Output: false (more than one separator)
//	is.Numeric(".")        // Output: false (no digit)
func Numeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	var hasDecimalSeparator, hasDigit bool
	for i, r := range s {
		// A sign is allowed only as the very first character.
		if i == 0 && (r == '+' || r == '-') {
			continue
		}

		// At most one decimal/grouping separator is allowed.
		if _, isSeparator := decimalSeparators[r]; isSeparator {
			if hasDecimalSeparator {
				return false // more than one decimal separator
			}
			hasDecimalSeparator = true
			continue
		}

		// Only true decimal digits (Unicode category Nd) count as numeric.
		if !unicode.IsDigit(r) {
			return false
		}
		hasDigit = true
	}

	// A number must contain at least one digit: "." or "-" alone is not one.
	return hasDigit
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

// Float returns true if the string represents a floating-point number,
// where the decimal separator is a dot, and all digits are ASCII digits (0-9).
// It supports an optional '+' or '-' sign at the beginning and at most one
// decimal point. It does not support exponential notation or other numeric
// formats.
//
// Example usage:
//
//	is.Float("123.456")     // Output: true
//	is.Float("-0.001")      // Output: true
//	is.Float("3.14")        // Output: true
//	is.Float("123")         // Output: true
//	is.Float("123.")        // Output: true
//	is.Float(".456")        // Output: true
//	is.Float("123.456.789") // Output: false
//	is.Float("abc")         // Output: false
//	is.Float("123a")        // Output: false
//	is.Float("123,456")     // Output: false
func Float(s string) bool {
	if len(s) == 0 {
		return false
	}

	var (
		hasDecimalPoint bool
		hasDigits       bool
		startIndex      int
	)

	// Check for optional sign at the beginning.
	if s[0] == '+' || s[0] == '-' {
		startIndex = 1
	}

	// Edge case: string contains only '+' or '-'.
	if startIndex >= len(s) {
		return false
	}

	for i := startIndex; i < len(s); i++ {
		r := rune(s[i])

		if r == '.' {
			if hasDecimalPoint {
				return false // more than one decimal point
			}
			hasDecimalPoint = true
			continue
		}

		if r >= '0' && r <= '9' {
			hasDigits = true
			continue
		}

		return false
	}

	// At least one digit is required.
	return hasDigits
}

// Alpha checks whether a string consists only of alphabetic
// characters (letters). It does not recognize digits, special
// characters.
//
// Example usage:
//
//	is.Alpha("Київ")  // Output: true
//	is.Alpha("abc")   // Output: true
//	is.Alpha("abc1")  // Output: false, contains a digit
//	is.Alpha("abc!")  // Output: false, contains a special character
//	is.Alpha("abcΔ")  // Output: true, Δ is a Unicode letter
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
// Example usage:
//
//	is.Alnum("abc123")      // Output: true
//	is.Alnum("abc")         // Output: true
//	is.Alnum("123")         // Output: true
//	is.Alnum("abc!")        // Output: false
//	is.Alnum("abcΔ")        // Output: true, Δ is a Unicode letter
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

// Title checks whether a string is a titlecased string: every word starts
// with an upper- (or title-) case letter and continues with lowercase
// letters.
//
// Word boundaries are any non-letter runes. This includes spaces and digits,
// but also punctuation such as the hyphen and the apostrophe, so a name like
// "O'Brien" is treated as the two words "O" and "Brien" (both titlecased,
// hence valid), and "Mc-donald" splits into "Mc" and "donald" (the second is
// lowercase, hence invalid). Pre-normalize the input if you need a different
// word model.
//
// Example usage:
//
//	is.Title("Hello World")   // Output: true
//	is.Title("Hello world")   // Output: false, 'world' starts with a lowercase
//	is.Title("HELLO WORLD")   // Output: false, all letters are uppercase
//	is.Title("hELLO wORLD")   // Output: false, words start with a lowercase
//	is.Title("O'Brien")       // Output: true, "O" and "Brien" are separate
func Title(s string) bool {
	if len(s) == 0 {
		return false
	}

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// A string with no letter word (e.g. "123" or "!!!") is not titlecased.
	if len(words) == 0 {
		return false
	}

	for _, word := range words {
		for index, char := range word {
			if index == 0 {
				// The leading letter may be upper- or titlecase; a titlecase
				// digraph such as 'ǅ' (category Lt) is a valid word start.
				if !unicode.IsUpper(char) && !unicode.IsTitle(char) {
					return false
				}
			} else {
				if !unicode.IsLower(char) {
					return false
				}
			}
		}
	}
	return true
}

// Space checks whether a string consists only of whitespace characters.
// Whitespace characters includes spaces, tabs, newlines, and other
// Unicode whitespace characters.
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
