package notification_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/contracts/notification"
)

func TestGiftCardIssuedEventRoundTrip(t *testing.T) {
	issuedAt := time.Date(2026, 7, 27, 1, 2, 3, 0, time.UTC)
	event := notification.GiftCardIssuedEvent{
		IssuanceID: "gift-1", RecipientEmail: "customer@example.com",
		RecipientName: "Customer", SenderName: "Sender",
		Amount:  common.Money{AmountMinor: 20_000, Currency: "AUD"},
		Message: "Enjoy", ClaimCode: strings.Repeat("a", 32), Locale: "en",
		IssuedAt: issuedAt,
	}
	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal gift-card event: %v", err)
	}
	if strings.Contains(string(payload), "payment_status") || strings.Contains(string(payload), "issuance_status") {
		t.Fatalf("gift-card event recreated removed bridge state: %s", payload)
	}
	var decoded notification.GiftCardIssuedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal gift-card event: %v", err)
	}
	if decoded.IssuanceID != event.IssuanceID || decoded.Amount != event.Amount || !decoded.IssuedAt.Equal(issuedAt) {
		t.Fatalf("gift-card event did not round-trip: %+v", decoded)
	}
}
