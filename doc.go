// Package is provides a comprehensive set of validation functions for data
// formats commonly used in web applications, financial systems, and general
// software development.
//
// # Overview
//
// The package offers validation functions for various data types and formats:
//
//   - Identity and Access: email addresses, nicknames, variable names;
//   - Financial: bank card numbers, IBAN accounts;
//   - Geographic: latitude/longitude coordinates;
//   - Network: IPv4/IPv6 addresses;
//   - Telecommunications: phone numbers in E.164 format;
//   - String Types: alphanumeric, numeric, hexadecimal, binary;
//   - Numbers: even/odd, positive/negative, natural numbers;
//   - Encodings: Base64, Base64URL, JWT tokens;
//   - Data formats: MD5 hashes.
//
// # Design Philosophy
//
// The package follows several key principles:
//
//   - No data cleaning - functions validate data as-is without modification;
//   - Zero dependencies - minimal external dependencies;
//   - Type safety - extensive use of Go generics where appropriate;
//   - Performance - optimized for speed and minimal memory allocation.
//
// # Usage
//
// Most functions follow a simple pattern returning a boolean
// indicating validity:
//
//	if is.Email("user@example.com") {
//	    // Valid email address
//	}
//
//	if is.BankCard("4111111111111111") {
//	    // Valid bank card number
//	}
//
// Some functions support additional validation modes:
//
//	// Strict mode validation
//	if is.Nickname("user123", true) {
//	    // Valid nickname in strict mode
//	}
//
// # Data Cleaning
//
// This package does not perform data cleaning. If input needs cleaning,
// use appropriate tools beforehand. For example, with the companion 'g' package:
//
//	// Clean IBAN before validation
//	raw := "GB82 WEST 1234 5698 7654 32"    // contains spaces
//	iban := g.Weed(raw, g.Whitespaces)      // remove spaces
//	if is.IBAN(iban) {
//	    // Valid IBAN
//	}
//
// # Common Validation Patterns
//
// Account Validation:
//
//	is.Email("user@example.com")      // Email address
//	is.Nickname("user123")            // Username
//	is.VariableName("myVar")          // Programming variable name
//
// Financial Validation:
//
//	is.BankCard("4111111111111111")   // Bank card number
//	is.IBAN("DE89370400440532013000") // IBAN account number
//
// Geographic Validation:
//
//	is.Latitude(45.423)               // Latitude coordinate
//	is.Longitude(-75.935)             // Longitude coordinate
//
// Network Address Validation:
//
//	is.IPv4("192.168.0.1")           // IPv4 address
//	is.IPv6("2001:db8::1")           // IPv6 address
//
// Phone Number Validation:
//
//	is.E164("+1234567890")           // E.164 format
//	is.Phone("+1 (234) 567-890")     // General phone format
//
// String Type Validation:
//
//	is.Alpha("ABC")                   // Alphabetic characters
//	is.Digit("123")                   // Numeric digits
//	is.Alnum("ABC123")                // Alphanumeric characters
//	is.HexColor("#FF5733")            // Hexadecimal color
//
// Number Validation:
//
//	is.Even(4)                        // Even number
//	is.Odd(3)                         // Odd number
//	is.Natural(5)                     // Natural number
//	is.Positive(1.5)                  // Positive number
//
// Format Validation:
//
//	is.Base64("SGVsbG8=")            // Base64 encoding
//	is.MD5("d41d8cd98f00b204...")    // MD5 hash
//	is.JWT("eyJhbGciOiJIUzI...")     // JWT token
//
// For more detailed information about specific validation functions,
// see their respective documentation.
package is
