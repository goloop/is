// Package is includes validation functions of various data such as e-mail,
// bank card number, bank account IBAN, phone number and others.
//
// The package provides validation functions that do not pre-clean the data.
// If the validation data needs to be cleaned, it must be cleaned beforehand,
// for example using the g module.
//
// Example:
//
//	data := "UA 90 305299 2990004149123456789" // contains spaces
//	iban := g.Weed(data, g.Whitespaces)        // remove spaces
//	is.Iban(iban) // Output: true
//	is.Iban(data) // Output: false
package is

import "github.com/goloop/g"

// Numerable is an interface type that is satisfied by all numeric types
// in Go, both integer and floating point. This includes int, int8, int16,
// int32, int64, uint, uint8, uint16, uint32, uint64, float32, and float64.
// It allows functions to operate on any of these types where numerical
// operations such as addition, subtraction, multiplication, and division
// are needed. It enables generic programming techniques for numeric types.
type Numerable interface {
	g.Numerable
}

// Verifiable is an interface type that is satisfied by classical types
// like numeric types and strings in Go.
//
// The purpose of the Verifiable interface is to enable generic programming
// techniques for numeric types and strings. Functions can use this interface
// as a constraint to operate on any of these types where numerical operations
// or string operations are needed.
type Verifiable interface {
	g.Verifiable
}
