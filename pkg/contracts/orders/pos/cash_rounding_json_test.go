package pos_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pricing/quote/quote_enums"
)

func TestCashRoundingSnapshotKeepsConsiderationExact(t *testing.T) {
	appliedAt := time.Date(2026, 8, 12, 5, 6, 7, 0, time.UTC)
	// 1-2 cents round down, 3-4 up to 5, 6-7 down to 5, 8-9 up to 10, and 0
	// and 5 are unchanged. The adjustment is recorded separately so the
	// order consideration and its tax stay at exact minor units.
	cases := []struct {
		exact      int64
		cash       int64
		adjustment int64
	}{
		{1001, 1000, -1}, {1002, 1000, -2},
		{1003, 1005, 2}, {1004, 1005, 1},
		{1006, 1005, -1}, {1007, 1005, -2},
		{1008, 1010, 2}, {1009, 1010, 1},
		{1000, 1000, 0}, {1005, 1005, 0},
	}
	for _, tc := range cases {
		value := pos.CashRoundingSnapshot{
			ExactAmountDue:     money.Money{AmountMinor: tc.exact, Currency: "AUD"},
			RoundingAdjustment: money.Money{AmountMinor: tc.adjustment, Currency: "AUD"},
			CashAmountDue:      money.Money{AmountMinor: tc.cash, Currency: "AUD"},
			IncrementMinor:     5,
			Mode:               quote_enums.RoundingModeCashIncrement,
			AppliedAt:          appliedAt,
		}
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal cash rounding: %v", err)
		}
		var decoded pos.CashRoundingSnapshot
		if err := json.Unmarshal(payload, &decoded); err != nil {
			t.Fatalf("unmarshal cash rounding: %v", err)
		}
		if decoded.ExactAmountDue.AmountMinor+decoded.RoundingAdjustment.AmountMinor != decoded.CashAmountDue.AmountMinor {
			t.Fatalf("rounding adjustment must reconcile exact to cash: %+v", decoded)
		}
		if decoded.IncrementMinor != 5 || decoded.Mode != quote_enums.RoundingModeCashIncrement {
			t.Fatalf("cash increment evidence changed: %+v", decoded)
		}
		if !strings.Contains(string(payload), `"exact_amount_due"`) || !strings.Contains(string(payload), `"cash_amount_due"`) {
			t.Fatalf("cash rounding JSON = %s", payload)
		}
	}
}
