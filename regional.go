package is

import (
	"regexp"
	"strings"
)

// CountryCode checks an assigned ISO 3166-1 alpha-2 code in uppercase.
// Reserved/user-assigned codes (including UK and XK) are not ISO assignments.
func CountryCode(s string) bool { return len(s) == 2 && strings.Contains(countryCodes, " "+s+" ") }

// CurrencyCode checks an uppercase ISO 4217 code in SIX List One, including
// fund, metal, testing (XTS) and no-currency (XXX) codes. It does not imply tender.
func CurrencyCode(s string) bool { return len(s) == 3 && strings.Contains(currencyCodes, " "+s+" ") }

// BIC checks ISO 9362 identifier syntax and its ISO country code. XK is also
// accepted for Kosovo. It does not query the directory of assigned identifiers.
func BIC(s string) bool { return bicPattern.MatchString(s) && (CountryCode(s[4:6]) || s[4:6] == "XK") }

var bicPattern = regexp.MustCompile(`^[A-Z0-9]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`)

// PostalCode checks the regional syntax of a postal code. Country must be an
// uppercase ISO alpha-2 code. Unsupported countries return false. This does not
// establish that a postcode is assigned or an address is deliverable.
func PostalCode(s, country string) bool {
	p, ok := postalPatterns[country]
	return ok && p.MatchString(s)
}

var postalPatterns = map[string]*regexp.Regexp{
	"US": regexp.MustCompile(`^[0-9]{5}(-[0-9]{4})?$`),
	"GB": regexp.MustCompile(`^(GIR 0AA|([A-PR-UWYZ][0-9][0-9]?|[A-PR-UWYZ][A-HK-Y][0-9][0-9]?|[A-PR-UWYZ][0-9][A-HJKPSTUW]|[A-PR-UWYZ][A-HK-Y][0-9][ABEHMNPRVWXY]) [0-9][ABD-HJLNP-UW-Z]{2})$`),
	"DE": regexp.MustCompile(`^[0-9]{5}$`),
	"PL": regexp.MustCompile(`^[0-9]{2}-[0-9]{3}$`),
	"UA": regexp.MustCompile(`^[0-9]{5}$`),
	"FR": regexp.MustCompile(`^[0-9]{5}$`),
	"ES": regexp.MustCompile(`^(0[1-9]|[1-4][0-9]|5[0-2])[0-9]{3}$`),
	"IT": regexp.MustCompile(`^[0-9]{5}$`),
	"PT": regexp.MustCompile(`^[1-9][0-9]{3}-[0-9]{3}$`),
	"NL": regexp.MustCompile(`^[1-9][0-9]{3} [A-Z]{2}$`),
	"CZ": regexp.MustCompile(`^[1-7][0-9]{2} [0-9]{2}$`),
}

// ISO 3166-1 alpha-2 assignments; snapshot 2026-10-24.
const countryCodes = " AD AE AF AG AI AL AM AO AQ AR AS AT AU AW AX AZ BA BB BD BE BF BG BH BI BJ BL BM BN BO BQ BR BS BT BV BW BY BZ CA CC CD CF CG CH CI CK CL CM CN CO CR CU CV CW CX CY CZ DE DJ DK DM DO DZ EC EE EG EH ER ES ET FI FJ FK FM FO FR GA GB GD GE GF GG GH GI GL GM GN GP GQ GR GS GT GU GW GY HK HM HN HR HT HU ID IE IL IM IN IO IQ IR IS IT JE JM JO JP KE KG KH KI KM KN KP KR KW KY KZ LA LB LC LI LK LR LS LT LU LV LY MA MC MD ME MF MG MH MK ML MM MN MO MP MQ MR MS MT MU MV MW MX MY MZ NA NC NE NF NG NI NL NO NP NR NU NZ OM PA PE PF PG PH PK PL PM PN PR PS PT PW PY QA RE RO RS RU RW SA SB SC SD SE SG SH SI SJ SK SL SM SN SO SR SS ST SV SX SY SZ TC TD TF TG TH TJ TK TL TM TN TO TR TT TV TW TZ UA UG UM US UY UZ VA VC VE VG VI VN VU WF WS YE YT ZA ZM ZW "

// SIX List One, published 2026-09-17.
const currencyCodes = " AED AFN ALL AMD AOA ARS AUD AWG AZN BAM BBD BDT BHD BIF BMD BND BOB BOV BRL BSD BTN BWP BYN BZD CAD CDF CHE CHF CHW CLF CLP CNY COP COU CRC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GHS GIP GMD GNF GTQ GYD HKD HNL HTG HUF IDR ILS INR IQD IRR ISK JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LYD MAD MDL MGA MKD MMK MNT MOP MRU MUR MVR MWK MXN MXV MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLE SOS SRD SSP STN SVC SYP SZL THB TJS TMT TND TOP TRY TTD TWD TZS UAH UGX USD USN UYI UYU UYW UZS VED VES VND VUV WST XAD XAF XAG XAU XBA XBB XBC XBD XCD XCG XDR XOF XPD XPF XPT XSU XTS XUA XXX YER ZAR ZMW ZWG "
