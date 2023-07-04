package is

import (
	"regexp"

	"github.com/goloop/g"
)

var (
	// The anyCreditCard is a regular expression for matching any credit card.
	anyCreditCard = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`)

	// Visa is a regular expression for matching Visa bank card.
	Visa = regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`)

	// MasterCard is a regular expression for matching MasterCard bank card.
	MasterCard = regexp.MustCompile(`^5[1-5][0-9]{14}$`)

	// AmericanExpress is a regular expression for matching
	// American Express bank card.
	AmericanExpress = regexp.MustCompile(`^3[47][0-9]{13}$`)

	// DiscoverCard is a regular expression for matching
	// Discover Card bank card.
	DiscoverCard = regexp.MustCompile(`^6(?:011\d{12}|5\d{14}|4[4-9]\d{13}|22(?:1(?:2[6-9]|[3-9]\d)|[2-8]\d{2}|9(?:[01]\d|2[0-5]))\d{10})$`)

	// DCI is a regular expression for matching DCI bank card.
	DCI = regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`)

	// UnionPay is a regular expression for matching UnionPay bank card.
	UnionPay = regexp.MustCompile("^62[0-5]\\d{13,16}$")

	// JCB is a regular expression for matching JCB bank card.
	JCB = regexp.MustCompile(`^(?:2131|1800|35[0-9]{3})[0-9]{11}$`)

	// Argencard is a regular expression for matching Argencard bank card.
	Argencard = regexp.MustCompile(`^501105\d{10}$`)

	// Cabal is a regular expression for matching Cabal bank card.
	Cabal = regexp.MustCompile(`^6042(0[1-9]|10|1[1-9])\d{6}$`)

	// Cencosud is a regular expression for matching Cencosud bank card.
	Cencosud = regexp.MustCompile(`^603493\d{10}$`)

	// ChinaUnionPay is a regular expression for matching
	// China UnionPay bank card.
	ChinaUnionPay = regexp.MustCompile(`^62[0-9]{14,17}$`)

	// DinersClubCarteBlanche is a regular expression for matching
	// Diners Club Carte Blanche bank card.
	DinersClubCarteBlanche = regexp.MustCompile(`^30[0-5][0-9]{11}$`)

	// DinersClubInternational is a regular expression for matching
	// Diners Club International bank card.
	DinersClubInternational = regexp.MustCompile(`^36[0-9]{12}$`)

	// DinersClubUSAndCanada is a regular expression for matching
	// Diners Club US & Canada bank card.
	DinersClubUSAndCanada = regexp.MustCompile(`^5[45][0-9]{14}$`)

	// DinersClub is a regular expression for matching Diners Club bank card.
	DinersClub = regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`)

	// InstaPayment is a regular expression for matching
	// InstaPayment bank card.
	InstaPayment = regexp.MustCompile(`^63[7-9][0-9]{13}$`)

	// Laser is a regular expression for matching Laser bank card.
	Laser = regexp.MustCompile(`^(6304|670[69]|6771)[0-9]{12,15}$`)

	// Maestro is a regular expression for matching Maestro bank card.
	Maestro = regexp.MustCompile(`^(5018|5020|5038|6304|6759|676[1-3])[0-9]{8,15}$`)

	// VisaElectron is a regular expression for matching
	// Visa Electron bank card.
	VisaElectron = regexp.MustCompile(`^(4026|417500|4508|4844|491[37])[0-9]{12}$`)

	// Dankort is a regular expression for matching Dankort bank card.
	Dankort = regexp.MustCompile(`^(5019)[0-9]{12}$`)

	// RuPay is a regular expression for matching RuPay bank card.
	RuPay = regexp.MustCompile(`^(508[5-9][0-9]{1}|60698|60699|607[0-8][0-9]{1}|6079[0-7]|60798[0-4]|608[0-4][0-9]{1}|608500)[0-9]{6,9}$`)

	// InterPayment is a regular expression for matching
	// InterPayment bank card.
	InterPayment = regexp.MustCompile(`^636[0-9]{12,15}$`)

	// Troy is a regular expression for matching Troy bank card.
	Troy = regexp.MustCompile(`^9792[0-9]{12}$`)

	// MIR is a regular expression for matching MIR bank card.
	MIR = regexp.MustCompile(`^220[0-9]{13}$`)

	// UATP is a regular expression for matching UATP bank card.
	UATP = regexp.MustCompile(`^1[0-9]{14}$`)

	// Hipercard is a regular expression for matching Hipercard bank card.
	Hipercard = regexp.MustCompile(`^(606282\d{10}(\d{3})?)|(3841\d{02}\d{10})$`)

	// Naranja is a regular expression for matching Naranja bank card.
	Naranja = regexp.MustCompile(`^589562\d{10}$`)

	// TarjetaShopping is a regular expression for matching
	// Tarjeta Shopping bank card.
	TarjetaShopping = regexp.MustCompile(`^603488\d{10}$`)

	// ELO is a regular expression for matching ELO bank card.
	ELO = regexp.MustCompile(`^(401178|401179|431274|438935|451416|457393|457631|457632|504175|627780|636297|636368|636369)\d{10}$`)
)

// BankCard validates a bank card number based on provided card types.
// If no type is given, it checks against any type.
//
// BankCard uses Luhn algorithm to validate the card number
// and tests as Credit Card and Debit Card.
//
// Example usage:
//
//	is.BankCard("4111111111111111")
//	// Output: true
//	// As it's a valid card number, checked against any type.
//
//	is.BankCard("4111111111111111", is.MasterCard)
//	// Output: false
//	// It's a valid card but not of type MasterCard.
//
//	is.BankCard("4111111111111111", is.Visa)
//	// Output: true
//	// It's a valid Visa card.
//
//	is.BankCard("4111111111111111", is.Visa, is.MasterCard)
//	// Output: true
//	// It's a valid card of either Visa or MasterCard type.
//
//	is.BankCard("1234567812345678")
//	// Output: false
//	// Not a valid card number.
func BankCard(str string, kinds ...*regexp.Regexp) bool {
	if len(kinds) == 0 {
		return cardChecker(str, anyCreditCard)
	}

	for _, kind := range kinds {
		if cardChecker(str, kind) {
			return true
		}
	}

	return false
}

// The cardChecker checks whether a string matches a given card regex and
// validates its checksum according to the Luhn algorithm.
// It's used internally in the BankCard and CreditCard functions.
// This function isn't exported, so it can't be used outside the 'is' package.
//
// Here's how it works:
//  1. It ignores spaces and hyphens in the input string.
//  2. It calculates the Luhn checksum of the card number.
//  3. It checks that the sum of the digits is a multiple of 10.
//  4. It validates the card number against the provided regular expression.
//
// It returns true if all checks pass, and false otherwise.
func cardChecker(n string, regex *regexp.Regexp) bool {
	var sum int64

	const (
		zero = 48
		ten  = 57
	)

	// We remove only the delimiters that can be in the card:
	// a space and a dash (-) symbol.
	clean := g.Weed(n, " -")

	// Luhn algorithm.
	p := len(clean) % 2
	for i, d := range clean {
		if d < zero || d > ten {
			return false
		}

		d = d - zero

		if i%2 == p {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += int64(d)
	}

	if sum%10 != 0 {
		return false
	}

	// Additionally check for a regular expression, because different
	// banks can have different initial codes.
	if !regex.MatchString(clean) {
		return false
	}

	return true
}
