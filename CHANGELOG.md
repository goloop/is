# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0]

This is a major release published under the module path
`github.com/goloop/is/v2`. It fixes a reachable panic and several false
positives, and modernizes the implementation. Go 1.24+ is required.

### Added
- `SHA1`, `SHA256`, `SHA512` — fixed-length hash digest validators.
- `UUID` — canonical 8-4-4-4-12 UUID validator.
- `URL`, `Hostname`, `Domain`, `MAC` — network identifier validators.
- `IBANCountry` — returns the country code of a valid IBAN.
- `CardKind` type and brand constants for `BankCard`.
- Identifier configs for C#, Dart, Bash, Elixir, Erlang, Julia,
  Objective-C, VB.NET, COBOL, Fortran, Prolog, Eiffel, and Assembly, so
  `VariableNameFor` supports them properly.
- Adversarial regression tests, runnable examples, and fuzz targets.

### Fixed
- `VariableNameFor`/`VarFor` no longer panic with a nil-pointer dereference
  for languages that had reserved words but no identifier config (15
  languages, including C#).
- `Numeric` no longer reports letter-based numeral systems (Hebrew, Greek,
  Armenian, …), CJK/Roman numerals, or a lone separator/sign as numeric.
- `Base64` no longer accepts a backslash as a valid character.
- `MD5` no longer accepts a `0x`/`#` prefix.
- `IPv4` rejects leading zeros, and its documentation now matches.
- `IMSI`/`IMEI` accept digits only; a leading `+`/`-` is rejected.
- `BankCard`'s Hipercard pattern is fully anchored.
- String `Latitude`/`Longitude` reject scientific notation, hexadecimal
  floats, and digit separators.

### Changed (breaking)
- Module path is now `github.com/goloop/is/v2`.
- `BankCard(string, ...CardKind)` replaces `BankCard(string, ...*regexp.Regexp)`;
  the per-brand regular expressions are no longer exported.
- `CalculateIBANChecksum` returns `int` (the MOD-97 check value, `-1` on an
  invalid character) instead of `*big.Int`.
- `ErrLanguageNotSupported` is now a sentinel `error` value (was a
  `func(string) error`); test for it with `errors.Is`. The error message
  format changed accordingly.
- `Nickname` no longer trims surrounding whitespace.

### Changed
- `Phone` now ignores hyphens and dots in addition to spaces and parentheses,
  so grouped numbers like `+1-234-567-8900` validate; it also caps the number
  at 15 digits (the E.164 maximum).
- `IPv6` rejects a scoped address carrying a zone identifier
  (e.g. `fe80::1%eth0`).
- `Email` rejects consecutive `.`/`-` separators within a label
  (e.g. `a..b@example.com`).
- `Title`'s word-boundary behavior (any non-letter, including apostrophe and
  hyphen) is now documented explicitly.

### Performance
- IBAN validation uses incremental integer MOD-97 arithmetic instead of
  `math/big`, removing per-call allocations and package-level cache state.
- `IMSI`/`IMEI`/`Latitude`/`Longitude` avoid intermediate string allocations.
- `BankCard` skips the cleanup pass when the input has no spaces or hyphens.
- Strict `VariableName` looks reserved words up in a single precomputed set
  (built once via `sync.OnceValue`) instead of scanning every language.
