package is

import (
	"testing"
)

// TestOdd tests the Odd function.
func TestOdd(t *testing.T) {
	// Test cases for integer values.
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

// TestEven tests the Even function.
func TestEven(t *testing.T) {
	// Test cases for integer values.
	tests := []struct {
		value    float64
		expected bool
	}{
		{4, true},
		{3, false},
		{-6, true},
		{-5, false},
	}

	for _, test := range tests {
		result := Even(test.value)
		if result != test.expected {
			t.Errorf("Even(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values.
	testsFloat := []struct {
		value    float64
		expected bool
		f        []bool
	}{
		{4.2, false, []bool{}},
		{4.2, true, []bool{true}},
		{3.8, false, []bool{true, false}},
		{-5.5, false, []bool{}},
		{-6.5, true, []bool{true}},
	}

	for _, test := range testsFloat {
		result := Even(test.value, test.f...)
		if result != test.expected {
			t.Errorf("Even(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

// TestWhole tests the Whole function.
func TestWhole(t *testing.T) {
	// Test cases for integer values
	tests := []struct {
		value    int
		expected bool
	}{
		{4, true},
		{3, true},
		{-6, true},
		{-5, true},
		{-10, true},
		{0, true},
		{100, true},
		{5, true},
		{5, true},
		{-2, true},
	}

	for _, test := range tests {
		result := Whole(test.value)
		if result != test.expected {
			t.Errorf("Whole(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values
	testsFloat := []struct {
		value    float64
		expected bool
	}{
		{4.2, false},
		{3.8, false},
		{-5.5, false},
		{-11.0, true},
		{0.0, true},
		{100.0, true},
		{5.0, true},
		{-2.0, true},
		{-2.1, false},
	}

	for _, test := range testsFloat {
		result := Whole(test.value)
		if result != test.expected {
			t.Errorf("Whole(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

// TestNegative tests the Negative function.
func TestNegative(t *testing.T) {
	// Test cases for int values.
	testsInt := map[int]bool{
		5:  false,
		-5: true,
		0:  false,
	}

	for k, v := range testsInt {
		result := Negative(k)
		if result != v {
			t.Errorf("Negative(%v): expected %v, but got %v",
				k, v, result)
		}
	}

	// Test cases for float values.
	testsFloat := map[float64]bool{
		3.14:  false,
		-3.14: true,
	}

	for k, v := range testsFloat {
		result := Negative(k)
		if result != v {
			t.Errorf("Negative(%v): expected %v, but got %v",
				k, v, result)
		}
	}
}

// TestPositive tests the Positive function.
func TestPositive(t *testing.T) {
	// Test cases for int values.
	testsInt := map[int]bool{
		5:  true,
		-5: false,
		0:  false,
	}

	for k, v := range testsInt {
		result := Positive(k)
		if result != v {
			t.Errorf("Positive(%v): expected %v, but got %v",
				k, v, result)
		}
	}

	// Test cases for int values.
	testsFloat := map[float64]bool{
		3.14:  true,
		-3.14: false,
	}

	for k, v := range testsFloat {
		result := Positive(k)
		if result != v {
			t.Errorf("Positive(%v): expected %v, but got %v",
				k, v, result)
		}
	}
}

// TestZero tests the Zero function.
func TestZero(t *testing.T) {
	// Test cases for int values.
	testsInt := map[int]bool{
		5:  false,
		-5: false,
		0:  true,
	}

	for k, v := range testsInt {
		result := Zero(k)
		if result != v {
			t.Errorf("Zero(%v): expected %v, but got %v",
				k, v, result)
		}
	}

	// Test cases for float values.
	testsFloat := map[float64]bool{
		0.0:   true,
		3.14:  false,
		-3.14: false,
	}

	for k, v := range testsFloat {
		result := Zero(k)
		if result != v {
			t.Errorf("Zero(%v): expected %v, but got %v",
				k, v, result)
		}
	}
}

// TestNatural tests the Natural function.
func TestNatural(t *testing.T) {
	// Test cases for int values.
	testsInt := map[int]bool{
		5:  true,
		-5: false,
		0:  false,
	}

	for k, v := range testsInt {
		result := Natural(k)
		if result != v {
			t.Errorf("Natural(%v): expected %v, but got %v",
				k, v, result)
		}
	}

	// Test cases for float values.
	testsFloat := map[float64]bool{
		5.0:   true,
		-5.0:  false,
		0.0:   false,
		3.14:  false,
		-3.14: false,
	}

	for k, v := range testsFloat {
		result := Natural(k)
		if result != v {
			t.Errorf("Natural(%v): expected %v, but got %v",
				k, v, result)
		}
	}
}
