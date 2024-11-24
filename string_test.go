package is

import "testing"

// TestDigit tests the Digit function.
func TestDigit(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"0", true},
		{"1234", true},
		{"Ⅳ", false},
		{"1234abc", false},
		{"", false},
		{"1.23", false},
		{"1,23", false},
		{"0000", true},
		{" 1234", false},
		{"1234 ", false},
		{"0b10001", false},
		{"0o21", false},
		{"0x11", false},
		{"0", true},
		{"⅕", false},
		{"一", false},
	}

	for _, tc := range testCases {
		got := Digit(tc.input)
		if got != tc.want {
			t.Errorf("Digit(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

// TestNumeric tests the Numeric function.
func TestNumeric(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"0", true},
		{"1234", true},
		{"Ⅳ", true},
		{"1234abc", false},
		{"", false},
		{"1.23", true},
		{"1,23", true},
		{"三・十四", true},
		{"0000", true},
		{" 1234", false},
		{"1234 ", false},
		{"0b10001", false},
		{"0o21", false},
		{"0x11", false},
		{"0", true},
		{"⅕", true},
		{"一", true},
	}

	for _, tc := range testCases {
		got := Numeric(tc.input)
		if got != tc.want {
			t.Errorf("Numeric(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

// TestDecimal tests the Decimal function.
func TestDecimal(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"0", true},
		{"1234", true},
		{"Ⅳ", false},
		{"1234abc", false},
		{"", false},
		{"1.23", false},
		{"1,23", false},
		{"0000", true},
		{" 1234", false},
		{"1234 ", false},
		{"0b10001", false},
		{"0o21", false},
		{"0x11", false},
		{"0", true},
		{"⅕", false},
		{"一", false},
	}

	for _, tc := range testCases {
		got := Decimal(tc.input)
		if got != tc.want {
			t.Errorf("Decimal(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

// TestFloat tests the Float function.
func TestFloat(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"0", true},
		{"123", true},
		{"123.456", true},
		{"+123.456", true},
		{"-123.456", true},
		{"123.", true},
		{".456", true},
		{"0.0", true},
		{"-.456", true},
		{"+.456", true},
		{"123.456.789", false},
		{"abc", false},
		{"123a", false},
		{"", false},
		{".", false},
		{"123,456", false},
		{"123.456e7", false}, // exponential notation not supported
		{"-0", true},
		{"-0.0", true},
		{"--123.456", false},
		{"123..456", false},
		{"12.34.56", false},
		{" 123.456", false},
		{"123.456 ", false},
		{"123.456.789.0", false},
		{"+", false},
		{"-", false},
	}

	for _, tc := range testCases {
		got := Float(tc.input)
		if got != tc.want {
			t.Errorf("Float(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

// TestAlpha tests the Alpha function.
func TestAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Київ", true},
		{"abc", true},
		{"abc1", false},
		{"abc!", false},
		{"abcΔ", true},
		{"", false},
		{"123", false},
		{"!@#$%^&*", false},
	}

	for _, test := range tests {
		result := Alpha(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t",
				test.input, test.expected, result)
		}
	}
}

// TestAlnum tests the Alnum function.
func TestAlnum(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Дніпро1", true},
		{"abc123", true},
		{"abc", true},
		{"123", true},
		{"abc!", false},
		{"abcΔ", true},
		{"", false},
		{"!@#$%^&*", false},
	}

	for _, test := range tests {
		result := Alnum(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t",
				test.input, test.expected, result)
		}
	}
}

// TestLower tests the Lower function.
func TestLower(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},       // contains lowercase letters only
		{"Abc", false},      // contains an uppercase letter
		{"abc123", false},   // contains a number
		{"abc!", false},     // contains a special character
		{"", false},         // empty string
		{"ABC", false},      // contains uppercase letters only
		{"abcABC", false},   // contains both lowercase and uppercase letters
		{"abc ", false},     // contains whitespace
		{"abc\txyz", false}, // contains whitespace
		{"\nabc", false},    // contains whitespace
	}

	for _, test := range tests {
		result := Lower(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t",
				test.input, test.expected, result)
		}
	}
}

// TestUpper tests the Upper function.
func TestUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},       // contains uppercase letters only
		{"AbC", false},      // contains a lowercase letter
		{"ABC123", false},   // contains a number
		{"ABC!", false},     // contains a special character
		{"", false},         // empty string
		{"abc", false},      // contains lowercase letters only
		{"ABCabc", false},   // contains both uppercase and lowercase letters
		{"ABC ", false},     // contains whitespace
		{"ABC\txyz", false}, // contains whitespace
		{"\nABC", false},    // contains whitespace
	}

	for _, test := range tests {
		result := Upper(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t",
				test.input, test.expected, result)
		}
	}
}

// TestTitle tests the Title function.
func TestTitle(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello World", true},
		{"Hello world", false},
		{"HELLO WORLD", false},
		{"hELLO wORLD", false},
		{"", false},
		{"Hello World!", true},
		{"Hello123 World", true},
		{"Hello-World", true},
		{"Hello_World", true},
		{"Hello, World", true},
		{"Hello	World", true},
		{"Hello\nWorld", true},
		{"Hello.World", true},
		{"Hello:World", true},
		{"Hello;World", true},
		{"Hello?World", true},
		{"Hello!World", true},
		{"Hello( )World", true},
		{"Hello( )World", true},
		{"Hello[ ]World", true},
		{"Hello< >World", true},
	}

	for _, test := range tests {
		result := Title(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t",
				test.input, test.expected, result)
		}
	}
}

// TestSpace tests the Space function.
func TestSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{" \t\n", true},      // whitespace characters
		{"Hello", false},     // non-whitespace characters
		{" ", true},          // single space
		{"\n\t ", true},      // whitespace characters with newline and tab
		{"", false},          // empty string
		{"   ", true},        // multiple spaces
		{"\t\t\t", true},     // multiple tabs
		{"\n\n\n", true},     // multiple newlines
		{"\t \n", true},      // combination of whitespace characters
		{"\tHello\n", false}, // non-whitespace character in the middle
	}

	for _, test := range tests {
		result := Space(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t",
				test.input, test.expected, result)
		}
	}
}
