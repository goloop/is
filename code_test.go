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

// TestSelectorName tests the SelectorName function.
func TestSelectorName(t *testing.T) {
	cases := []struct {
		in     string
		strict bool
		want   bool
	}{
		{"div", false, true},
		{".myClass", false, true},
		{"#myID", false, true},
		{"#myID", true, false},
		{"div.myClass", false, false},
		{"#my_id.my-class > a", false, false},
		{"div.myClass", true, false},
		{"", false, false},
		{"invalid@name", false, false},
		{"#123", false, false},
		{".123", false, false},
		{"123", false, false},
	}

	for _, tc := range cases {
		got := Sel(tc.in, tc.strict)
		if got != tc.want {
			t.Errorf("Selector(%q) = %v; want %v", tc.in, got, tc.want)
		}
	}
}
