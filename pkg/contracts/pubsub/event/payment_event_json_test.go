package event_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/payments/payment/payment_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pubsub/event"
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
