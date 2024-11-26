package is

import (
	"testing"
	"unicode"
)

// TestVariableName tests the VariableName function.
func TestVariableName(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		strict bool
		want   bool
	}{
		{
			name: "Standard variable name",
			in:   "myVar",
			want: true,
		},
		{
			name: "Variable name starting with underscore",
			in:   "_myVar",
			want: true,
		},
		{
			name: "Variable name with underscore",
			in:   "my_var",
			want: true,
		},
		{
			name: "Variable name with underscore and number",
			in:   "my_var_2",
			want: true,
		},
		{
			name: "Starts with a number",
			in:   "2myVar",
			want: false,
		},
		{
			name: "Contains space",
			in:   "my var",
			want: false,
		},
		{
			name: "Contains hyphen",
			in:   "my-var",
			want: false,
		},
		{
			name: "Contains special character",
			in:   "my@var",
			want: false,
		},
		{
			name: "Empty string",
			in:   "",
			want: false,
		},
		{
			name: "String with spaces",
			in:   "   ",
			want: false,
		},
		{
			name: "Starts with a number, has underscore",
			in:   "_2",
			want: true,
		},
		{
			name: "Contains period",
			in:   "my.var",
			want: false,
		},
		{
			name: "Reserved keyword in Go but strong false",
			in:   "return",
			want: true,
		},
		{
			name:   "Reserved keyword in Go",
			in:     "return",
			strict: true,
			want:   false,
		},
		{
			name: "Reserved keyword in Python",
			in:   "while",
			want: true,
		},
		{
			name:   "Reserved keyword in Python but strong false",
			in:     "while",
			strict: true,
			want:   false,
		},
		{
			name: "Capitalized variable name",
			in:   "MyVar",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Var(tt.in, tt.strict); got != tt.want {
				t.Errorf("IsVarName() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestVar tests the Var function as synonym for VariableName.
func TestVar(t *testing.T) {
	// Test one case to ensure Var works the same as VariableName.
	got := Var("myVar", true)
	want := true

	if got != want {
		t.Errorf("Var(%q, true) = %v; want %v",
			"myVar", got, want)
	}
}

// TestVariableNameFor tests the VariableNameFor function.
func TestVariableNameFor(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{
			name:     "Valid - Go, Non-reserved",
			v:        "myVar",
			language: "go",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Go, Reserved",
			v:        "func",
			language: "go",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Valid - Python, Non-reserved",
			v:        "myVar",
			language: "python",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Python, Reserved",
			v:        "return",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Valid - Java, Non-reserved",
			v:        "myVar",
			language: "java",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Java, Reserved",
			v:        "class",
			language: "java",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Valid - Ruby method with ?",
			v:        "valid?",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - Ruby method with !",
			v:        "save!",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Ruby method with ? in middle",
			v:        "va?lid",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Ruby method with ! in middle",
			v:        "sa!ve",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Valid - Ruby method with _ and ?",
			v:        "is_valid?",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - Ruby instance variable",
			v:        "@name",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - Ruby class variable",
			v:        "@@count",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - Ruby global variable",
			v:        "$global",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - JavaScript (js alias)",
			v:        "myVar",
			language: "js",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - Python (py alias)",
			v:        "myVar",
			language: "py",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Valid - SQL non-reserved",
			v:        "mySelect",
			language: "sql",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - SQL reserved",
			v:        "SELECT",
			language: "sql",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Special characters",
			v:        "my@Var",
			language: "go",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Valid - Python Unicode",
			v:        "μεταβλητή",
			language: "python",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Empty string",
			v:        "",
			language: "go",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Invalid - Unsupported Language",
			v:        "myVar",
			language: "unknown",
			want:     false,
			wantErr:  ErrLanguageNotSupported("unknown"),
		},
		{
			name:     "Invalid - Empty Language",
			v:        "myVar",
			language: "",
			want:     false,
			wantErr:  ErrLanguageNotSupported(""),
		},
		{
			name:     "Valid - PHP variable with $",
			v:        "$myVar",
			language: "php",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Invalid - PHP variable without $",
			v:        "myVar",
			language: "php",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby valid instance variable",
			v:        "@var",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Ruby valid class variable",
			v:        "@@var",
			language: "ruby",
			want:     true,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf(
					"VariableNameFor(%q, %q) error = %v; wantErr %v",
					tt.v,
					tt.language,
					err,
					tt.wantErr,
				)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf(
						"VariableNameFor(%q, %q) error = %v; wantErr %v",
						tt.v,
						tt.language,
						err,
						tt.wantErr,
					)
					return
				}
			}

			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v; want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestVarFor tests the VarFor function as synonym for VariableNameFor.
func TestVarFor(t *testing.T) {
	// Test one case to ensure VarFor works the same as VariableNameFor.
	got, err := VarFor("myVar", "go")
	want := true
	wantErr := error(nil)

	if err != wantErr {
		t.Errorf("VarFor(%q, %q, true) error = %v; want %v",
			"myVar", "go", err, wantErr)
	}

	if got != want {
		t.Errorf("VarFor(%q, %q, true) = %v; want %v",
			"myVar", "go", got, want)
	}
}

// TestLanguageAliases tests that language aliases work correctly.
func TestLanguageAliases(t *testing.T) {
	aliases := map[string]string{
		"js":     "javascript",
		"py":     "python",
		"rb":     "ruby",
		"golang": "go",
	}

	for alias, full := range aliases {
		t.Run(alias, func(t *testing.T) {
			gotAlias, errAlias := VariableNameFor("myVar", alias)
			gotFull, errFull := VariableNameFor("myVar", full)

			if errAlias != errFull {
				t.Errorf(
					"Different errors for alias %q and full name %q: %v vs %v",
					alias, full, errAlias, errFull)
			}

			if gotAlias != gotFull {
				t.Errorf(
					"Different results for alias %q and full name %q: %v vs %v",
					alias, full, gotAlias, gotFull)
			}
		})
	}
}

// TestCaseSensitivity tests case sensitivity handling for different languages.
func TestCaseSensitivity(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
	}{
		{"SQL case insensitive - lowercase", "select", "sql", false},
		{"SQL case insensitive - uppercase", "SELECT", "sql", false},
		{"SQL case insensitive - mixed", "SeLeCt", "sql", false},
		{"Python case sensitive", "Print", "python", true},
		{"Python case sensitive reserved", "print", "python", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v; want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestSelectorName tests the SelectorName function.
func TestSelectorName(t *testing.T) {
	tests := []struct {
		name   string
		v      string
		strict bool
		want   bool
	}{
		{"Valid - Element", "div", false, true},
		{"Valid - Class", ".myClass", false, true},
		{"Valid - ID", "#myID", false, true},
		{"Invalid - Complex Selector", "div .myClass", false, false},
		{"Invalid - Complex Selector", "#myID .myClass", false, false},
		{"Invalid - Complex Selector", "div > .myClass", false, false},
		{"Invalid - Empty String", "", false, false},
		{"Invalid - Complex Name", "div.myClass", false, false},
		{"Invalid - Strict Mode with ID", "#myID", true, false},
		{"Invalid - Strict Mode with Class", ".myClass", true, false},
		{"Valid - Strict Mode without Prefix", "myID", true, true},
		{"Valid - Strict Mode without Prefix", "myClass", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectorName(tt.v, tt.strict)
			if got != tt.want {
				t.Errorf("SelectorName(%q, strict=%v) = %v; want %v",
					tt.v, tt.strict, got, tt.want)
			}
		})
	}
}

// TestSel tests the Sel function as synonym for SelectorName.
func TestSel(t *testing.T) {
	tests := []struct {
		name   string
		v      string
		strict bool
		want   bool
	}{
		{
			name:   "Valid Element Selector",
			v:      "div",
			strict: false,
			want:   true,
		},
		{
			name:   "Valid Class Selector",
			v:      ".myClass",
			strict: false,
			want:   true,
		},
		{
			name:   "Valid ID Selector",
			v:      "#myID",
			strict: false,
			want:   true,
		},
		{
			name:   "Invalid Complex Selector",
			v:      "div .myClass",
			strict: false,
			want:   false,
		},
		{
			name:   "Invalid Complex Selector",
			v:      "#myID .myClass",
			strict: false,
			want:   false,
		},
		{
			name:   "Invalid Complex Selector",
			v:      "div> .myClass",
			strict: false,
			want:   false,
		},
		{
			name:   "Invalid Empty Selector",
			v:      "",
			strict: false,
			want:   false,
		},
		{
			name:   "Valid Strict Mode Selector",
			v:      "myID",
			strict: true,
			want:   true,
		},
		{
			name:   "Invalid Strict Mode with Class Prefix",
			v:      ".myClass",
			strict: true,
			want:   false,
		},
		{
			name:   "Invalid Strict Mode with ID Prefix",
			v:      "#myID",
			strict: true,
			want:   false,
		},
		{
			name:   "Reserved CSS Selector",
			v:      "inherit",
			strict: true,
			want:   false, // 'inherit' is a reserved CSS keyword
		},
		{
			name:   "Non-Reserved CSS Selector",
			v:      "myCustomClass",
			strict: true,
			want:   true,
		},
		{
			name:   "Reserved CSS Selector with Hash",
			v:      "#inherit",
			strict: true,
			want:   false,
		},
		{
			name:   "Reserved CSS Selector with Dot",
			v:      ".inherit",
			strict: true,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var strict []bool
			if tt.strict {
				strict = []bool{true}
			}
			got := Sel(tt.v, strict...)
			if got != tt.want {
				t.Errorf("Sel(%q, strict=%v) = %v; want %v",
					tt.v, tt.strict, got, tt.want)
			}
		})
	}
}

// TestLanguageSpecificCases tests cases specific to certain languages.
func TestLanguageSpecificCases(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		// TypeScript
		{
			name:     "TypeScript - valid with $",
			v:        "$myVar",
			language: "typescript",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "TypeScript - unicode allowed",
			v:        "змінна",
			language: "typescript",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "TypeScript - starts with number",
			v:        "1var",
			language: "typescript",
			want:     false,
			wantErr:  nil,
		},

		// Ruby
		{
			name:     "Ruby - empty after @@",
			v:        "@@",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - invalid after @@",
			v:        "@@1var",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - empty after @",
			v:        "@",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - invalid after @",
			v:        "@1var",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - empty after $",
			v:        "$",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - invalid after $",
			v:        "$1var",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - question mark in middle",
			v:        "is?valid",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Ruby - exclamation mark in middle",
			v:        "save!now",
			language: "ruby",
			want:     false,
			wantErr:  nil,
		},

		// Perl
		{
			name:     "Perl - unicode not allowed",
			v:        "変数",
			language: "perl",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Perl - valid variable",
			v:        "var",
			language: "perl",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Perl - invalid with $",
			v:        "$var",
			language: "perl",
			want:     false,
			wantErr:  nil,
		},

		// Lua
		{
			name:     "Lua - unicode not allowed",
			v:        "змінна",
			language: "lua",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Lua - valid underscore start",
			v:        "_test",
			language: "lua",
			want:     true,
			wantErr:  nil,
		},

		// Rust
		{
			name:     "Rust - unicode not allowed",
			v:        "変数",
			language: "rust",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Rust - valid underscore",
			v:        "_rust_var",
			language: "rust",
			want:     true,
			wantErr:  nil,
		},

		// C++
		{
			name:     "C++ - unicode not allowed",
			v:        "変数",
			language: "cpp",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "C++ - valid identifier",
			v:        "_cpp_var",
			language: "cpp",
			want:     true,
			wantErr:  nil,
		},

		// C
		{
			name:     "C - unicode not allowed",
			v:        "変数",
			language: "c",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "C - valid identifier",
			v:        "_c_var",
			language: "c",
			want:     true,
			wantErr:  nil,
		},

		// Kotlin
		{
			name:     "Kotlin - unicode allowed",
			v:        "переменная",
			language: "kotlin",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Kotlin - valid identifier",
			v:        "_kotlinVar",
			language: "kotlin",
			want:     true,
			wantErr:  nil,
		},

		// Swift
		{
			name:     "Swift - unicode allowed",
			v:        "変数",
			language: "swift",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Swift - valid identifier",
			v:        "_swiftVar",
			language: "swift",
			want:     true,
			wantErr:  nil,
		},

		// HTML
		{
			name:     "HTML - unicode not allowed",
			v:        "変数",
			language: "html",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "HTML - valid with hyphen",
			v:        "custom-element",
			language: "html",
			want:     true,
			wantErr:  nil,
		},

		// CSS
		{
			name:     "CSS - unicode not allowed",
			v:        "変数",
			language: "css",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "CSS - valid with hyphen and @",
			v:        "@my-variable",
			language: "css",
			want:     true,
			wantErr:  nil,
		},

		// R
		{
			name:     "R - unicode allowed",
			v:        "変数",
			language: "r",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "R - valid with dot",
			v:        ".my.var",
			language: "r",
			want:     true,
			wantErr:  nil,
		},

		// MATLAB
		{
			name:     "MATLAB - unicode not allowed",
			v:        "変数",
			language: "matlab",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "MATLAB - must start with letter",
			v:        "_var",
			language: "matlab",
			want:     false,
			wantErr:  nil,
		},

		// Haskell
		{
			name:     "Haskell - must start with lowercase",
			v:        "Var",
			language: "haskell",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Haskell - valid with apostrophe",
			v:        "var'",
			language: "haskell",
			want:     true,
			wantErr:  nil,
		},

		// Scala
		{
			name:     "Scala - unicode allowed",
			v:        "変数",
			language: "scala",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Scala - valid identifier",
			v:        "_scalaVar",
			language: "scala",
			want:     true,
			wantErr:  nil,
		},

		// PHP
		{
			name:     "PHP - $ only",
			v:        "$",
			language: "php",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "PHP - invalid after $",
			v:        "$1var",
			language: "php",
			want:     false,
			wantErr:  nil,
		},

		// Unknown
		{
			name:     "Unknown language - default config",
			v:        "myVar",
			language: "unknown",
			want:     false,
			wantErr:  ErrLanguageNotSupported("unknown"),
		},

		// Aliases
		{
			name:     "C++ alias - cxx",
			v:        "myVar",
			language: "cxx",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "C++ alias - hpp",
			v:        "myVar",
			language: "hpp",
			want:     true,
			wantErr:  nil,
		},
		{
			name:     "Go alias - golang",
			v:        "myVar",
			language: "golang",
			want:     true,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}

			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("VariableNameFor() error = %v, wantErr %v",
						err, tt.wantErr)
					return
				}
			}

			if got != tt.want {
				t.Errorf("VariableNameFor(%s, %s) = %v, want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestRubySpecificCases tests Ruby-specific cases.
func TestRubySpecificCases(t *testing.T) {
	tests := []struct {
		name    string
		v       string
		want    bool
		wantErr error
	}{
		{
			name:    "Ruby - starts with letter",
			v:       "variable",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - starts with underscore",
			v:       "_variable",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - starts with @",
			v:       "@variable",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - starts with @@",
			v:       "@@variable",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - starts with $",
			v:       "$variable",
			want:    true,
			wantErr: nil,
		},

		{
			name:    "Ruby - @ with number",
			v:       "@123",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - @@ with number",
			v:       "@@123",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - $ with number",
			v:       "$123",
			want:    false,
			wantErr: nil,
		},

		{
			name:    "Ruby - unicode variable",
			v:       "перемінна",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - unicode with @",
			v:       "@перемінна",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - unicode with @@",
			v:       "@@перемінна",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - unicode with $",
			v:       "$перемінна",
			want:    true,
			wantErr: nil,
		},

		{
			name:    "Ruby - method with ?",
			v:       "empty?",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - method with !",
			v:       "delete!",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - invalid ? in middle",
			v:       "is?valid",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - invalid ! in middle",
			v:       "delete!all",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - both ? and !",
			v:       "valid?!",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - ! before ?",
			v:       "valid!?",
			want:    false,
			wantErr: nil,
		},

		{
			name:    "Ruby - empty string",
			v:       "",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - single @",
			v:       "@",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - single @@",
			v:       "@@",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - single $",
			v:       "$",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - single ?",
			v:       "?",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Ruby - single !",
			v:       "!",
			want:    false,
			wantErr: nil,
		},

		{
			name:    "Ruby - @ with ?",
			v:       "@empty?",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - @@ with ?",
			v:       "@@empty?",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - $ with ?",
			v:       "$empty?",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - @ with !",
			v:       "@save!",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - @@ with !",
			v:       "@@save!",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - $ with !",
			v:       "$save!",
			want:    true,
			wantErr: nil,
		},

		{
			name:    "Ruby - valid with numbers",
			v:       "var123",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - valid @ with numbers",
			v:       "@var123",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - valid @@ with numbers",
			v:       "@@var123",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Ruby - valid $ with numbers",
			v:       "$var123",
			want:    true,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, "ruby")
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf(
					"VariableNameFor(%q, ruby) error = %v, wantErr %v",
					tt.v, err, tt.wantErr,
				)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf(
						"VariableNameFor(%q, ruby) error = %v, wantErr %v",
						tt.v, err, tt.wantErr,
					)
					return
				}
			}
			if got != tt.want {
				t.Errorf(
					"VariableNameFor(%q, ruby) = %v, want %v",
					tt.v, got, tt.want,
				)
			}
		})
	}
}

// TestMATLABSpecificCases tests MATLAB-specific cases.
func TestMATLABSpecificCases(t *testing.T) {
	tests := []struct {
		name    string
		v       string
		want    bool
		wantErr error
	}{
		{
			name:    "MATLAB - starts with letter lowercase",
			v:       "var",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - starts with letter uppercase",
			v:       "Var",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - starts with underscore",
			v:       "_var",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - starts with number",
			v:       "1var",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - letters only",
			v:       "variable",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - letters and numbers",
			v:       "var123",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - letters and underscore",
			v:       "var_name",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - letters, numbers and underscore",
			v:       "var_123_name",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - invalid hyphen",
			v:       "var-name",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - invalid dot",
			v:       "var.name",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - invalid space",
			v:       "var name",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - unicode letter",
			v:       "변수",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - unicode mixed with ASCII",
			v:       "var변수",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - unicode number",
			v:       "var١٢٣",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - single letter",
			v:       "x",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - empty string",
			v:       "",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - only underscore",
			v:       "_",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - only number",
			v:       "123",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - uppercase",
			v:       "VAR",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - mixed case",
			v:       "varNAME",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "MATLAB - special characters",
			v:       "var@name",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - brackets",
			v:       "var[1]",
			want:    false,
			wantErr: nil,
		},
		{
			name:    "MATLAB - parentheses",
			v:       "var()",
			want:    false,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, "matlab")
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf(
					"VariableNameFor(%q, matlab) error = %v, wantErr %v",
					tt.v, err, tt.wantErr,
				)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf(
						"VariableNameFor(%q, matlab) error = %v, wantErr %v",
						tt.v, err, tt.wantErr,
					)
					return
				}
			}
			if got != tt.want {
				t.Errorf(
					"VariableNameFor(%q, matlab) = %v, want %v",
					tt.v, got, tt.want,
				)
			}
		})
	}
}

