package notifications_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/metadata"
	notifications "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/notifications"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/notifications/notification_enums"
)

func TestNotificationPreferencesSupportsBackendDefinedTopicsAndDestinationScopedSocialConsent(t *testing.T) {
	changedAt := time.Date(2026, 8, 24, 2, 3, 4, 0, time.UTC)
	prefs := notifications.NotificationPreferences{
		ID: "notification-preference-1", UserID: "user-1", CustomerNumber: "customer-1", Revision: 7, UpdatedAt: changedAt,
		Topics: []notifications.NotificationTopicPreference{{
			TopicCode: "seasonal_restock", // backend-created, not a shared enum value.
			Channels: []notifications.NotificationChannelPreference{{
				Channel: notification_enums.NotificationChannelSocialMedia, Enabled: true,
				DestinationCodes: []string{"destination-opaque-1"},
			}},
		}},
		Consents: []notifications.NotificationChannelConsent{{
			Channel: notification_enums.NotificationChannelSocialMedia, DestinationCode: "destination-opaque-1", Granted: true, ChangedAt: changedAt,
		}},
	}

	body, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("marshal preferences: %v", err)
	}
	for _, want := range []string{`"topic_code":"seasonal_restock"`, `"channel":"social_media"`, `"destination_codes":["destination-opaque-1"]`, `"destination_code":"destination-opaque-1"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("preferences JSON missing %s: %s", want, body)
		}
	}
}

func TestNotificationTopicAndDeliveryRemainProviderNeutral(t *testing.T) {
	now := time.Date(2026, 8, 24, 5, 6, 7, 0, time.UTC)
	topic := notifications.NotificationTopic{
		ID: "topic-1", Code: "campaign_reminder", Active: true, Revision: 2,
		Name:     []localization.LocalizedName{{Language: "en", Name: "Campaign reminder"}},
		Channels: []notifications.NotificationTopicChannel{{Channel: notification_enums.NotificationChannelSocialMedia, ConsentRequired: true}},
	}
	notification := notifications.Notification{
		ID: "notification-1", TopicCode: topic.Code, Recipient: notifications.NotificationRecipient{UserID: "user-1"},
		Status: notification_enums.NotificationStatusPending,
		Deliveries: []notifications.NotificationDelivery{{
			ID: "delivery-1", Channel: notification_enums.NotificationChannelSocialMedia, DestinationCode: "destination-opaque-1",
			Status: notification_enums.NotificationDeliveryStatusPending,
			SocialMedia: &notifications.SocialMediaNotification{
				ProviderCode: "configured_provider", MessageMode: "direct_message", RecipientReference: "provider-recipient-ref",
				Body: "A campaign is ready", Metadata: metadata.Metadata{"template_key": "campaign-reminder"},
			},
		}},
		AuditFields: audit.AuditFields{CreatedAt: now, UpdatedAt: now},
	}

	body, err := json.Marshal(struct {
		Topic        notifications.NotificationTopic `json:"topic"`
		Notification notifications.Notification      `json:"notification"`
	}{Topic: topic, Notification: notification})
	if err != nil {
		t.Fatalf("marshal notification: %v", err)
	}
	for _, forbidden := range []string{"webhook_secret", "access_token"} {
		if strings.Contains(strings.ToLower(string(body)), forbidden) {
			t.Fatalf("notification contract serialized provider-specific or secret field %q: %s", forbidden, body)
		}
	}
	if !strings.Contains(string(body), `"destination_code":"destination-opaque-1"`) || !strings.Contains(string(body), `"provider_code":"configured_provider"`) {
		t.Fatalf("notification JSON lost opaque endpoint or open provider code: %s", body)
	}
}

func TestPublishedNotificationExcludesProtectedDeliveryData(t *testing.T) {
	now := time.Date(2026, 8, 24, 7, 8, 9, 0, time.UTC)
	published := notifications.PublishedNotification{
		ID: "notification-1", TopicCode: "order_update", Title: "Order update", Body: "Your order moved.",
		Status: notification_enums.InAppNotificationStatusUnread, CreatedAt: now,
	}
	body, err := json.Marshal(published)
	if err != nil {
		t.Fatalf("marshal published notification: %v", err)
	}
	for _, forbidden := range []string{"recipient", "destination", "provider", "delivery", "error", "email", "sms", "social_media"} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("published notification leaked %q: %s", forbidden, body)
		}
	}
}
