# is — reference

The full reference for the `is` package: the validation contract, every
function grouped by topic, the `CardKind` brands, and practical recipes.

Ukrainian version: **[DOC.UK.md](DOC.UK.md)**.

## Contents

- [The validation contract](#the-validation-contract)
- [Type constraints](#type-constraints)
- [Account and identity](#account-and-identity)
- [Financial](#financial)
- [Geographic](#geographic)
- [Network and communication](#network-and-communication)
- [String character classes](#string-character-classes)
- [Numbers](#numbers)
- [Encoding and format](#encoding-and-format)
- [Mobile and telecom](#mobile-and-telecom)
- [Errors](#errors)
- [Recipes and tips](#recipes-and-tips)

## The validation contract

Every function in `is` answers a single yes/no question about the **format** of
its input and returns a `bool` (a few return an extra value or an error). Three
rules hold throughout:

1. **No cleaning, no normalisation.** `is` validates exactly what you pass in.
   It never trims whitespace, strips separators or lower-cases anything. If the
   input needs cleaning, clean it first (for example with the `g` package's
   `Weed`/`Preserve`), then validate.
2. **Strictness is explicit.** Where a looser and a stricter reading both make
   sense, the stricter one is opt-in through a trailing `strict ...bool`
   argument, so the default never surprises you.
3. **No false positives by design.** v2 tightened many validators (`Numeric`,
   `MD5`, `Base64`, `IPv4`, string `Latitude`/`Longitude`, `IMEI`/`IMSI`) so
   that ordinary words and malformed values no longer slip through.

```go
import "github.com/goloop/is/v2"
```

## Type constraints

Two constraints appear in the generic validators:

| Constraint | Permits |
|------------|---------|
| `Numerable`  | every integer and float type |
| `Verifiable` | integer, float, `string`, `rune` |

Some validators are generic over `string | float64` (the geographic helpers) or
`string | int64` (`IMEI`), so you can validate either a parsed number or its
textual form with the same call.

## Account and identity

```go
func Email(email string) bool
func Nickname(nickname string, strict ...bool) bool
func VariableName(v string, strict ...bool) bool          // alias: Var
func VariableNameFor(v string, language string) (bool, error) // alias: VarFor
func SelectorName(v string, strict ...bool) bool          // alias: Sel
```

`Email` validates the address format (and rejects consecutive `.`/`-`
separators inside a label). `Nickname` validates a username; in strict mode only
a conservative character set is allowed. It does **not** trim surrounding
whitespace.

`VariableName` checks a general programming identifier; `VariableNameFor`
validates against a specific language's rules **and reserved words** (Go,
Python, C#, Dart, Bash, Elixir, Erlang, Julia, Objective-C, VB.NET, COBOL,
Fortran, Prolog, Eiffel, Assembly and more), returning
[`ErrLanguageNotSupported`](#errors) for an unknown language. `SelectorName`
validates a CSS-style selector name. `Var`, `VarFor` and `Sel` are short
aliases.

```go
is.Email("user@example.com")            // true
is.Nickname("user@123", true)           // false (strict)
ok, err := is.VariableNameFor("class", "python") // false, nil — reserved word
```

## Financial

```go
func BankCard(str string, kinds ...CardKind) bool
func IBAN(iban string, strict ...bool) bool               // alias: Iban
func IBANCountry(iban string, strict ...bool) (string, bool)
func CalculateIBANChecksum(iban string) int
```

`BankCard` validates a card number (length and Luhn); pass one or more
`CardKind` constants to also require a specific brand. The brands are opaque
constants rather than mutable regexes, so validation can't be silently broken
process-wide:

| Kind constants (selection) |
|----------------------------|
| `Visa`, `VisaElectron`, `MasterCard`, `Maestro`, `DiscoverCard`, `JCB`, `UnionPay`, `ChinaUnionPay`, `DinersClub` (+ `DinersClubCarteBlanche`, `DinersClubInternational`, `DinersClubUSAndCanada`), … |

`IBAN` validates an International Bank Account Number (MOD-97 checksum); strict
mode also enforces the exact per-country length. `IBANCountry` returns the
two-letter country code when the IBAN is valid. `CalculateIBANChecksum` returns
the MOD-97 value as an `int` (`-1` on an invalid character).

```go
is.BankCard("4111111111111111")            // true
is.BankCard("4111111111111111", is.Visa)   // true
is.IBAN("DE89370400440532013000")          // true
code, ok := is.IBANCountry("DE89370400440532013000") // "DE", true
```

## Geographic

```go
func Latitude[T string | float64](lat T) bool
func Longitude[T string | float64](lon T) bool
func Coordinates[T string | float64](lat, lon T) bool
```

Validate a latitude (`-90..90`), longitude (`-180..180`), or a pair. The string
path accepts plain decimals only — scientific notation, hexadecimal floats and
digit separators are rejected; the `float64` path validates the numeric range.

```go
is.Coordinates(51.5074, -0.1278) // true
is.Latitude("91.0")              // false
```

## Network and communication

```go
func IPv4(ip string) bool     func IPv6(ip string) bool     func IP(ip string) bool
func MAC(v string) bool
func URL(v string) bool
func Hostname(v string) bool  func Domain(v string) bool
func Phone(phone string) bool func E164(v string) bool
```

`IPv4` matches dotted-quad addresses and (like `net/netip`) rejects leading
zeros; `IPv6` matches IPv6; `IP` accepts either. `MAC` validates an IEEE 802
address. `URL` requires an absolute URL (scheme + host). `Hostname` follows
RFC 1123; `Domain` is a hostname with a letter TLD. `Phone` accepts common
separators (spaces, hyphens, dots, parentheses); `E164` requires the strict
`+` and up to 15 digits.

```go
is.IP("2001:db8::1")               // true
is.Domain("example.com")           // true
is.Phone("+1 (234) 567-8900")      // true
is.E164("+12345678900")            // true
```

## String character classes

```go
func Alpha(s string) bool    func Alnum(s string) bool    func Digit(s string) bool
func Lower(s string) bool    func Upper(s string) bool    func Title(s string) bool
func Space(s string) bool
func Numeric(s string) bool  func Decimal(s string) bool  func Float(s string) bool
```

Character-class predicates over the whole string. `Alpha` is letters only,
`Alnum` letters and digits, `Digit` ASCII digits, `Lower`/`Upper`/`Title` case
checks, `Space` whitespace. `Numeric` accepts only true decimal digits (Unicode
category `Nd`) — letter-based numerals, CJK/Roman numerals and lone signs are
rejected. `Decimal` and `Float` validate numeric literals.

```go
is.Alnum("abc123")   // true
is.Numeric("Ⅻ")      // false (Roman numeral, not Nd)
is.Float("3.14")     // true
```

## Numbers

```go
func Even[T Numerable](v T, f ...bool) bool
func Odd[T Numerable](v T, f ...bool) bool
func Whole[T Numerable](v T) bool
func Natural[T Numerable](v T) bool
func Positive[T Numerable](v T) bool
func Negative[T Numerable](v T) bool
func Zero[T Numerable](v T) bool
```

Numeric property checks for any integer or float type. `Whole` reports an
integral value, `Natural` a positive whole number, and `Zero`/`Positive`/
`Negative` the sign.

```go
is.Even(4)       // true
is.Natural(5)    // true
is.Positive(-1)  // false
```

## Encoding and format

```go
func Base64(v string) bool   func Base64URL(v string) bool
func Hex(v string) bool      func Bin(v string) bool
func HexColor(v string) bool func RGBColor(v string) bool
func MD5(v string) bool      func SHA1(v string) bool
func SHA256(v string) bool   func SHA512(v string) bool
func UUID(v string) bool     func JWT(v string) bool
```

Format validators. `Base64` uses the standard alphabet (`A–Z a–z 0–9 + /`);
`Base64URL` the URL-safe alphabet. `Hex`/`Bin` validate the digit set.
`HexColor` (`#RRGGBB`) and `RGBColor` validate colours. The hash validators
require exactly the right number of hexadecimal nibbles (MD5 = 32, SHA-1 = 40,
SHA-256 = 64, SHA-512 = 128) with no prefix. `UUID` validates the canonical
`8-4-4-4-12` form; `JWT` the three-segment token shape.

```go
is.Base64("SGVsbG8=")   // true
is.HexColor("#FF5733")  // true
is.UUID("550e8400-e29b-41d4-a716-446655440000") // true
```

## Mobile and telecom

```go
func IMEI[T string | int64](imei T) bool
func IMSI(imsi string) bool
```

`IMEI` validates the 15-digit equipment identity (with Luhn check) from a string
or an `int64`; `IMSI` the subscriber identity. Both accept digits only — a
leading sign is rejected.

```go
is.IMEI("490154203237518") // true
is.IMSI("310150123456789") // true
```

## Errors

```go
var ErrLanguageNotSupported = errors.New("programming language is not supported")
```

Returned by `VariableNameFor`/`VarFor` for an unknown language. It is a sentinel,
so match it with `errors.Is`:

```go
if _, err := is.VariableNameFor(name, lang); errors.Is(err, is.ErrLanguageNotSupported) {
    // unsupported language
}
```

## Recipes and tips

**Clean, then validate.** `is` never cleans input. Pair it with `g` to strip
noise first:

```go
raw := "GB82 WEST 1234 5698 7654 32"
iban := g.Weed(raw, g.Whitespaces) // remove spaces
ok := is.IBAN(iban)                // validate
```

**Use strict mode at trust boundaries.** For user-supplied nicknames and
identifiers, pass `true` so only the conservative character set is accepted;
keep the lenient default for internal data you already trust.

**Require a brand only when you must.** `BankCard(num)` accepts any valid brand;
add `CardKind` constants only when your flow is limited to specific networks —
over-constraining rejects otherwise-valid cards.

**Validate the value, not its text, when you already have a number.** The
generic `Latitude`/`Longitude`/`Coordinates` and `IMEI` accept the parsed type
directly, avoiding a string round-trip.
