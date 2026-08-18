package event_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	notification "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pubsub/event"
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

// TestOrderPaidEventLegacyPayloadCarriesNoQualificationEvidence pins the
// fail-closed contract: an event published before v27.3.0 carries no subtotal,
// discount, or tag keys, and decodes to an empty currency and a nil slice. A
// consumer must treat that as "no evidence" and refuse to qualify, never as a
// zero subtotal or a zero discount.
func TestOrderPaidEventLegacyPayloadCarriesNoQualificationEvidence(t *testing.T) {
	const legacyWire = `{"order_id":"order_1","order_number":"MAMA260703ABC123",` +
		`"amount_paid":{"amount_minor":10000,"currency":"AUD"},` +
		`"paid_at":"2026-08-16T00:00:00Z"}`

	var decoded notification.OrderPaidEvent
	if err := json.Unmarshal([]byte(legacyWire), &decoded); err != nil {
		t.Fatalf("unmarshal legacy order paid event: %v", err)
	}

	payload, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal legacy order paid event: %v", err)
	}
	if strings.Contains(string(payload), `"tags"`) {
		t.Fatalf("absent order tags must stay absent on the wire: %s", payload)
	}

	roundTripped := map[string]json.RawMessage{}
	if err := json.Unmarshal(payload, &roundTripped); err != nil {
		t.Fatalf("decode re-marshalled legacy order paid event: %v", err)
	}
	for _, key := range []string{"subtotal", "discount_amount"} {
		raw, present := roundTripped[key]
		if !present {
			t.Fatalf("%s must remain a declared money field: %s", key, payload)
		}
		if !strings.Contains(string(raw), `"currency":""`) {
			t.Fatalf("legacy %s must decode to an empty currency so consumers fail closed: %s", key, raw)
		}
	}
}
