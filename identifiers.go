package is

import "strings"

// GTIN checks the length, digits and GS1 modulo-10 check digit of a GTIN-8,
// GTIN-12, GTIN-13 or GTIN-14. It does not check allocation or ownership.
func GTIN(s string) bool {
	switch len(s) {
	case 8, 12, 13, 14:
	default:
		return false
	}
	sum, weight := 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		sum += int(s[i]-'0') * weight
		weight = 4 - weight
	}
	return sum%10 == 0
}

// EAN checks EAN-8 or EAN-13 digits and checksum.
func EAN(s string) bool { return (len(s) == 8 || len(s) == 13) && GTIN(s) }

// ISBN checks ISBN-10 or ISBN-13 digits and checksum. Single spaces or hyphens
// between groups are accepted, but group allocation and hyphen placement are
// not checked. ISBN-13 requires 978 or 979, excluding the ISMN prefix 9790.
func ISBN(s string) bool {
	var b strings.Builder
	separator := true
	for i := range s {
		if s[i] == ' ' || s[i] == '-' {
			if separator {
				return false
			}
			separator = true
			continue
		}
		b.WriteByte(s[i])
		separator = false
	}
	if separator {
		return false
	}
	s = b.String()
	if len(s) == 13 {
		return (strings.HasPrefix(s, "978") || (strings.HasPrefix(s, "979") && s[3] != '0')) && GTIN(s)
	}
	if len(s) != 10 {
		return false
	}
	sum := 0
	for i := range s {
		d := int(s[i] - '0')
		if i == 9 && s[i] == 'X' {
			d = 10
		} else if s[i] < '0' || s[i] > '9' {
			return false
		}
		sum += d * (10 - i)
	}
	return sum%11 == 0
}
