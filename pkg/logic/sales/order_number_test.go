package saleslogic

import "testing"

func TestIsValidOrderNumber(t *testing.T) {
	valid := []string{
		"MAMA260703ABC123",
		"MAMA260101A1B2C3",
		"MAMA991231ZZZZZZ",
		"MAMA000000000000",
	}
	for _, s := range valid {
		if !IsValidOrderNumber(s) {
			t.Fatalf("IsValidOrderNumber(%q) = false, want true", s)
		}
	}

	invalid := []string{
		"",
		"ORD-20260703-ABC123",
		"MAMA260703abc123",   // lowercase suffix
		"MAMA26073ABC123",    // 5 date digits
		"MAMA260703ABC12",    // 5-char suffix
		"MAMA260703ABC1234",  // 7-char suffix
		" MAMA260703ABC123",  // leading space
		"MAMA260703ABC123 ",  // trailing space
		"XMAMA260703ABC123",  // wrong prefix
		"MAMA-260703-ABC123", // separators
	}
	for _, s := range invalid {
		if IsValidOrderNumber(s) {
			t.Fatalf("IsValidOrderNumber(%q) = true, want false", s)
		}
	}
}
