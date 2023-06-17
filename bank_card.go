package is

import (
	"regexp"
	"strings"
)

var (
	creditCardRegex      = regexp.MustCompile("^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$")
	visaRegex            = regexp.MustCompile("^4[0-9]{12}(?:[0-9]{3})?$")
	masterCardRegex      = regexp.MustCompile("^(5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12})$")
	americanExpressRegex = regexp.MustCompile("^3[47][0-9]{13}$")
	discoverRegex        = regexp.MustCompile("^6(?:011|5[0-9][0-9])[0-9]{12}$")
	dinersClubRegex      = regexp.MustCompile("^3(?:0[0-5]|[68][0-9])[0-9]{11}$")
	jcbRegex             = regexp.MustCompile("^(?:2131|1800|35\\d{3})\\d{11}$")
	maestroRegex         = regexp.MustCompile("^(?:5[0678]\\d\\d|6304|6390|67\\d\\d)\\d{8,15}$")
	unionPayRegex        = regexp.MustCompile("^62[0-5]\\d{13,16}$")
	mirRegex             = regexp.MustCompile("^220[0-4]\\d{12,15}$")
)

// BancCard checks if the string is a valid bank card number.
func BancCard(str string) bool {
	return cardChecker(str, creditCardRegex)
}

// CreditCard checks if the string is a valid credit card number.
func CreditCard(str string) bool {
	return cardChecker(str, creditCardRegex)
}

// Visa checks if the string is a valid VISA card number.
func Visa(n string) bool {
	return cardChecker(n, visaRegex)
}

// MasterCard checks if the string is a valid MasterCard card number.
func MasterCard(n string) bool {
	return cardChecker(n, masterCardRegex)
}

// AmericanExpress checks if the string is a valid American Express card number.
func AmericanExpress(n string) bool {
	return cardChecker(n, americanExpressRegex)
}

// Discover checks if the string is a valid Discover card number.
func Discover(n string) bool {
	return cardChecker(n, discoverRegex)
}

// DinersClub checks if the string is a valid Diners Club card number.
func DinersClub(n string) bool {
	return cardChecker(n, dinersClubRegex)
}

// JCB checks if the string is a valid JCB card number.
func JCB(n string) bool {
	return cardChecker(n, jcbRegex)
}

// Maestro checks if the string is a valid Maestro card number.
func Maestro(n string) bool {
	return cardChecker(n, maestroRegex)
}

// UnionPay checks if the string is a valid UnionPay card number.
func UnionPay(n string) bool {
	return cardChecker(n, unionPayRegex)
}

// Mir checks if the string is a valid Mir card number.
func Mir(n string) bool {
	return cardChecker(n, mirRegex)
}

// cardChecker checks if the string is a valid card number.
func cardChecker(n string, regex *regexp.Regexp) bool {
	// Check the regex
	if !regex.MatchString(n) {
		return false
	}

	// Normalize the number
	n = strings.ReplaceAll(n, "-", "")
	n = strings.ReplaceAll(n, " ", "")

	// Perform the Luhn Algorithm
	var sum int
	alt := false
	for i := len(n) - 1; i >= 0; i-- {
		if alt {
			temp := int(n[i]-'0') * 2
			if temp > 9 {
				temp -= 9
			}
			sum += temp
		} else {
			sum += int(n[i] - '0')
		}
		alt = !alt
	}
	return sum%10 == 0
}
