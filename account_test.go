package is

import "testing"

// TestNickname tests the Nickname function.
func TestNickname(t *testing.T) {
	tests := []struct {
		name     string
		nickname string
		want     bool
	}{
		{
			name:     "Valid Nickname - Alphanumeric",
			nickname: "User123",
			want:     true,
		},
		{
			name:     "Valid Nickname - Underscores",
			nickname: "User_123",
			want:     true,
		},
		{
			name:     "Invalid Nickname - Special Characters",
			nickname: "User!123",
			want:     false,
		},
		{
			name:     "Invalid Nickname - Too Long",
			nickname: "ThisUsernameIsWayTooLong1234567890",
			want:     false,
		},
		{
			name:     "Invalid Nickname - Empty String",
			nickname: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nickname(tt.nickname); got != tt.want {
				t.Errorf("Nickname() = %v, want %v", got, tt.want)
			}
		})
	}
}
