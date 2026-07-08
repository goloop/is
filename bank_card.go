package is

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// CardKind identifies a bank card brand accepted by BankCard.
//
// In contrast to exporting the underlying regular expressions, an opaque
// kind keeps the matching patterns immutable: callers select a brand by
// constant and cannot replace a global regex and silently break validation
// for the whole process.
//
// The zero value is not a valid kind. Use the exported constants below.
type CardKind int

const (
	// Visa matches Visa cards.
	Visa CardKind = iota + 1

	// MasterCard matches MasterCard cards.
	MasterCard

	// AmericanExpress matches American Express cards.
	AmericanExpress

	// DiscoverCard matches Discover cards.
	DiscoverCard

	// DCI matches Diners Club International cards.
	DCI

	// UnionPay matches UnionPay cards.
	UnionPay

	// JCB matches JCB cards.
	JCB

	// Argencard matches Argencard cards.
	Argencard

	// Cabal matches Cabal cards.
	Cabal

	// Cencosud matches Cencosud cards.
	Cencosud

	// ChinaUnionPay matches China UnionPay cards.
	ChinaUnionPay

	// DinersClubCarteBlanche matches Diners Club Carte Blanche cards.
	DinersClubCarteBlanche

	// DinersClubInternational matches Diners Club International cards.
	DinersClubInternational

	// DinersClubUSAndCanada matches Diners Club US & Canada cards.
	DinersClubUSAndCanada

	// DinersClub matches Diners Club cards.
	DinersClub

	// InstaPayment matches InstaPayment cards.
	InstaPayment

	// Laser matches Laser cards.
	Laser

	// Maestro matches Maestro cards.
	Maestro

	// VisaElectron matches Visa Electron cards.
	VisaElectron

	// Dankort matches Dankort cards.
	Dankort

	// RuPay matches RuPay cards.
	RuPay

	// InterPayment matches InterPayment cards.
	InterPayment

	// Troy matches Troy cards.
	Troy

	// MIR matches MIR cards.
	MIR

	// UATP matches UATP cards.
	UATP

	// Hipercard matches Hipercard cards.
	Hipercard

	// Naranja matches Naranja cards.
	Naranja

	// TarjetaShopping matches Tarjeta Shopping cards.
	TarjetaShopping

	// ELO matches Elo cards.
	ELO
)

// cardKindNames maps each card kind to a human-readable name, used by
// the String method (handy in test failures and logs).
var cardKindNames = map[CardKind]string{
	Visa:                    "Visa",
	MasterCard:              "MasterCard",
	AmericanExpress:         "AmericanExpress",
	DiscoverCard:            "DiscoverCard",
	DCI:                     "DCI",
	UnionPay:                "UnionPay",
	JCB:                     "JCB",
	Argencard:               "Argencard",
	Cabal:                   "Cabal",
	Cencosud:                "Cencosud",
	ChinaUnionPay:           "ChinaUnionPay",
	DinersClubCarteBlanche:  "DinersClubCarteBlanche",
	DinersClubInternational: "DinersClubInternational",
	DinersClubUSAndCanada:   "DinersClubUSAndCanada",
	DinersClub:              "DinersClub",
	InstaPayment:            "InstaPayment",
	Laser:                   "Laser",
	Maestro:                 "Maestro",
	VisaElectron:            "VisaElectron",
	Dankort:                 "Dankort",
	RuPay:                   "RuPay",
	InterPayment:            "InterPayment",
	Troy:                    "Troy",
	MIR:                     "MIR",
	UATP:                    "UATP",
	Hipercard:               "Hipercard",
	Naranja:                 "Naranja",
	TarjetaShopping:         "TarjetaShopping",
	ELO:                     "ELO",
}

// String returns the human-readable name of the card kind, or "CardKind(n)"
// for an unknown value.
func (k CardKind) String() string {
	if name, ok := cardKindNames[k]; ok {
		return name
	}
	return "CardKind(" + strconv.Itoa(int(k)) + ")"
}

// anyCreditCard returns the pattern used by BankCard when no specific kind is
// requested: the union of every brand pattern in cardPatterns. Building it from
// cardPatterns keeps a single source of truth, so a new or corrected brand is
// automatically part of the "any brand" check. It is built once, lazily.
var anyCreditCard = sync.OnceValue(func() *regexp.Regexp {
	// Sort the brand bodies for a deterministic combined pattern (map
	// iteration order is randomized). Order does not affect matching.
	bodies := make([]string, 0, len(cardPatterns))
	for _, re := range cardPatterns {
		body := strings.TrimSuffix(strings.TrimPrefix(re.String(), "^"), "$")
		bodies = append(bodies, "(?:"+body+")")
	}
	sort.Strings(bodies)
	return regexp.MustCompile("^(?:" + strings.Join(bodies, "|") + ")$")
})

