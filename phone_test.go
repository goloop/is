package is

import "testing"

// TestIMSI tests IMSI function.
func TestIMSI(t *testing.T) {
	tests := []struct {
		name  string
		imsi  string
		valid bool
	}{
		{
			name:  "Valid IMSI",
			imsi:  "310150123456789",
			valid: true,
		},
		{
			name:  "Valid IMSI with different MCC",
			imsi:  "460001234567890",
			valid: true,
		},
		{
			name:  "Invalid IMSI - length exceeds 15 digits",
			imsi:  "1234567890123456",
			valid: false,
		},
		{
			name:  "Invalid IMSI - invalid characters in MSIN",
			imsi:  "310150abc123456",
			valid: false,
		},
		{
			name:  "Invalid IMSI - empty string",
			imsi:  "",
			valid: false,
		},
		{
			name:  "Invalid MCC with non-numeric characters",
			imsi:  "31a150123456789",
			valid: false,
		},
		{
			name:  "Invalid MCC with less than three digits",
			imsi:  "31150123456789",
			valid: false,
		},
		{
			name:  "Invalid MNC with non-numeric characters",
			imsi:  "31015a123456789",
			valid: false,
		},
		{
			name:  "Valid IMSI",
			imsi:  "310150123456789",
			valid: true,
		},
		{
			name:  "Valid IMSI with three digits MNC",
			imsi:  "310150112345678", // Assume 1501 is a valid MNC
			valid: true,
		},
		{
			name:  "Invalid MNC with non-numeric characters",
			imsi:  "3101a2145678901", // Assume 1 is not a valid MNC
			valid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			valid := IMSI(test.imsi)
			if valid != test.valid {
				t.Errorf("IMSI(%s) = %v; want %v",
					test.imsi, valid, test.valid)
			}
		})
	}
}

// TestIMEI tests IMEI function.
func TestIMEI(t *testing.T) {
	tests := []struct {
		name  string
		inStr string
		inInt int64
		want  bool
	}{
		{
			name:  "Valid IMEI",
			inStr: "305114708429270",
			want:  true,
		},
		{
			name:  "Not valid IMEI",
			inStr: "532593572995861",
			want:  false,
		},
		{
			name:  "Not valid IMEI with spaces",
			inStr: "4901 5420 3237 518",
			want:  false,
		},
		{
			name:  "Not valid IMEI with dashes",
			inStr: "4901-5420-3237-518",
			want:  false,
		},
		{
			name:  "Valid IMEI int64",
			inInt: 305114708429270,
			want:  true,
		},
		{
			name:  "Not valid IMEI int64",
			inInt: 532593572995861,
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got bool

			if test.inStr != "" {
				got = IMEI(test.inStr)
			} else {
				got = IMEI(test.inInt)
			}

			if got != test.want {
				t.Errorf("IMEI(%q) = %v; want %v",
					test.inStr, got, test.want)
			}
		})
	}
}

// TestE164 tests E164 function.
func TestE164(t *testing.T) {
	tests := []struct {
		name     string
		phoneNum string
		want     bool
	}{
		{
			name:     "Valid E.164 phone number",
			phoneNum: "+123456789",
			want:     true,
		},
		{
			name:     "Valid E.164 phone number with leading zero",
			phoneNum: "+0123456789",
			want:     true,
		},
		{
			name:     "Invalid E.164 phone number - no digits after plus sign",
			phoneNum: "+",
			want:     false,
		},
		{
			name:     "Invalid E.164 phone number - non-digit character",
			phoneNum: "+1234567890a",
			want:     false,
		},
		{
			name:     "Invalid E.164 phone number - no plus sign",
			phoneNum: "1234567890",
			want:     false,
		},
		{
			name:     "Invalid E.164 phone number - empty string",
			phoneNum: "",
			want:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := E164(test.phoneNum)

			if got != test.want {
				t.Errorf("E164(%q) = %v; want %v",
					test.phoneNum, got, test.want)
			}
		})
	}
}

// TestPhone tests Phone function.
func TestPhone(t *testing.T) {
	tests := []struct {
		name     string
		phoneNum string
		want     bool
	}{
		{
			name:     "Valid phone number with country code in parentheses",
			phoneNum: "+3 (8096) 46 200 00",
			want:     true,
		},
		{
			name:     "Valid phone number without spaces",
			phoneNum: "+380961234567",
			want:     true,
		},
		{
			name:     "Invalid phone number - no plus sign",
			phoneNum: "123456789",
			want:     false,
		},
		{
			name:     "Invalid phone number - empty string",
			phoneNum: "",
			want:     false,
		},
		{
			name:     "Invalid phone number - alphabetic characters",
			phoneNum: "+3 (abc) 46 200 00",
			want:     false,
		},
		{
			name:     "Invalid phone number - special characters",
			phoneNum: "+3 (@#$) 46 200 00",
			want:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Phone(test.phoneNum)

			if got != test.want {
				t.Errorf("Phone(%q) = %v; want %v",
					test.phoneNum, got, test.want)
			}
		})
	}
}
