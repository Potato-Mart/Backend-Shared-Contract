package template

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/core/notification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/email"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/push"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/sms"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/notification/template/template_enums"
)

// These fixtures pin representation and round trips, not Notification's
// validation, authorization, hashing or translation behavior. Consumers may
// reuse testdata with their own behavioral checks and synthetic values.
func TestNotificationContentJSONFixtures(t *testing.T) {
	var fixtures []struct {
		Name    string              `json:"name"`
		Content NotificationContent `json:"content"`
	}
	assertJSONFixtureRoundTrip(t, "testdata/notification_content.json", &fixtures)
	seen := make(map[string]bool)
	for _, fixture := range fixtures {
		t.Run(fixture.Name, func(t *testing.T) {
			content := fixture.Content
			key := content.Channel.String() + "/" + content.SourceLanguage
			if fixture.Name != key || seen[key] {
				t.Fatalf("duplicate or misnamed fixture %q", fixture.Name)
			}
			seen[key] = true
			if content.SchemaVersion != 1 || len(content.Variants) != 3 || len(content.Placeholders) == 0 {
				t.Fatalf("incomplete fixture %q", fixture.Name)
			}
			languages := make(map[string]bool)
			for _, variant := range content.Variants {
				if languages[variant.Language] {
					t.Fatalf("duplicate language %q", variant.Language)
				}
				languages[variant.Language] = true
				if variant.Language == content.SourceLanguage {
					if variant.Origin != template_enums.TemplateContentOriginAuthored || variant.Review != nil {
						t.Fatal("source must retain authored origin and absent review")
					}
				} else if variant.Origin != template_enums.TemplateContentOriginTranslated || variant.Review == nil ||
					variant.Review.SourceFingerprint != content.SourceFingerprint || variant.Review.TargetFingerprint == "" ||
					variant.Review.ReviewedAt.IsZero() || variant.Review.ReviewedBy == "" {
					t.Fatal("target review evidence lost")
				}
				var text string
				switch content.Channel {
				case notification_enums.NotificationChannelEmail:
					if variant.Email == nil || variant.SMS != nil || variant.Push != nil || content.EmailTheme == nil {
						t.Fatal("email arm/theme changed")
					}
					blocks := variant.Email.Blocks
					if len(blocks) != 7 || blocks[0].Level == nil || *blocks[0].Level != 1 ||
						blocks[2].SectionCode != "order_summary" || blocks[3].ActionURL != "{{order.url}}" ||
						blocks[5].AltText == nil || *blocks[5].AltText != "" || blocks[6].Text != nil {
						t.Fatal("structured email, protected section, URL or explicit decorative alt changed")
					}
					text = *blocks[1].Text
				case notification_enums.NotificationChannelSMS:
					if variant.SMS == nil || variant.Email != nil || variant.Push != nil || content.EmailTheme != nil {
						t.Fatal("sms arm changed")
					}
					text = variant.SMS.Body
				case notification_enums.NotificationChannelPush:
					if variant.Push == nil || variant.Email != nil || variant.SMS != nil || content.EmailTheme != nil {
						t.Fatal("push arm changed")
					}
					if variant.Push.ActionURL != "{{order.url}}" || variant.Push.Image.Code != "fixture-logo" {
						t.Fatal("push nontranslatable values changed")
					}
					text = variant.Push.Body
				default:
					t.Fatalf("unexpected fixture channel %q", content.Channel)
				}
				if !strings.Contains(text, "{{customer.given_name}}") {
					t.Fatal("literal scalar placeholder was not preserved")
				}
			}
			for _, language := range []string{"en", "zh-TW", "zh-CN"} {
				if !languages[language] {
					t.Fatalf("missing language %q", language)
				}
			}
		})
	}
	for _, channel := range []string{"email", "sms", "push"} {
		for _, source := range []string{"en", "zh-TW", "zh-CN"} {
			if !seen[channel+"/"+source] {
				t.Errorf("missing any-source fixture %s/%s", channel, source)
			}
		}
	}
}

func TestCountryTemplateRecordJSONFixtures(t *testing.T) {
	var fixtures []struct {
		Template  NotificationTemplate          `json:"template"`
		Version   NotificationTemplateVersion   `json:"version"`
		Reference NotificationTemplateReference `json:"reference"`
		Binding   NotificationTemplateBinding   `json:"binding"`
	}
	assertJSONFixtureRoundTrip(t, "testdata/country_template_records.json", &fixtures)
	if len(fixtures) != 2 {
		t.Fatal("expected independent AU and TW fixture identities")
	}
	for index, country := range []string{"AU", "TW"} {
		fixture := fixtures[index]
		if string(fixture.Template.CountryCode) != country || string(fixture.Version.CountryCode) != country ||
			string(fixture.Reference.CountryCode) != country || string(fixture.Binding.CountryCode) != country ||
			fixture.Binding.Template != fixture.Reference {
			t.Fatal("country must survive every identity/reference/binding")
		}
		if fixture.Template.Code != "order-update" || fixture.Version.TemplateCode != fixture.Template.Code ||
			fixture.Reference.TemplateCode != fixture.Template.Code {
			t.Fatal("same business code must remain representable independently in both countries")
		}
		if fixture.Template.Revision != 7 || fixture.Version.SourceRevision != 5 ||
			fixture.Version.PublishedVersion != 2 || fixture.Reference.PublishedVersion != 2 ||
			fixture.Binding.Revision != 3 || fixture.Version.Content.SchemaVersion != 1 {
			t.Fatal("draft, publication, binding and schema counters were conflated")
		}
	}
	if fixtures[0].Reference == fixtures[1].Reference {
		t.Fatal("country-qualified references must remain distinct")
	}
	if fixtures[0].Binding.PurposeCode != "" || fixtures[1].Binding.PurposeCode != "fixture-purpose" {
		t.Fatal("absent versus explicit open purpose changed")
	}
}

