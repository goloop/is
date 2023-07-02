package is

import (
	"regexp"
)

var (
	// The base64Regex is a regex pattern used to validate
	// standard Base64 encoded strings.
	// It matches strings that:
	//  - consist of sequences of 4 valid Base64 characters (A-Za-z0-9+/),
	//    repeated 0 or more times.
	//  - optionally, end with a sequence of 2 valid Base64 characters
	//    followed by '==', or 3 valid Base64 characters followed by '=',
	//    or 4 valid Base64 characters.
	base64Regex = regexp.MustCompile(
		`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|` +
			`[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`,
	)

	// The hexRegex is a regex pattern used to validate hexadecimal strings.
	hexRegex = regexp.MustCompile(`^(#|0x)?[0-9a-fA-F]+$`)

	// The binRegex is a regex pattern used to validate binary strings.
	binRegex = regexp.MustCompile(`^(0b)?[01]+$`)

	// The hexColorRegex is a regex pattern used to validate
	// hexadecimal color strings.
	hexColorRegex = regexp.MustCompile(`^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)

	// The rgbRegex is a regex pattern used to validate RGB color strings.
	rgbColorRegex = regexp.MustCompile(
		`^(rgb|RGB)\(\s*([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\s*,` +
			`\s*([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\s*,` +
			`\s*([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\s*\)$`,
	)
)

// Base64 validates whether a given string 'v' is a valid Base64
// encoded string.
//
// Base64 is a binary-to-text encoding scheme that is commonly used
// to encode binary data, notably when that data needs to be stored and
// transferred over media designed to handle text. This encoding helps to
// ensure that the data remains intact without modification during transport.
//
// This function uses a regular expression to verify that the input string
// conforms to the format of a Base64 encoded string. It checks if the string
// can be evenly divided by 4, and only contains valid Base64 characters
// (A-Z, a-z, 0-9, +, /, and = for padding). The padding at the end of
// Base64 string, which is one or two '=' characters, is also checked for.
//
// If the input string matches this format, the function returns true,
// indicating that the string is a valid Base64 encoded string.
// Otherwise, it returns false.
//
// Example usage:
//
//	is.Base64("SGVsbG8sIHdvcmxkIQ==") // Returns: true
//	is.Base64("SGVsbG8sIHdvcmxkIQ")   // Returns: false
//
// Note: This function does not validate the content of the encoded data,
// just the format of Base64 strings.
func Base64(v string) bool {
	return base64Regex.MatchString(v)
}

// Hex validates whether a given string 'v' is a valid hexadecimal value.
//
// A hexadecimal value is a number that includes digits from 0 to 9 and
// letters from A to F (either lower case or upper case). Hexadecimal values
// are often used in computing to represent numbers in a human-readable format.
//
// This function does not require '0x' or '#' prefix to be present in the
// input string, although it will accept strings with these prefixes.
//
// The function uses a regular expression to verify that the input string
// conforms to the format of a hexadecimal number.
//
// If the input string matches this format, the function returns true,
// indicating that the string is a valid hexadecimal value.
// Otherwise, it returns false.
//
// Example usage:
//
//	is.Hex("deadBEEF") // Returns: true
//	is.Hex("#c0ffee")  // Returns: true
//	is.Hex("nothex")   // Returns: false
//
// Note: This function does not validate semantic correctness,
// just the format of hexadecimal values.
func Hex(v string) bool {
	return hexRegex.MatchString(v)
}

// Bin validates whether a given string 'v' is a valid binary value.
//
// A binary value is a number that includes only the digits 0 and 1.
// Binary values are the most basic unit of data in computing and digital
// communications.
//
// This function does not require '0b' prefix to be present in the
// input string, although it will accept strings with this prefix.
//
// The function uses a regular expression to verify that the input string
// conforms to the format of a binary number.
//
// If the input string matches this format, the function returns true,
// indicating that the string is a valid binary value.
// Otherwise, it returns false.
//
// Example usage:
//
//	is.Bin("101010") // Returns: true
//	is.Bin("0b1101")  // Returns: true
//	is.Bin("10201")   // Returns: false
//
// Note: This function does not validate semantic correctness,
// just the format of binary values.
func Bin(v string) bool {
	return binRegex.MatchString(v)
}

// HexColor validates whether a given string 'v' is a valid hexadecimal
// RGB color value. Hexadecimal RGB color values are defined as '#<color>',
// where <color> is either a  3-digit or a 6-digit hexadecimal number.
// Each digit is in the range 0-9 and A-F (either lower case or upper case).
//
// The function uses a regular expression to verify that the input string
// conforms to this format.
//
// In the case of a 3-digit color, each digit represents a color value for red,
// green, and blue respectively. Each digit is equivalent to repeating it twice
// in a 6-digit color. For example, "#123" in 3-digit color is equivalent to
// "#112233" in 6-digit color.
//
// In the case of a 6-digit color, the first two digits represent red, the next
// two represent green, and the last two represent blue.
//
// If the input string matches this format, the function returns true,
// indicating that the string is a valid hexadecimal RGB color value.
// Otherwise, it returns false.
//
// Example usage:
//
//	is.HexColor("#fff")    // Returns: true
//	is.HexColor("#efef01") // Returns: true
//	is.HexColor("not")     // Returns: false
//
// Note: The function does not validate color semantic correctness,
// just the format of hexadecimal RGB colors.
func HexColor(v string) bool {
	return hexColorRegex.MatchString(v)
}

// RGBColor validates whether a given string 'v' is a valid RGB color value.
// RGB color values are defined as 'rgb(<red>, <green>, <blue>)', where each
// of <red>, <green>, and <blue> is an integer in the range 0-255.
// The function uses a regular expression to verify that the input
// string conforms to this format.
//
// The function does not account for leading or trailing spaces in the input
// string. The values for red, green, and blue must be separated by commas.
// These can be followed by spaces, tabs, or newline characters, all of which
// are considered valid. The entire string must be enclosed within 'rgb()'
// with no spaces between 'rgb' and the opening parenthesis.
//
// If the input string matches this format and all color values are in the
// range 0-255, the function returns true, indicating that the string is a
// valid RGB color value. Otherwise, it returns false.
//
// Example usage:
//
//	is.RGBColor("rgb(255, 255, 255)") // Returns: true
//	is.RGBColor("rgb(255, 255, 256)") // Returns: false
//	is.RGBColor("rgb(255, 255)")      // Returns: false
//
// Note: The function doesn't validate color semantic correctness,
// just the syntax and range of values.
func RGBColor(v string) bool {
	return rgbColorRegex.MatchString(v)
}
