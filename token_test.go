package is

import "testing"

func TestTokenOctet(t *testing.T) {
	valid := map[byte]bool{}
	for c := byte('0'); c <= '9'; c++ {
		valid[c] = true
	}
	for c := byte('A'); c <= 'Z'; c++ {
		valid[c] = true
	}
	for c := byte('a'); c <= 'z'; c++ {
		valid[c] = true
	}
	for _, c := range []byte("!#$%&'*+-.^_`|~") {
		valid[c] = true
	}

	for i := 0; i < 256; i++ {
		b := byte(i)
		got := TokenOctet(b)
		want := valid[b]
		if got != want {
			t.Errorf("TokenOctet(%q / 0x%02x) = %v; want %v",
				b, b, got, want)
		}
	}
}

func TestTokenOctetExamples(t *testing.T) {
	tests := []struct {
		name string
		in   byte
		want bool
	}{
		{name: "letter", in: 'A', want: true},
		{name: "digit", in: '7', want: true},
		{name: "allowed symbol", in: '_', want: true},
		{name: "space", in: ' ', want: false},
		{name: "tab", in: '\t', want: false},
		{name: "separator slash", in: '/', want: false},
		{name: "separator colon", in: ':', want: false},
		{name: "non ascii", in: 0x80, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenOctet(tt.in); got != tt.want {
				t.Errorf("TokenOctet(%q) = %v; want %v",
					tt.in, got, tt.want)
			}
		})
	}
}

func TestIsToken(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{name: "field name", in: "Content-Type", want: true},
		{name: "method", in: "GET", want: true},
		{name: "all symbols", in: "!#$%&'*+-.^_`|~", want: true},
		{name: "digits", in: "12345", want: true},
		{name: "single char", in: "a", want: true},
		{name: "empty", in: "", want: false},
		{name: "inner space", in: "a b", want: false},
		{name: "leading space", in: " GET", want: false},
		{name: "trailing space", in: "GET ", want: false},
		{name: "separator colon", in: "Host:", want: false},
		{name: "separator slash", in: "a/b", want: false},
		{name: "control byte", in: "a\tb", want: false},
		{name: "non ascii", in: "café", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsToken(tt.in); got != tt.want {
				t.Errorf("IsToken(%q) = %v; want %v",
					tt.in, got, tt.want)
			}
		})
	}
}

// TestIsTokenMatchesOctet cross-checks IsToken against TokenOctet: a
// single-byte string is a token exactly when that byte is a token octet.
func TestIsTokenMatchesOctet(t *testing.T) {
	for i := 0; i < 128; i++ {
		b := byte(i)
		if got, want := IsToken(string(b)), TokenOctet(b); got != want {
			t.Errorf("IsToken(%q) = %v; TokenOctet = %v", b, got, want)
		}
	}
}
