package payments_test

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/payments"
)

func TestRefundRequestClosedEventFixtureRoundTrip(t *testing.T) {
	raw, err := os.ReadFile("testdata/refund_request_closed.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded event.RefundRequestClosedEvent
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	want := event.RefundRequestClosedEvent{
		RefundRequestID: "rr_42",
		OrderNumber:     "CUSTOM-0042/A",
		Status:          "rejected",
		Station:         true,
		DiscrepancyID:   "shortage-1",
		ItemIDs:         []string{"item-1", "item-2"},
		ClosedBy:        "staff-1",
		ClosedAt:        time.Date(2026, 11, 14, 1, 2, 3, 0, time.UTC),
		MarketCode:      "AU",
		CountryCode:     "AU",
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("fixture decoded to %+v, want %+v", decoded, want)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var again event.RefundRequestClosedEvent
	if err := json.Unmarshal(encoded, &again); err != nil || !reflect.DeepEqual(again, want) {
		t.Fatalf("round trip changed the event: %v %+v", err, again)
	}
}

// A cancelled non-station request keeps station=false on the wire and omits
// the optional discrepancy and geography.
func TestRefundRequestClosedEventOmitsAbsentOptionalFields(t *testing.T) {
	encoded, err := json.Marshal(event.RefundRequestClosedEvent{
		RefundRequestID: "rr_1", OrderNumber: "SO-1", Status: "cancelled",
		ItemIDs: []string{"item-1"}, ClosedBy: "customer-1",
		ClosedAt: time.Date(2026, 11, 14, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var fields map[string]any
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, absent := range []string{"discrepancy_id", "market_code", "country_code"} {
		if _, ok := fields[absent]; ok {
			t.Fatalf("absent %s must be omitted: %s", absent, encoded)
		}
	}
	if station, ok := fields["station"]; !ok || station != false {
		t.Fatalf("station must always be on the wire: %s", encoded)
	}
	if !strings.Contains(string(encoded), `"status":"cancelled"`) {
		t.Fatalf("status lost: %s", encoded)
	}
}

func TestRefundCompletedEventGiftCardReversals(t *testing.T) {
	completedAt := time.Date(2026, 11, 14, 2, 0, 0, 0, time.UTC)
	want := event.RefundCompletedEvent{
		RefundID: "refund-1", OrderNumber: "POS-1", CompletedAt: completedAt,
		GiftCardReversals: []event.RefundGiftCardReversal{{LineID: "gift-line-1", IssuanceID: "issuance-1", State: "reversed"}},
	}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(encoded), `"gift_card_reversals":[{"line_id":"gift-line-1","issuance_id":"issuance-1","state":"reversed"}]`) {
		t.Fatalf("gift card reversal wire shape changed: %s", encoded)
	}
	var got event.RefundCompletedEvent
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("round trip changed the event: %v %+v", err, got)
	}

	encoded, err = json.Marshal(event.RefundCompletedEvent{RefundID: "refund-2", OrderNumber: "SO-2", CompletedAt: completedAt})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(encoded), "gift_card_reversals") {
		t.Fatalf("a refund without sold gift cards must omit gift_card_reversals: %s", encoded)
	}
}

func TestRefundFailedEventReasonCodeIsOptional(t *testing.T) {
	failedAt := time.Date(2026, 11, 14, 3, 0, 0, 0, time.UTC)
	encoded, err := json.Marshal(event.RefundFailedEvent{RefundID: "refund-1", OrderNumber: "POS-1", Reason: "gift card already used", ReasonCode: "GIFT_CARD_ALREADY_USED", FailedAt: failedAt})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(encoded), `"reason":"gift card already used","reason_code":"GIFT_CARD_ALREADY_USED"`) {
		t.Fatalf("reason code wire shape changed: %s", encoded)
	}
	var got event.RefundFailedEvent
	if err := json.Unmarshal(encoded, &got); err != nil || got.ReasonCode != "GIFT_CARD_ALREADY_USED" {
		t.Fatalf("reason code lost: %v %+v", err, got)
	}

	encoded, err = json.Marshal(event.RefundFailedEvent{RefundID: "refund-2", OrderNumber: "SO-2", FailedAt: failedAt})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(encoded), "reason_code") {
		t.Fatalf("an absent reason code must be omitted: %s", encoded)
	}
}
