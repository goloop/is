package is

import (
	"errors"
	"strings"
	"testing"
)

// This file collects adversarial "from the attacker's side" regression tests
// for the bugs fixed in v2, plus coverage for the validators added in v2.
// Each block names the bug it guards against so a future regression is easy
// to trace.

// TestBase64RejectsBackslash guards BUG-01: the standard Base64 alphabet has
// no backslash. A 4-char group built from a backslash must not validate, and
// must not survive into a real decoder.
func TestBase64RejectsBackslash(t *testing.T) {
	bad := []string{
		`ab\=`, // 4 chars: a b \ =
		`\\\\`, // all backslashes
		`AB\/`, // backslash next to slash
		`A\BC`, // backslash in the middle
		`====`, // padding only
		`ab=c`, // padding in the wrong place
	}
	for _, s := range bad {
		if Base64(s) {
			t.Errorf("Base64(%q) = true; want false", s)
		}
	}

	good := []string{
		"SGVsbG8sIHdvcmxkIQ==",
		"SGVsbG8=",
		"Zm9vYmFy",
		"AB+/", // '+' and '/' ARE in the standard alphabet
	}
	for _, s := range good {
		if !Base64(s) {
			t.Errorf("Base64(%q) = false; want true", s)
		}
	}
}

// TestMD5RejectsPrefix guards BUG-09: a real MD5 digest is exactly 32 hex
// nibbles with no '0x'/'#' decoration. A 32-char string that only reaches 32
// by counting a prefix must be rejected.
func TestMD5RejectsPrefix(t *testing.T) {
	bad := []string{
		"0xd41d8cd98f00b204e9800998ecf842",  // 0x + 30 hex = 32 chars
		"#d41d8cd98f00b204e9800998ecf8427",  // #  + 31 hex = 32 chars
		"d41d8cd98f00b204e9800998ecf8427",   // 31 chars
		"d41d8cd98f00b204e9800998ecf8427ee", // 33 chars
		"g41d8cd98f00b204e9800998ecf8427e",  // 'g' is not hex
	}
	for _, s := range bad {
		if MD5(s) {
			t.Errorf("MD5(%q) = true; want false", s)
		}
	}

	if !MD5("d41d8cd98f00b204e9800998ecf8427e") {
		t.Error("MD5 rejected a valid 32-nibble digest")
	}
	if !MD5("D41D8CD98F00B204E9800998ECF8427E") {
		t.Error("MD5 rejected a valid uppercase digest")
	}
}

// TestSHAFamily checks the fixed-length SHA validators, including the
// off-by-one boundaries and the absence-of-prefix rule.
func TestSHAFamily(t *testing.T) {
	hex := func(n int) string { return strings.Repeat("a", n) }

	cases := []struct {
		fn   func(string) bool
		name string
		n    int // exact valid length
	}{
		{SHA1, "SHA1", 40},
		{SHA256, "SHA256", 64},
		{SHA512, "SHA512", 128},
	}
	for _, c := range cases {
		if !c.fn(hex(c.n)) {
			t.Errorf("%s rejected a valid %d-nibble digest", c.name, c.n)
		}
		if c.fn(hex(c.n - 1)) {
			t.Errorf("%s accepted %d nibbles (one short)", c.name, c.n-1)
		}
		if c.fn(hex(c.n + 1)) {
			t.Errorf("%s accepted %d nibbles (one long)", c.name, c.n+1)
		}
		if c.fn("0x" + hex(c.n-2)) {
			t.Errorf("%s accepted a 0x-prefixed value", c.name)
		}
	}
}

