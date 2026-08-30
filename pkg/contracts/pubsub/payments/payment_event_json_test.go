package payments_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/payments"
)

func TestReceiptGeneratedEventRoundTrip(t *testing.T) {
	issuedAt := time.Date(2026, 8, 13, 1, 2, 3, 0, time.UTC)
	generated := event.ReceiptGeneratedEvent{
		OrderNumber:          "SO-123",
		RetailCustomerNumber: "RC-1",
		DocumentKind:         payment_enums.DocumentKindReceipt,
		Revision:             2,
		IssuedAt:             issuedAt,
		RequestID:            "req-1",
	}

	payload, err := json.Marshal(generated)
	if err != nil {
		t.Fatalf("marshal receipt generated event: %v", err)
	}
	for _, field := range []string{
		`"order_number":"SO-123"`,
		`"retail_customer_number":"RC-1"`,
		`"document_kind":"receipt"`,
		`"revision":2`,
		`"issued_at":"2026-08-13T01:02:03Z"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("receipt generated event missing %s: %s", field, payload)
		}
	}
	if strings.Contains(string(payload), "organisation_access_id") {
		t.Fatalf("absent organisation_access_id must be omitted: %s", payload)
	}

	var decoded event.ReceiptGeneratedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal receipt generated event: %v", err)
	}
	if decoded.OrderNumber != generated.OrderNumber || decoded.DocumentKind != payment_enums.DocumentKindReceipt ||
		decoded.Revision != 2 || !decoded.IssuedAt.Equal(issuedAt) {
		t.Fatalf("receipt generated event did not round-trip: %+v", decoded)
	}
}

func TestPaymentAndRefundFactJSONUseFactIdentityAndTimestamp(t *testing.T) {
	factOccurredAt := time.Date(2026, 8, 30, 1, 2, 3, 0, time.UTC)
	for _, value := range []any{
		event.PaymentFact{
			FactID: "fact_payment_1", PaymentID: "payment_1", OrderNumber: "SO-1", Status: "captured",
			Amount: money.Money{AmountMinor: 1000, Currency: "AUD"}, FactOccurredAt: factOccurredAt,
		},
		event.RefundFact{
			FactID: "fact_refund_1", RefundID: "refund_1", OrderNumber: "SO-1", Status: "completed",
			Amount: money.Money{AmountMinor: 1000, Currency: "AUD"}, FactOccurredAt: factOccurredAt,
		},
	} {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %T: %v", value, err)
		}
		if !strings.Contains(string(payload), `"fact_id"`) || !strings.Contains(string(payload), `"fact_occurred_at"`) {
			t.Fatalf("%T JSON misses fact identity: %s", value, payload)
		}
		for _, forbidden := range []string{`"event_id"`, `"occurred_at"`} {
			if strings.Contains(string(payload), forbidden) {
				t.Fatalf("%T JSON retained %s: %s", value, forbidden, payload)
			}
		}
	}
}
