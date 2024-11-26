package is

import "testing"

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
			got := SelectorName(tt.v, strict...)
			if got != tt.want {
				t.Errorf("SelectorName(%q, strict=%v) = %v; want %v", tt.v, tt.strict, got, tt.want)
			}
		})
	}
}
