package account

import (
	"encoding/json"
	"strings"
	"testing"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
)

func TestUserProfileJSONIncludesObjectMediaAvatarWhenPresent(t *testing.T) {
	body, err := json.Marshal(UserProfile{
		ID:          "usr_1",
		Email:       "customer@example.test",
		DisplayName: "Customer",
		Avatar: &security.ObjectMedia{
			ID:  "med_avatar",
			URL: "https://cdn.example.test/avatar.png",
		},
		Active: true,
	})
	if err != nil {
		t.Fatalf("marshal user profile: %v", err)
	}
	text := string(body)
	if !strings.Contains(text, `"avatar":{"id":"med_avatar","url":"https://cdn.example.test/avatar.png"}`) {
		t.Fatalf("UserProfile JSON = %s, want nested object_media avatar", text)
	}
	for _, legacy := range []string{"avatar_media_id", "avatar_url"} {
		if strings.Contains(text, legacy) {
			t.Fatalf("UserProfile JSON retained legacy %s: %s", legacy, text)
		}
	}
}

func TestUserProfileJSONOmitsEmptyObjectMediaAvatar(t *testing.T) {
	body, err := json.Marshal(UserProfile{
		ID:     "usr_1",
		Email:  "customer@example.test",
		Active: true,
	})
	if err != nil {
		t.Fatalf("marshal user profile: %v", err)
	}
	text := string(body)
	if strings.Contains(text, "avatar") {
		t.Fatalf("empty avatar should be omitted, got %s", text)
	}
}