func TestEmailBlockOptionalFieldPresence(t *testing.T) {
	var blocks []EmailTemplateBlock
	raw := []byte(`[{"id":"decorative","kind":"image","media":{"code":"decoration"},"alt_text":""},{"id":"line","kind":"divider"},{"id":"copy","kind":"paragraph","text":""}]`)
	if err := json.Unmarshal(raw, &blocks); err != nil {
		t.Fatal(err)
	}
	if blocks[0].AltText == nil || *blocks[0].AltText != "" || blocks[1].AltText != nil ||
		blocks[1].Text != nil || blocks[2].Text == nil || *blocks[2].Text != "" {
		t.Fatal("absent and explicit empty leaves must remain distinguishable")
	}
	got, err := json.Marshal(blocks)
	if err != nil || !bytes.Equal(got, raw) {
		t.Fatalf("optional field round trip = %s, err=%v", got, err)
	}
}

func TestExistingDeliveryAndLocalizationJSONUnchanged(t *testing.T) {
	cases := []struct {
		value any
		want  string
	}{
		{email.EmailNotification{Subject: "Subject", Body: "Body"}, `{"subject":"Subject","body":"Body"}`},
		{sms.SMSNotification{Body: "Body"}, `{"body":"Body"}`},
		{push.PushNotification{Title: "Title", Body: "Body"}, `{"title":"Title","body":"Body"}`},
		{localization.LocalizedText{Language: "zh-TW", Text: "訂單"}, `{"language":"zh-TW","text":"訂單"}`},
	}
	for _, tc := range cases {
		got, err := json.Marshal(tc.value)
		if err != nil || string(got) != tc.want {
			t.Errorf("%T JSON = %s, err=%v; want %s", tc.value, got, err, tc.want)
		}
	}
}

func TestRequiredIdentityAndRevisionFieldsAreNotOmitted(t *testing.T) {
	cases := []struct {
		value any
		keys  []string
	}{
		{NotificationTemplate{}, []string{"country_code", "code", "revision", "active", "content"}},
		{NotificationTemplateVersion{}, []string{"country_code", "template_code", "published_version", "source_revision", "published_at", "published_by"}},
		{NotificationTemplateReference{}, []string{"country_code", "template_code", "published_version"}},
		{NotificationTemplateBinding{}, []string{"country_code", "topic_code", "channel", "template", "active", "revision"}},
		{NotificationContent{}, []string{"schema_version", "channel", "source_language", "source_fingerprint", "placeholders", "variants"}},
		{TemplatePlaceholder{}, []string{"key", "value_type", "sensitive"}},
		{TranslationReview{}, []string{"source_fingerprint", "target_fingerprint", "reviewed_at", "reviewed_by"}},
	}
	for _, tc := range cases {
		body, err := json.Marshal(tc.value)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(body, &fields); err != nil {
			t.Fatal(err)
		}
		for _, key := range tc.keys {
			if _, present := fields[key]; !present {
				t.Errorf("%T must serialize %q even when zero; service validates its value", tc.value, key)
			}
		}
	}
}

func TestTranslationEvidenceDoesNotConflateFingerprints(t *testing.T) {
	// Synthetic evidence only: deliberately stale source and independently
	// edited target must remain distinguishable for owner freshness checks.
	raw := []byte(`{"language":"zh-TW","sms":{"body":"已編輯 {{customer.given_name}}"},"origin":"translated","review":{"source_fingerprint":"sha256:previous-source","target_fingerprint":"sha256:confirmed-target","reviewed_at":"2026-09-22T00:00:00Z","reviewed_by":"fixture-reviewer"}}`)
	var variant LocalizedNotificationContent
	if err := json.Unmarshal(raw, &variant); err != nil {
		t.Fatal(err)
	}
	if variant.Review == nil || variant.Review.SourceFingerprint != "sha256:previous-source" ||
		variant.Review.TargetFingerprint != "sha256:confirmed-target" {
		t.Fatal("independent source/target evidence was lost")
	}
	variant.Review = nil
	body, err := json.Marshal(variant)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(body, []byte(`"review"`)) {
		t.Fatal("missing review must be omitted rather than implied by an origin value")
	}
}

func assertJSONFixtureRoundTrip(t *testing.T, path string, value any) {
	t.Helper()
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	decoder := json.NewDecoder(bytes.NewReader(want))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(value); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
	got, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var wantValue, gotValue any
	if err := json.Unmarshal(want, &wantValue); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(got, &gotValue); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("%s JSON shape or values changed on round trip", path)
	}
}
