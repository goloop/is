package is

import (
	"testing"

	"github.com/goloop/g"
)

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
			t.Errorf("%s, expected validity of IBAN %s (strict: %v) to be %v, but got %v",
				test.name, test.iban, test.strict, test.valid, result)
		}
	}
}
