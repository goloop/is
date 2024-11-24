package is

import (
	"testing"

	"github.com/goloop/g"
)

// TestEmail tests the Email function.
func TestEmail(t *testing.T) {
	// Test cases.
	tests := map[string]bool{
		"":                                 false,
		"name.@example.com":                false,
		".name@example.com":                false,
		"name-@example.com":                false,
		"-name@example.com":                false,
		"name+surname@example.com":         false,
		"!def!xyz%abc@test.org":            false,
		"+1~1+@test.org":                   false,
		"+@b.c":                            false,
		"0932910@example.c":                false,
		"0932910@example-.com":             false,
		"0932910@-example.com":             false,
		"0932910@example@com":              false,
		"0932910@com":                      false,
		"bob@example.com ":                 false,
		" bob@example.com":                 false,
		"@example.com":                     false,
		"test@example@example.com":         false,
		"test@wrong domain.com":            false,
		"é&ààà@example.com":                false,
		"not-a-valid-email":                false,
		"test test@example.com":            false,
		"michel @example.com":              false,
		"example@example.c":                false,
		g.Trim(" example@example.com "):    true,
		g.Weed("twolines\t@example.com\n"): true,
		"michel@hotmail.com":               true,
		"support@g2mail.com":               true,
		"MICHEL@hotmail.com":               true,
		"michel@HOTMAIL.COM":               true,
		"support@G2MAIL.com":               true,
		"test@912-wrong-domain902.com":     true,
		"admin@notarealdomain12345.com":    true,
		"a@example.xyz":                    true,
		"2023@знищена-московія.укр":        true,
		"プーチンはクソだ@example.com":             true,
		"二ノ宮@黒川.日本":                        true,
		"我買@屋企.香港":                         true,
	}

	for k, v := range tests {
		result := Email(k)
		if result != v {
			t.Errorf("Email(\"%v\"): expected %v, but got %v",
				k, v, result)
		}
	}
}
