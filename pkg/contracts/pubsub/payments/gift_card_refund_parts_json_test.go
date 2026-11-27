package payments_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/payments"
)

// V53.54 lets one refund restore several gift card benefit reservations.
// The parts list is optional, so refund events written before the release
// read back unchanged.
func TestV5354GiftCardRefundPartsFields(t *testing.T) {
	cases := []struct {
		model any
		field string
		json  string
		typ   reflect.Type
	}{
		{event.RefundCompletedEvent{}, "GiftCardRefundParts", "gift_card_refund_parts,omitempty", reflect.TypeOf([]event.GiftCardRefundPart(nil))},
		{event.RefundCompletedEvent{}, "BenefitReservationID", "benefit_reservation_id,omitempty", reflect.TypeOf("")},
		{event.RefundCompletedEvent{}, "GiftCardRefundAmount", "gift_card_refund_amount,omitempty", reflect.TypeOf((*money.Money)(nil))},
		{event.GiftCardRefundPart{}, "BenefitReservationID", "benefit_reservation_id", reflect.TypeOf("")},
		{event.GiftCardRefundPart{}, "GiftCardRefundAmount", "gift_card_refund_amount", reflect.TypeOf(money.Money{})},
	}
	for _, tc := range cases {
		field, ok := reflect.TypeOf(tc.model).FieldByName(tc.field)
		if !ok || field.Type != tc.typ || field.Tag.Get("json") != tc.json {
			t.Errorf("%T.%s = %+v, want %s `json:%q`", tc.model, tc.field, field, tc.typ, tc.json)
		}
	}
	if got := reflect.TypeOf(event.GiftCardRefundPart{}).NumField(); got != 2 {
		t.Fatalf("GiftCardRefundPart has %d fields, want 2", got)
	}
}

func TestV5354RefundCompletedEventGiftCardRefundPartsRoundTrip(t *testing.T) {
	completedAt := time.Date(2026, 11, 27, 2, 0, 0, 0, time.UTC)
	first := money.Money{AmountMinor: 1500, Currency: "AUD"}
	want := event.RefundCompletedEvent{
		RefundID: "refund-1", OrderNumber: "SO-1", CompletedAt: completedAt,
		Amount:               money.Money{AmountMinor: 4000, Currency: "AUD"},
		BenefitReservationID: "reservation-a",
		GiftCardRefundAmount: &first,
		GiftCardRefundParts: []event.GiftCardRefundPart{
			{BenefitReservationID: "reservation-a", GiftCardRefundAmount: first},
			{BenefitReservationID: "reservation-b", GiftCardRefundAmount: money.Money{AmountMinor: 2500, Currency: "AUD"}},
		},
	}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	const parts = `"gift_card_refund_parts":[{"benefit_reservation_id":"reservation-a","gift_card_refund_amount":{"amount_minor":1500,"currency":"AUD"}},{"benefit_reservation_id":"reservation-b","gift_card_refund_amount":{"amount_minor":2500,"currency":"AUD"}}]`
	if !strings.Contains(string(encoded), parts) {
		t.Fatalf("gift card refund parts wire shape changed: %s", encoded)
	}
	if !strings.Contains(string(encoded), `"benefit_reservation_id":"reservation-a","gift_card_refund_amount":{"amount_minor":1500,"currency":"AUD"},"gift_card_refund_parts"`) {
		t.Fatalf("single fields must mirror the first part: %s", encoded)
	}
	var got event.RefundCompletedEvent
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("round trip changed the event: %v %+v", err, got)
	}
}

func TestV5354LegacyRefundCompletedEventHasNoParts(t *testing.T) {
	const legacy = `{"refund_id":"refund-2","order_number":"SO-2","amount":{"amount_minor":1000,"currency":"AUD"},"benefit_reservation_id":"reservation-a","gift_card_refund_amount":{"amount_minor":400,"currency":"AUD"},"completed_at":"2026-11-27T02:00:00Z"}`
	var decoded event.RefundCompletedEvent
	if err := json.Unmarshal([]byte(legacy), &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.GiftCardRefundParts != nil {
		t.Fatalf("a legacy event must decode without parts: %+v", decoded.GiftCardRefundParts)
	}
	if decoded.BenefitReservationID != "reservation-a" || decoded.GiftCardRefundAmount == nil || decoded.GiftCardRefundAmount.AmountMinor != 400 {
		t.Fatalf("legacy single fields lost: %+v", decoded)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(encoded), "gift_card_refund_parts") {
		t.Fatalf("an event without parts must omit gift_card_refund_parts: %s", encoded)
	}
}
