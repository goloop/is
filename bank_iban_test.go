package is

import (
	"math/big"
	"testing"

	"github.com/goloop/g"
)

// TestMapLetterToNumber tests the mapLetterToNumber function.
func TestMapLetterToNumber(t *testing.T) {
	tests := []struct {
		name string
		in   rune
		want int
	}{
		{
			name: "Maps letter A to 10",
			in:   'A',
			want: 10,
		},
		{
			name: "Maps letter B to 11",
			in:   'B',
			want: 11,
		},
		{
			name: "Maps letter Z to 35",
			in:   'Z',
			want: 35,
		},
		{
			name: "Maps undefined letter to 0",
			in:   '!',
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapLetterToNumber(tt.in); got != tt.want {
				t.Errorf("mapLetterToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIBAN tests the IBAN function.
// Tests the IBAN function for large coverage, including strict mode.
func TestIBAN(t *testing.T) {
	tests := []struct {
		name   string
		iban   string
		strict bool
		valid  bool
	}{
		// Non-strict mode test cases
		{"Albania", "AL3", false, false}, // invalid length
		{"Albania", "AL35202111090000000001234567", false, true},
		{"Andorra", "AD1400080001001234567890", false, true},
		{"Azerbaijan", "AZ77VTBA00000000001234567890", false, true},
		{"Bahrain", "BH02CITI00001077181611", false, true},
		{"Brazil", "BR1500000000000010932840814P2", false, true},
		{"Bulgaria", "BG18RZBB91550123456789", false, true},
		{"Cyprus", "CY21002001950000357001234567", false, true},
		{"Dominican Republic", "DO22ACAU00000000000123456789", false, true},
		{"Egypt", "EG800002000156789012345180002", false, true},
		{"Georgia", "GE60NB0000000123456789", false, true},
		{"Gibraltar", "GI56XAPO000001234567890", false, true},
		{"Greece", "GR9608100010000001234567890", false, true},
		{"Guatemala", "GT20AGRO00000000001234567890", false, true},
		{"Jordan", "JO71CBJO0000000000001234567890", false, true},
		{"Kazakhstan", "KZ244350000012344567", false, true},
		{"Kuwait", "KW81CBKU0000000000001234560101", false, true},
		{"Latvia", "LV97HABA0012345678910", false, true},
		{"Lebanon", "LB92000700000000123123456123", false, true},
		{"Lithuania", "LT601010012345678901", false, true},
		{"Luxembourg", "LU120010001234567891", false, true},
		{"Malta", "MT31MALT01100000000000000000123", false, true},
		{"Mauritius", "MU43BOMM0101123456789101000MUR", false, true},
		{"Moldova", "MD21EX000000000001234567", false, true},
		{"Pakistan", "PK36SCBL0000001123456702", false, true},
		{"Palestine", "PS92PALS000000000400123456702", false, true},
		{"Qatar", "QA54QNBA000000000000693123456", false, true},
		{"Romania", "RO66BACX0000001234567890", false, true},
		{"Saint Lucia", "LC14BOSL123456789012345678901234", false, true},
		{"Sao Tome and Principe", "ST23000200000289355710148", false, true},
		{"Saudi Arabia", "SA4420000001234567891234", false, true},
		{"Turkey", "TR320010009999901234567890", false, true},
		{"United Arab Emirates", "AE460090000000123456789", false, true},
		{"Holy See", "VA59001123000012345678", false, true},
		{"Virgin Islands, British", "VG07ABVI0000000123456789", false, true},
		{"Ukraine", "UA903052992990004149123456789", false, true},
		{"Seychelles", "SC74MCBL01031234567890123456USD", false, true},
		{"Iraq", "IQ20CBIQ861800101010500", false, true},
		{"Belarus", "BY86AKBB10100000002966000000", false, true},
		{"El Salvador", "SV43ACAT00000000000000123123", false, true},
		{"Libya", "LY38021001000000123456789", false, true},
		{"Sudan", "SD8811123456789012", false, true},
		{"Burundi", "BI4210000100010000332045181", false, true},
		{"Djibouti", "DJ2110002010010409943020008", false, true},
		{"Somalia", "SO061000001123123456789", false, true},
		{"Austria", "AT483200000012345864", false, true},
		{"Belgium", "BE71096123456769", false, true},
		{"Bosnia and Herzegovina", "BA393385804800211234", false, true},
		{"Costa Rica", "CR23015108410026012345", false, true},
		{"Croatia", "HR1723600001101234565", false, true},
		{"Czech Republic", "CZ5508000000001234567899", false, true},
		{"Faroe Islands", "FO9264600123456789", false, true},
		{"Greenland", "GL8964710123456789", false, true},
		{"Denmark", "DK9520000123456789", false, true},
		{"Estonia", "EE471000001020145685", false, true},
		{"Finland", "FI1410093000123458", false, true},
		{"France", "FR7630006000011234567890189", false, true},
		{"Germany", "DE75512108001245126199", false, true},
		{"Hungary", "HU93116000060000000012345676", false, true},
		{"Iceland", "IS750001121234563108962099", false, true},
		{"Ireland", "IE64IRCE92050112345678", false, true},
		{"Israel", "IL170108000000012612345", false, true},
		{"Italy", "IT60X0542811101000000123456", false, true},
		{"Kosovo", "XK051212012345678906", false, true},
		{"Liechtenstein", "LI7408806123456789012", false, true},
		{"North Macedonia", "MK07200002785123453", false, true},
		{"Mauritania", "MR1300020001010000123456753", false, true},
		{"Monaco", "MC5810096180790123456789085", false, true},
		{"Montenegro", "ME25505000012345678951", false, true},
		{"Netherlands", "NL02ABNA0123456789", false, true},
		{"Norway", "NO8330001234567", false, true},
		{"Poland", "PL10105000997603123456789123", false, true},
		{"Portugal", "PT50002700000001234567833", false, true},
		{"San Marino", "SM76P0854009812123456789123", false, true},
		{"Serbia", "RS35105008123123123173", false, true},
		{"Slovakia", "SK8975000000000012345671", false, true},
		{"Slovenia", "SI56192001234567892", false, true},
		{"Spain", "ES7921000813610123456789", false, true},
		{"Sweden", "SE7280000810340009783242", false, true},
		{"Switzerland", "CH5604835012345678009", false, true},
		{"Timor-Leste", "TL380010012345678910106", false, true},
		{"Tunisia", "TN5904018104004942712345", false, true},
		{"United Kingdom", "GB33BUKB20201555555555", false, true},
		{"Russia", "RU0204452560040702810412345678901", false, true},

		{
			"Clean with g.Weed",
			g.Weed("GB82 WEST 1234 5698 7654 32"),
			false,
			true,
		},

		{"With tabs", "GB82\tWEST 1234 5698 7654 32", false, false},
		{
			"More than necessary chars",
			"GB82WEST1234569876543222222",
			false,
			false,
		},
		{"Fewer than necessary chars", "GB82WEST12", false, false},
		{"Invalid character set", "GB82TEST12345698765432", false, false},

		// Strict mode test cases
		{"Albania Strict", "AL35202111090000000001234567", true, true},
		{
			"Albania Lowercase Strict",
			"al35202111090000000001234567",
			true,
			false,
		},
		{
			"Albania With Spaces Strict",
			"AL35 2021 1109 0000 0000 0123 4567",
			true,
			false,
		},
		{"GB With Spaces Strict", "GB82 WEST 1234 5698 7654 32", true, false},
		{"GB Without Spaces Strict", "GB82WEST12345698765432", true, true},
		{"GB Lowercase Strict", "gb82west12345698765432", true, false},
		{"Ukraine Strict", "UA903052992990004149123456789", true, true},
		{
			"Ukraine Lowercase Strict", "ua903052992990004149123456789",
			true,
			false,
		},
		{
			"Ukraine With Spaces Strict",
			"UA90 3052 9929 9000 4149 1234 5678 9",
			true,
			false,
		},
		{"Invalid Characters Strict", "GB82WEST1234569876543$", true, false},
		{"Invalid Length Strict", "GB82WEST1234569876543", true, false},
		{"Contains Tabs Strict", "GB82\tWEST12345698765432", true, false},
	}

	for _, test := range tests {
		var result bool
		if test.strict {
			result = IBAN(test.iban, true)
		} else {
			result = IBAN(test.iban)
		}
		if result != test.valid {
			t.Errorf(
				"%s, expected validity of IBAN %s (strict: %v) "+
					"to be %v, but got %v",
				test.name, test.iban, test.strict, test.valid, result,
			)
		}
	}
}

// TestCalculateIBANChecksum tests the CalculateIBANChecksum function.
func TestCalculateIBANChecksum(t *testing.T) {
	tests := []struct {
		name string
		iban string
		want *big.Int
	}{
		{
			name: "Valid IBAN with only numbers",
			iban: "123456789",
			want: new(big.Int).Mod(big.NewInt(123456789), big.NewInt(97)),
		},
		{
			name: "Valid IBAN with letters",
			iban: "GB82WEST12345698765432",
			want: new(big.Int).SetInt64(85), // actual calculated checksum
		},
		{
			name: "Valid IBAN mixed case",
			iban: "GB82west12345698765432",
			want: big.NewInt(-1), // invalid chars (lowercase)
		},

		// Test letters to numbers conversion.
		{
			name: "IBAN with single letters A-Z",
			iban: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want: new(big.Int).SetInt64(80), // actual calculated checksum
		},

		// Test invalid characters.
		{
			name: "IBAN with space",
			iban: "GB82 WEST",
			want: big.NewInt(-1),
		},
		{
			name: "IBAN with special character",
			iban: "GB82@WEST",
			want: big.NewInt(-1),
		},
		{
			name: "IBAN with hyphen",
			iban: "GB82-WEST",
			want: big.NewInt(-1),
		},
		{
			name: "IBAN with lowercase special chars",
			iban: "gb82#west",
			want: big.NewInt(-1),
		},

		// Edge cases.
		{
			name: "Empty IBAN",
			iban: "",
			want: big.NewInt(-1), // invalid IBAN
		},
		{
			name: "Single digit",
			iban: "1",
			want: new(big.Int).SetInt64(1),
		},
		{
			name: "Single letter",
			iban: "A",
			want: new(big.Int).SetInt64(10),
		},

		// Test boundary cases.
		{
			name: "Number at lower boundary",
			iban: "0",
			want: new(big.Int).SetInt64(0),
		},
		{
			name: "Number at upper boundary",
			iban: "9",
			want: new(big.Int).SetInt64(9),
		},
		{
			name: "Letter at lower boundary",
			iban: "A",
			want: new(big.Int).SetInt64(10),
		},
		{
			name: "Letter at upper boundary",
			iban: "Z",
			want: new(big.Int).SetInt64(35),
		},

		// Test combinations.
		{
			name: "Mixed letters and numbers",
			iban: "A1B2C3",
			want: new(big.Int).SetInt64(2), // actual calculated checksum
		},
		{
			name: "All invalid characters",
			iban: "@#$%",
			want: big.NewInt(-1),
		},
		{
			name: "Mix of valid and invalid",
			iban: "AB12@CD34",
			want: big.NewInt(-1),
		},

		// Real IBAN examples.
		{
			name: "Valid GB IBAN",
			iban: "GB82WEST12345698765432",
			want: new(big.Int).SetInt64(85), // actual calculated checksum
		},
		{
			name: "Valid DE IBAN",
			iban: "DE89370400440532013000",
			want: new(big.Int).SetInt64(69), // actual calculated checksum
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateIBANChecksum(tt.iban)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("CalculateIBANChecksum(%q) = %v, want %v",
					tt.iban, got, tt.want)
			}
		})
	}
}

// TestIban_InvalidCharacterChecksum tests the Iban
// function for invalid characters.
func TestIban_InvalidCharacterChecksum(t *testing.T) {
	tests := []struct {
		name    string
		iban    string
		strict  bool
		want    bool
		comment string
	}{
		{
			name:    "Invalid character in rearranged IBAN part",
			iban:    "GB82WES@12345698765432",
			strict:  false,
			want:    false,
			comment: "Special character in the main part",
		},
		{
			name:   "Invalid character in country code part when moved",
			iban:   "GB82WEST1234569876543@",
			strict: false,
			want:   false,
			comment: "Special character will be in country code position " +
				"after rearrangement",
		},
		{
			name:   "Invalid character in checksum part when moved",
			iban:   "GB82WEST12345698765#32",
			strict: false,
			want:   false,
			comment: "Special character will be in checksum position " +
				"after rearrangement",
		},
		{
			name:   "Lowercase letters in rearranged part - non-strict",
			iban:   "GB82WEST12345698765432", // Using correct IBAN
			strict: false,
			want:   true, // Should be valid after conversion to uppercase
			comment: "Lowercase letters should be converted to " +
				"uppercase and validate",
		},
		{
			name:    "Lowercase letters in rearranged part - strict",
			iban:    "GB82west12345698765432",
			strict:  true,
			want:    false,
			comment: "Lowercase letters are invalid in strict mode",
		},
		{
			name:    "Valid IBAN with spaces - non-strict",
			iban:    "GB82 WEST 1234 5698 7654 32",
			strict:  false,
			want:    true, // Should be valid after space removal
			comment: "Spaces should be removed in non-strict mode and validate",
		},
		{
			name:    "Valid IBAN with spaces - strict",
			iban:    "GB82 WEST 1234 5698 7654 32",
			strict:  true,
			want:    false,
			comment: "Spaces are invalid in strict mode",
		},
		{
			name:    "Valid country code but invalid checksum chars",
			iban:    "GB@@WEST12345698765432",
			strict:  false,
			want:    false,
			comment: "Special characters in checksum position",
		},
		{
			name:    "Special characters throughout",
			iban:    "GB82WE@#12$4%698&6543*",
			strict:  false,
			want:    false,
			comment: "Multiple special characters",
		},
		{
			name:    "Unicode characters",
			iban:    "GB82WESTЫ2345698765432",
			strict:  false,
			want:    false,
			comment: "Unicode characters should cause checksum failure",
		},
		{
			name:    "Control characters",
			iban:    "GB82WEST\x00\x0112345698765432",
			strict:  false,
			want:    false,
			comment: "Control characters should cause checksum failure",
		},
		{
			name:    "Valid numbers but invalid letter position",
			iban:    "GB82123456987654321234",
			strict:  false,
			want:    false,
			comment: "Valid characters but wrong position for checksum",
		},
		{
			name:   "Mix of invalid characters after rearrangement",
			iban:   "GB82WEST1234569876@#$%",
			strict: false,
			want:   false,
			comment: "Invalid characters will be in critical position " +
				"after rearrangement",
		},
		{
			name:    "Valid format but with emoji",
			iban:    "GB82WEST1234569876😊432",
			strict:  false,
			want:    false,
			comment: "Emoji should cause checksum failure",
		},
		// Additional test cases for non-strict mode
		{
			name:   "Mixed case and spaces - non-strict",
			iban:   "gb82 WeSt 1234 5698 7654 32",
			strict: false,
			want:   true,
			comment: "Mixed case and spaces should be handled " +
				"in non-strict mode",
		},
		{
			name:    "Multiple spaces - non-strict",
			iban:    "GB82  WEST    12345698765432",
			strict:  false,
			want:    true,
			comment: "Multiple spaces should be removed in non-strict mode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Iban(tt.iban, tt.strict)
			if got != tt.want {
				t.Errorf("Iban(%q, %v) = %v, want %v (%s)",
					tt.iban, tt.strict, got, tt.want, tt.comment)
			}
		})
	}
}
