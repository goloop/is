package is_test

import (
	"errors"
	"fmt"

	"github.com/goloop/is/v2"
)

func ExampleEmail() {
	fmt.Println(is.Email("user@example.com"))
	fmt.Println(is.Email(" user@example.com")) // not trimmed
	// Output:
	// true
	// false
}

func ExampleBankCard() {
	fmt.Println(is.BankCard("4111111111111111"))          // any brand
	fmt.Println(is.BankCard("4111111111111111", is.Visa)) // as Visa
	fmt.Println(is.BankCard("4111111111111111", is.MasterCard))
	// Output:
	// true
	// true
	// false
}

func ExampleIBAN() {
	fmt.Println(is.IBAN("GB82 WEST 1234 5698 7654 32"))       // spaces ok
	fmt.Println(is.IBAN("GB82 WEST 1234 5698 7654 32", true)) // strict: no spaces
	// Output:
	// true
	// false
}

func ExampleNumeric() {
	fmt.Println(is.Numeric("3.14")) // decimal
	fmt.Println(is.Numeric("٣٫١٤")) // Arabic-Indic digits
	fmt.Println(is.Numeric("一二三"))  // CJK numerals are not digits
	fmt.Println(is.Numeric("."))    // no digit
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleIPv4() {
	fmt.Println(is.IPv4("192.168.0.1"))
	fmt.Println(is.IPv4("192.168.0.01")) // leading zero is rejected
	// Output:
	// true
	// false
}

func ExampleMD5() {
	fmt.Println(is.MD5("d41d8cd98f00b204e9800998ecf8427e"))
	fmt.Println(is.MD5("0xd41d8cd98f00b204e9800998ecf842")) // no prefix allowed
	// Output:
	// true
	// false
}

func ExampleUUID() {
	fmt.Println(is.UUID("550e8400-e29b-41d4-a716-446655440000"))
	fmt.Println(is.UUID("550e8400e29b41d4a716446655440000")) // needs hyphens
	// Output:
	// true
	// false
}

func ExampleURL() {
	fmt.Println(is.URL("https://example.com"))
	fmt.Println(is.URL("example.com")) // no scheme
	// Output:
	// true
	// false
}

func ExampleDomain() {
	fmt.Println(is.Domain("sub.example.com"))
	fmt.Println(is.Domain("localhost")) // no TLD
	// Output:
	// true
	// false
}

func ExampleIBANCountry() {
	country, ok := is.IBANCountry("DE89370400440532013000")
	fmt.Println(country, ok)
	// Output:
	// DE true
}

func ExampleLatitude() {
	fmt.Println(is.Latitude(45.5))
	fmt.Println(is.Latitude("45.5"))
	fmt.Println(is.Latitude("1e1")) // scientific notation rejected
	fmt.Println(is.Latitude(90.1))  // out of range
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleVariableNameFor() {
	ok, _ := is.VariableNameFor("myVar", "go")
	fmt.Println(ok)

	// csharp is now fully supported (previously this call panicked).
	ok, _ = is.VariableNameFor("myVar", "csharp")
	fmt.Println(ok)

	_, err := is.VariableNameFor("x", "klingon")
	fmt.Println(err)
	fmt.Println(errors.Is(err, is.ErrLanguageNotSupported))
	// Output:
	// true
	// true
	// programming language is not supported: "klingon"
	// true
}
