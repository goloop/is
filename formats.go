package is

import (
	"encoding/base32"
	"encoding/hex"
	"encoding/json"
	"mime"
	"net/netip"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var datePattern = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
var timestampPattern = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}[Tt][0-9]{2}:[0-9]{2}:[0-9]{2}(\.[0-9]+)?([Zz]|[+-][0-9]{2}:[0-9]{2})$`)

// Date checks a Gregorian YYYY-MM-DD date, including leap years (years 0001-9999).
func Date(s string) bool {
	if !datePattern.MatchString(s) || s[:4] == "0000" {
		return false
	}
	_, err := time.Parse(time.DateOnly, s)
	return err == nil
}

// UTCOffset checks Z or a signed HH:MM offset with hours 00-23 and minutes 00-59.
// -00:00 is accepted as RFC 3339's unknown local offset.
func UTCOffset(s string) bool {
	if s == "Z" {
		return true
	}
	if len(s) != 6 || (s[0] != '+' && s[0] != '-') || s[3] != ':' {
		return false
	}
	for _, i := range []int{1, 2, 4, 5} {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return s[1:3] <= "23" && s[4:6] <= "59"
}

// RFC3339 checks an RFC 3339 timestamp in Go's time.Time range. Lowercase t/z
// are accepted; leap-second notation is not supported by time.Time.
func RFC3339(s string) bool {
	if !timestampPattern.MatchString(s) || !Date(s[:10]) {
		return false
	}
	s = strings.ToUpper(s)
	if s[len(s)-1] != 'Z' && !UTCOffset(s[len(s)-6:]) {
		return false
	}
	_, err := time.Parse(time.RFC3339Nano, s)
	return err == nil
}

// Duration checks Go duration syntax, not ISO 8601 calendar periods.
func Duration(s string) bool { _, err := time.ParseDuration(s); return err == nil }

// CIDR checks an IPv4 or IPv6 prefix. Host bits are allowed; zones are not.
func CIDR(s string) bool { _, err := netip.ParsePrefix(s); return err == nil }

// Port checks an ASCII decimal port in 0-65535, without signs or whitespace.
func Port(s string) bool {
	if len(s) == 0 || len(s) > 5 {
		return false
	}
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	n, err := strconv.ParseUint(s, 10, 16)
	return err == nil && n <= 65535
}

// IPPort checks an IP literal and port, with brackets around IPv6. Zones are rejected.
func IPPort(s string) bool {
	p, err := netip.ParseAddrPort(s)
	i := strings.LastIndexByte(s, ':')
	return err == nil && p.Addr().Zone() == "" && i >= 0 && Port(s[i+1:])
}

// JSON checks one complete UTF-8 JSON value, including scalar values.
func JSON(s string) bool { return utf8.ValidString(s) && json.Valid([]byte(s)) }

// MIME checks a concrete type/subtype and optional parameters. It checks syntax,
// not IANA registration; wildcard media ranges and surrounding spaces are rejected.
func MIME(s string) bool {
	if strings.TrimSpace(s) != s || strings.ContainsAny(s, "\r\n") {
		return false
	}
	typ, _, err := mime.ParseMediaType(s)
	parts := strings.Split(typ, "/")
	return err == nil && len(parts) == 2 && parts[0] != "" && parts[1] != "" && !strings.Contains(typ, "*")
}

var semverPattern = regexp.MustCompile(`^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-([0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*))?(\+([0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*))?$`)

// SemVer checks Semantic Versioning 2.0.0, without a leading v.
func SemVer(s string) bool {
	if !semverPattern.MatchString(s) {
		return false
	}
	core := strings.SplitN(s, "+", 2)[0]
	if i := strings.IndexByte(core, '-'); i >= 0 {
		for _, part := range strings.Split(core[i+1:], ".") {
			numeric := true
			for j := range part {
				if part[j] < '0' || part[j] > '9' {
					numeric = false
					break
				}
			}
			if numeric && len(part) > 1 && part[0] == '0' {
				return false
			}
		}
	}
	return true
}

// Base32 checks nonempty canonical RFC 4648 uppercase Base32. Padding is required
// by default; passing false selects the unpadded form. Nonzero pad bits fail.
func Base32(s string, padded ...bool) bool {
	if s == "" {
		return false
	}
	enc := base32.StdEncoding
	if len(padded) > 0 && !padded[0] {
		enc = enc.WithPadding(base32.NoPadding)
	}
	b, err := enc.DecodeString(s)
	return err == nil && enc.EncodeToString(b) == s
}

// SHA384 checks exactly 96 hexadecimal characters (either case).
func SHA384(s string) bool {
	if len(s) != 96 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}

// UUIDv4 checks the UUID text form, version 4, and RFC 9562 variant bits.
func UUIDv4(s string) bool { return uuidVersion(s, '4') }

// UUIDv7 checks the UUID text form, version 7, and RFC 9562 variant bits.
func UUIDv7(s string) bool { return uuidVersion(s, '7') }
func uuidVersion(s string, v byte) bool {
	return UUID(s) && s[14] == v && strings.ContainsRune("89aAbB", rune(s[19]))
}
