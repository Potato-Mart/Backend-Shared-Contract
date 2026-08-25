package notifications_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/metadata"
	notifications "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/notifications"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/notifications/notification_enums"
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

func TestNotificationDeliveryOneOfPayloadGoldenAndNegativeCases(t *testing.T) {
	valid := map[notification_enums.NotificationChannel]notifications.NotificationDelivery{
		notification_enums.NotificationChannelEmail: {
			ID: "email-1", Channel: notification_enums.NotificationChannelEmail, DestinationCode: "destination-email", Email: &notifications.EmailNotification{Subject: "Subject", Body: "Body"},
		},
		notification_enums.NotificationChannelPush: {
			ID: "push-1", Channel: notification_enums.NotificationChannelPush, DestinationCode: "destination-push", Push: &notifications.PushNotification{Title: "Title", Body: "Body"},
		},
		notification_enums.NotificationChannelSMS: {
			ID: "sms-1", Channel: notification_enums.NotificationChannelSMS, DestinationCode: "destination-sms", SMS: &notifications.SMSNotification{Body: "Body"},
		},
		notification_enums.NotificationChannelInApp: {
			ID: "in-app-1", Channel: notification_enums.NotificationChannelInApp, DestinationCode: "destination-in-app", InApp: &notifications.InAppNotification{Title: "Title", Body: "Body"},
		},
		notification_enums.NotificationChannelSocialMedia: {
			ID: "social-1", Channel: notification_enums.NotificationChannelSocialMedia, DestinationCode: "destination-social", SocialMedia: &notifications.SocialMediaNotification{ProviderCode: "provider", MessageMode: "direct_message", Body: "Body"},
		},
	}
	for channel, delivery := range valid {
		t.Run(string(channel), func(t *testing.T) {
			if !hasMatchingSinglePayload(delivery) {
				t.Fatalf("%s golden delivery does not have one matching payload", channel)
			}
		})
	}

	for name, delivery := range map[string]notifications.NotificationDelivery{
		"no-payload":         {ID: "none-1", Channel: notification_enums.NotificationChannelEmail, DestinationCode: "destination"},
		"multiple-payloads":  {ID: "multiple-1", Channel: notification_enums.NotificationChannelEmail, DestinationCode: "destination", Email: &notifications.EmailNotification{}, SMS: &notifications.SMSNotification{}},
		"mismatched-payload": {ID: "mismatch-1", Channel: notification_enums.NotificationChannelEmail, DestinationCode: "destination", Push: &notifications.PushNotification{}},
	} {
		t.Run(name, func(t *testing.T) {
			if hasMatchingSinglePayload(delivery) {
				t.Fatal("invalid one-of fixture was classified as valid")
			}
		})
	}
}

func TestSocialPreferencesRequireDestinationSelectionAndSeparateConsent(t *testing.T) {
	now := time.Date(2026, 8, 24, 5, 0, 0, 0, time.UTC)
	emptySelection := notifications.NotificationPreferences{
		UserID: "user-1", Topics: []notifications.NotificationTopicPreference{{
			TopicCode: "campaign_reminder", Channels: []notifications.NotificationChannelPreference{{Channel: notification_enums.NotificationChannelSocialMedia, Enabled: true}},
		}},
	}
	if len(emptySelection.Topics[0].Channels[0].DestinationCodes) != 0 {
		t.Fatal("empty social allow-list fixture is invalid")
	}

	prefs := notifications.NotificationPreferences{
		UserID: "user-1", Topics: []notifications.NotificationTopicPreference{{
			TopicCode: "campaign_reminder", Channels: []notifications.NotificationChannelPreference{{Channel: notification_enums.NotificationChannelSocialMedia, Enabled: true, DestinationCodes: []string{"destination-1", "destination-2"}}},
		}},
	}
	if socialDestinationEnabled(prefs, "campaign_reminder", "destination-1") {
		t.Fatal("newly selected destination is enabled without separate consent")
	}
	prefs.Consents = []notifications.NotificationChannelConsent{{Channel: notification_enums.NotificationChannelSocialMedia, DestinationCode: "destination-1", Granted: true, ChangedAt: now}}
	if !socialDestinationEnabled(prefs, "campaign_reminder", "destination-1") {
		t.Fatal("selected destination with granted consent is disabled")
	}
	if socialDestinationEnabled(prefs, "campaign_reminder", "destination-2") {
		t.Fatal("destination without consent is enabled")
	}
	invalidConsent := notifications.NotificationPreferences{UserID: "user-1", Consents: []notifications.NotificationChannelConsent{{Channel: notification_enums.NotificationChannelSocialMedia, Granted: true}}}
	if invalidConsent.Consents[0].DestinationCode != "" {
		t.Fatal("social consent fixture must demonstrate missing destination code")
	}
}

func hasMatchingSinglePayload(delivery notifications.NotificationDelivery) bool {
	payloads := 0
	if delivery.Email != nil {
		payloads++
	}
	if delivery.Push != nil {
		payloads++
	}
	if delivery.SMS != nil {
		payloads++
	}
	if delivery.InApp != nil {
		payloads++
	}
	if delivery.SocialMedia != nil {
		payloads++
	}
	if payloads != 1 {
		return false
	}
	return (delivery.Channel == notification_enums.NotificationChannelEmail && delivery.Email != nil) ||
		(delivery.Channel == notification_enums.NotificationChannelPush && delivery.Push != nil) ||
		(delivery.Channel == notification_enums.NotificationChannelSMS && delivery.SMS != nil) ||
		(delivery.Channel == notification_enums.NotificationChannelInApp && delivery.InApp != nil) ||
		(delivery.Channel == notification_enums.NotificationChannelSocialMedia && delivery.SocialMedia != nil)
}

func socialDestinationEnabled(preferences notifications.NotificationPreferences, topicCode, destinationCode string) bool {
	selected := false
	for _, topic := range preferences.Topics {
		if topic.TopicCode != topicCode {
			continue
		}
		for _, channel := range topic.Channels {
			if channel.Channel != notification_enums.NotificationChannelSocialMedia || !channel.Enabled {
				continue
			}
			for _, code := range channel.DestinationCodes {
				if code == destinationCode {
					selected = true
				}
			}
		}
	}
	if !selected {
		return false
	}
	for _, consent := range preferences.Consents {
		if consent.Channel == notification_enums.NotificationChannelSocialMedia && consent.DestinationCode == destinationCode && consent.Granted {
			return true
		}
	}
	return false
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
