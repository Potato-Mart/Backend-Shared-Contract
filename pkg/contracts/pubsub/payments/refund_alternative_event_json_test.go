package payments_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/payments"
)

func TestRefundAlternativeEventFixtureRoundTrip(t *testing.T) {
	raw, err := os.ReadFile("testdata/refund_alternative.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var decoded event.RefundAlternativeEvent
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	want := event.RefundAlternativeEvent{
		RefundRequestID:      "rr_1",
		ProposalRevision:     2,
		OrderNumber:          "CUSTOM-0042/A",
		RetailCustomerNumber: "RCABCDEF",
		Locale:               "zh-TW",
		Lines: []event.RefundAlternativeLine{{
			OriginalLineRef: "line-1", OriginalSKUCode: "SKU-1", AlternativeSKUCode: "SKU-2",
			AlternativeName: "Alternative", Quantity: 2, UnitPrice: money.Money{AmountMinor: 450, Currency: "AUD"},
		}},
		OriginalPayable: money.Money{AmountMinor: 1000, Currency: "AUD"},
		ProposedPayable: money.Money{AmountMinor: 900, Currency: "AUD"},
		ContactNote:     "Reply to the store",
		MarketCode:      "AU",
		OfferedBy:       "staff-1",
		OfferedAt:       time.Date(2026, 11, 1, 1, 2, 3, 0, time.UTC),
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("fixture decoded to %+v, want %+v", decoded, want)
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var again event.RefundAlternativeEvent
	if err := json.Unmarshal(encoded, &again); err != nil || !reflect.DeepEqual(again, want) {
		t.Fatalf("round trip changed the event: %v %+v", err, again)
	}
}

func TestRefundAlternativeEventOmitsAbsentOptionalFields(t *testing.T) {
	encoded, err := json.Marshal(event.RefundAlternativeEvent{
		RefundRequestID: "rr_1", ProposalRevision: 1, OrderNumber: "SO-1",
		Lines:           []event.RefundAlternativeLine{},
		OriginalPayable: money.Money{AmountMinor: 1, Currency: "AUD"},
		ProposedPayable: money.Money{AmountMinor: 1, Currency: "AUD"},
		OfferedAt:       time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var fields map[string]any
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, absent := range []string{"customer_id", "retail_customer_number", "locale", "contact_note", "market_code", "country_code", "offered_by"} {
		if _, ok := fields[absent]; ok {
			t.Fatalf("absent %s must be omitted: %s", absent, encoded)
		}
	}
}
