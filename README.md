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
	// Validate a phone number
	phoneNum := "+3 (8096) 46 200 00"
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
}
```

This example demonstrates how to use the Go Is package to validate a phone number, email address, and geographic coordinates. You can import the github.com/goloop/is package and call the corresponding validation functions for the specific data you want to validate. The functions will return a boolean value indicating whether the data is valid or not.
