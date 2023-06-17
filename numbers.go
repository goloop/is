package is

import (
	"github.com/goloop/g"
)

// Even checks if a value is an even number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. If the `f` argument is provided
// and set to true, the function ignores the fractional part
// of the value when checking for evenness. For integer types,
// it checks if the value is divisible by 2 without a remainder.
// For floating-point types, it considers only the integer part
// of the value and determines the parity of the integer part.
// If the value has a non-zero fractional part and `f` is true,
// it returns false since an even number cannot have a fractional part.
//
// Example usage:
//
//	even := is.Even(6)
//	fmt.Println(even)  // Output: true
//
//	odd := is.Even(7)
//	fmt.Println(odd)  // Output: false
//
//	floatingPoint := is.Even(6.6)
//	fmt.Println(floatingPoint)  // Output: false
//
//	floatingPoint = is.Even(6.6, true)
//	fmt.Println(floatingPoint)  // Output: true
func Even[T Numerable](v T, f ...bool) bool {
	return g.IsEven(v, f...)
}

// Odd checks if a value is an odd number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. If the `f` argument is provided
// and set to true, the function ignores the fractional part
// of the value when checking for oddness. For integer types,
// it checks if the value is not divisible by 2 without a remainder.
// For floating-point types, it considers only the integer part
// of the value and determines the parity of the integer part.
// If the value has a non-zero fractional part and `f` is true,
// it returns true since an odd number cannot have a fractional part.
// Otherwise, it returns the negation of the IsEven function.
//
// Example usage:
//
//	odd := is.Odd(7)
//	fmt.Println(odd)  // Output: true
//
//	even := is.Odd(6)
//	fmt.Println(even)  // Output: false
//
//	floatingPoint := is.Odd(7.7)
//	fmt.Println(floatingPoint)  // Output: false
//
//	floatingPoint = is.Odd(7.7, true)
//	fmt.Println(floatingPoint)  // Output: true
func Odd[T Numerable](v T, f ...bool) bool {
	return g.IsOdd(v, f...)
}

// Whole checks if a value is a whole number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. It first checks if the value has
// a non-zero fractional part. If it does, it returns false
// since a whole number cannot have a fractional part.
// If the value does not have a fractional part, it returns true.
//
// Example usage:
//
//	whole := is.Whole(5)
//	fmt.Println(whole)  // Output: true
//
//	notWhole := is.Whole(5.5)
//	fmt.Println(notWhole)  // Output: false
//
//	zero := is.Whole(0)
//	fmt.Println(zero)  // Output: true
//
//	negative := is.Whole(-3)
//	fmt.Println(negative)  // Output: true
func Whole[T Numerable](v T) bool {
	return g.IsWhole(v)
}

// Negative verifies if the numerable value v is less than zero.
// The function uses generic type T, which must satisfy the Numerable
// type constraint. This function is useful for validating the sign of
// numerical data.
//
// Example usage:
//
//	is.Negative(5)      // Output: false
//	is.Negative(-5)     // Output: true
//	is.Negative(0)      // Output: false
//	is.Negative(3.14)   // Output: false
//	is.Negative(-3.14)  // Output: true
//
// Please note that zero is neither positive nor negative. It lies
// between positive and negative numbers on the number line and does
// not change the value of another number when it is added to or
// subtracted from it.
func Negative[T Numerable](v T) bool {
	return v < 0
}

// Positive verifies if the numerable value v is greater than zero.
// The function uses generic type T, which must satisfy the Numerable
// type constraint. This function is useful for validating the sign of
// numerical data.
//
// Example usage:
//
//	is.Positive(5)      // Output: true
//	is.Positive(-5)     // Output: false
//	is.Positive(0)      // Output: false
//	is.Positive(3.14)   // Output: true
//	is.Positive(-3.14)  // Output: false
//
// Please note that zero is neither positive nor negative. It lies
// between positive and negative numbers on the number line and does
// not change the value of another number when it is added to or
// subtracted from it.
func Positive[T Numerable](v T) bool {
	return v > 0
}

// Zero checks if the numerable value v is equal to zero.
// The function uses generic type T, which must satisfy the Numerable
// type constraint. This function is useful for validating whether the
// numerical data equals zero.
//
// Example usage:
//
//	is.Zero(5)      // Output: false
//	is.Zero(-5)     // Output: false
//	is.Zero(0)      // Output: true
//	is.Zero(3.14)   // Output: false
//	is.Zero(-3.14)  // Output: false
//
// Please note that zero is a unique number that is neither positive
// nor negative. It lies between positive and negative numbers on the
// number line and does not change the value of another number when it
// is added to or subtracted from it.
func Zero[T Numerable](v T) bool {
	return v == 0
}

// Natural checks if the numerable value v is a natural number.
// A natural number is a positive integer starting from 1. Zero is not
// considered a natural number. The function uses generic type T, which
// must satisfy the Numerable type constraint.
//
// Example usage:
//
//	is.Natural(5)      // Output: true
//	is.Natural(5.0)    // Output: true
//	is.Natural(-5)     // Output: false
//	is.Natural(0)      // Output: false
//	is.Natural(3.14)   // Output: false
//
// Please note that this function checks if the value is a whole number
// and is positive.
func Natural[T Numerable](v T) bool {
	return Whole(v) && Positive(v)
}