// cardPatterns holds the brand-specific regular expression for every
// CardKind. The patterns are unexported and immutable from the outside.
var cardPatterns = map[CardKind]*regexp.Regexp{
	Visa:                    regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
	MasterCard:              regexp.MustCompile(`^5[1-5][0-9]{14}$`),
	AmericanExpress:         regexp.MustCompile(`^3[47][0-9]{13}$`),
	DiscoverCard:            regexp.MustCompile(`^6(?:011\d{12}|5\d{14}|4[4-9]\d{13}|22(?:1(?:2[6-9]|[3-9]\d)|[2-8]\d{2}|9(?:[01]\d|2[0-5]))\d{10})$`),
	DCI:                     regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`),
	UnionPay:                regexp.MustCompile(`^62[0-5]\d{13,16}$`),
	JCB:                     regexp.MustCompile(`^(?:2131|1800|35[0-9]{3})[0-9]{11}$`),
	Argencard:               regexp.MustCompile(`^501105\d{10}$`),
	Cabal:                   regexp.MustCompile(`^6042(0[1-9]|1[0-9])\d{10}$`),
	Cencosud:                regexp.MustCompile(`^603493\d{10}$`),
	ChinaUnionPay:           regexp.MustCompile(`^62[0-9]{14,17}$`),
	DinersClubCarteBlanche:  regexp.MustCompile(`^30[0-5][0-9]{11}$`),
	DinersClubInternational: regexp.MustCompile(`^36[0-9]{12}$`),
	DinersClubUSAndCanada:   regexp.MustCompile(`^5[45][0-9]{14}$`),
	DinersClub:              regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`),
	InstaPayment:            regexp.MustCompile(`^63[7-9][0-9]{13}$`),
	Laser:                   regexp.MustCompile(`^(6304|670[69]|6771)[0-9]{12,15}$`),
	Maestro:                 regexp.MustCompile(`^(5018|5020|5038|6304|6759|676[1-3])[0-9]{8,15}$`),
	VisaElectron:            regexp.MustCompile(`^(?:(?:4026|4508|4844|491[37])\d{12}|417500\d{10})$`),
	Dankort:                 regexp.MustCompile(`^(5019)[0-9]{12}$`),
	RuPay:                   regexp.MustCompile(`^(?:(?:508[5-9]\d|60698|60699|607[0-8]\d|6079[0-7]|608[0-4]\d)\d{11}|(?:60798[0-4]|608500)\d{10})$`),
	InterPayment:            regexp.MustCompile(`^636[0-9]{12,15}$`),
	Troy:                    regexp.MustCompile(`^9792[0-9]{12}$`),
	MIR:                     regexp.MustCompile(`^220[0-4][0-9]{12,15}$`),
	UATP:                    regexp.MustCompile(`^1[0-9]{14}$`),
	// BUG-07 fix: anchor the whole alternation, not just one branch, and
	// write \d{2} instead of the odd \d{02}.
	Hipercard:       regexp.MustCompile(`^(?:606282\d{10}(?:\d{3})?|3841\d{2}\d{10})$`),
	Naranja:         regexp.MustCompile(`^589562\d{10}$`),
	TarjetaShopping: regexp.MustCompile(`^603488\d{10}$`),
	ELO:             regexp.MustCompile(`^(401178|401179|431274|438935|451416|457393|457631|457632|504175|627780|636297|636368|636369)\d{10}$`),
}

// BankCard validates a bank card number, optionally restricting it to one
// or more card brands.
//
// Validation always combines two independent checks: the Luhn checksum and a
// brand pattern. If no kind is given, the number is checked against the union
// of all supported brands. If one or more kinds are given, the number is
// valid when it passes the Luhn check AND matches at least one of them.
//
// Spaces and hyphens in the input are ignored, so grouped numbers such as
// "4111 1111 1111 1111" validate the same as "4111111111111111".
//
// Example usage:
//
//	is.BankCard("4111111111111111")
//	// Output: true (valid number, any brand)
//
//	is.BankCard("4111111111111111", is.MasterCard)
//	// Output: false (valid number, but not a MasterCard)
//
//	is.BankCard("4111111111111111", is.Visa)
//	// Output: true (valid Visa)
//
//	is.BankCard("4111111111111111", is.Visa, is.MasterCard)
//	// Output: true (matches Visa or MasterCard)
//
//	is.BankCard("1234567812345678")
//	// Output: false (fails the Luhn check)
func BankCard(str string, kinds ...CardKind) bool {
	if len(kinds) == 0 {
		return cardChecker(str, anyCreditCard())
	}

	for _, kind := range kinds {
		if regex, ok := cardPatterns[kind]; ok && cardChecker(str, regex) {
			return true
		}
	}

	return false
}

// The cardChecker checks whether a string matches a given card regex and
// validates its checksum according to the Luhn algorithm.
//
// Here's how it works:
//  1. It ignores spaces and hyphens in the input string.
//  2. It calculates the Luhn checksum of the card number.
//  3. It checks that the sum of the digits is a multiple of 10.
//  4. It validates the card number against the provided regular expression.
//
// It returns true if all checks pass, and false otherwise.
func cardChecker(n string, regex *regexp.Regexp) bool {
	// Fast path: only allocate a cleaned copy if there is something to
	// clean. Typical input is already free of spaces and hyphens.
	clean := n
	if strings.ContainsAny(n, " -") {
		clean = strings.Map(func(r rune) rune {
			if r == ' ' || r == '-' {
				return -1 // remove the character
			}
			return r
		}, n)
	}

	if clean == "" {
		return false
	}

	var sum int
	parity := len(clean) % 2

	for i, r := range clean {
		// Only ASCII digits are valid in a card number.
		if r < '0' || r > '9' {
			return false
		}

		d := int(r - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
	}

	// Check if the sum is a multiple of 10.
	if sum%10 != 0 {
		return false
	}

	// Check the regular expression match.
	return regex.MatchString(clean)
}
