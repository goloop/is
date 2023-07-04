package is

import "unicode"

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

// Var validates if the given string can be used as a variable
// name in most programming languages. The function checks if the
// name starts with a letter or underscore, if it doesn't contain
// any special characters or spaces, and if it's not a reserved
// word. Returns true if the given name is valid, false otherwise.
//
// Example usage:
//
//	// Test a valid variable name
//	fmt.Println(Var("myVar"))  // Output: true
//
//	// Test a name starting with a digit
//	fmt.Println(Var("9myVar"))  // Output: false
//
//	// Test a name containing a space
//	fmt.Println(Var("my Var"))  // Output: false
//
//	// Test a reserved word
//	fmt.Println(Var("return"))  // Output: false
func Var(v string) bool {
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
	if _, ok := varReservedWords[v]; ok {
		return false
	}

	return true
}
