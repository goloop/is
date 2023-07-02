package is

import "testing"

// TestBase64 tests Base64 function.
func TestBase64(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		strict []bool
		want   bool
	}{
		/*{
			name:   "Valid standard base64",
			input:  "SGVsbG8sIHdvcmxkIQ==",
			strict: []bool{},
			want:   true,
		},
		{
			name:   "Valid standard base64 with strict mode",
			input:  "SGVsbG8sIHdvcmxkIQ==",
			strict: []bool{true},
			want:   true,
		},
		{
			name:   "Valid standard base64 with strict mode as many params",
			input:  "SGVsbG8sIHdvcmxkIQ==",
			strict: []bool{true, true},
			want:   true,
		},
		{
			name:   "Valid standard base64 with safe mode",
			input:  "SGVsbG8sIHdvcmxkIQ==",
			strict: []bool{false},
			want:   true,
		},
		{
			name:   "Valid standard base64 with safe mode as many params",
			input:  "SGVsbG8sIHdvcmxkIQ==",
			strict: []bool{true, true, false}, // any false is false exp.
			want:   true,
		},
		{
			name:   "Not valid standard base64 (missing padding)",
			input:  "SGVsbG8sIHdvcmxkIQ",
			strict: []bool{},
			want:   false,
		},*/
		{
			name:   "Valid base64 with missing padding for safe mode",
			input:  "SGVsbG8sIHdvcmxkIQ",
			strict: []bool{false},
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Base64(test.input, test.strict...)
			if got != test.want {
				t.Errorf("Base64(%q, %v) = %v; want %v",
					test.input, test.strict, got, test.want)
			}
		})
	}
}
