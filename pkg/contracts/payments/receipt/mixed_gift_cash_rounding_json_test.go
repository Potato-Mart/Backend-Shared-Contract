package receipt_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/receipt"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

func TestMixedGiftCashRoundingReceiptJSON(t *testing.T) {
	value := receipt.MixedGiftCashRoundingSnapshot{
		OrderTotal:         money.Money{AmountMinor: 10003, Currency: "AUD"},
		GiftCardAmount:     money.Money{AmountMinor: 3001, Currency: "AUD"},
		ExactCashAmountDue: money.Money{AmountMinor: 7002, Currency: "AUD"},
		RoundingAdjustment: money.Money{AmountMinor: -2, Currency: "AUD"},
		CashAmountDue:      money.Money{AmountMinor: 7000, Currency: "AUD"},
		IncrementMinor:     5,
		Mode:               quote_enums.RoundingModeCashIncrement,
		AppliedAt:          time.Date(2026, 10, 31, 1, 2, 3, 0, time.UTC),
	}
	payload, err := json.Marshal(receipt.ReceiptSnapshot{MixedGiftCashRounding: &value})
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	if _, exists := fields["cash_rounding"]; exists {
		t.Fatal("mixed receipt must not serialize all-cash evidence")
	}
	var actual map[string]json.RawMessage
	if err := json.Unmarshal(fields["mixed_gift_cash_rounding"], &actual); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"order_total", "gift_card_amount", "exact_cash_amount_due", "rounding_adjustment", "cash_amount_due", "increment_minor", "mode", "applied_at"} {
		if _, ok := actual[name]; !ok {
			t.Fatalf("mixed rounding JSON missing %q", name)
		}
	}
	if len(actual) != 8 {
		t.Fatalf("unexpected mixed rounding fields: %s", fields["mixed_gift_cash_rounding"])
	}
	var got receipt.ReceiptSnapshot
	if err := json.Unmarshal(payload, &got); err != nil || got.MixedGiftCashRounding == nil || !reflect.DeepEqual(*got.MixedGiftCashRounding, value) {
		t.Fatalf("mixed rounding changed on round trip: %+v, %v", got.MixedGiftCashRounding, err)
	}
}

func TestAllCashReceiptOmitsMixedGiftCashRounding(t *testing.T) {
	value := receipt.ReceiptSnapshot{CashRounding: &receipt.CashRoundingSnapshot{
		ExactAmountDue:     money.Money{AmountMinor: 10003, Currency: "AUD"},
		RoundingAdjustment: money.Money{AmountMinor: 2, Currency: "AUD"},
		CashAmountDue:      money.Money{AmountMinor: 10005, Currency: "AUD"},
	}}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	if _, ok := fields["cash_rounding"]; !ok {
		t.Fatal("all-cash evidence missing")
	}
	if _, ok := fields["mixed_gift_cash_rounding"]; ok {
		t.Fatal("all-cash receipt serialized mixed evidence")
	}
}
