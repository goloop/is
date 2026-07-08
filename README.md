![Go Report Card](https://img.shields.io/badge/go%20report-retired-lightgrey) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/is/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://pkg.go.dev/github.com/goloop/is/v2) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# is

`is` is a comprehensive set of format-validation functions for the data types
you meet in web, financial and general Go applications — emails, phone numbers,
bank cards, IBANs, coordinates, IP/host/URL, hashes, UUIDs, character classes
and number properties.

Every function answers a single yes/no question about the **format** of its
input. The package does not clean or normalise data: it validates exactly what
you pass in. Need to strip spaces or separators first? Clean the input (for
example with the [`g`](https://github.com/goloop/g) package), then validate.

## Features

- **Identity** — `Email`, `Nickname`, `VariableName`, `VariableNameFor`
  (language-aware, reserved words included).
- **Financial** — `BankCard` (Luhn + optional brand), `IBAN`/`IBANCountry`.
- **Geographic** — `Latitude`, `Longitude`, `Coordinates` (string or `float64`).
- **Network** — `IPv4`/`IPv6`/`IP`, `MAC`, `URL`, `Hostname`, `Domain`,
  `Phone`, `E164`.
- **Character classes** — `Alpha`, `Alnum`, `Digit`, `Numeric`, `Decimal`,
  `Float`, `Lower`/`Upper`/`Title`, `Space`.
- **Numbers** — `Even`/`Odd`, `Whole`, `Natural`, `Positive`/`Negative`/`Zero`.
- **Encoding & format** — `Base64`/`Base64URL`, `Hex`/`Bin`, `HexColor`/
  `RGBColor`, `MD5`/`SHA1`/`SHA256`/`SHA512`, `UUID`, `JWT`.
- **Telecom** — `IMEI`, `IMSI`.

## Installation

```shell
go get -u github.com/goloop/is/v2
```

```go
import "github.com/goloop/is/v2"
```

Requires Go 1.24 or newer. The package has no third-party dependencies.

## Quick start

```go
package main

import (
    "fmt"

    "github.com/goloop/is/v2"
)

func main() {
    fmt.Println(is.Email("user@example.com"))         // true
    fmt.Println(is.Phone("+1 (234) 567-8900"))        // true
    fmt.Println(is.E164("+12345678900"))              // true
    fmt.Println(is.Coordinates(51.5074, -0.1278))     // true

    // BankCard: any brand, or require a specific one.
    fmt.Println(is.BankCard("4111111111111111"))          // true
    fmt.Println(is.BankCard("4111111111111111", is.Visa)) // true
    fmt.Println(is.IBAN("DE89370400440532013000"))        // true

    // Formats and hashes.
    fmt.Println(is.HexColor("#FF5733"))                             // true
    fmt.Println(is.UUID("550e8400-e29b-41d4-a716-446655440000"))    // true
}
```

Because `is` never cleans input, strip noise first when needed:

```go
raw := "GB82 WEST 1234 5698 7654 32"   // contains spaces
iban := g.Weed(raw, g.Whitespaces)     // remove them
ok := is.IBAN(iban)                    // validate
```

## Documentation

- Full reference and recipes: [DOC.md](DOC.md) · [DOC.UK.md](DOC.UK.md)
- Package API: [pkg.go.dev/github.com/goloop/is/v2](https://pkg.go.dev/github.com/goloop/is/v2)
- Changes between versions: [CHANGELOG.md](CHANGELOG.md)

## Contributing

Contributions are welcome. Please run `go test ./...`, `go vet ./...` and
`gofmt -l .` before submitting a pull request.

## License

`is` is released under the MIT License. See [LICENSE](LICENSE).