// TestIsValidIdentifier tests the isValidIdentifier function.
func TestIsValidIdentifier(t *testing.T) {
	tests := []struct {
		name   string
		v      string
		config languageConfig
		want   bool
	}{
		{
			name: "Basic valid identifier",
			v:    "myVar",
			config: languageConfig{
				language:      "unknown",
				caseSensitive: true,
				allowUnicode:  false,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) || unicode.IsNumber(r)
				},
			},
			want: true,
		},
		{
			name: "Invalid first character",
			v:    "1var",
			config: languageConfig{
				language:      "unknown",
				caseSensitive: true,
				allowUnicode:  false,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) || unicode.IsNumber(r)
				},
			},
			want: false,
		},
		{
			name: "Ruby valid instance variable",
			v:    "@var",
			config: languageConfig{
				language:      "ruby",
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r) || r == '_' || r == '@' || r == '$'
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
				},
			},
			want: true,
		},
		{
			name: "Ruby valid class variable",
			v:    "@@var",
			config: languageConfig{
				language:      "ruby",
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r) || r == '_' || r == '@' || r == '$'
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
				},
			},
			want: true,
		},
		{
			name: "Unicode allowed",
			v:    "变量",
			config: languageConfig{
				language:      "unknown",
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r)
				},
			},
			want: true,
		},
		{
			name: "Unicode not allowed",
			v:    "变量",
			config: languageConfig{
				language:      "typescript",
				caseSensitive: true,
				allowUnicode:  false,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r)
				},
			},
			want: false,
		},
		{
			name: "Empty string",
			v:    "",
			config: languageConfig{
				language:      "unknown",
				caseSensitive: true,
				allowUnicode:  false,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r)
				},
			},
			want: false,
		},
		{
			name: "Special chars allowed",
			v:    "$var",
			config: languageConfig{
				language:      "ruby",
				caseSensitive: true,
				allowUnicode:  false,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r) || r == '$'
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) || r == '$'
				},
			},
			want: true,
		},
		{
			name: "Special chars not allowed",
			v:    "$var",
			config: languageConfig{
				language:      "ruby",
				caseSensitive: true,
				allowUnicode:  false,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r)
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidIdentifier(tt.v, tt.config); got != tt.want {
				t.Errorf("isValidIdentifier(%q) = %v, want %v",
					tt.v, got, tt.want)
			}
		})
	}
}

