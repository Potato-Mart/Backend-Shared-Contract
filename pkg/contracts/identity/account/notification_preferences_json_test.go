package account_test

import (
	"encoding/json"
	"strings"
	"testing"

	identity "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/account"
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
