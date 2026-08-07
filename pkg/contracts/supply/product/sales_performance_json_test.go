package product_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
)

func TestSalesPerformanceJSONShape(t *testing.T) {
	asOf := time.Date(2026, 7, 18, 0, 0, 0, 0, time.UTC)
	stats := product.SalesPerformanceStats{
		Last7Days:  product.SalesWindowStats{WindowDays: 7, SalesTotals: product.SalesTotals{PaidOrderCount: 2, GrossUnits: 8, RefundedUnits: 1, NetUnits: 7}},
		Last30Days: product.SalesWindowStats{WindowDays: 30},
		Last90Days: product.SalesWindowStats{WindowDays: 90},
		Lifetime:   product.SalesTotals{PaidOrderCount: 10, GrossUnits: 40, RefundedUnits: 2, NetUnits: 38},
		CategoryRanks: []product.CategorySalesRank{{
			CategoryTagID: "tag_1", Rank: 1, Population: 25, WindowDays: 30, NetUnits: 18,
		}},
		AsOf: asOf, Timezone: "Australia/Melbourne",
	}
	payload, err := json.Marshal(stats)
	if err != nil {
		t.Fatalf("marshal sales performance: %v", err)
	}
	for _, want := range []string{`"last_7_days"`, `"refunded_units":1`, `"category_ranks"`, `"timezone":"Australia/Melbourne"`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("sales performance JSON = %s, want %s", payload, want)
		}
	}
}
