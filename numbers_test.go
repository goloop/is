package is

import (
	"testing"
)

// TestOdd tests the Odd function.
func TestOdd(t *testing.T) {
	// Test cases for integer values.
	// Expected results: true if the value is odd, false otherwise.
	tests := []struct {
		value    int
		expected bool
	}{
		{4, false},   // even number
		{3, true},    // odd number
		{-6, false},  // negative even number
		{-5, true},   // negative odd number
		{-10, false}, // negative even number
	}

	for _, test := range tests {
		result := Odd(test.value)
		if result != test.expected {
			t.Errorf("Odd(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values.
	// Expected results: true if the integer part is odd, false otherwise.
	testsFloat := []struct {
		value    float64
		expected bool
		f        []bool
	}{
		{3.2, false, []bool{}},
		{3.2, true, []bool{true}},
		{3.0, true, []bool{}},
		{-5.5, true, []bool{true}},
		{-6.5, false, []bool{true, false}},
		{-11.0, true, []bool{false}},
	}

	for _, test := range testsFloat {
		result := Odd(test.value, test.f...)
		if result != test.expected {
			t.Errorf("Odd(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}
