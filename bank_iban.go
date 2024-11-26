package is

/* IBAN DATA
Albania*AL*false*28*false*true*AL35202111090000000001234567
Andorra*AD*true*24*false*true*AD1400080001001234567890
Azerbaijan*AZ*false*28*false*false*AZ77VTBA00000000001234567890
Bahrain*BH*false*22*false*false*BH02CITI00001077181611
Brazil*BR*false*29*false*true*BR1500000000000010932840814P2
Bulgaria*BG*true*22*false*true*BG18RZBB91550123456789
Cyprus*CY*true*28*false*true*CY21002001950000357001234567
Dominican Republic*DO*false*28*false*false*DO22ACAU00000000000123456789
Egypt*EG*false*29*false*false*EG800002000156789012345180002
Georgia*GE*false*22*false*false*GE60NB0000000123456789
Gibraltar*GI*true*23*false*true*GI56XAPO000001234567890
Greece*GR*true*27*false*true*GR9608100010000001234567890
Guatemala*GT*false*28*false*false*GT20AGRO00000000001234567890
Jordan*JO*false*30*false*true*JO71CBJO0000000000001234567890
Kazakhstan*KZ*false*20*false*false*KZ244350000012344567
Kuwait*KW*false*30*false*false*KW81CBKU0000000000001234560101
Latvia*LV*true*21*false*false*LV97HABA0012345678910
Lebanon*LB*false*28*false*false*LB92000700000000123123456123
Lithuania*LT*true*20*false*true*LT601010012345678901
Luxembourg*LU*true*20*false*false*LU120010001234567891
Malta*MT*true*31*false*true*MT31MALT01100000000000000000123
Mauritius*MU*false*30*false*true*MU43BOMM0101123456789101000MUR
Moldova*MD*false*24*false*false*MD21EX000000000001234567
Pakistan*PK*false*24*false*false*PK36SCBL0000001123456702
Palestine*PS*false*29*false*false*PS92PALS000000000400123456702
Qatar*QA*false*29*false*false*QA54QNBA000000000000693123456
Romania*RO*true*24*false*false*RO66BACX0000001234567890
Saint Lucia*LC*false*32*false*false*LC14BOSL123456789012345678901234
Sao Tome and Principe*ST*false*25*false*false*ST23000200000289355710148
Saudi Arabia*SA*false*24*false*false*SA4420000001234567891234
Turkey*TR*false*26*false*true*TR320010009999901234567890
United Arab Emirates*AE*false*23*false*false*AE460090000000123456789
Holy See*VA*true*22*false*false*VA59001123000012345678
Virgin Islands, British*VG*false*24*false*false*VG07ABVI0000000123456789
Ukraine*UA*false*29*false*true*UA903052992990004149123456789
Seychelles*SC*false*31*false*false*SC74MCBL01031234567890123456USD
Iraq*IQ*false*23*false*false*IQ20CBIQ861800101010500
Belarus*BY*false*28*false*false*BY86AKBB10100000002966000000
El Salvador*SV*false*28*false*false*SV43ACAT00000000000000123123
Libya*LY*false*25*false*false*LY38021001000000123456789
Sudan*SD*false*18*false*false*SD8811123456789012
Burundi*BI*false*27*false*false*BI4210000100010000332045181
Djibouti*DJ*false*27*false*false*DJ2110002010010409943020008
Somalia*SO*false*23*false*true*SO061000001123123456789
Austria*AT*true*20*true*true*AT483200000012345864
Belgium*BE*true*16*true*true*BE71096123456769
Bosnia and Herzegovina*BA*false*20*true*true*BA393385804800211234
Costa Rica*CR*false*22*true*false*CR23015108410026012345
Croatia*HR*true*21*true*false*HR1723600001101234565
Czech Republic*CZ*true*24*true*false*CZ5508000000001234567899
Faroe Islands*FO*false*18*true*true*FO9264600123456789
Greenland*GL*false*18*true*true*GL8964710123456789
Denmark*DK*true*18*true*true*DK9520000123456789
Estonia*EE*true*20*true*true*EE471000001020145685
Finland*FI*true*18*true*true*FI1410093000123458
France*FR*true*27*true*true*FR7630006000011234567890189
Germany*DE*true*22*true*true*DE75512108001245126199
Hungary*HU*true*28*true*true*HU93116000060000000012345676
Iceland*IS*true*26*true*true*IS750001121234563108962099
Ireland*IE*true*22*true*true*IE64IRCE92050112345678
Israel*IL*false*23*true*true*IL170108000000012612345
Italy*IT*true*27*true*true*IT60X0542811101000000123456
Kosovo*XK*false*20*true*true*XK051212012345678906
Liechtenstein*LI*true*21*true*true*LI7408806123456789012
North Macedonia*MK*false*19*true*false*MK07200002785123453
Mauritania*MR*false*27*true*true*MR1300020001010000123456753
Monaco*MC*true*27*true*true*MC5810096180790123456789085
Montenegro*ME*false*22*true*false*ME25505000012345678951
Netherlands*NL*true*18*true*true*NL02ABNA0123456789
Norway*NO*true*15*true*true*NO8330001234567
Poland*PL*true*28*true*true*PL10105000997603123456789123
Portugal*PT*true*25*true*true*PT50002700000001234567833
San Marino*SM*true*27*true*true*SM76P0854009812123456789123
Serbia*RS*false*22*true*false*RS35105008123123123173
Slovakia*SK*true*24*true*false*SK8975000000000012345671
Slovenia*SI*true*19*true*true*SI56192001234567892
Spain*ES*true*24*true*true*ES7921000813610123456789
Sweden*SE*true*24*true*true*SE7280000810340009783242
Switzerland*CH*true*21*true*true*CH5604835012345678009
Timor-Leste*TL*false*23*true*false*TL380010012345678910106
Tunisia*TN*false*24*true*true*TN5904018104004942712345
United Kingdom*GB*true*22*true*true*GB33BUKB20201555555555
Russia*RU*false*33*true*true*RU0204452560040702810412345678901
*/

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/goloop/g"
)

