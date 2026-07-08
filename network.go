package is

import (
	"net"
	"net/netip"
	"net/url"
	"regexp"
	"strings"
)

// hostnameRegex matches an RFC 1123 hostname: one or more dot-separated
// labels, each 1-63 characters of letters/digits/hyphens, not starting or
// ending with a hyphen. The total length is bounded separately (<= 253).
var hostnameRegex = regexp.MustCompile(
	`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)*` +
		`[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`,
)

// IPv4 checks if the string is a valid representation of an IPv4 address.
// An IPv4 address consists of four decimal octets in the range 0-255,
// separated by dots (dotted-decimal notation).
//
// Leading zeros are NOT allowed: "192.168.0.01" is rejected, because a
// leading zero is ambiguous (it could be read as octal) and the canonical
// textual form has none. This matches the parsing rules of the standard
// library's net/netip. Empty strings, IPv6 addresses, and strings with the
// wrong number of octets are not valid IPv4 addresses.
//
// Example usage:
//
//	is.IPv4("127.0.0.1")       // Returns: true
//	is.IPv4("192.168.0.1")     // Returns: true
//	is.IPv4("0.0.0.0")         // Returns: true
//	is.IPv4("255.255.255.255") // Returns: true
//
//	is.IPv4("192.168.0.01")  // Returns: false, leading zero
//	is.IPv4("256.0.0.1")     // Returns: false, octet exceeds the range
//	is.IPv4("192.168.0")     // Returns: false, only three octets
//	is.IPv4("192.168.0.1.1") // Returns: false, more than four octets
//	is.IPv4("192.168.0.one") // Returns: false, non-numeric characters
//	is.IPv4("")              // Returns: false, empty string
//
// This function can be used to validate user input to ensure
// an IPv4 address entered is in the correct format before
// attempting to use it in network operations.
func IPv4(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	return err == nil && addr.Is4()
}

// IPv6 checks if the string is a valid representation of an IPv6 address.
// An IPv6 address consists of eight groups of four hexadecimal digits,
// each group representing 16 bits. The groups are separated by colons (:).
//
// This function uses net/netip from the Go standard library to parse the
// address and then checks that it is an IPv6 address (including IPv4-mapped
// IPv6 addresses), which excludes plain dotted-decimal IPv4 input. A scoped
// address carrying a zone identifier (for example "fe80::1%eth0") is
// rejected: the zone is a routing scope, not part of the address itself.
//
// Example usage:
//
//	is.IPv6("2001:0db8:85a3:0000:0000:8a2e:0370:7334") // Returns: true
//	is.IPv6("2001:db8:85a3:0:0:8a2e:370:7334")         // Returns: true
//	is.IPv6("2001:db8:85a3::8a2e:370:7334")            // Returns: true
//	is.IPv6("::1")                                     // Returns: true
//	is.IPv6("::")                                      // Returns: true
//
//	// The group "37023" exceeds 16 bits.
//	is.IPv6("2001:db8::8a2e:37023:7334") // Returns: false
//
//	// Only one "::" is allowed in an IPv6 address.
//	is.IPv6("2001::25de::cade") // Returns: false
//
//	// This is an IPv4 address.
//	is.IPv6("192.168.0.1") // Returns: false,
//	is.IPv6("") // Returns: false, empty string
//
// This function can be used to validate user input to ensure an IPv6
// address entered is in the correct format before attempting to use
// it in network operations.
func IPv6(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	return err == nil && addr.Is6() && addr.Zone() == ""
}

// IP checks if the string is a valid representation of an IP address.
// The IP address can be either IPv4 or IPv6.
//
// This function first checks if the string is a valid IPv4 address using
// the IPv4 function, if that check fails it then checks if the string is
// a valid IPv6 address using the IPv6 function.
//
// Example usage:
//
//	is.IP("127.0.0.1")       // Returns: true, valid IPv4
//	is.IP("::1")             // Returns: true, valid IPv6
//	is.IP("2001:db8::8a2e")  // Returns: true, valid IPv6
//
//	is.IP("256.0.0.1")     // Returns: false, invalid IPv4
//	is.IP("192.168.0")     // Returns: false, invalid IPv4
//	is.IP("2001::25de::cade") // Returns: false, invalid IPv6
//	is.IP("")              // Returns: false, empty string
//
// This function can be used to validate user input to ensure
// an IP address entered is in the correct format before
// attempting to use it in network operations.
func IP(ip string) bool {
	return IPv4(ip) || IPv6(ip)
}

