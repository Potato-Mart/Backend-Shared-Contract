package message_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/marketing/message"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/marketing/message/message_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/notifications/notification_enums"
)

func TestMarketingMessageUsesCanonicalNotificationChannelWithoutRecipientData(t *testing.T) {
	payload, err := json.Marshal(message.MarketingMessage{
		Code: "msg-1", Channel: notification_enums.NotificationChannelEmail,
		NotificationTopicCode: "seasonal_offer", SegmentCode: "vip",
		Status: message_enums.MarketingMessageStatusScheduled,
	})
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}
	if !strings.Contains(string(payload), `"channel":"email"`) || strings.Contains(string(payload), "recipient") || strings.Contains(string(payload), "provider") {
		t.Fatalf("unexpected message contract: %s", payload)
	}
}