// TestUUID checks the canonical 8-4-4-4-12 form and common rejects.
func TestUUID(t *testing.T) {
	good := []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"550E8400-E29B-41D4-A716-446655440000",
		"00000000-0000-0000-0000-000000000000",
	}
	for _, s := range good {
		if !UUID(s) {
			t.Errorf("UUID(%q) = false; want true", s)
		}
	}
	bad := []string{
		"550e8400e29b41d4a716446655440000",       // no hyphens
		"550e8400-e29b-41d4-a716-44665544000",    // 11 in last group
		"550e8400-e29b-41d4-a716-4466554400000",  // 13 in last group
		"{550e8400-e29b-41d4-a716-446655440000}", // braces
		"550e8400-e29b-41d4-a716-44665544000g",   // non-hex
		"",
	}
	for _, s := range bad {
		if UUID(s) {
			t.Errorf("UUID(%q) = true; want false", s)
		}
	}
}

// TestIMSIRejectsSign guards BUG-05: an IMSI is 15 ASCII digits only; a
// leading or embedded sign must not slip through.
func TestIMSIRejectsSign(t *testing.T) {
	bad := []string{
		"+10150123456789", // 15 chars, first is '+'
		"-10150123456789",
		"31015012345678 ",  // trailing space
		"3101501234567890", // 16 digits
		"31015012345678",   // 14 digits
		"31015012345678a",
	}
	for _, s := range bad {
		if IMSI(s) {
			t.Errorf("IMSI(%q) = true; want false", s)
		}
	}
	if !IMSI("310150123456789") {
		t.Error("IMSI rejected a valid 15-digit identifier")
	}
}

// TestIMEINegativeInt64 guards the int64 path: a negative number renders with
// a '-', which is not a digit, so it must be rejected rather than mis-summed.
func TestIMEINegativeInt64(t *testing.T) {
	if IMEI(int64(-22593572995861)) {
		t.Error("IMEI accepted a negative int64")
	}
	if !IMEI(int64(522593572995861)) {
		t.Error("IMEI rejected a valid int64 IMEI")
	}
	if !IMEI("522593572995861") {
		t.Error("IMEI rejected a valid string IMEI")
	}
	if IMEI("+22593572995861") { // sign in the string form
		t.Error("IMEI accepted a signed string")
	}
}

// TestIPv4LeadingZeros guards BUG-04: leading zeros are ambiguous and are
// rejected (matching net/netip), which is also what the docs now promise.
func TestIPv4LeadingZeros(t *testing.T) {
	bad := []string{
		"192.168.0.01",
		"010.0.0.1",
		"1.2.3.04",
		"00.0.0.0",
		"256.0.0.1",
		"192.168.0",
		"192.168.0.1.1",
		"::1", // IPv6 is not IPv4
		"",
	}
	for _, s := range bad {
		if IPv4(s) {
			t.Errorf("IPv4(%q) = true; want false", s)
		}
	}
	good := []string{"0.0.0.0", "127.0.0.1", "192.168.0.1", "255.255.255.255"}
	for _, s := range good {
		if !IPv4(s) {
			t.Errorf("IPv4(%q) = false; want true", s)
		}
	}
}

// TestHipercardAnchored guards BUG-07: the Hipercard pattern must anchor the
// whole alternation, so a valid prefix with trailing garbage cannot match.
func TestHipercardAnchored(t *testing.T) {
	rx := cardPatterns[Hipercard]
	if rx.MatchString("6062821234567890YYY") {
		t.Error("Hipercard pattern matched trailing garbage (anchor bug)")
	}
	if rx.MatchString("3841001234567890extra") {
		t.Error("Hipercard pattern matched a non-anchored right branch")
	}
	if !rx.MatchString("6062821234567890") {
		t.Error("Hipercard pattern rejected a well-formed number")
	}
}

