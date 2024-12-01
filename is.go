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

// The init initializes this package.
func init() {
	// Initialize the letterToNumberCache map.
	// See IBAN validation for more details.
	for _, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letterToNumberCache[letter] = int(letter-'A') + 10
	}
}
