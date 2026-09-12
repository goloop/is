package is

import "testing"

// TestCardPatternsSixteenDigits guards BUG-01/02/05/08: brand patterns accept
// real 16-digit numbers and reject the short/oversized ones the old patterns
// allowed.
func TestCardPatternsSixteenDigits(t *testing.T) {
	cases := []struct {
		name string
		card string
		kind CardKind
		want bool
	}{
		{"RuPay 16-digit valid", "6069812345678901", RuPay, true},
		{"RuPay 6-digit prefix valid", "6085001234567892", RuPay, true},
		{"RuPay 14-digit rejected", "60698012345672", RuPay, false},
		{"Cabal 16-digit valid", "6042011234567893", Cabal, true},
		{"Cabal 12-digit rejected", "604201123451", Cabal, false},
		{"VisaElectron 417500 16-digit valid", "4175001234567898", VisaElectron, true},
		{"VisaElectron 417500 18-digit rejected", "417500123456789016", VisaElectron, false},
		{"MIR 2204 valid", "2204123456789015", MIR, true},
		{"MIR 2205 rejected", "2205123456789014", MIR, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := BankCard(c.card, c.kind); got != c.want {
				t.Errorf("BankCard(%q, %v) = %v, want %v", c.card, c.kind, got, c.want)
			}
		})
	}
}

// TestBankCardUnionIncludesAllBrands guards BUG-03: the no-kind call is the
// union of every brand, so a valid MIR/Maestro number validates without an
// explicit kind.
func TestBankCardUnionIncludesAllBrands(t *testing.T) {
	mir := "2204123456789015"
	if !BankCard(mir, MIR) {
		t.Fatalf("precondition: BankCard(%q, MIR) should be true", mir)
	}
	if !BankCard(mir) {
		t.Errorf("BankCard(%q) with no kind = false, want true (union of all brands)", mir)
	}
}

// TestEmailAtextLocalPart guards BUG-04: plus-addressing and underscores are
// accepted while leading/trailing and doubled separators stay rejected.
func TestEmailAtextLocalPart(t *testing.T) {
	cases := map[string]bool{
		"user+tag@example.com":   true,
		"user_name@example.com":  true,
		"first.last@example.com": true,
		"a-b@example.com":        true,
		"-name@example.com":      false,
		"name-@example.com":      false,
		"a--b@example.com":       false,
		"a.-b@example.com":       false,
	}
	for addr, want := range cases {
		if got := Email(addr); got != want {
			t.Errorf("Email(%q) = %v, want %v", addr, got, want)
		}
	}
}

// TestIbanNewCountries guards BUG-06: IBANs for countries added to the SWIFT
// registry in 2023-2024 validate.
func TestIbanNewCountries(t *testing.T) {
	valid := []string{
		"FK4912345678901234",           // Falkland Islands, 18
		"MN181234567890123456",         // Mongolia, 20
		"NI31123456789012345678901234", // Nicaragua, 28
		"OM311234567890123456789",      // Oman, 23
	}
	for _, iban := range valid {
		if !Iban(iban) {
			t.Errorf("Iban(%q) = false, want true", iban)
		}
	}
}

// TestIbanRegistryMissingCountries guards the HN/YE gap: Honduras (len 28) and
// Yemen (len 30) are ISO 13616 registry members that were absent from
// ibanLenPatterns, so every valid Honduran/Yemeni IBAN was rejected. The
// official-registry example IBANs must validate, while a mutated check digit
// (same length) and a wrong length must still be rejected, confirming the
// mod-97 checksum is enforced and not just the length.
func TestIbanRegistryMissingCountries(t *testing.T) {
	cases := []struct {
		name string
		iban string
		want bool
	}{
		{"Honduras official example", "HN88CABF00000000000250005469", true},
		{"Yemen official example", "YE15CBYE0001018861234567891234", true},
		{"Honduras bad check digit", "HN89CABF00000000000250005469", false},
		{"Yemen bad check digit", "YE16CBYE0001018861234567891234", false},
		{"Honduras wrong length", "HN88CABF0000000000025000546", false},
		{"Yemen wrong length", "YE15CBYE00010188612345678912340", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Iban(c.iban); got != c.want {
				t.Errorf("Iban(%q) = %v, want %v", c.iban, got, c.want)
			}
		})
	}
}

// TestDomainPunycode guards BUG-07: an IDN A-label (punycode) TLD is accepted.
func TestDomainPunycode(t *testing.T) {
	if !Domain("example.xn--p1ai") {
		t.Error("Domain(example.xn--p1ai) = false, want true")
	}
	if Domain("example.123") {
		t.Error("Domain(example.123) = true, want false (numeric TLD)")
	}
}

// TestAlphaUnicodeLetters guards BUG-09: any Unicode letter counts for Alpha
// and Alnum.
func TestAlphaUnicodeLetters(t *testing.T) {
	if !Alpha("abcΔ") {
		t.Error("Alpha(abcΔ) = false, want true")
	}
	if !Alnum("abcΔ1") {
		t.Error("Alnum(abcΔ1) = false, want true")
	}
}

// TestTitleEdgeCases guards BUG-10: no-letter strings are not titlecased and a
// titlecase digraph is a valid word start.
func TestTitleEdgeCases(t *testing.T) {
	if Title("123") || Title("!!!") {
		t.Error("Title of a letterless string = true, want false")
	}
	if !Title("ǅungla") { // 'ǅ' is a titlecase letter (category Lt)
		t.Error("Title(ǅungla) = false, want true")
	}
}

// TestE164NoLeadingZero guards BUG-11: the country code cannot start with zero.
func TestE164NoLeadingZero(t *testing.T) {
	if E164("+0123456789") {
		t.Error("E164(+0123456789) = true, want false")
	}
	if !E164("+123456789") {
		t.Error("E164(+123456789) = false, want true")
	}
}

// TestSelectorKeywordAllowed guards BUG-12: a selector name coinciding with a
// CSS keyword is legal.
func TestSelectorKeywordAllowed(t *testing.T) {
	if !SelectorName("hover", true) || !SelectorName("color", true) {
		t.Error("SelectorName rejected a valid class name matching a CSS keyword")
	}
	// Strict still rejects a leading '.' or '#'.
	if SelectorName("#inherit", true) {
		t.Error("SelectorName(#inherit, strict) = true, want false")
	}
}

// TestRubySigilVariables guards BUG-13: sigilled Ruby names are valid even when
// the bare identifier is a keyword.
func TestRubySigilVariables(t *testing.T) {
	cases := []string{"@@class", "@def", "$if"}
	for _, name := range cases {
		ok, err := VariableNameFor(name, "ruby")
		if err != nil {
			t.Fatalf("VariableNameFor(%q, ruby) error: %v", name, err)
		}
		if !ok {
			t.Errorf("VariableNameFor(%q, ruby) = false, want true", name)
		}
	}
}
