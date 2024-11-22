package is

import (
	"net"
	"strconv"
	"strings"
)

// IPv4 checks if the string is a valid representation of an IPv4 address.
// An IPv4 address consists of four numbers separated by dots.
// Each number should be within the range of 0 to 255.
// Leading zeros are allowed. Empty strings or strings with more than
// four parts are not considered as valid IPv4 addresses.
//
// Example usage:
//
//	is.IPv4("127.0.0.1")       // Returns: true
//	is.IPv4("192.168.0.1")     // Returns: true
//	is.IPv4("0.0.0.0")         // Returns: true
//	is.IPv4("255.255.255.255") // Returns: true
//
//	is.IPv4("256.0.0.1")     // Returns: false, numbers exceed the range
//	is.IPv4("192.168.0")     // Returns: false, only three parts
//	is.IPv4("192.168.0.1.1") // Returns: false, more than four parts
//	is.IPv4("192.168.0.one") // Returns: false, non-numeric characters
//	is.IPv4("")              // Returns: false, empty string
//
// This function can be used to validate user input to ensure
// an IPv4 address entered is in the correct format before
// attempting to use it in network operations.
func IPv4(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if len(part) == 0 || (len(part) > 1 && part[0] == '0') {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

// IPv6 checks if the string is a valid representation of an IPv6 address.
// An IPv6 address consists of eight groups of four hexadecimal digits,
// each group representing 16 bits. The groups are separated by colons (:).
//
// This function uses the net.ParseIP function from the Go standard library
// which returns a valid IP address (either IPv6 or IPv4). The function then
// checks if the string contains a colon, which is a requirement for it to be
// an IPv6 address.
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
	return net.ParseIP(ip) != nil && strings.Contains(ip, ":")
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