// MAC checks if the string is a valid IEEE 802 MAC address. It accepts the
// standard textual forms parsed by the standard library: six groups of two
// hex digits separated by colons or hyphens (EUI-48), the four-group
// dotted form, and the eight-group EUI-64 variants.
//
// Example usage:
//
//	is.MAC("00:1b:63:84:45:e6") // Returns: true
//	is.MAC("00-1B-63-84-45-E6") // Returns: true
//	is.MAC("001b.6384.45e6")    // Returns: true
//	is.MAC("00:1b:63:84:45")    // Returns: false, too few groups
//	is.MAC("")                  // Returns: false, empty string
func MAC(v string) bool {
	_, err := net.ParseMAC(v)
	return err == nil
}

// URL checks if the string is a valid absolute URL: it must have both a
// scheme (such as "http" or "https") and a host. Relative references and
// scheme-only or host-less inputs are rejected.
//
// The function validates structure, not reachability: it does not perform
// any network request.
//
// Example usage:
//
//	is.URL("https://example.com")          // Returns: true
//	is.URL("http://example.com:8080/path") // Returns: true
//	is.URL("ftp://files.example.com")      // Returns: true
//	is.URL("example.com")                  // Returns: false, no scheme
//	is.URL("/relative/path")               // Returns: false, no scheme/host
//	is.URL("https://")                     // Returns: false, no host
//	is.URL("")                             // Returns: false, empty string
func URL(v string) bool {
	if v == "" {
		return false
	}

	u, err := url.Parse(v)
	if err != nil {
		return false
	}

	return u.Scheme != "" && u.Host != ""
}

// Hostname checks if the string is a valid RFC 1123 hostname: one or more
// dot-separated labels of letters, digits, and hyphens, where no label
// starts or ends with a hyphen, each label is at most 63 characters, and the
// whole name is at most 253 characters.
//
// It validates the textual format only; it does not resolve the name.
//
// Example usage:
//
//	is.Hostname("example.com")     // Returns: true
//	is.Hostname("sub.example.com") // Returns: true
//	is.Hostname("localhost")       // Returns: true
//	is.Hostname("-bad.example")    // Returns: false, label starts with '-'
//	is.Hostname("a..b")            // Returns: false, empty label
//	is.Hostname("")                // Returns: false, empty string
func Hostname(v string) bool {
	if len(v) == 0 || len(v) > 253 {
		return false
	}

	return hostnameRegex.MatchString(v)
}

// Domain checks if the string is a valid domain name: a hostname (see
// Hostname) with at least two labels and a top-level domain of two or more
// letters. Unlike Hostname, a single label such as "localhost" and an
// all-numeric or one-letter TLD are rejected.
//
// It validates the textual format only; it does not resolve the name or
// consult a public-suffix list.
//
// Example usage:
//
//	is.Domain("example.com")     // Returns: true
//	is.Domain("sub.example.com") // Returns: true
//	is.Domain("localhost")       // Returns: false, no TLD
//	is.Domain("example.123")     // Returns: false, numeric TLD
//	is.Domain("example.c")       // Returns: false, one-letter TLD
//	is.Domain("example.xn--p1ai") // Returns: true, punycode (IDN) TLD
//	is.Domain("")                // Returns: false, empty string
func Domain(v string) bool {
	if !Hostname(v) {
		return false
	}

	// Require a dot-separated TLD of at least two ASCII letters.
	i := strings.LastIndexByte(v, '.')
	if i < 0 {
		return false
	}

	tld := v[i+1:]
	if len(tld) < 2 {
		return false
	}

	// A punycode (IDN A-label) TLD such as "xn--p1ai" (.рф) is valid: its label
	// characters were already checked by Hostname. Otherwise the TLD must be
	// ASCII letters only, which rejects a purely numeric TLD.
	if !hasACEPrefix(tld) {
		for j := 0; j < len(tld); j++ {
			c := tld[j]
			if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
				return false
			}
		}
	}

	return true
}

// hasACEPrefix reports whether s begins with the case-insensitive IDN ACE
// prefix "xn--".
func hasACEPrefix(s string) bool {
	return len(s) >= 4 &&
		(s[0] == 'x' || s[0] == 'X') &&
		(s[1] == 'n' || s[1] == 'N') &&
		s[2] == '-' && s[3] == '-'
}
