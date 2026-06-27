[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/is/v2)](https://goreportcard.com/report/github.com/goloop/is/v2) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/is/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://pkg.go.dev/github.com/goloop/is/v2) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# is

The `is` package provides a comprehensive set of validation functions for various data types commonly used in web applications, financial systems, and general software development. It offers a clean, efficient way to validate data formats ranging from basic types to complex financial and network identifiers.

Every function answers a single yes/no question about the *format* of its input. The package does not clean or normalize data: it validates what you pass in, exactly as you pass it in.

## Features

### Account & Identity Validation
- `Email(string) bool` - Email addresses
- `Nickname(string, ...bool) bool` - Usernames (with optional strict mode)
- `VariableName(string, ...bool) bool` - Programming variable names
- `VariableNameFor(string, string) (bool, error)` - Language-specific variable names

### Financial Validation
- `BankCard(string, ...CardKind) bool` - Credit/Debit card numbers with optional brand checking
- `IBAN(string, ...bool) bool` - International Bank Account Numbers
- `Iban(string, ...bool) bool` - Alias for IBAN validation
- `IBANCountry(string, ...bool) (string, bool)` - IBAN country code when valid

### Geographic Validation
- `Latitude[T string|float64](T) bool` - Latitude coordinates
- `Longitude[T string|float64](T) bool` - Longitude coordinates
- `Coordinates[T string|float64](lat, lon T) bool` - Coordinate pairs

### Network & Communication
- `IPv4(string) bool` - IPv4 addresses
- `IPv6(string) bool` - IPv6 addresses
- `IP(string) bool` - Any IP address (v4 or v6)
- `MAC(string) bool` - IEEE 802 MAC addresses
- `URL(string) bool` - Absolute URLs (scheme + host)
- `Hostname(string) bool` - RFC 1123 hostnames
- `Domain(string) bool` - Domain names (hostname with a letter TLD)
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
- `Numeric(string) bool` - Numeric strings (decimal digits of any script)
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
- `SHA1(string) bool` - SHA-1 hash strings
- `SHA256(string) bool` - SHA-256 hash strings
- `SHA512(string) bool` - SHA-512 hash strings
- `UUID(string) bool` - UUID (canonical 8-4-4-4-12 form)
- `JWT(string) bool` - JSON Web Tokens

### Mobile & Telecom
- `IMEI[T string|int64](T) bool` - International Mobile Equipment Identity
- `IMSI(string) bool` - International Mobile Subscriber Identity

## Installation

```shell
go get -u github.com/goloop/is/v2
```

```go
import "github.com/goloop/is/v2"
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

    "github.com/goloop/is/v2"
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
    fmt.Println(is.BankCard("4111111111111111", is.Visa)) // true
    fmt.Println(is.IBAN("DE89370400440532013000"))    // true
}
```

### Advanced Validation

```go
package main

import (
    "fmt"

    "github.com/goloop/is/v2"
)

func main() {
    // Variable name validation for different languages.
    ok, _ := is.VariableNameFor("myVar", "go")
    fmt.Println(ok)                                    // true
    ok, _ = is.VariableNameFor("class", "python")
    fmt.Println(ok)                                    // false (reserved word)

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
    fmt.Println(is.UUID("550e8400-e29b-41d4-a716-446655440000")) // true
}
```

## Migrating to v2

v2 is published under the module path `github.com/goloop/is/v2` and requires Go 1.24+. It tightens several validators that previously produced false positives, and reworks `BankCard` for safety. Update your import path and review the breaking changes below.

### Breaking changes

- **`BankCard` brands are now an opaque `CardKind`.** The card patterns are no longer exported as mutable `*regexp.Regexp` variables; pass kind constants instead.

  ```go
  // before (v1):
  is.BankCard(num, is.Visa)            // is.Visa was a *regexp.Regexp
  // after (v2):
  is.BankCard(num, is.Visa)            // is.Visa is a CardKind constant
  ```

- **`Numeric` accepts only decimal digits (Unicode category Nd).** Letter-based numeral systems (Hebrew, Greek, Armenian, …), CJK/Roman numerals, and a lone separator or sign are now rejected. ASCII, Arabic-Indic, Devanagari, and other true decimal digits still pass. This fixes false positives where ordinary words validated as numbers.

- **`MD5` no longer accepts a `0x`/`#` prefix.** It now requires exactly 32 hexadecimal nibbles, matching the new `SHA1`/`SHA256`/`SHA512`.

- **`Base64` no longer accepts a backslash.** The standard alphabet is `A-Z a-z 0-9 + /` only.

- **`IPv4` rejects leading zeros** (e.g. `192.168.0.01`), matching `net/netip`. The documentation now reflects this.

- **String `Latitude`/`Longitude` accept plain decimals only.** Scientific notation (`1e1`), hexadecimal floats (`0x1p4`), and digit separators (`1_0`) are rejected. The `float64` path is unchanged.

- **`IMSI`/`IMEI` accept digits only.** A leading `+`/`-` is no longer tolerated.

- **`Nickname` no longer trims surrounding whitespace.** This restores the package's no-cleaning contract; trim the input yourself if needed.

- **`CalculateIBANChecksum` returns `int`** (the MOD-97 check value, `-1` on an invalid character) instead of `*big.Int`.

- **`ErrLanguageNotSupported` is a sentinel `error`** (was a `func(string) error`). Match it with `errors.Is`:

  ```go
  if _, err := is.VariableNameFor(name, lang); errors.Is(err, is.ErrLanguageNotSupported) {
      // unsupported language
  }
  ```

### New in v2

- `SHA1`, `SHA256`, `SHA512`, `UUID`, `URL`, `Hostname`, `Domain`, `MAC`, `IBANCountry`.
- `VariableNameFor` no longer panics for any supported language and now fully supports C#, Dart, Bash, Elixir, Erlang, Julia, Objective-C, VB.NET, COBOL, Fortran, Prolog, Eiffel, and Assembly.
- `Phone` accepts hyphen and dot separators.
- `Email` rejects consecutive `.`/`-` separators in a label.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
