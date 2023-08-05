[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/is)](https://goreportcard.com/report/github.com/goloop/is) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/is/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/is) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# is

Is package provides various validation functions to check the correctness of different types of data. It offers a collection of functions to validate data such as numbers, strings, geographic coordinates, email addresses, e-mail, bank card number, bank account IBAN, phone and more.

## Note

The package provides validation functions that do not pre-clean the data. If the validation data needs to be cleaned, it must be cleaned beforehand, for example, using the [g package](https://github.com/goloop/g), as `g.Weed`, `g.Trim` etc.


## Installation

To use the Go Is package in your project, you need to install it using the go get command:

```shell
go get github.com/goloop/is
```

## Usage

Here's an example of how to use the Go Is package to validate different types of data:

```go
package main

import (
	"fmt"
	"github.com/goloop/is"
)

func main() {
	// Validate a phone number.
	phoneNum := "+380 (96) 46 200 00"
	if is.Phone(phoneNum) {
		fmt.Println("Valid phone number")
	} else {
		fmt.Println("Invalid phone number")
	}

	// Validate an email address
	email := "test@example.com"
	if is.Email(email) {
		fmt.Println("Valid email address")
	} else {
		fmt.Println("Invalid email address")
	}

	// Validate geographic coordinates
	latitude := 37.7749
	longitude := -122.4194
	if is.Coordinates(latitude, longitude) {
		fmt.Println("Valid coordinates")
	} else {
		fmt.Println("Invalid coordinates")
	}

	// Output:
	// Valid phone number
	// Valid email address
	// Valid coordinates
}
```

This example demonstrates how to use the Go Is package to validate a phone number, email address, and geographic coordinates. You can import the github.com/goloop/is package and call the corresponding validation functions for the specific data you want to validate. The functions will return a boolean value indicating whether the data is valid or not.


## Type definitions

Use the is.Instance function to check for type matching, like `is.Instance[type](object)`. The function returns true if the `object` type is similar to `type`.

The function can take an optional second parameter as a bool value. If it is true, it activates strict mode, where the is.Instance function performs exact type matching.

```go
package main

import (
	"fmt"

	"github.com/goloop/is"
)

// One is the helper struct.
type One struct {
	N string
}

// Two is helper struct too with the same field name as One.
type Two struct {
	N string
}

func main() {
	{
		// Sequences.
		// Sequences are arrays, slices and maps.
		var (
			a [3]int
			s []int
			m map[int]int
		)

		fmt.Println("\nSequences:")
		// In flexible mode, we simply determine whether an object is
		// of a certain shape (slice, array, map), ignoring the basic
		// type and dimensions.
		fmt.Printf("%-28s%v\n", "a is any array (flex):",
			is.Instance[[0]uint](a))
		fmt.Printf("%-28s%v\n", "s is any slice (flex):",
			is.Instance[[]uint](s))
		fmt.Printf("%-28s%v\n", "m is any map (flex):",
			is.Instance[map[uint]uint](m))

		fmt.Printf("\n%-28s%v\n", "a is array (flex):",
			is.Instance[[3]int](a))
		fmt.Printf("%-28s%v\n", "s is slice (flex):",
			is.Instance[[]int](s))
		fmt.Printf("%-28s%v\n", "m is map (flex):",
			is.Instance[map[int]int](m))

		// In strict mode, we determine whether an object is of a certain
		// shape (slice, array, map) and has the same basic type
		// and dimensions.
		fmt.Printf("\n%-28s%v\n", "a is any array (strict):",
			is.Instance[[0]uint](a, true))
		fmt.Printf("%-28s%v\n", "s is any slice (strict):",
			is.Instance[[]uint](s, true))
		fmt.Printf("%-28s%v\n", "m is any map (strict):",
			is.Instance[map[uint]uint](m, true))

		fmt.Printf("\n%-28s%v\n", "a is array (strict):",
			is.Instance[[3]int](a, true))
		fmt.Printf("%-28s%v\n", "s is slice (strict):",
			is.Instance[[]int](s, true))
		fmt.Printf("%-28s%v\n", "m is map (strict):",
			is.Instance[map[int]int](m, true))

		// Sequences:
		// a is any array (flex):   true
		// s is any slice (flex):   true
		// m is any map (flex):     true

		// a is array (flex):       true
		// s is slice (flex):       true
		// m is map (flex):         true

		// a is any array (strict): false
		// s is any slice (strict): false
		// m is any map (strict):   false

		// a is array (strict):     true
		// s is slice (strict):     true
		// m is map (strict):       true
	}

	{
		// Simple types and pointers for them.
		var (
			intVal     int
			float63Val float64
			boolVal    bool
			stringVal  string

			intPtr     *int
			float64Ptr *float64
			boolPtr    *bool
			stringPtr  *string
		)

		fmt.Println("\nSimple types:")

		// For simple types, there is no difference in strict or flexible mode.
		fmt.Printf("%-28s%v\n", "intVal is int:",
			is.Instance[int](intVal))
		fmt.Printf("%-28s%v\n", "float63Val is float64:",
			is.Instance[float64](float63Val))
		fmt.Printf("%-28s%v\n", "boolVal is bool:",
			is.Instance[bool](boolVal))
		fmt.Printf("%-28s%v\n", "stringVal is string:",
			is.Instance[string](stringVal))

		fmt.Printf("\n%-28s%v\n", "intVal is uint:",
			is.Instance[uint](intVal))
		fmt.Printf("%-28s%v\n", "float63Val is float32:",
			is.Instance[float32](float63Val))
		fmt.Printf("%-28s%v\n", "boolVal is int:",
			is.Instance[int](boolVal))
		fmt.Printf("%-28s%v\n", "stringVal is rune:",
			is.Instance[rune](stringVal))

		fmt.Printf("\n%-28s%v\n", "intPtr is *int:",
			is.Instance[*int](intPtr))
		fmt.Printf("%-28s%v\n", "float64Ptr is *float64:",
			is.Instance[*float64](float64Ptr))
		fmt.Printf("%-28s%v\n", "boolPtr is *bool:",
			is.Instance[*bool](boolPtr))
		fmt.Printf("%-28s%v\n", "stringPtr is *string:",
			is.Instance[*string](stringPtr))

		fmt.Printf("\n%-28s%v\n", "intPtr is *uint:",
			is.Instance[*uint](intPtr))
		fmt.Printf("%-28s%v\n", "float64Ptr is *float32:",
			is.Instance[*float32](float64Ptr))
		fmt.Printf("%-28s%v\n", "boolPtr is *int:",
			is.Instance[*int](boolPtr))
		fmt.Printf("%-28s%v\n", "stringPtr is *rune:",
			is.Instance[*rune](stringPtr))

		// Simple types:
		// intVal is int:           true
		// float63Val is float64:   true
		// boolVal is bool:         true
		// stringVal is string:     true

		// intVal is uint:          false
		// float63Val is float32:   false
		// boolVal is int:          false
		// stringVal is rune:       false

		// intPtr is *int:          true
		// float64Ptr is *float64:  true
		// boolPtr is *bool:        true
		// stringPtr is *string:    true

		// intPtr is *uint:         false
		// float64Ptr is *float32:  false
		// boolPtr is *int:         false
		// stringPtr is *rune:      false
	}

	{
		// Functions.
		sum := func(a, b int) int { return a + b }

		fmt.Println("\nFunc types:")

		// Flexibel mode determines whether the object is any function.
		fmt.Printf("%-28s%v\n", "sum is any func (flex):",
			is.Instance[func()](sum))

		// Strict mode determines whether the object is a function with
		// the same signature.
		fmt.Printf("%-28s%v\n", "sum is any func (strict):",
			is.Instance[func()](sum, true))
		fmt.Printf("%-28s%v\n", "sum is func(int, int) int:",
			is.Instance[func(int, int) int](sum, true))

		// Func types:
		// sum is any func (flex):     true
		// sum is any func (strict):   false
		// sum is func(int, int) int:  true
	}

	{
		// Structs.
		var one One
		fmt.Println("\nStruct types:")

		// Flexibel mode determines whether the object is any struct.
		fmt.Printf("%-28s%v\n", "one is One struct (flex):",
			is.Instance[One](one))
		fmt.Printf("%-28s%v\n", "one is any struct (flex):",
			is.Instance[struct{}](one))

		// Any struct is struct{} signature but not a different struct,
		// like Two or Three:
		fmt.Printf("%-28s%v\n", "one is Two struct (flex):",
			is.Instance[Two](one))

		// Strict mode determines whether the object is a struct with
		// the same signature.
		fmt.Printf("\n%-28s%v\n", "one is One struct (strict):",
			is.Instance[One](one, true))
		fmt.Printf("%-28s%v\n", "one is any struct (strict):",
			is.Instance[struct{}](one, true))

		// Struct types:
		// one is One struct (flex):   true
		// one is any struct (flex):   true
		// one is Two struct (flex):   false

		// one is One struct (strict): true
		// one is any struct (strict): false

	}
}
```