// TestRemainingCharactersValidation tests the validation
// of characters after the first in isValidIdentifier.
func TestRemainingCharactersValidation(t *testing.T) {
	tests := []struct {
		name   string
		v      string
		config languageConfig
		want   bool
	}{
		{
			name: "Valid characters after first",
			v:    "a123_xyz",
			config: languageConfig{
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) ||
						unicode.IsNumber(r) || r == '_'
				},
			},
			want: true,
		},
		{
			name: "Invalid character in the middle",
			v:    "abc@def",
			config: languageConfig{
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					// Explicitly specify that @ is not allowed.
					c := unicode.IsLetter(r) ||
						unicode.IsNumber(r) || r == '_'
					return c && r != '@'
				},
			},
			want: false,
		},
		{
			name: "Invalid character at the end",
			v:    "abc#",
			config: languageConfig{
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					// Explicitly specify that # is not allowed.
					c := unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
					return c && r != '#'
				},
			},
			want: false,
		},
		{
			name: "Only first character",
			v:    "a",
			config: languageConfig{
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) ||
						unicode.IsNumber(r) || r == '_'
				},
			},
			want: true,
		},

		// Additional test cases for special characters.
		{
			name: "Contains space",
			v:    "abc def",
			config: languageConfig{
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) ||
						unicode.IsNumber(r) || r == '_'
				},
			},
			want: false,
		},
		{
			name: "Contains hyphen",
			v:    "abc-def",
			config: languageConfig{
				caseSensitive: true,
				allowUnicode:  true,
				checkFirst: func(r rune) bool {
					return unicode.IsLetter(r)
				},
				validChars: func(r rune) bool {
					return unicode.IsLetter(r) ||
						unicode.IsNumber(r) || r == '_'
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidIdentifier(tt.v, tt.config); got != tt.want {
				t.Errorf("isValidIdentifier(%q) = %v, want %v",
					tt.v, got, tt.want)
			}
		})
	}
}

