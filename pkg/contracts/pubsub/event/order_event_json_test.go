package event_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	notification "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/event"
)

func TestOrderPaidEventQualificationFieldsRoundTrip(t *testing.T) {
	const wire = `{"order_id":"order_1","order_number":"MAMA260703ABC123",` +
		`"amount_paid":{"amount_minor":10000,"currency":"AUD"},` +
		`"subtotal":{"amount_minor":12000,"currency":"AUD"},` +
		`"discount_amount":{"amount_minor":2000,"currency":"AUD"},` +
		`"tags":["gift","staff_purchase"],` +
		`"paid_at":"2026-08-16T00:00:00Z"}`

	var decoded notification.OrderPaidEvent
	if err := json.Unmarshal([]byte(wire), &decoded); err != nil {
		t.Fatalf("unmarshal order paid event: %v", err)
	}

	payload, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal order paid event: %v", err)
	}
	for _, want := range []string{
		`"subtotal":{"amount_minor":12000,"currency":"AUD"}`,
		`"discount_amount":{"amount_minor":2000,"currency":"AUD"}`,
		`"tags":["gift","staff_purchase"]`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("order paid event JSON missing %s: %s", want, payload)
		}
	}
	if decoded.Subtotal != (money.Money{AmountMinor: 12_000, Currency: "AUD"}) ||
		decoded.DiscountAmount != (money.Money{AmountMinor: 2_000, Currency: "AUD"}) ||
		len(decoded.Tags) != 2 || decoded.Tags[0] != "gift" || decoded.Tags[1] != "staff_purchase" {
		t.Fatalf("order paid qualification fields did not round-trip: %+v", decoded)
	}
	if decoded.Subtotal.AmountMinor-decoded.DiscountAmount.AmountMinor != decoded.AmountPaid.AmountMinor {
		t.Fatalf("subtotal minus discount must reconcile with the amount paid: %+v", decoded)
	}
}
