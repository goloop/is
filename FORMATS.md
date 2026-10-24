# Structured formats

All functions validate syntax, ranges or checksums. They do not establish that
an account, identifier, postcode, media type or country-specific service exists.
No network access is performed.

## Dates and time

- `Date(string)` requires `YYYY-MM-DD`, years 0001-9999, with Gregorian leap days.
- `RFC3339(string)` accepts upper/lowercase T/Z and fractional seconds. It rejects
  leap-second `:60`, as Go's `time.Time` cannot represent leap seconds. Offsets
  require hours 00-23 and minutes 00-59; comma fractions and single-digit hours fail.
- `UTCOffset(string)` accepts uppercase `Z` or signed `HH:MM`, including `-00:00`.
  This validates RFC syntax, not whether a civil timezone uses that offset.
- `Duration(string)` uses `time.ParseDuration`, including signed values and `0`.
  ISO 8601 calendar periods such as `P1M` are outside this contract.

Source: [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339).

## Network and technical syntax

- `CIDR(string)` accepts IPv4/IPv6 prefixes with host bits; zones are rejected.
- `IPPort(string)` requires an IP literal, IPv6 brackets and decimal port.
  `Port(string)` accepts 0-65535 with up to five digits, including leading zeros.
- `JSON(string)` requires valid UTF-8 and one complete JSON value; scalars count.
- `MIME(string)` uses `mime.ParseMediaType` with a concrete type/subtype and optional
  parameters. No wildcards, CR/LF, leading/trailing spaces, or registry lookup.
- `SemVer(string)` follows [SemVer 2.0.0](https://semver.org/spec/v2.0.0.html).
  No leading `v` or leading zero in a numeric prerelease identifier.
- `Base32(string, padded ...bool)` accepts nonempty canonical uppercase RFC 4648
  encoding. Default is padded; `false` selects unpadded. Pad bits must be zero.
- `SHA384(string)` accepts exactly 96 hexadecimal characters, in either case.

## Identifiers and finance

`UUID` retains its existing broad textual-format contract. `UUIDv4` and `UUIDv7`
add version and variant checks from [RFC 9562](https://www.rfc-editor.org/rfc/rfc9562).
They do not check timestamp plausibility, randomness or uniqueness.

`ISBN` accepts ISBN-10 (modulo 11, uppercase X allowed only at the end) and ISBN-13
(GS1 modulo 10, prefix 978 or 979 excluding 9790/ISMN). Single spaces/hyphens may
separate groups; registration-group allocation and exact hyphen positions are
not verified. `GTIN` accepts 8/12/13/14 digits; `EAN` accepts 8/13 digits. Both
check the GS1 check digit, not allocation. References:
[ISBN Agency](https://www.isbn-international.org/index.php/content/what-isbn/10),
[GS1 check digit](https://www.gs1.org/services/how-calculate-check-digit-manually).

`IBAN` validates national BBAN character positions, length and MOD-97 in both
modes. Check digits must be 02-98. Non-strict mode strips ASCII spaces and
uppercases input; strict mode requires the electronic form. The catalog has
89 countries, including HN and YE, from [SWIFT registry release 101, December
2025](https://www.swift.com/sites/default/files/files/iban-registry-v101.pdf).
It is a pinned reference, not a claim to track newer registry releases live.
Domestic account checksums, bank assignments and account existence are outside
this contract. Reference examples are in `testdata/iban_swift_101.txt`.

`BIC` checks ISO 9362 syntax: four alphanumeric characters, two country letters,
two alphanumeric characters, then an optional three-character branch. Assigned
ISO country codes and XK are accepted; lowercase is not. The first four
characters may contain digits under ISO 9362, although SWIFT uses letters.
Reference: [SWIFT BIC structure](https://www.swift.com/standards/data-standards/bic-business-identifier-code).

## Regional catalogs

`CountryCode` checks 249 assigned uppercase ISO 3166-1 alpha-2 codes (snapshot
2026-10-24); reserved/user-assigned codes such as UK and XK are excluded.
Reference: [ISO online browsing platform](https://www.iso.org/obp/ui/#search/code/).

`CurrencyCode` checks 178 current uppercase codes from [SIX List One, published
2026-09-17](https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml).
This includes funds, metals, XTS and XXX. It is not a legal-tender check.

`PostalCode(value, country)` supports US, GB, DE, PL, UA, FR, ES, IT, PT, NL and CZ.
Country codes must be uppercase ISO alpha-2. Separators and letter case are
canonical: `00-001` (PL), `SW1A 1AA` (GB), `1000-001` (PT), `1012 AB` (NL),
`110 00` (CZ). US ZIP+4 is accepted. Unsupported countries return false.
These are regional shape checks, not a database of assigned codes or deliverable
addresses. Reference: [UPU addressing systems](https://www.upu.int/en/Postal-Solutions/Programmes-Services/Addressing-Solutions).

Catalog updates should pin the source date/version, compare changed assignments
and structures, and update positive/negative reference tests together.