// TestDefaultLanguageConfig tests behavior with missing language configuration.
func TestDefaultLanguageConfig(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{
			name:     "Non-existent language - valid identifier",
			v:        "myVar",
			language: "nonexistent",
			want:     false,
			wantErr:  ErrLanguageNotSupported("nonexistent"),
		},
		{
			name:     "Non-existent language - invalid start",
			v:        "1var",
			language: "nonexistent",
			want:     false,
			wantErr:  ErrLanguageNotSupported("nonexistent"),
		},
		{
			name:     "Non-existent language - empty string",
			v:        "",
			language: "nonexistent",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Non-existent language - special chars",
			v:        "var@name",
			language: "nonexistent",
			want:     false,
			wantErr:  ErrLanguageNotSupported("nonexistent"),
		},
		{
			name:     "Non-existent language - unicode",
			v:        "변수",
			language: "nonexistent",
			want:     false,
			wantErr:  ErrLanguageNotSupported("nonexistent"),
		},
		{
			name:     "Empty language name",
			v:        "myVar",
			language: "",
			want:     false,
			wantErr:  ErrLanguageNotSupported(""),
		},
		{
			name:     "Whitespace language name",
			v:        "myVar",
			language: "   ",
			want:     false,
			wantErr:  ErrLanguageNotSupported("   "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor(%q, %q) error = %v, wantErr %v",
					tt.v, tt.language, err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("VariableNameFor(%q, %q) error = %v, wantErr %v",
						tt.v, tt.language, err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v, want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestEmptyRunesHandling tests handling of empty strings.
func TestEmptyRunesHandling(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{
			name:     "Empty string",
			v:        "",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Only spaces",
			v:        "   ",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Only tabs",
			v:        "\t\t",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Only newlines",
			v:        "\n\n",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Mixed whitespace",
			v:        " \t\n",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Zero width space",
			v:        "\u200B",
			language: "python",
			want:     false,
			wantErr:  nil,
		},
		{
			name:     "Empty string with different language",
			v:        "",
			language: "javascript",
			want:     false,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor(%q, %q) error = %v, wantErr %v",
					tt.v, tt.language, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v, want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestIsLanguageSupportedWithAliases checks if the isLanguageSupported
// function returns the expected results when using aliases.
func TestIsLanguageSupportedWithAliases(t *testing.T) {
	tests := []struct {
		alias    string
		expected bool
	}{
		{"js", true},
		{"jsx", true},
		{"ts", true},
		{"py", true},
		{"rb", true},
		{"golang", true},
		{"c++", true},
		{"cxx", true},
		{"hpp", true},
		{"h", true},
		{"unknown", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.alias, func(t *testing.T) {
			supported := isLanguageSupported(tt.alias)
			if supported != tt.expected {
				t.Errorf("isLanguageSupported(%q) = %v; want %v",
					tt.alias, supported, tt.expected)
			}
		})
	}
}

// TestIsValidIdentifier_EmptyString checks the handling of empty strings.
func TestIsValidIdentifier_EmptyString(t *testing.T) {
	config := languageConfig{
		language:      "python",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	}

	if got := isValidIdentifier("", config); got {
		t.Errorf("isValidIdentifier(\"\") = %v; want false", got)
	}
}

// TestIsValidIdentifier_Ruby_InvalidFirstCharAfterPrefix checks for
// cases where the first char after the prefix is ​​invalid.
func TestIsValidIdentifier_Ruby_InvalidFirstCharAfterPrefix(t *testing.T) {
	tests := []struct {
		name string
		v    string
		want bool
	}{
		{"Ruby - @ with non-letter", "@1var", false},
		{"Ruby - @@ with non-letter", "@@1var", false},
		{"Ruby - $ with non-letter", "$1var", false},
	}

	config := languageConfig{
		language:      "ruby",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '@' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidIdentifier(tt.v, config)
			if got != tt.want {
				t.Errorf("isValidIdentifier(%q) = %v; want %v",
					tt.v, got, tt.want)
			}
		})
	}
}

// TestIsValidIdentifier_Ruby_InvalidMiddleChars checks
// that ? and ! can only be at the end.
func TestIsValidIdentifier_Ruby_InvalidMiddleChars(t *testing.T) {
	tests := []struct {
		name string
		v    string
		want bool
	}{
		{"Ruby - ? in middle", "is?valid", false},
		{"Ruby - ! in middle", "save!now", false},
		{"Ruby - ! and ? not at end", "valid?!", false},
		{"Ruby - ! before ?", "valid!?", false},
	}

	config := languageConfig{
		language:      "ruby",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '@' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidIdentifier(tt.v, config)
			if got != tt.want {
				t.Errorf("isValidIdentifier(%q) = %v; want %v",
					tt.v, got, tt.want)
			}
		})
	}
}

// TestIsValidIdentifier_SpecialCharacters checks the
// handling of special characters.
func TestIsValidIdentifier_SpecialCharacters(t *testing.T) {
	tests := []struct {
		name string
		v    string
		want bool
	}{
		{"Space in variable name", "my var", false},
		{"Hash in variable name", "my#var", false},
		{"Hyphen in variable name", "my-var", false},
		{"At symbol not allowed without prefix", "my@var", false},
		{"Underscore is allowed", "my_var", true},
	}

	config := languageConfig{
		language:      "unknown",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidIdentifier(tt.v, config)
			if got != tt.want {
				t.Errorf("isValidIdentifier(%q) = %v; want %v",
					tt.v, got, tt.want)
			}
		})
	}
}

