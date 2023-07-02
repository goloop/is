package is

import "testing"

// TestBase64 tests Base64 function.
func TestBase64(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid standard base64",
			input: "SGVsbG8sIHdvcmxkIQ==",
			want:  true,
		},
		{
			name:  "Not valid standard base64 (missing padding)",
			input: "SGVsbG8sIHdvcmxkIQ",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Base64(test.input)
			if got != test.want {
				t.Errorf("Base64(%q) = %v; want %v",
					test.input, got, test.want)
			}
		})
	}
}

// TestBase64URL tests Base64URL function.
func TestBase64URL(t *testing.T) {
	validBase64URL := "SGVsbG8sIHdvcmxkIQ"
	invalidBase64URL := "notabase64url"

	validResult := Base64URL(validBase64URL)
	if !validResult {
		t.Errorf("Expected valid Base64URL to return true, got false")
	}

	invalidResult := Base64URL(invalidBase64URL)
	if invalidResult {
		t.Errorf("Expected invalid Base64URL to return false, got true")
	}
}

// TestHex tests Hex function.
func TestHex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid hex string",
			input: "deadbeef",
			want:  true,
		},
		{
			name:  "Valid hex string with uppercase",
			input: "DEADBEEF",
			want:  true,
		},
		{
			name:  "Valid hex string with 0x prefix",
			input: "0xdeadbeef",
			want:  true,
		},
		{
			name:  "Invalid hex string",
			input: "deadbefg",
			want:  false,
		},
		{
			name:  "Invalid hex string with 0x prefix",
			input: "0xdeadbefg",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Hex(test.input)
			if got != test.want {
				t.Errorf("Hex(%q) = %v; want %v",
					test.input, got, test.want)
			}
		})
	}
}

// TestHexColor tests HexColor function.
func TestHexColor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid hex color",
			input: "#ff7377",
			want:  true,
		},
		{
			name:  "Valid hex color without # prefix",
			input: "ff7377",
			want:  true,
		},
		{
			name:  "Valid hex color with uppercase",
			input: "#FF7377",
			want:  true,
		},
		{
			name:  "Valid hex color but with 0x prefix",
			input: "0xff7377",
			want:  false,
		},
		{
			name:  "Invalid hex color",
			input: "#ff7377ff",
			want:  false,
		},
		{
			name:  "Invalid hex color with 0x prefix",
			input: "0xff7377ff",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := HexColor(test.input)
			if got != test.want {
				t.Errorf("HexColor(%q) = %v; want %v",
					test.input, got, test.want)
			}
		})
	}
}

// TestRGBColor tests RGBColor function.
func TestRGBColor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid RGB color",
			input: "rgb(255, 115, 119)",
			want:  true,
		},
		{
			name:  "Valid RGB color with spaces",
			input: "rgb( 255, 115, 119 )",
			want:  true,
		},
		{
			name:  "Valid RGB color with tabs",
			input: "rgb(\t255,\t115,\t119\t)",
			want:  true,
		},
		{
			name:  "Valid RGB color with uppercase",
			input: "RGB(255, 115, 119)",
			want:  true,
		},
		{
			name:  "Invalid RGB color",
			input: "rgb(255, 115, 119, 255)",
			want:  false,
		},
		{
			name:  "Invalid RGB color with spaces",
			input: "rgb( 255, 115, 119, 255 )",
			want:  false,
		},
		{
			name:  "Invalid RGB color with tabs",
			input: "rgb(\t255,\t115,\t119,\t255\t)",
			want:  false,
		},
		{
			name:  "Invalid RGB color with uppercase",
			input: "RGB(255, 115, 119, 255)",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RGBColor(test.input)
			if got != test.want {
				t.Errorf("RGBColor(%q) = %v; want %v",
					test.input, got, test.want)
			}
		})
	}
}

// TestBin tests Bin function.
func TestBin(t *testing.T) {
	testCases := []struct {
		name string
		v    string
		want bool
	}{
		{
			name: "Valid Binary String",
			v:    "1010",
			want: true,
		},
		{
			name: "Valid Binary String With Prefix",
			v:    "0b1010",
			want: true,
		},
		{
			name: "Invalid Binary String",
			v:    "10201",
			want: false,
		},
		{
			name: "String with Non-Binary Characters",
			v:    "abc101",
			want: false,
		},
		{
			name: "Empty String",
			v:    "",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Bin(tc.v)
			if got != tc.want {
				t.Errorf("Bin(%q) = %v; want %v", tc.v, got, tc.want)
			}
		})
	}
}

// TestJWT tests JWT function.
func TestJWT(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "Valid JWT with secret base64 encoded",
			in:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibWVzc2FnZSI6IlJ1c3NpYW5zIGFyZSBtYW5pYWNzIGFuZCBmcmVha3MiLCJpYXQiOjE1MTYyMzkwMjJ9.3tG1rYCCxE7umetsvJPEYAZWCMQy-uwY5hA7QGvVqnM",
			want: true,
		},
		{
			name: "Valid JWT",
			in:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibWVzc2FnZSI6IlB1dGluIGlzIGFic29sdXRlIHNoaXQiLCJpYXQiOjE1MTYyMzkwMjJ9.wkLWA5GtCpWdxNOrRse8yHZgORDgf8TpJp73WUQb910",
			want: true,
		},
		{
			name: "Not valid JWT",
			in:   "eyHHGJHgJciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ZSI6IlB1dGl",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := JWT(tc.in)
			if got != tc.want {
				t.Errorf("JWT(%q) = %v; want %v", tc.in, got, tc.want)
			}
		})
	}
}

// TestMD5 tests MD5 function.
func TestMD5(t *testing.T) {
	validMD5 := "d41d8cd98f00b204e9800998ecf8427e"
	invalidMD5 := "notamd5hash"

	validResult := MD5(validMD5)
	if !validResult {
		t.Errorf("Expected valid MD5 to return true, got false")
	}

	invalidResult := MD5(invalidMD5)
	if invalidResult {
		t.Errorf("Expected invalid MD5 to return false, got true")
	}
}
