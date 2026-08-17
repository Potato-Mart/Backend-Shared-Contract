package account_test

import (
	"encoding/json"
	"strings"
	"testing"

	identity "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/account"
)

func TestUserNotificationTopicsJSONShape(t *testing.T) {
	prefs := identity.UserNotificationPreferences{
		Channels: identity.UserNotificationChannels{Email: true},
		Topics: identity.UserNotificationTopics{
			AccountUpdates:  true,
			SecurityAlerts:  true,
			OrderUpdates:    true,
			DeliveryUpdates: true,
			InvoiceUpdates:  true,
			Promotions:      true,
			SystemAlerts:    true,
		},
	}

	raw, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("marshal notification preferences: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal notification preferences: %v", err)
	}
	if strings.Contains(string(raw), "quiet_hours") {
		t.Fatalf("removed quiet-hours field reappeared: %s", raw)
	}
	topics, ok := got["topics"].(map[string]any)
	if !ok {
		t.Fatalf("topics missing or wrong type: %#v", got["topics"])
	}

	for _, key := range []string{
		"account_updates",
		"security_alerts",
		"order_updates",
		"delivery_updates",
		"invoice_updates",
		"promotions",
		"system_alerts",
	} {
		if topics[key] != true {
			t.Fatalf("topics[%q]=%#v, want true in %#v", key, topics[key], topics)
		}
	}
}

func TestUserNotificationTopicsRoundTrip(t *testing.T) {
	raw := []byte(`{"channels":{"email":true},"topics":{"delivery_updates":true,"invoice_updates":true}}`)

	var prefs identity.UserNotificationPreferences
	if err := json.Unmarshal(raw, &prefs); err != nil {
		t.Fatalf("unmarshal notification preferences: %v", err)
	}
	if !prefs.Topics.DeliveryUpdates || !prefs.Topics.InvoiceUpdates {
		t.Fatalf("new topic fields did not round-trip: %+v", prefs.Topics)
	}
}

func TestUserNotificationTopicGroupsJSONShape(t *testing.T) {
	prefs := identity.UserNotificationPreferences{
		Channels: identity.UserNotificationChannels{Email: true},
		TopicGroups: &identity.UserNotificationTopicGroupPreferences{
			Payment:        identity.NotificationTopicGroupChannels{Email: true},
			Order:          identity.NotificationTopicGroupChannels{Email: true, Push: true},
			Receipt:        identity.NotificationTopicGroupChannels{Email: true, SMS: true},
			ProductUpdates: identity.NotificationTopicGroupChannels{Push: true},
			Promotions:     identity.NotificationTopicGroupChannels{Email: true, Push: true, SMS: true},
		},
	}

	raw, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("marshal notification preferences: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal notification preferences: %v", err)
	}
	topicGroups, ok := got["topic_groups"].(map[string]any)
	if !ok {
		t.Fatalf("topic_groups missing or wrong type: %#v", got["topic_groups"])
	}
	for _, key := range []string{
		"payment",
		"order",
		"receipt",
		"product_updates",
		"promotions",
	} {
		group, ok := topicGroups[key].(map[string]any)
		if !ok {
			t.Fatalf("topic_groups[%q] missing or wrong type in %#v", key, topicGroups)
		}
		for _, channel := range []string{"email", "push", "sms"} {
			if _, ok := group[channel]; !ok {
				t.Fatalf("topic_groups[%q][%q] missing in %#v", key, channel, group)
			}
		}
	}
	if promotions := topicGroups["promotions"].(map[string]any); promotions["sms"] != true {
		t.Fatalf("topic_groups promotions sms = %#v, want true", promotions["sms"])
	}

	var decoded identity.UserNotificationPreferences
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal notification preferences round-trip: %v", err)
	}
	if decoded.TopicGroups == nil || !decoded.TopicGroups.Receipt.SMS || !decoded.TopicGroups.ProductUpdates.Push {
		t.Fatalf("topic groups did not round-trip: %+v", decoded.TopicGroups)
	}
}

func TestUserNotificationTopicGroupsOmittedWhenAbsent(t *testing.T) {
	prefs := identity.UserNotificationPreferences{
		Channels: identity.UserNotificationChannels{Email: true},
	}

	raw, err := json.Marshal(prefs)
	if err != nil {
		t.Fatalf("marshal notification preferences: %v", err)
	}
	if strings.Contains(string(raw), "topic_groups") {
		t.Fatalf("absent topic_groups must be omitted: %s", raw)
	}

	var decoded identity.UserNotificationPreferences
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal notification preferences: %v", err)
	}
	if decoded.TopicGroups != nil {
		t.Fatalf("absent topic_groups decoded non-nil: %+v", decoded.TopicGroups)
	}
}
