package is

import "testing"

func TestIban(t *testing.T) {
	tests := []struct {
		iban    string
		isValid bool
	}{
		{"GB82WEST12345698765432", true},
		{"GR1601101250000000012300695", true},
		{"GB82WEST1234569876543222222", false}, // more than necessary characters
		{"GB82WEST12", false},                  // fewer than necessary characters
		{"GB82TEST12345698765432", false},      // invalid character set
	}

	for _, test := range tests {
		if isValid := Iban(test.iban); isValid != test.isValid {
			t.Errorf("Expected validity of IBAN %s to be %v, but got %v",
				test.iban, test.isValid, isValid)
		}
	}
}