// TestGeoStringRejectsGoFloatSyntax guards BUG-10: a coordinate parsed from a
// string must be a plain decimal, never a Go-specific float literal.
func TestGeoStringRejectsGoFloatSyntax(t *testing.T) {
	bad := []string{
		"0x1p4",   // hex float == 16
		"0x1.8p1", // hex float == 3
		"1_0",     // digit separator == 10
		"1e1",     // scientific notation == 10
		"4.5e1",   // scientific notation
		"  45.0",  // surrounding space
		"45.0 ",
		"+",
		".",
		"",
		"NaN",
		"Inf",
	}
	for _, s := range bad {
		if Latitude(s) {
			t.Errorf("Latitude(%q) = true; want false", s)
		}
		if Longitude(s) {
			t.Errorf("Longitude(%q) = true; want false", s)
		}
	}

	// Plain decimals within range pass; out of range fails.
	if !Latitude("45.5") || !Latitude("-90") || !Latitude("90") {
		t.Error("Latitude rejected a valid decimal string")
	}
	if Latitude("90.0001") || Latitude("-90.1") {
		t.Error("Latitude accepted an out-of-range value")
	}
	if !Longitude("-180") || !Longitude("180") || !Longitude(".5") {
		t.Error("Longitude rejected a valid decimal string")
	}
	if Longitude("180.1") {
		t.Error("Longitude accepted an out-of-range value")
	}

	// The float64 path is unaffected and allocation-free.
	if !Latitude(0.0) || !Latitude(-90.0) || Latitude(90.1) {
		t.Error("Latitude float64 path misbehaved")
	}
}

// TestVariableNameForNoPanic guards BUG-06: every language present in the
// reserved-word tables must have an identifier config, so no documented call
// can panic with a nil function. Previously 15 languages (csharp, dart, …)
// dereferenced a nil checkFirst.
func TestVariableNameForNoPanic(t *testing.T) {
	for lang := range reservedWords {
		// Must not panic, and a supported language must not report
		// ErrLanguageNotSupported.
		_, err := VariableNameFor("x", lang)
		if err != nil {
			t.Errorf("VariableNameFor(\"x\", %q) returned error %v for a "+
				"supported language", lang, err)
		}
	}

	// Mainstream languages added in v2 accept a plain identifier.
	for _, lang := range []string{"csharp", "dart", "bash", "elixir",
		"erlang", "julia", "objectivec", "vbnet", "cobol", "fortran",
		"prolog", "eiffel", "assembly"} {
		ok, err := VariableNameFor("myVar", lang)
		if err != nil || !ok {
			t.Errorf("VariableNameFor(\"myVar\", %q) = (%v, %v); want (true, nil)",
				lang, ok, err)
		}
	}

	// An unsupported language is a clean, matchable error, not a panic.
	if ok, err := VariableNameFor("x", "klingon"); ok ||
		!errors.Is(err, ErrLanguageNotSupported) {
		t.Errorf("VariableNameFor for an unknown language = (%v, %v); "+
			"want (false, ErrLanguageNotSupported)", ok, err)
	}

	// markdown/regex were pruned: they are no longer "languages".
	for _, lang := range []string{"markdown", "regex"} {
		if _, err := VariableNameFor("x", lang); err == nil {
			t.Errorf("VariableNameFor accepted pruned pseudo-language %q", lang)
		}
	}
}

// TestMAC checks the standard textual MAC forms.
func TestMAC(t *testing.T) {
	good := []string{
		"00:1b:63:84:45:e6",
		"00-1B-63-84-45-E6",
		"001b.6384.45e6",
		"00:00:00:00:00:00",
	}
	for _, s := range good {
		if !MAC(s) {
			t.Errorf("MAC(%q) = false; want true", s)
		}
	}
	bad := []string{
		"00:1b:63:84:45",    // too few groups
		"00:1b:63:84:45:zz", // non-hex
		"001b6384 45e6",     // space
		"",
	}
	for _, s := range bad {
		if MAC(s) {
			t.Errorf("MAC(%q) = true; want false", s)
		}
	}
}

// TestURL checks the scheme+host requirement.
func TestURL(t *testing.T) {
	good := []string{
		"https://example.com",
		"http://example.com:8080/path?q=1#frag",
		"ftp://files.example.com",
		"https://192.168.0.1",
	}
	for _, s := range good {
		if !URL(s) {
			t.Errorf("URL(%q) = false; want true", s)
		}
	}
	bad := []string{
		"example.com",    // no scheme
		"/relative/path", // no scheme/host
		"https://",       // no host
		"mailto:a@b.com", // no host
		"://example.com", // no scheme
		"",
	}
	for _, s := range bad {
		if URL(s) {
			t.Errorf("URL(%q) = true; want false", s)
		}
	}
}

