[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/is)](https://goreportcard.com/report/github.com/goloop/is) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/is/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/is) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# is

The `is` package provides a comprehensive set of validation functions for various data types commonly used in web applications, financial systems, and general software development. It offers a clean, efficient way to validate data formats ranging from basic types to complex financial and network identifiers.

## Features

### Account & Identity Validation
- `Email(string) bool` - Email addresses
- `Nickname(string, ...bool) bool` - Usernames (with optional strict mode)
- `VariableName(string, ...bool) bool` - Programming variable names
- `VariableNameFor(string, string) (bool, error)` - Language-specific variable names

### Financial Validation
- `BankCard(string, ...*regexp.Regexp) bool` - Credit/Debit card numbers with optional card type checking
- `IBAN(string, ...bool) bool` - International Bank Account Numbers
- `Iban(string, ...bool) bool` - Alias for IBAN validation

### Geographic Validation
- `Latitude[T string|float64](T) bool` - Latitude coordinates
- `Longitude[T string|float64](T) bool` - Longitude coordinates
- `Coordinates[T string|float64](lat, lon T) bool` - Coordinate pairs

### Network & Communication
- `IPv4(string) bool` - IPv4 addresses
- `IPv6(string) bool` - IPv6 addresses
- `IP(string) bool` - Any IP address (v4 or v6)
- `Phone(string) bool` - Phone numbers
- `E164(string) bool` - E.164 format phone numbers

### String Type Validation
- `Alpha(string) bool` - Alphabetic characters
- `Alnum(string) bool` - Alphanumeric characters
- `Digit(string) bool` - Numeric digits only
- `Lower(string) bool` - Lowercase letters
- `Upper(string) bool` - Uppercase letters
- `Title(string) bool` - Title case text
- `Space(string) bool` - Whitespace characters
- `Numeric(string) bool` - Numeric strings (including decimals)
- `Decimal(string) bool` - Decimal numbers
- `Float(string) bool` - Floating-point numbers

### Number Type Validation
- `Even[T Numerable](T, ...bool) bool` - Even numbers
- `Odd[T Numerable](T, ...bool) bool` - Odd numbers
- `Whole[T Numerable](T) bool` - Whole numbers
- `Natural[T Numerable](T) bool` - Natural numbers
- `Positive[T Numerable](T) bool` - Positive numbers
- `Negative[T Numerable](T) bool` - Negative numbers
- `Zero[T Numerable](T) bool` - Zero value check

### Encoding & Format Validation
- `Base64(string) bool` - Base64 encoding
- `Base64URL(string) bool` - URL-safe Base64 encoding
- `Hex(string) bool` - Hexadecimal strings
- `Bin(string) bool` - Binary strings
- `HexColor(string) bool` - Hexadecimal color codes
- `RGBColor(string) bool` - RGB color format
- `MD5(string) bool` - MD5 hash strings
- `JWT(string) bool` - JSON Web Tokens

### Mobile & Telecom
- `IMEI[T string|int64](T) bool` - International Mobile Equipment Identity
- `IMSI(string) bool` - International Mobile Subscriber Identity

## Installation

```shell
go get -u github.com/goloop/is
```

## Important Note

The package provides validation functions that do not pre-clean the data. If the validation data needs to be cleaned, it must be cleaned beforehand, for example, using the [g package](https://github.com/goloop/g).

Example:
```go
raw := "GB82 WEST 1234 5698 7654 32"    // contains spaces
iban := g.Weed(raw, g.Whitespaces)      // remove spaces
valid := is.IBAN(iban)                  // validate
```

## Usage Examples

### Basic Validation

```go
package main

import (
    "fmt"
    "github.com/goloop/is"
)

func main() {
    // Email validation.
    fmt.Println(is.Email("user@example.com"))         // true
    fmt.Println(is.Email("invalid-email"))            // false

    // Phone number validation.
    fmt.Println(is.Phone("+1 (234) 567-8900"))        // true
    fmt.Println(is.E164("+12345678900"))              // true

    // Geographic coordinates.
    fmt.Println(is.Coordinates(51.5074, -0.1278))     // true
    fmt.Println(is.Coordinates(91.0, 0.0))            // false

    // Financial validation.
    fmt.Println(is.BankCard("4111111111111111"))      // true
    fmt.Println(is.IBAN("DE89370400440532013000"))    // true
}
```

### Advanced Validation

```go
package main

import (
    "fmt"
    "github.com/goloop/is"
)

func main() {
    // Variable name validation for different languages.
    fmt.Println(is.VariableNameFor("myVar", "go"))     // true
    fmt.Println(is.VariableNameFor("class", "python")) // false (reserved word)

    // Strict mode validation.
    fmt.Println(is.Nickname("user123", true))          // true
    fmt.Println(is.Nickname("user@123", true))         // false

    // Number validation.
    fmt.Println(is.Even(4))                            // true
    fmt.Println(is.Natural(5))                         // true
    fmt.Println(is.Positive(-1))                       // false

    // Format validation.
    fmt.Println(is.HexColor("#FF5733"))                // true
    fmt.Println(is.Base64("SGVsbG8="))                 // true
    fmt.Println(is.JWT("eyJhbGciOiJIUzI..."))          // true
}
```

## Contributing

Contributions are welcome! Here are a few ways you can help:

1. Report bugs by opening an issue
2. Suggest new features or improvements
3. Submit pull requests
4. Improve documentation
5. Share the package with others

Please make sure to:
- Write tests for new features
- Follow Go coding standards
- Update documentation as needed
- Add comments to explain complex logic

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

Created and maintained by [Goloop Team](https://github.com/goloop)

## Related Projects

- [g](https://github.com/goloop/g) - Go utility functions package
