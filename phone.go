package is

import (
	"regexp"
	"strconv"

	"github.com/goloop/g/v2"
)

var (
	// The e164Regex is regular expression for E.164.
	e164Regex = regexp.MustCompile(`^\+\d{1,15}$`)

	// The phoneRegex matches a '+' followed by 1 to 15 digits, the maximum
	// length of an international number under E.164, applied after the
	// grouping separators have been removed.
	phoneRegex = regexp.MustCompile(`^\+\d{1,15}$`)
)

// isASCIIDigits reports whether s is non-empty and consists solely of
// ASCII digits '0'..'9'. Unlike strconv.Atoi it rejects a leading '+'/'-'
// sign and underscores, which is exactly what identifier formats (IMSI,
// IMEI) require.
func isASCIIDigits(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// IMSI checks if the given value is a valid International Mobile Subscriber
// Identity (IMSI). The IMSI is a unique identifier associated with a mobile
// network subscriber. It consists of three parts: MCC (Mobile Country Code),
// MNC (Mobile Network Code), and MSIN (Mobile Subscriber Identification
// Number).
//
// This function validates the IMSI structurally: it must be exactly 15
// ASCII digits (0-9). The MCC, MNC, and MSIN segments are positional, but
// only their digit-ness is checked here, not their assignment to a real
// operator. A sign, space, or any non-digit makes the value invalid.
//
// Example usage:
//
//	is.IMSI("310150123456789")  // Returns: true
//	is.IMSI("460001234567890")  // Returns: true
//	is.IMSI("1234567890123456") // Returns: false, length exceeds 15 digits
//	is.IMSI("310150abc123456")  // Returns: false, invalid characters in MSIN
//	is.IMSI("+10150123456789")  // Returns: false, sign is not a digit
func IMSI(imsi string) bool {
	// An IMSI is exactly 15 ASCII digits — no sign, no separators.
	// A single pass rejects everything else and allocates nothing.
	return len(imsi) == 15 && isASCIIDigits(imsi)
}

// IMEI validates whether a given string is a valid International Mobile
// Equipment Identity (IMEI). It uses the Luhn algorithm for verification.
//
// The value cannot contain characters other than numbers.
//
// The function takes as input a string 'v' representing a potential IMEI
// number, which could include spaces, dots, hyphens, or newline characters
// as separators. These are removed from the input string to obtain the raw
// number.
//
// The Luhn algorithm is then applied to this raw number to check its validity.
// This involves iterating over each digit in the number. If the digit is
// at an even-indexed position in the string (where the first position is
// considered index 1), its value is added directly to a running sum.
// If the digit is at an odd-indexed position, it is doubled.
// If the result of this doubling is greater than 9, 9 is subtracted from it.
// This resultant value is then added to the sum.
//
// After all digits have been processed, the function checks if the total
// sum is a multiple of 10. If it is, the function returns true, indicating
// that the input string is a valid IMEI.
// Otherwise, it returns false.
//
// Example usage:
//
//	is.IMEI("522593572995861") // Returns: true
//	is.IMEI("532593572995861") // Returns: false
//
// Note: An IMEI number is a unique identifier assigned to each mobile device.
// It is a 15-digit number used for tracking and identifying the device.
// The last digit is a check digit, computed according to the Luhn algorithm.
func IMEI[T string | int64](imei T) bool {
	var v string
	switch x := any(imei).(type) {
	case string:
		v = x
	case int64:
		v = strconv.FormatInt(x, 10)
	}

	// An IMEI is exactly 15 ASCII digits. The explicit digit check also
	// rejects a leading '-' from a negative int64 (FormatInt would keep it).
	if len(v) != 15 || !isASCIIDigits(v) {
		return false
	}

	sum := 0
	for i := 0; i < len(v); i++ {
		digit := int(v[i] - '0')

		// If the digit is in an even-indexed position (where the
		// first position is 1), double its value.
		if i%2 == 0 {
			sum += digit
		} else {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
			sum += digit
		}
	}

	// Check if the total sum is a multiple of 10.
	// If it is, then the IMEI is valid.
	return sum%10 == 0
}

// E164 checks if the string is a valid representation of a phone number
// according to the E.164 standard.
//
// The E.164 standard defines a numbering plan for the international
// public telecommunication numbering system. It specifies the format
// for international phone numbers and assigns a unique number to each
// country or region.
//
// This function uses a regular expression to match the input string against
// the E.164 phone number pattern. The pattern requires the string to start
// with a plus sign (+) followed by one or more digits. It allows a maximum
// length of 15 digits (excluding the plus sign).
//
// Example usage:
//
//	is.E164("+123456789")   // Returns: true
//	is.E164("+0123456789")  // Returns: true
//	is.E164("+")            // Returns: false, no digits after plus sign
//	is.E164("+1234567890a") // Returns: false, non-digit character
//	is.E164("1234567890")   // Returns: false, no plus sign
//	is.E164("")             // Returns: false, empty string
//
// This function can be used to validate user input or data to ensure
// it follows the E.164 standard for phone numbers.
func E164(v string) bool {
	return e164Regex.MatchString(v)
}

// Phone checks if the given string is a valid phone number.
// The phone number can have the following format:
// - It starts with a plus sign (+) followed by the country code.
// - The country code can be enclosed in parentheses.
// - Digits may be separated by spaces, hyphens, or dots.
//
// These common grouping separators are ignored before validation; once they
// are removed, the remaining value must be a '+' followed by 1 to 15 digits
// (the E.164 maximum).
//
// Example usage:
//
//	is.Phone("+380 (96) 00 00 000") // Returns: true
//	is.Phone("+1-234-567-8900")     // Returns: true
//	is.Phone("+380961234567")       // Returns: true
//	is.Phone("123456789")           // Returns: false, no plus sign
//	is.Phone("+1234567890123456")   // Returns: false, more than 15 digits
//	is.Phone("")                    // Returns: false, empty string
//
// This function can be used to validate user input or data to ensure
// it follows the specified format for phone numbers.
func Phone(phone string) bool {
	return phoneRegex.MatchString(g.Weed(phone, " ()-."))
}
