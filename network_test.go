package is

import (
	"testing"
)

// TestIPv4 tests IPv4 function.
func TestIPv4(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Valid IPv4",
			ip:   "127.0.0.1",
			want: true,
		},
		{
			name: "Invalid IPv4",
			ip:   "256.0.0.1",
			want: false,
		},
		{
			name: "Invalid IPv4 format",
			ip:   "127.0.0",
			want: false,
		},
		{
			name: "Invalid IPv4 format",
			ip:   "invalid",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IPv4(test.ip); got != test.want {
				t.Errorf("IsIPv4(%v) = %v; want %v", test.ip, got, test.want)
			}
		})
	}
}

// TestIPv6 tests IPv6 function.
func TestIPv6(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Valid IPv6",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: true,
		},
		{
			name: "Invalid IPv6",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:733g",
			want: false,
		},
		{
			name: "Invalid IPv6 format",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370",
			want: false,
		},
		{
			name: "Invalid IPv6 format",
			ip:   "invalid",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IPv6(test.ip); got != test.want {
				t.Errorf("IsIPv6(%v) = %v; want %v", test.ip, got, test.want)
			}
		})
	}
}

// TestIP tests IP function.
func TestIP(t *testing.T) {
	var tests = []struct {
		ip   string
		want bool
	}{
		{"127.0.0.1", true},
		{"::1", true},
		{"2001:db8::8a2e", true},
		{"256.0.0.1", false},
		{"192.168.0", false},
		{"2001::25de::cade", false},
		{"", false},
	}

	for _, test := range tests {
		if got := IP(test.ip); got != test.want {
			t.Errorf("is.IP(%q) = %v; want %v", test.ip, got, test.want)
		}
	}
}