// TestHostname checks RFC 1123 hostname rules and length bounds.
func TestHostname(t *testing.T) {
	good := []string{
		"example.com",
		"sub.example.com",
		"localhost",
		"a-b.example",
		"xn--d1acufc.example", // punycode label
	}
	for _, s := range good {
		if !Hostname(s) {
			t.Errorf("Hostname(%q) = false; want true", s)
		}
	}
	bad := []string{
		"-bad.example",                  // label starts with '-'
		"bad-.example",                  // label ends with '-'
		"a..b",                          // empty label
		".example.com",                  // leading dot
		strings.Repeat("a", 64),         // label too long
		strings.Repeat("a.", 130) + "a", // whole name too long
		"exa mple.com",                  // space
		"",
	}
	for _, s := range bad {
		if Hostname(s) {
			t.Errorf("Hostname(%q) = true; want false", s)
		}
	}
}

// TestBankCardKindIsolation checks that requesting a specific kind does not
// accidentally validate a number of another brand, and that the Luhn check
// gates everything.
func TestBankCardKindIsolation(t *testing.T) {
	// A valid Visa is not a MasterCard.
	if BankCard("4111111111111111", MasterCard) {
		t.Error("Visa number validated as MasterCard")
	}
	// Luhn must hold even for a brand-matching prefix.
	if BankCard("4111111111111112", Visa) {
		t.Error("BankCard accepted a number that fails the Luhn check")
	}
	// Multiple kinds: matches if any one matches.
	if !BankCard("4111111111111111", MasterCard, Visa) {
		t.Error("BankCard rejected a number matching one of several kinds")
	}
	// An unknown (zero) CardKind matches nothing.
	if BankCard("4111111111111111", CardKind(0)) {
		t.Error("BankCard matched against an invalid zero kind")
	}
	// String() is stable for known and unknown kinds.
	if Visa.String() != "Visa" || CardKind(0).String() != "CardKind(0)" {
		t.Errorf("CardKind.String() unexpected: %q / %q",
			Visa.String(), CardKind(0).String())
	}
}

// TestEmailRejectsConsecutiveSeparators checks the tightened slug rule: a
// label run may not contain two '.'/'-' in a row, nor start/end with one.
func TestEmailRejectsConsecutiveSeparators(t *testing.T) {
	bad := []string{
		"a..b@example.com",
		"a--b@example.com",
		"a.-b@example.com",
		"user@ex..ample.com",
		"user@example..com",
		"user@-example.com",
		"user@example-.com",
	}
	for _, s := range bad {
		if Email(s) {
			t.Errorf("Email(%q) = true; want false", s)
		}
	}
	good := []string{
		"a.b@example.com",
		"a-b@ex-ample.com",
		"first.last@sub.example.com",
	}
	for _, s := range good {
		if !Email(s) {
			t.Errorf("Email(%q) = false; want true", s)
		}
	}
}

// TestDomain checks the TLD requirement that separates Domain from Hostname.
func TestDomain(t *testing.T) {
	good := []string{"example.com", "sub.example.com", "a.bc", "xn--p1ai.com"}
	for _, s := range good {
		if !Domain(s) {
			t.Errorf("Domain(%q) = false; want true", s)
		}
	}
	bad := []string{
		"localhost",    // no TLD
		"example.123",  // numeric TLD
		"example.c",    // one-letter TLD
		"example.co-m", // non-letter in TLD
		"-bad.com",     // invalid hostname
		"",
	}
	for _, s := range bad {
		if Domain(s) {
			t.Errorf("Domain(%q) = true; want false", s)
		}
	}
	// Every valid Domain is also a valid Hostname, but not vice versa.
	if Domain("localhost") || !Hostname("localhost") {
		t.Error("Domain/Hostname relationship broken for localhost")
	}
}

