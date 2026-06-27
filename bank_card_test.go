package is

import (
	"regexp"
	"testing"
)

// TestCardChecker tests the cardChecker function.
func TestCardChecker(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "Card number with non-numeric characters",
			in:   "1234-5678-9012-abcd",
			want: false,
		},
		{
			name: "Card number with numeric characters out of range",
			in:   "1234-5678-9012-123:4",
			want: false,
		},
		{
			name: "Card number with slash",
			in:   "1234-5678-9012-123/4",
			want: false,
		},
		{
			name: "Card number with characters less than zero ASCII code",
			in:   "1234-5678-9012-123/",
			want: false,
		},
		{
			name: "Card number with characters greater than ten ASCII code",
			in:   "1234-5678-9012-123:",
			want: false,
		},
		{
			name: "Valid card number",
			in:   "4111-1111-1111-1111",
			want: true,
		},
	}

	regex := regexp.MustCompile(`^\d{16}$`)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cardChecker(tc.in, regex)
			if got != tc.want {
				t.Errorf("cardChecker(%q) = %v; want %v", tc.in, got, tc.want)
			}
		})
	}
}

// TestBankCard tests the BankCard function.
func TestBankCard(t *testing.T) {
	tests := []struct {
		name     string
		card     string
		kinds    []CardKind
		expected bool
	}{
		{
			name:     "Valid Visa card",
			card:     "4111111111111111",
			kinds:    []CardKind{Visa},
			expected: true,
		},
		{
			name:     "Invalid MasterCard card",
			card:     "4111111111111111",
			kinds:    []CardKind{MasterCard},
			expected: false,
		},
		{
			name:     "Invalid card number",
			card:     "1234567812345678",
			kinds:    nil,
			expected: false,
		},
		{
			name:     "Valid card of any type",
			card:     "4111111111111111",
			kinds:    nil,
			expected: true,
		},
		{
			name:     "Valid Amex card",
			card:     "342883359122187",
			kinds:    []CardKind{AmericanExpress},
			expected: true,
		},
		{
			name:     "Invalid Amex card",
			card:     "378282246310007",
			kinds:    []CardKind{AmericanExpress},
			expected: false,
		},
		{
			name:     "Valid China Union Pay card",
			card:     "6250941006528599",
			kinds:    []CardKind{ChinaUnionPay},
			expected: true,
		},
		{
			name:     "Valid Discover card",
			card:     "60115564485789458",
			kinds:    []CardKind{DiscoverCard},
			expected: false,
		},
		{
			name:     "Valid JCB card",
			card:     "3566000020000410",
			kinds:    []CardKind{JCB},
			expected: true,
		},
		{
			name:     "Invalid JCB card",
			card:     "3530111333300000",
			kinds:    []CardKind{JCB},
			expected: true,
		},
		{
			name:     "Valid Mastercard card",
			card:     "5425233430109903",
			kinds:    []CardKind{MasterCard},
			expected: true,
		},
		{
			name:     "Valid Visa card",
			card:     "4263982640269299",
			kinds:    []CardKind{Visa},
			expected: true,
		},
		{
			name:     "Invalid Visa card",
			card:     "491748458989713",
			kinds:    []CardKind{Visa},
			expected: false,
		},
		{
			name:     "Valid Visa card",
			card:     "4001919257537193",
			kinds:    []CardKind{Visa},
			expected: true,
		},
		{
			name:     "Valid ELO card",
			card:     "6362970000457013",
			kinds:    []CardKind{ELO},
			expected: true,
		},
		{
			name:     "Valid Hipercard card",
			card:     "6062826786276634",
			kinds:    []CardKind{Hipercard},
			expected: true,
		},
		{
			name:     "Valid Argencard card",
			card:     "5011054488597827",
			kinds:    []CardKind{Argencard},
			expected: true,
		},
		{
			name:     "Valid Cabal card",
			card:     "604211212211",
			kinds:    []CardKind{Cabal},
			expected: true,
		},
		{
			name:     "Valid Cencosud card",
			card:     "6034932528973614",
			kinds:    []CardKind{Cencosud},
			expected: true,
		},
		{
			name:     "Valid Naranja card",
			card:     "5895626746595650",
			kinds:    []CardKind{Naranja},
			expected: true,
		},
		{
			name:     "Valid Tarjeta Shopping card",
			card:     "6034883265619896",
			kinds:    []CardKind{TarjetaShopping},
			expected: true,
		},
		{
			name:     "Valid Visa card with spaces",
			card:     "4111 1111 1111 1111",
			kinds:    []CardKind{Visa},
			expected: true,
		},
		{
			name:     "Valid Visa card with hyphens",
			card:     "4111-1111-1111-1111",
			kinds:    []CardKind{Visa},
			expected: true,
		},
		{
			name:     "Valid MasterCard card with spaces",
			card:     "5425 2334 3010 9903",
			kinds:    []CardKind{MasterCard},
			expected: true,
		},
		{
			name:     "Valid MasterCard card with hyphens",
			card:     "5425-2334-3010-9903",
			kinds:    []CardKind{MasterCard},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BankCard(tt.card, tt.kinds...)
			if result != tt.expected {
				t.Errorf("BankCard(%q, %v) = %v; want %v",
					tt.card, tt.kinds, result, tt.expected)
			}
		})
	}
}
