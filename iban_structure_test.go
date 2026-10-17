package is

import "testing"

func TestIBANRejectsAlphabeticCheckDigits(t *testing.T) {
	for a := byte('A'); a <= 'Z'; a++ {
		for b := byte('A'); b <= 'Z'; b++ {
			iban := "GB" + string([]byte{a, b}) + "BUKB20201555555555"
			if IBAN(iban) {
				t.Fatalf("accepted alphabetic check digits: %s", iban)
			}
		}
	}
}
