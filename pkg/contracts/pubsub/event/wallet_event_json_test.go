package event_test

import (
	"encoding/json"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	notification "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pubsub/event"
)

func TestGiftCardIssuedEventRoundTrip(t *testing.T) {
	issuedAt := time.Date(2026, 7, 27, 1, 2, 3, 0, time.UTC)
	event := notification.GiftCardIssuedEvent{
		IssuanceID: "gift-1", DenominationPolicyVersion: 2,
		Amount: money.Money{AmountMinor: 50_000, Currency: "AUD"}, IssuedAt: issuedAt,
	}
	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal gift-card event: %v", err)
	}
	for _, forbidden := range []string{
		"payment_status", "issuance_status", "recipient_email", "recipient_name",
		"sender_name", "message", "claim_code", "locale",
	} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("gift-card event exposed retired delivery field %q: %s", forbidden, payload)
		}
	}
	var decoded notification.GiftCardIssuedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal gift-card event: %v", err)
	}
	if decoded.IssuanceID != event.IssuanceID || decoded.DenominationPolicyVersion != 2 ||
		decoded.Amount != event.Amount || !decoded.IssuedAt.Equal(issuedAt) {
		t.Fatalf("gift-card event did not round-trip: %+v", decoded)
	}
}

// TestGiftCardIssuedEventCarriesBonusSeparately keeps Amount as the face and
// purchase evidence and reports any promotional bonus separately, so the
// charged amount and the issued balance stay independently auditable.
func TestGiftCardIssuedEventCarriesBonusSeparately(t *testing.T) {
	const wire = `{"issuance_id":"gift-1","amount":{"amount_minor":100000,"currency":"AUD"},` +
		`"bonus_amount_minor":5000,"issued_at":"2026-08-16T00:00:00Z"}`

	var decoded notification.GiftCardIssuedEvent
	if err := json.Unmarshal([]byte(wire), &decoded); err != nil {
		t.Fatalf("unmarshal gift-card event: %v", err)
	}

	payload, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal gift-card event: %v", err)
	}
	if !strings.Contains(string(payload), `"amount":{"amount_minor":100000,"currency":"AUD"}`) {
		t.Fatalf("gift-card event must keep the charged face amount: %s", payload)
	}
	if !strings.Contains(string(payload), `"bonus_amount_minor":5000`) {
		t.Fatalf("gift-card event JSON missing bonus_amount_minor: %s", payload)
	}
	for _, forbidden := range []string{"recipient_email", "recipient_name", "sender_name", "claim_code"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("gift-card event must not expose delivery material %q: %s", forbidden, payload)
		}
	}

	noBonus, err := json.Marshal(notification.GiftCardIssuedEvent{IssuanceID: "gift-2"})
	if err != nil {
		t.Fatalf("marshal gift-card event without bonus: %v", err)
	}
	if strings.Contains(string(noBonus), "bonus_amount_minor") {
		t.Fatalf("an absent bonus must stay off the wire: %s", noBonus)
	}
}
