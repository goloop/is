package is

import (
	"testing"
)

// TestLatitude tests Latitude function.
func TestLatitude(t *testing.T) {
	tests := []struct {
		name    string
		inStr   string
		inFloat float64
		want    bool
	}{
		{
			name:    "Valid Latitude float",
			inFloat: 45.0,
			want:    true,
		},
		{
			name:    "Invalid Latitude float",
			inFloat: 91.0,
			want:    false,
		},
		{
			name:  "Valid Latitude string",
			inStr: "45.0",
			want:  true,
		},
		{
			name:  "Invalid Latitude string",
			inStr: "91.0",
			want:  false,
		},
		{
			name:  "Invalid Latitude string format",
			inStr: "invalid",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got bool

			if test.inStr != "" {
				got = Latitude(test.inStr)
			} else {
				got = Latitude(test.inFloat)
			}

			if got != test.want {
				t.Errorf("Latitude(%v) = %v; want %v",
					test.inStr, got, test.want)
			}
		})
	}
}

// TestLongitude tests Longitude function.
func TestLongitude(t *testing.T) {
	tests := []struct {
		name    string
		inStr   string
		inFloat float64
		want    bool
	}{
		{
			name:    "Valid Longitude float",
			inFloat: 45.0,
			want:    true,
		},
		{
			name:    "Invalid Longitude float",
			inFloat: 181.0,
			want:    false,
		},
		{
			name:  "Valid Longitude string",
			inStr: "45.0",
			want:  true,
		},
		{
			name:  "Invalid Longitude string",
			inStr: "181.0",
			want:  false,
		},
		{
			name:  "Invalid Longitude string format",
			inStr: "invalid",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got bool

			if test.inStr != "" {
				got = Longitude(test.inStr)
			} else {
				got = Longitude(test.inFloat)
			}

			if got != test.want {
				t.Errorf("Longitude(%v) = %v; want %v",
					test.inStr, got, test.want)
			}
		})
	}
}

// TestCoordinates tests Coordinates function.
func TestCoordinates(t *testing.T) {
	tests := []struct {
		name     string
		latStr   string
		lonStr   string
		latFloat float64
		lonFloat float64
		want     bool
	}{
		{
			name:     "Valid Coordinates float",
			latFloat: 45.0,
			lonFloat: -123.1,
			want:     true,
		},
		{
			name:     "Invalid Coordinates float",
			latFloat: 90.1,
			lonFloat: 180.1,
			want:     false,
		},
		{
			name:   "Valid Coordinates string",
			latStr: "45.0",
			lonStr: "-123.1",
			want:   true,
		},
		{
			name:   "Invalid Coordinates string",
			latStr: "90.1",
			lonStr: "180.1",
			want:   false,
		},
		{
			name:   "Invalid Coordinates string format",
			latStr: "invalid",
			lonStr: "180.1",
			want:   false,
		},
		{
			name:   "Invalid Coordinates string format",
			latStr: "90.1",
			lonStr: "invalid",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got bool

			if test.latStr != "" && test.lonStr != "" {
				got = Coordinates(test.latStr, test.lonStr)
			} else {
				got = Coordinates(test.latFloat, test.lonFloat)
			}

			if got != test.want {
				t.Errorf("Coordinates(%v, %v) = %v; want %v",
					test.latStr, test.lonStr, got, test.want)
			}
		})
	}
}
