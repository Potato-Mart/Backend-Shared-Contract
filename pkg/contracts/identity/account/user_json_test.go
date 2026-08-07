package account

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUserProfileJSONIncludesAvatarFieldsWhenPresent(t *testing.T) {
	body, err := json.Marshal(UserProfile{
		ID:            "usr_1",
		Email:         "customer@example.test",
		DisplayName:   "Customer",
		AvatarMediaID: "med_avatar",
		AvatarURL:     "https://cdn.example.test/avatar.png",
		Active:        true,
	})
	if err != nil {
		t.Fatalf("marshal user profile: %v", err)
	}
	text := string(body)
	for _, want := range []string{
		`"avatar_media_id":"med_avatar"`,
		`"avatar_url":"https://cdn.example.test/avatar.png"`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("UserProfile JSON = %s, want %s", text, want)
		}
	}
}

func TestUserProfileJSONOmitsEmptyAvatarFields(t *testing.T) {
	body, err := json.Marshal(UserProfile{
		ID:     "usr_1",
		Email:  "customer@example.test",
		Active: true,
	})
	if err != nil {
		t.Fatalf("marshal user profile: %v", err)
	}
	text := string(body)
	if strings.Contains(text, "avatar_media_id") || strings.Contains(text, "avatar_url") {
		t.Fatalf("empty avatar fields should be omitted, got %s", text)
	}
}
