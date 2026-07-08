package is

// TokenOctet reports whether b is a valid HTTP token character ("tchar").
//
// The HTTP token grammar (RFC 9110, section 5.6.2) allows the ASCII letters
// and digits plus a fixed set of symbols, and nothing else:
//
//	tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//	        "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
//
// These are the octets that may appear, unescaped, in HTTP field names,
// request-method names, cache-directive names and similar bare tokens. The
// delimiters - space, tab, control bytes and the separators
// ()<>@,;:\"/[]?={} - are not token characters, so TokenOctet returns false
// for them, and for every non-ASCII byte (0x80-0xFF).
func TokenOctet(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}

	switch b {
	case '!', '#', '$', '%', '&', '\'', '*', '+', '-',
		'.', '^', '_', '`', '|', '~':
		return true
	}

	return false
}

// IsToken reports whether s is a valid HTTP token: one or more token
// characters and nothing else (RFC 9110, section 5.6.2):
//
//	token = 1*tchar
//
// The empty string is not a token, so IsToken("") is false. Each byte is
// checked with [TokenOctet], so a token never contains whitespace, control
// bytes, non-ASCII bytes or any HTTP delimiter. This is the rule that HTTP
// field names and request-method names must satisfy, which makes IsToken a
// convenient guard when reading or forwarding those values.
//
//	is.IsToken("Content-Type") // true
//	is.IsToken("GET")          // true
//	is.IsToken("a b")          // false - space is a delimiter
//	is.IsToken("")             // false - a token needs at least one octet
//
// The check is byte-oriented on purpose: tokens are ASCII by definition, so a
// multi-byte UTF-8 sequence can never be a valid token and is rejected at its
// first non-ASCII byte.
func IsToken(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !TokenOctet(s[i]) {
			return false
		}
	}

	return true
}