// TestIBANCountry checks that a country code is returned only for a valid
// IBAN, and is normalized in non-strict mode.
func TestIBANCountry(t *testing.T) {
	if c, ok := IBANCountry("DE89370400440532013000"); !ok || c != "DE" {
		t.Errorf("IBANCountry(valid DE) = (%q, %v); want (DE, true)", c, ok)
	}
	// Non-strict normalizes case and spaces.
	if c, ok := IBANCountry("gb82 west 1234 5698 7654 32"); !ok || c != "GB" {
		t.Errorf("IBANCountry(spaced gb) = (%q, %v); want (GB, true)", c, ok)
	}
	// Invalid checksum yields no country.
	if c, ok := IBANCountry("DE0037040044053201300"); ok || c != "" {
		t.Errorf("IBANCountry(invalid) = (%q, %v); want (\"\", false)", c, ok)
	}
	// Strict mode rejects the spaced form, so no country.
	if c, ok := IBANCountry("GB82 WEST 1234 5698 7654 32", true); ok || c != "" {
		t.Errorf("IBANCountry(spaced, strict) = (%q, %v); want (\"\", false)",
			c, ok)
	}
}

// TestIPv6Zone checks that a zoned address is rejected while ordinary and
// IPv4-mapped IPv6 addresses are accepted.
func TestIPv6Zone(t *testing.T) {
	if IPv6("fe80::1%eth0") {
		t.Error("IPv6 accepted a zoned address")
	}
	good := []string{"::1", "2001:db8::1", "::ffff:1.2.3.4"}
	for _, s := range good {
		if !IPv6(s) {
			t.Errorf("IPv6(%q) = false; want true", s)
		}
	}
	// A zoned address must not count as an IP at all.
	if IP("fe80::1%eth0") {
		t.Error("IP accepted a zoned IPv6 address")
	}
}

// TestPhoneLength checks the E.164 upper bound and that separators do not
// inflate the digit count.
func TestPhoneLength(t *testing.T) {
	if !Phone("+123456789012345") { // exactly 15 digits
		t.Error("Phone rejected a 15-digit number")
	}
	if Phone("+1234567890123456") { // 16 digits
		t.Error("Phone accepted a 16-digit number")
	}
	if !Phone("+1-23-456-789-012-345") { // 15 digits with separators
		t.Error("Phone rejected 15 digits split by separators")
	}
	if Phone("+") { // no digits
		t.Error("Phone accepted a lone plus sign")
	}
}

// TestInternalEdgeBranches exercises a few defensive branches that production
// callers never reach (length guards are checked earlier), keeping behavior
// pinned even if the call sites change.
func TestInternalEdgeBranches(t *testing.T) {
	// isASCIIDigits rejects the empty string.
	if isASCIIDigits("") {
		t.Error("isASCIIDigits(\"\") = true; want false")
	}
	// cardChecker: an input made only of separators cleans down to "".
	if BankCard("---", Visa) || BankCard("   ") {
		t.Error("BankCard accepted an all-separator input")
	}
	// parseCoordinate: a value that matches the decimal regex but overflows
	// float64 must fail (ParseFloat returns ErrRange), not panic.
	huge := strings.Repeat("9", 400)
	if Latitude(huge) || Longitude(huge) {
		t.Error("coordinate accepted an overflowing decimal string")
	}
}

// TestNickRejectsWhitespace pins the no-cleaning contract for Nickname.
func TestNickRejectsWhitespace(t *testing.T) {
	for _, s := range []string{" user", "user ", "us er", "\tuser", "user\n"} {
		if Nickname(s) {
			t.Errorf("Nickname(%q) = true; want false (no trimming)", s)
		}
	}
	if !Nickname("user_123") {
		t.Error("Nickname rejected a valid identifier")
	}
}
