package is

import "testing"

// TestNickname tests the Nickname function.
func TestNickname(t *testing.T) {
	tests := []struct {
		name     string
		nickname string
		strict   bool
		want     bool
	}{
		// Non-strict mode tests
		{
			name:     "Valid Nickname - Alphanumeric",
			nickname: "User123",
			strict:   false,
			want:     true,
		},
		{
			name:     "Valid Nickname - Underscores",
			nickname: "User_123",
			strict:   false,
			want:     true,
		},
		{
			name:     "Valid Nickname - Unicode Japanese",
			nickname: "ユーザー名",
			strict:   false,
			want:     true,
		},
		{
			name:     "Valid Nickname - Unicode Ukraine",
			nickname: "Користувач",
			strict:   false,
			want:     true,
		},
		{
			name:     "Valid Nickname - Unicode Chinese",
			nickname: "用户名",
			strict:   false,
			want:     true,
		},
		{
			name:     "Invalid Nickname - Contains Emoji",
			nickname: "😀User",
			strict:   false,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Contains Space",
			nickname: "User Name",
			strict:   false,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Special Characters",
			nickname: "User!123",
			strict:   false,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Empty String",
			nickname: "",
			strict:   false,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Too Long",
			nickname: "ThisUsernameIsWayTooLong1234567890",
			strict:   false,
			want:     false,
		},
		// Strict mode tests
		{
			name:     "Valid Nickname - Alphanumeric (Strict)",
			nickname: "User123",
			strict:   true,
			want:     true,
		},
		{
			name:     "Valid Nickname - Underscores (Strict)",
			nickname: "User_123",
			strict:   true,
			want:     true,
		},
		{
			name:     "Invalid Nickname - Unicode Japanese (Strict)",
			nickname: "ユーザー名",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Unicode Ukraine (Strict)",
			nickname: "Користувач",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Unicode Chinese (Strict)",
			nickname: "用户名",
			strict:   true,
			want:     false,
		},
		{
			name:     "Valid Nickname - Alphanumeric (Strict)",
			nickname: "UserName",
			strict:   true,
			want:     true,
		},
		{
			name:     "Invalid Nickname - Contains Space (Strict)",
			nickname: "User Name",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Special Characters (Strict)",
			nickname: "User-Name",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Contains Dot (Strict)",
			nickname: "User.Name",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Contains Emoji (Strict)",
			nickname: "User😀",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Empty String (Strict)",
			nickname: "",
			strict:   true,
			want:     false,
		},
		{
			name:     "Invalid Nickname - Too Long (Strict)",
			nickname: "ThisUsernameIsWayTooLong1234567890",
			strict:   true,
			want:     false,
		},
		{
			name:     "Valid Nickname - Trimmed Spaces",
			nickname: "   User123   ",
			strict:   false,
			want:     true,
		},
		{
			name:     "Valid Nickname - Trimmed Spaces (Strict)",
			nickname: "   User123   ",
			strict:   true,
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			if tt.strict {
				got = Nickname(tt.nickname, true)
			} else {
				got = Nickname(tt.nickname)
			}
			if got != tt.want {
				t.Errorf(
					"Nickname(%q, strict=%v) = %v, want %v",
					tt.nickname,
					tt.strict,
					got,
					tt.want,
				)
			}
		})
	}
}
