package is

import (
	"strings"
	"testing"
	"unicode"
)

// The fuzz targets below assert invariants rather than just "does not panic":
// a validator that returns true must produce output a downstream consumer can
// rely on. They also exercise the panic-prone paths (VariableNameFor) across
// every supported language.

// FuzzEmail: a string accepted as an email must contain exactly one '@' and
// stay within the documented length bounds.
func FuzzEmail(f *testing.F) {
	for _, s := range []string{"user@example.com", "", "a@b.cd", "@@", "x"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if !Email(s) {
			return
		}
		if strings.Count(s, "@") != 1 {
			t.Errorf("Email(%q) accepted a string without a single '@'", s)
		}
		if len(s) < 6 || len(s) > 254 {
			t.Errorf("Email(%q) accepted an out-of-bounds length %d", s, len(s))
		}
	})
}

// FuzzBase64: an accepted standard Base64 string must have a length that is a
// multiple of four and must not contain a backslash (the BUG-01 invariant).
func FuzzBase64(f *testing.F) {
	for _, s := range []string{"SGVsbG8=", "ab\\=", "", "===="} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if !Base64(s) {
			return
		}
		if len(s)%4 != 0 {
			t.Errorf("Base64(%q) accepted a non-multiple-of-4 length", s)
		}
		if strings.ContainsRune(s, '\\') {
			t.Errorf("Base64(%q) accepted a backslash", s)
		}
	})
}

// FuzzNumeric: an accepted numeric string must contain at least one decimal
// digit (the BUG-03 invariant) and at most one decimal separator.
func FuzzNumeric(f *testing.F) {
	for _, s := range []string{"123", "-1.5", ".", "1.2.3", "٣٫١٤", "abc"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if !Numeric(s) {
			return
		}
		hasDigit := false
		seps := 0
		for _, r := range s {
			if _, ok := decimalSeparators[r]; ok {
				seps++
			} else if unicode.IsDigit(r) {
				hasDigit = true
			}
		}
		if !hasDigit {
			t.Errorf("Numeric(%q) accepted a string with no digit", s)
		}
		if seps > 1 {
			t.Errorf("Numeric(%q) accepted more than one separator", s)
		}
	})
}

// FuzzIban: Iban must never panic, and a valid IBAN must agree with the
// MOD-97 check value of its rearranged form.
func FuzzIban(f *testing.F) {
	for _, s := range []string{
		"GB82WEST12345698765432",
		"gb82 west 1234 5698 7654 32",
		"",
		"GB82WEST1234569876543@",
	} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if !Iban(s) {
			return
		}
		norm := strings.ToUpper(strings.ReplaceAll(s, " ", ""))
		rearranged := norm[4:] + norm[0:4]
		if CalculateIBANChecksum(rearranged) != 1 {
			t.Errorf("Iban(%q) accepted but check value != 1", s)
		}
	})
}

// FuzzVariableNameFor is the strongest BUG-06 guard: for ANY input string and
// EVERY supported language, the call must return without panicking.
func FuzzVariableNameFor(f *testing.F) {
	langs := make([]string, 0, len(reservedWords))
	for lang := range reservedWords {
		langs = append(langs, lang)
	}
	f.Add("x")
	f.Add("@@class")
	f.Add("$global?")
	f.Add("")
	f.Fuzz(func(t *testing.T, name string) {
		for _, lang := range langs {
			// Must not panic for any (name, language) pair.
			if _, err := VariableNameFor(name, lang); err != nil {
				t.Errorf("VariableNameFor(%q, %q) errored on a supported "+
					"language: %v", name, lang, err)
			}
		}
	})
}

// FuzzIPv4: IPv4 and IPv6 are mutually exclusive, and any address either
// accepts as IP.
func FuzzIPv4(f *testing.F) {
	for _, s := range []string{"127.0.0.1", "::1", "192.168.0.01", "", "x"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		v4, v6 := IPv4(s), IPv6(s)
		if v4 && v6 {
			t.Errorf("%q reported as both IPv4 and IPv6", s)
		}
		if (v4 || v6) != IP(s) {
			t.Errorf("IP(%q) disagrees with IPv4||IPv6", s)
		}
	})
}
