package is

import "testing"

// TestVar tests the Var function.
func TestVar(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		strong bool
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
			strong: true,
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
			strong: true,
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
			if got := Var(tt.in, tt.strong); got != tt.want {
				t.Errorf("IsVarName() = %v, want %v", got, tt.want)
			}
		})
	}
}
