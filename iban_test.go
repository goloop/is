package is

import (
	"testing"

	"github.com/goloop/g"
)

func TestIban(t *testing.T) {
	tests := []struct {
		name  string
		iban  string
		valid bool
	}{
		{"Albania", "AL35202111090000000001234567", true},
		{"Andorra", "AD1400080001001234567890", true},
		{"Azerbaijan", "AZ77VTBA00000000001234567890", true},
		{"Bahrain", "BH02CITI00001077181611", true},
		{"Brazil", "BR1500000000000010932840814P2", true},
		{"Bulgaria", "BG18RZBB91550123456789", true},
		{"Cyprus", "CY21002001950000357001234567", true},
		{"Dominican Republic", "DO22ACAU00000000000123456789", true},
		{"Egypt", "EG800002000156789012345180002", true},
		{"Georgia", "GE60NB0000000123456789", true},
		{"Gibraltar", "GI56XAPO000001234567890", true},
		{"Greece", "GR9608100010000001234567890", true},
		{"Guatemala", "GT20AGRO00000000001234567890", true},
		{"Jordan", "JO71CBJO0000000000001234567890", true},
		{"Kazakhstan", "KZ244350000012344567", true},
		{"Kuwait", "KW81CBKU0000000000001234560101", true},
		{"Latvia", "LV97HABA0012345678910", true},
		{"Lebanon", "LB92000700000000123123456123", true},
		{"Lithuania", "LT601010012345678901", true},
		{"Luxembourg", "LU120010001234567891", true},
		{"Malta", "MT31MALT01100000000000000000123", true},
		{"Mauritius", "MU43BOMM0101123456789101000MUR", true},
		{"Moldova", "MD21EX000000000001234567", true},
		{"Pakistan", "PK36SCBL0000001123456702", true},
		{"Palestine", "PS92PALS000000000400123456702", true},
		{"Qatar", "QA54QNBA000000000000693123456", true},
		{"Romania", "RO66BACX0000001234567890", true},
		{"Saint Lucia", "LC14BOSL123456789012345678901234", true},
		{"Sao Tome and Principe", "ST23000200000289355710148", true},
		{"Saudi Arabia", "SA4420000001234567891234", true},
		{"Turkey", "TR320010009999901234567890", true},
		{"United Arab Emirates", "AE460090000000123456789", true},
		{"Holy See", "VA59001123000012345678", true},
		{"Virgin Islands, British", "VG07ABVI0000000123456789", true},
		{"Ukraine", "UA903052992990004149123456789", true},
		{"Seychelles", "SC74MCBL01031234567890123456USD", true},
		{"Iraq", "IQ20CBIQ861800101010500", true},
		{"Belarus", "BY86AKBB10100000002966000000", true},
		{"El Salvador", "SV43ACAT00000000000000123123", true},
		{"Libya", "LY38021001000000123456789", true},
		{"Sudan", "SD8811123456789012", true},
		{"Burundi", "BI4210000100010000332045181", true},
		{"Djibouti", "DJ2110002010010409943020008", true},
		{"Somalia", "SO061000001123123456789", true},
		{"Austria", "AT483200000012345864", true},
		{"Belgium", "BE71096123456769", true},
		{"Bosnia and Herzegovina", "BA393385804800211234", true},
		{"Costa Rica", "CR23015108410026012345", true},
		{"Croatia", "HR1723600001101234565", true},
		{"Czech Republic", "CZ5508000000001234567899", true},
		{"Faroe Islands", "FO9264600123456789", true},
		{"Greenland", "GL8964710123456789", true},
		{"Denmark", "DK9520000123456789", true},
		{"Estonia", "EE471000001020145685", true},
		{"Finland", "FI1410093000123458", true},
		{"France", "FR7630006000011234567890189", true},
		{"Germany", "DE75512108001245126199", true},
		{"Hungary", "HU93116000060000000012345676", true},
		{"Iceland", "IS750001121234563108962099", true},
		{"Ireland", "IE64IRCE92050112345678", true},
		{"Israel", "IL170108000000012612345", true},
		{"Italy", "IT60X0542811101000000123456", true},
		{"Kosovo", "XK051212012345678906", true},
		{"Liechtenstein", "LI7408806123456789012", true},
		{"North Macedonia", "MK07200002785123453", true},
		{"Mauritania", "MR1300020001010000123456753", true},
		{"Monaco", "MC5810096180790123456789085", true},
		{"Montenegro", "ME25505000012345678951", true},
		{"Netherlands", "NL02ABNA0123456789", true},
		{"Norway", "NO8330001234567", true},
		{"Poland", "PL10105000997603123456789123", true},
		{"Portugal", "PT50002700000001234567833", true},
		{"San Marino", "SM76P0854009812123456789123", true},
		{"Serbia", "RS35105008123123123173", true},
		{"Slovakia", "SK8975000000000012345671", true},
		{"Slovenia", "SI56192001234567892", true},
		{"Spain", "ES7921000813610123456789", true},
		{"Sweden", "SE7280000810340009783242", true},
		{"Switzerland", "CH5604835012345678009", true},
		{"Timor-Leste", "TL380010012345678910106", true},
		{"Tunisia", "TN5904018104004942712345", true},
		{"United Kingdom", "GB33BUKB20201555555555", true},
		{"Russia", "RU0204452560040702810412345678901", true},

		{"Clean with g.Weed", g.Weed("GB82 WEST 1234 5698 7654 32"), true},

		{"With tabs", "GB82\tWEST 1234 5698 7654 32", false},
		{"More than necessary chars", "GB82WEST1234569876543222222", false},
		{"Fewer than necessary chars", "GB82WEST12", false},
		{"Invalid character set", "GB82TEST12345698765432", false},
	}

	for _, test := range tests {
		if v := IBAN(test.iban); v != test.valid {
			t.Errorf("%s, expected validity of IBAN %s to be %v, but got %v",
				test.name, test.iban, test.valid, v)
		}
	}
}