// TestIsValidIdentifier_SpecialHandlingAt checks the handling
// of the @ character when it is not allowed.
func TestIsValidIdentifier_SpecialHandlingAt(t *testing.T) {
	tests := []struct {
		name string
		v    string
		want bool
	}{
		{"At symbol not allowed", "@var", false},
		{"At symbol in the middle", "var@name", false},
	}

	config := languageConfig{
		language:      "unknown",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidIdentifier(tt.v, config)
			if got != tt.want {
				t.Errorf("isValidIdentifier(%q) = %v; want %v",
					tt.v, got, tt.want)
			}
		})
	}
}

// TestVariableNameFor_DefaultConfig checks to use the default
// configuration when the language is not supported.
func TestVariableNameFor_DefaultConfig(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{
			"Default config - valid identifier", "myVar", "unknown",
			false, ErrLanguageNotSupported("unknown"),
		},
		{
			"Default config - invalid identifier", "1var", "unknown",
			false, ErrLanguageNotSupported("unknown"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v",
					tt.v, tt.language, err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v",
						tt.v, tt.language, err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v; want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestVariableNameFor_Ruby_SpecialChars tests character
// handling ? and ! in Ruby.
func TestVariableNameFor_Ruby_SpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{"Ruby - method with ?", "empty?", "ruby", true, nil},
		{"Ruby - method with !", "delete!", "ruby", true, nil},
		{"Ruby - method with ? in middle", "is?valid", "ruby", false, nil},
		{"Ruby - method with ! in middle", "save!now", "ruby", false, nil},
		{"Ruby - both ? and !", "valid?!", "ruby", false, nil},
		{"Ruby - ! before ?", "valid!?", "ruby", false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v",
					tt.v, tt.language, err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v",
						tt.v, tt.language, err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v; want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestVariableNameFor_PHP_Prefix tests PHP's handling of the $ prefix.
func TestVariableNameFor_PHP_Prefix(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{"PHP - valid variable with $", "$var", "php", true, nil},
		{"PHP - invalid variable without $", "var", "php", false, nil},
		{"PHP - only $", "$", "php", false, nil},
		{"PHP - $ with invalid character", "$var-name", "php", false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v", tt.v, tt.language, err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v", tt.v, tt.language, err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v; want %v", tt.v, tt.language, got, tt.want)
			}
		})
	}
}

// TestVariableNameFor_SQL_ReservedWords tests reserved words in SQL.
func TestVariableNameFor_SQL_ReservedWords(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		language string
		want     bool
		wantErr  error
	}{
		{"SQL - non-reserved word", "mySelect", "sql", true, nil},
		{"SQL - reserved word lowercase", "select", "sql", false, nil},
		{"SQL - reserved word uppercase", "SELECT", "sql", false, nil},
		{"SQL - reserved word mixed case", "SeLeCt", "sql", false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VariableNameFor(tt.v, tt.language)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v",
					tt.v, tt.language, err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("VariableNameFor(%q, %q) error = %v; wantErr %v",
						tt.v, tt.language, err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("VariableNameFor(%q, %q) = %v; want %v",
					tt.v, tt.language, got, tt.want)
			}
		})
	}
}