var ibanLenPatterns = map[string]int{
	"AL": 28, "AD": 24, "AZ": 28, "BH": 22, "BR": 29, "BG": 22, "CY": 28,
	"DO": 28, "EG": 29, "GE": 22, "GI": 23, "GR": 27, "GT": 28, "JO": 30,
	"KZ": 20, "KW": 30, "LV": 21, "LB": 28, "LT": 20, "LU": 20, "MT": 31,
	"MU": 30, "MD": 24, "PK": 24, "PS": 29, "QA": 29, "RO": 24, "LC": 32,
	"ST": 25, "SA": 24, "TR": 26, "AE": 23, "VA": 22, "VG": 24, "UA": 29,
	"SC": 31, "IQ": 23, "BY": 28, "SV": 28, "LY": 25, "SD": 18, "BI": 27,
	"DJ": 27, "SO": 23, "AT": 20, "BE": 16, "BA": 20, "CR": 22, "HR": 21,
	"CZ": 24, "FO": 18, "GL": 18, "DK": 18, "EE": 20, "FI": 18, "FR": 27,
	"DE": 22, "HU": 28, "IS": 26, "IE": 22, "IL": 23, "IT": 27, "XK": 20,
	"LI": 21, "MK": 19, "MR": 27, "MC": 27, "ME": 22, "NL": 18, "NO": 15,
	"PL": 28, "PT": 25, "SM": 27, "RS": 22, "SK": 24, "SI": 19, "ES": 24,
	"SE": 24, "CH": 21, "TL": 23, "TN": 24, "GB": 22, "RU": 33,
}

// letterToNumberCache is a cache for the mapLetterToNumber function.
// Is initialized in init().
var letterToNumberCache = make(map[rune]int)

// The mapLetterToNumber converts a letter to a number as per IBAN
// specifications. For example, 'A' is mapped to 10, 'B' to 11 etc.
func mapLetterToNumber(letter rune) int {
	if value, ok := letterToNumberCache[letter]; ok {
		return value
	}

	return 0
}

