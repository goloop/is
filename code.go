package is

import (
	"regexp"
	"unicode"

	"github.com/goloop/g"
)

var (
	// The selectorStrictRgex ...
	selectorStrictRgex = regexp.MustCompile(`^[a-zA-Z_-][a-zA-Z\d_-]*$`)

	// The selectorRgex ...
	selectorRgex = regexp.MustCompile(`^(#|\.)?[a-zA-Z_-][a-zA-Z\d_-]*$`)
)

// The varReservedWords holds a map of reserved words in Python, Go and C++.
// These words cannot be used as variable names.
var varReservedWords = map[string]struct{}{
	"False": {}, "None": {}, "True": {}, "and": {}, "as": {}, "assert": {},
	"async": {}, "await": {}, "break": {}, "class": {}, "continue": {},
	"def": {}, "del": {}, "elif": {}, "else": {}, "except": {}, "finally": {},
	"for": {}, "from": {}, "global": {}, "if": {}, "import": {}, "in": {},
	"is": {}, "lambda": {}, "nonlocal": {}, "not": {}, "or": {}, "pass": {},
	"raise": {}, "return": {}, "try": {}, "while": {}, "with": {}, "yield": {},
	"default": {}, "func": {}, "interface": {}, "select": {}, "case": {},
	"defer": {}, "go": {}, "map": {}, "struct": {}, "chan": {}, "goto": {},
	"package": {}, "switch": {}, "const": {}, "fallthrough": {}, "range": {},
	"type": {}, "var": {}, "asm": {}, "auto": {}, "bool": {}, "catch": {},
	"char": {}, "const_cast": {}, "delete": {}, "do": {}, "double": {},
	"dynamic_cast": {}, "enum": {}, "explicit": {}, "export": {}, "extern": {},
	"float": {}, "friend": {}, "inline": {}, "int": {}, "long": {},
	"mutable": {}, "namespace": {}, "new": {}, "operator": {}, "private": {},
	"protected": {}, "public": {}, "register": {}, "reinterpret_cast": {},
	"short": {}, "signed": {}, "sizeof": {}, "static": {}, "static_cast": {},
	"template": {}, "this": {}, "throw": {}, "typedef": {}, "typeid": {},
	"typename": {}, "union": {}, "unsigned": {}, "using": {}, "virtual": {},
	"void": {}, "volatile": {}, "wchar_t": {},
}

// VariableName validates if the given string can be used as a variable
// name in most programming languages. The function checks if the
// name starts with a letter or underscore, if it doesn't contain
// any special characters or spaces.
//
// The second optional argument can check if the specified name
// is reserved word (as return, false, if etc.).
//
// Returns true if the given name is valid, false otherwise.
//
// Example usage:
//
//	// Test a valid variable name.
//	fmt.Println(VariableName("myVar"))  // Output: true
//
//	// Test a name starting with a digit.
//	fmt.Println(VariableName("9myVar"))  // Output: false
//
//	// Test a name containing a space.
//	fmt.Println(VariableName("my VariableName"))  // Output: false
//
//	// Test a reserved word.
//	fmt.Println(VariableName("return"))  // Output: true
//
//	// Test a reserved word (strong mode).
//	fmt.Println(VariableName("return", true))  // Output: false
func VariableName(v string, strict ...bool) bool {
	// Empty string is not a valid variable name.
	if v == "" {
		return false
	}

	// Split the string into runes.
	r := []rune(v)

	// Check if the name starts with a letter or underscore.
	if !unicode.IsLetter(r[0]) && r[0] != '_' {
		return false
	}

	// Check if the name doesn't contain any special characters or spaces.
	for _, c := range r {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_' {
			return false
		}
	}

	// Check if the name is not a reserved word.
	if s := g.All(strict...); s {
		if _, ok := varReservedWords[v]; ok {
			return false
		}
	}

	return true
}

// Var is synonym for VariableName function.
func Var(v string, strict ...bool) bool {
	return VariableName(v, strict...)
}

// SelectorName checks if the given string is a valid CSS selector name.
// It supports simple selectors including elements (e.g., "div"), classes
// (e.g., ".myClass"), and IDs (e.g., "#myID"). It does not support
// more complex or combined selectors (e.g., "div .myClass",
// "#myID .myClass", "div > .myClass").
//
// The function removes the first character if it's a '.' or '#' symbol,
// as these are used in CSS syntax but are not part of the actual
// class or ID name. After this, it checks if the string is empty or
// doesn't match the valid CSS selector pattern.
//
// The second optional parameter starts the check in strict mode,
// i.e. the leading characters # and . are inadmissible.
//
// Example Usage:
//
//	is.SelectorName("div")         // Returns: true
//	is.SelectorName(".myClass")    // Returns: true
//	is.SelectorName("#myID")       // Returns: true
//	is.SelectorName("#myID", true) // Returns: false // strict mode
//	is.SelectorName("div.myClass") // Returns: false
//	is.SelectorName("")            // Returns: false
func SelectorName(v string, strict ...bool) bool {
	// Check for empty string.
	if len(v) == 0 {
		return false
	}

	rgex := g.If(g.All(strict...), selectorStrictRgex, selectorRgex)
	return rgex.MatchString(v)
}

// Sel is synonym for SelectorName function.
func Sel(v string, strict ...bool) bool {
	return SelectorName(v, strict...)
}
