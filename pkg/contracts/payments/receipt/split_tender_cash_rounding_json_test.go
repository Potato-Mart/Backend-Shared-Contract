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

// 100.03 paid 60.00 by EFTPOS and 40.03 in cash, rounded to 40.05.
func splitTenderCashRounding() receipt.SplitTenderCashRoundingSnapshot {
	return receipt.SplitTenderCashRoundingSnapshot{
		OrderTotal:         money.Money{AmountMinor: 10003, Currency: "AUD"},
		NonCashAmount:      money.Money{AmountMinor: 6000, Currency: "AUD"},
		ExactCashAmountDue: money.Money{AmountMinor: 4003, Currency: "AUD"},
		RoundingAdjustment: money.Money{AmountMinor: 2, Currency: "AUD"},
		CashAmountDue:      money.Money{AmountMinor: 4005, Currency: "AUD"},
		IncrementMinor:     5,
		Mode:               quote_enums.RoundingModeCashIncrement,
		AppliedAt:          time.Date(2026, 11, 15, 1, 2, 3, 0, time.UTC),
		TenderPartID:       "part-2",
	}
}

func TestSplitTenderCashRoundingReceiptJSON(t *testing.T) {
	value := splitTenderCashRounding()
	payload, err := json.Marshal(receipt.ReceiptSnapshot{Total: value.OrderTotal, SplitTenderCashRounding: &value})
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	for _, other := range []string{"cash_rounding", "mixed_gift_cash_rounding"} {
		if _, exists := fields[other]; exists {
			t.Fatalf("split-tender receipt must not serialize %s", other)
		}
	}
	var actual map[string]json.RawMessage
	if err := json.Unmarshal(fields["split_tender_cash_rounding"], &actual); err != nil {
		t.Fatal(err)
	}
	names := []string{"order_total", "non_cash_amount", "exact_cash_amount_due", "rounding_adjustment", "cash_amount_due", "increment_minor", "mode", "applied_at", "tender_part_id"}
	for _, name := range names {
		if _, ok := actual[name]; !ok {
			t.Fatalf("split rounding JSON missing %q", name)
		}
	}
	if len(actual) != len(names) {
		t.Fatalf("unexpected split rounding fields: %s", fields["split_tender_cash_rounding"])
	}
	var got receipt.ReceiptSnapshot
	if err := json.Unmarshal(payload, &got); err != nil || got.SplitTenderCashRounding == nil || !reflect.DeepEqual(*got.SplitTenderCashRounding, value) {
		t.Fatalf("split rounding changed on round trip: %+v, %v", got.SplitTenderCashRounding, err)
	}
	if got.CashRounding != nil || got.MixedGiftCashRounding != nil {
		t.Fatal("split rounding decoded into another rounding record")
	}
}

// The documented arithmetic holds for the reference record.
func TestSplitTenderCashRoundingArithmetic(t *testing.T) {
	value := splitTenderCashRounding()
	currency := value.OrderTotal.Currency
	for _, amount := range []money.Money{value.NonCashAmount, value.ExactCashAmountDue, value.RoundingAdjustment, value.CashAmountDue} {
		if amount.Currency != currency {
			t.Fatalf("split rounding mixes currency %s with %s", amount.Currency, currency)
		}
	}
	if value.ExactCashAmountDue.AmountMinor != value.OrderTotal.AmountMinor-value.NonCashAmount.AmountMinor || value.ExactCashAmountDue.AmountMinor <= 0 {
		t.Fatal("exact cash due must be the positive total less non-cash tenders")
	}
	if value.CashAmountDue.AmountMinor != value.ExactCashAmountDue.AmountMinor+value.RoundingAdjustment.AmountMinor || value.CashAmountDue.AmountMinor <= 0 {
		t.Fatal("cash due must be the positive exact cash due plus the adjustment")
	}
	if value.CashAmountDue.AmountMinor%value.IncrementMinor != 0 || !value.Mode.IsValid() || value.TenderPartID == "" {
		t.Fatal("cash due must land on the increment and name its tender part")
	}
}

// At most one of the three rounding records is present on a receipt; the
// existing two keep their meaning (all cash, and gift cards plus cash).
func TestReceiptCarriesAtMostOneCashRounding(t *testing.T) {
	split := splitTenderCashRounding()
	receipts := map[string]receipt.ReceiptSnapshot{
		"none": {},
		"all_cash": {CashRounding: &receipt.CashRoundingSnapshot{
			ExactAmountDue: money.Money{AmountMinor: 10003, Currency: "AUD"}, RoundingAdjustment: money.Money{AmountMinor: 2, Currency: "AUD"}, CashAmountDue: money.Money{AmountMinor: 10005, Currency: "AUD"},
		}},
		"gift_and_cash": {MixedGiftCashRounding: &receipt.MixedGiftCashRoundingSnapshot{
			OrderTotal: money.Money{AmountMinor: 10003, Currency: "AUD"}, GiftCardAmount: money.Money{AmountMinor: 3001, Currency: "AUD"},
		}},
		"split_tender": {SplitTenderCashRounding: &split},
	}
	for name, value := range receipts {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(payload, &fields); err != nil {
			t.Fatal(err)
		}
		present := 0
		for _, key := range []string{"cash_rounding", "mixed_gift_cash_rounding", "split_tender_cash_rounding"} {
			if _, ok := fields[key]; ok {
				present++
			}
		}
		if want := map[bool]int{true: 0, false: 1}[name == "none"]; present != want {
			t.Fatalf("%s receipt serialized %d rounding records, want %d", name, present, want)
		}
	}
}