// CalculateIBANChecksum calculates the checksum of an IBAN as per the
// specifications. This is used to verify the validity of an IBAN number.
func CalculateIBANChecksum(iban string) *big.Int {
	mapped := ""

	// Cycle through each IBAN character.
	for _, letter := range iban {
		if letter >= 'A' && letter <= 'Z' {
			// Convert the letter to a number.
			mapped += fmt.Sprintf("%d", mapLetterToNumber(letter))
		} else if letter >= '0' && letter <= '9' {
			// Append digits as they are.
			mapped += string(letter)
		} else {
			// Invalid character found.
			return big.NewInt(-1)
		}
	}

	number := new(big.Int)
	_, ok := number.SetString(mapped, 10)
	if !ok {
		// Error converting string to number.
		return big.NewInt(-1)
	}

	// Return the remainder from division by 97.
	return new(big.Int).Mod(number, big.NewInt(97))
}

// Iban checks if a given IBAN (International Bank Account Number)
// has a valid format. If the input matches the pattern for the corresponding
// country and passes the checksum validation, the function returns true.
//
// By default, the function operates in non-strict mode. In non-strict mode:
//   - Spaces are removed from the input string.
//   - The input string is converted to uppercase.
//
// This allows the function to validate IBANs that may have spaces
// or lowercase letters.
//
// If the optional parameter `strict` is provided and set to true,
// the function operates in strict mode:
//   - The input string is not modified.
//   - Any deviations from the standard IBAN format (such as spaces,
//     lowercase letters, or special characters) will cause the function
//     to return false.
//
// Example usage:
//
//	// Non-strict mode (default):
//	is.Iban("GB82 WEST 1234 5698 7654 32")    // Returns: true
//	is.Iban("ua903052992990004149123456789")  // Returns: true
//
//	// Strict mode:
//	is.Iban("GB82WEST12345698765432", true)       // Returns: true
//	is.Iban("GB82 WEST 1234 5698 7654 32", true)  // Returns: false
//	is.Iban("ua903052992990004149123456789", true)// Returns: false (lowercase)
//
// Note:
//   - In strict mode, you should ensure that the input IBAN is properly
//     formatted:
//   - No spaces or special characters.
//   - All letters are uppercase.
//
// - You can preprocess the IBAN before validation if needed.
//
// For example, to validate an IBAN with spaces in strict mode:
//
//	iban := "GB82 WEST 1234 5698 7654 32"
//	iban = strings.ReplaceAll(iban, " ", "")
//	iban = strings.ToUpper(iban)
//	is.Iban(iban, true) // Returns: true
func Iban(iban string, strict ...bool) bool {
	if !g.All(strict...) {
		// Remove spaces and convert to uppercase.
		iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))
	} else {
		// In strict mode, check for invalid characters.
		for _, ch := range iban {
			if ch < '0' || ch > '9' && ch < 'A' || ch > 'Z' {
				return false
			}
		}
		// Convert to uppercase for checksum calculation.
		iban = strings.ToUpper(iban)
	}

	// IBAN must be at least 4 characters long to extract
	// country code and checksum.
	if len(iban) < 4 {
		return false
	}

	// Check if the length matches the expected length for the country.
	countryCode := iban[0:2]
	if l, ok := ibanLenPatterns[countryCode]; !ok || l != len(iban) {
		return false
	}

	// Move the first four characters to the end of the string.
	rearrangedIban := iban[4:] + iban[0:4]

	// Calculate the checksum
	remainder := CalculateIBANChecksum(rearrangedIban)
	if remainder.Sign() < 0 {
		// Invalid character encountered in checksum calculation.
		return false
	}
	return remainder.Cmp(big.NewInt(1)) == 0
}

// IBAN is a synonym for the Iban function. This naming approach adheres
// to Go's conventions for abbreviations in function names. The function
// takes an IBAN (International Bank Account Number) as a string parameter
// and returns a boolean value indicating whether the input IBAN is valid
// according to the defined pattern and checksum.
func IBAN(iban string, strict ...bool) bool {
	return Iban(iban, strict...)
}
