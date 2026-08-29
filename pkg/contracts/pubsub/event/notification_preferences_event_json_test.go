package event_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notifications/notification_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/event"
)

func TestNotificationPreferencesChangedEventContainsOnlyChangedIdentifiers(t *testing.T) {
	changedAt := time.Date(2026, 8, 24, 8, 9, 10, 0, time.UTC)
	payload, err := json.Marshal(event.NotificationPreferencesChangedEvent{
		UserID: "user-1", AccountID: "account-1", CustomerNumber: "customer-1", PreferencesRevision: 4,
		ChangedTopicCodes: []string{"order_update"}, ChangedChannels: []notification_enums.NotificationChannel{notification_enums.NotificationChannelPush},
		Source: "preference_centre", ChangedAt: changedAt, RequestID: "request-1",
	})
	if err != nil {
		t.Fatalf("marshal preferences-changed event: %v", err)
	}
	for _, want := range []string{`"preferences_revision":4`, `"changed_topic_codes":["order_update"]`, `"changed_channels":["push"]`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("event JSON missing %s: %s", want, payload)
		}
	}
	for _, forbidden := range []string{"destination", "email_opt_in", "sms_opt_in"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("event leaked %q: %s", forbidden, payload)
		}
	}
}
