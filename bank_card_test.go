package is

import (
	"fmt"
	"regexp"
	"testing"
)

// TestBankCard tests the BankCard function.
func TestBankCard(t *testing.T) {
	tests := []struct {
		name     string
		card     string
		kinds    []*regexp.Regexp
		expected bool
	}{
		{
			name:     "Valid Visa card",
			card:     "4111111111111111",
			kinds:    []*regexp.Regexp{Visa},
			expected: true,
		},
		{
			name:     "Invalid MasterCard card",
			card:     "4111111111111111",
			kinds:    []*regexp.Regexp{MasterCard},
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
			kinds:    []*regexp.Regexp{AmericanExpress},
			expected: true,
		},
		{
			name:     "Invalid Amex card",
			card:     "378282246310007",
			kinds:    []*regexp.Regexp{AmericanExpress},
			expected: false,
		},
		{
			name:     "Valid China Union Pay card",
			card:     "6250941006528599",
			kinds:    []*regexp.Regexp{ChinaUnionPay},
			expected: true,
		},
		{
			name:     "Valid Discover card",
			card:     "60115564485789458",
			kinds:    []*regexp.Regexp{DiscoverCard},
			expected: false,
		},
		{
			name:     "Valid JCB card",
			card:     "3566000020000410",
			kinds:    []*regexp.Regexp{JCB},
			expected: true,
		},
		{
			name:     "Invalid JCB card",
			card:     "3530111333300000",
			kinds:    []*regexp.Regexp{JCB},
			expected: true,
		},
		{
			name:     "Valid Mastercard card",
			card:     "5425233430109903",
			kinds:    []*regexp.Regexp{MasterCard},
			expected: true,
		},
		{
			name:     "Valid Visa card",
			card:     "4263982640269299",
			kinds:    []*regexp.Regexp{Visa},
			expected: true,
		},
		{
			name:     "Invalid Visa card",
			card:     "491748458989713",
			kinds:    []*regexp.Regexp{Visa},
			expected: false,
		},
		{
			name:     "Valid Visa card",
			card:     "4001919257537193",
			kinds:    []*regexp.Regexp{Visa},
			expected: true,
		},
		{
			name:     "Valid ELO card",
			card:     "6362970000457013",
			kinds:    []*regexp.Regexp{ELO},
			expected: true,
		},
		{
			name:     "Valid Hipercard card",
			card:     "6062826786276634",
			kinds:    []*regexp.Regexp{Hipercard},
			expected: true,
		},
		{
			name:     "Valid Argencard card",
			card:     "5011054488597827",
			kinds:    []*regexp.Regexp{Argencard},
			expected: true,
		},
		{
			name:     "Valid Cabal card",
			card:     "604211212211",
			kinds:    []*regexp.Regexp{Cabal},
			expected: true,
		},
		{
			name:     "Valid Cencosud card",
			card:     "6034932528973614",
			kinds:    []*regexp.Regexp{Cencosud},
			expected: true,
		},
		{
			name:     "Valid Naranja card",
			card:     "5895626746595650",
			kinds:    []*regexp.Regexp{Naranja},
			expected: true,
		},
		{
			name:     "Valid Tarjeta Shopping card",
			card:     "6034883265619896",
			kinds:    []*regexp.Regexp{TarjetaShopping},
			expected: true,
		},
		{
			name:     "Valid Visa card with spaces",
			card:     "4111 1111 1111 1111",
			kinds:    []*regexp.Regexp{Visa},
			expected: true,
		},
		{
			name:     "Valid Visa card with hyphens",
			card:     "4111-1111-1111-1111",
			kinds:    []*regexp.Regexp{Visa},
			expected: true,
		},
		{
			name:     "Valid MasterCard card with spaces",
			card:     "5425 2334 3010 9903",
			kinds:    []*regexp.Regexp{MasterCard},
			expected: true,
		},
		{
			name:     "Valid MasterCard card with hyphens",
			card:     "5425-2334-3010-9903",
			kinds:    []*regexp.Regexp{MasterCard},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BankCard(tt.card, tt.kinds...)

			ks := ""
			for _, k := range tt.kinds {
				ks += fmt.Sprintf("\"%s\", ", k.String())
			}

			if result != tt.expected {
				t.Errorf("Expected CreditCard(\"%s\", %s) must %v, got %v",
					tt.card, ks, tt.expected, result)
			}
		})
	}
}
