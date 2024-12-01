package is

import (
	"testing"
)

// Benchmark variables to prevent compiler optimizations
var (
	resultBool   bool
	resultString string
	resultInt    int
	resultFloat  float64
)

// Account validation benchmarks
func BenchmarkNickname(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		strict   bool
		expected bool
	}{
		{"ValidSimple", "User123", false, true},
		{"ValidUnicode", "用户123", false, true},
		{"ValidStrict", "User123", true, true},
		{"Invalid", "User@123", false, false},
		{"InvalidStrict", "用户123", true, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Nickname(tc.input, tc.strict)
			}
		})
	}
}

// Bank Card validation benchmarks
func BenchmarkBankCard(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidVisa", "4111111111111111", true},
		{"ValidMasterCard", "5555555555554444", true},
		{"ValidAmEx", "378282246310005", true},
		{"Invalid", "1234567812345678", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = BankCard(tc.input)
			}
		})
	}
}

// IBAN validation benchmarks
func BenchmarkIBAN(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		strict   bool
		expected bool
	}{
		{"ValidGB", "GB82 WEST 1234 5698 7654 32", false, true},
		{"ValidDE", "DE89370400440532013000", true, true},
		{"Invalid", "Invalid IBAN", false, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = IBAN(tc.input, tc.strict)
			}
		})
	}
}

// Code validation benchmarks
func BenchmarkVariableName(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		strict   bool
		expected bool
	}{
		{"ValidSimple", "myVariable", false, true},
		{"ValidStrict", "myVariable", true, true},
		{"Invalid", "1variable", false, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = VariableName(tc.input, tc.strict)
			}
		})
	}
}

// Email validation benchmarks
func BenchmarkEmail(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "test@example.com", true},
		{"ValidComplex", "test.name+filter@sub.example.com", true},
		{"Invalid", "not-an-email", false},
		{"InvalidFormat", "test@", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Email(tc.input)
			}
		})
	}
}

// Geographic validation benchmarks
func BenchmarkLatitude(b *testing.B) {
	cases := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"ValidPositive", 45.0, true},
		{"ValidNegative", -45.0, true},
		{"Invalid", 91.0, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Latitude(tc.input)
			}
		})
	}
}

func BenchmarkLongitude(b *testing.B) {
	cases := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"ValidPositive", 120.0, true},
		{"ValidNegative", -120.0, true},
		{"Invalid", 181.0, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Longitude(tc.input)
			}
		})
	}
}

// Hash validation benchmarks
func BenchmarkMD5(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "d41d8cd98f00b204e9800998ecf8427e", true},
		{"Invalid", "not-an-md5-hash", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = MD5(tc.input)
			}
		})
	}
}

// Network validation benchmarks
func BenchmarkIPv4(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "192.168.0.1", true},
		{"ValidLocalhost", "127.0.0.1", true},
		{"Invalid", "256.256.256.256", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = IPv4(tc.input)
			}
		})
	}
}

func BenchmarkIPv6(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"ValidShort", "::1", true},
		{"Invalid", "not-an-ipv6", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = IPv6(tc.input)
			}
		})
	}
}

// Number validation benchmarks
func BenchmarkEven(b *testing.B) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"ValidEven", 2, true},
		{"InvalidOdd", 3, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Even(tc.input)
			}
		})
	}
}

func BenchmarkOdd(b *testing.B) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"ValidOdd", 3, true},
		{"InvalidEven", 2, false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Odd(tc.input)
			}
		})
	}
}

// Phone validation benchmarks
func BenchmarkE164(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "+12345678901", true},
		{"Invalid", "12345678901", false},
		{"InvalidFormat", "+1234abcd", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = E164(tc.input)
			}
		})
	}
}

func BenchmarkPhone(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "+380961234567", true},
		{"ValidWithSpaces", "+380 96 123 45 67", true},
		{"Invalid", "1234567890", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Phone(tc.input)
			}
		})
	}
}

// String validation benchmarks
func BenchmarkDigit(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "12345", true},
		{"Invalid", "12345a", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Digit(tc.input)
			}
		})
	}
}

func BenchmarkNumeric(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidInteger", "12345", true},
		{"ValidDecimal", "123.45", true},
		{"ValidNegative", "-123.45", true},
		{"Invalid", "12345a", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Numeric(tc.input)
			}
		})
	}
}

func BenchmarkAlpha(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "abcdef", true},
		{"Invalid", "abc123", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Alpha(tc.input)
			}
		})
	}
}

func BenchmarkAlnum(b *testing.B) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid", "abc123", true},
		{"Invalid", "abc123!", false},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBool = Alnum(tc.input)
			}
		})
	}
}
