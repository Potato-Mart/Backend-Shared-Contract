package event_test

import (
	"encoding/json"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	notification "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pubsub/event"
)

func TestGiftCardIssuedEventRoundTrip(t *testing.T) {
	issuedAt := time.Date(2026, 7, 27, 1, 2, 3, 0, time.UTC)
	event := notification.GiftCardIssuedEvent{
		IssuanceID: "gift-1", DenominationPolicyVersion: 2, RecipientEmail: "customer@example.com",
		RecipientName: "Customer", SenderName: "Sender",
		Amount:  money.Money{AmountMinor: 50_000, Currency: "AUD"},
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
	if decoded.IssuanceID != event.IssuanceID || decoded.DenominationPolicyVersion != 2 ||
		decoded.Amount != event.Amount || !decoded.IssuedAt.Equal(issuedAt) {
		t.Fatalf("gift-card event did not round-trip: %+v", decoded)
	}
}

func TestVoucherClaimIssuedEventContainsHandleNotClaimMaterial(t *testing.T) {
	event := notification.VoucherClaimIssuedEvent{
		IssuanceID: "voucher-issuance-1", DeliveryHandle: "delivery-handle-1",
	}

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal voucher claim event: %v", err)
	}
	for _, forbidden := range []string{"claim_code", "claim_token", "recipient_email", "recipient_customer_number"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("voucher claim event exposed %s: %s", forbidden, payload)
		}
	}

	var decoded notification.VoucherClaimIssuedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal voucher claim event: %v", err)
	}
	if decoded.IssuanceID != event.IssuanceID || decoded.DeliveryHandle != event.DeliveryHandle {
		t.Fatalf("voucher claim event did not round-trip: %+v", decoded)
	}
}
