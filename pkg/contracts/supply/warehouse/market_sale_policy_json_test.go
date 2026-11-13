package warehouse_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestMarketSalePolicyJSONShape(t *testing.T) {
	payload, err := json.Marshal(warehouse.MarketSalePolicy{
		MarketCode: "AU-VIC", SellPastDateMark: warehouse_enums.SellPastDateMarkBestBefore, Stored: true, Revision: 2,
	})
	if err != nil {
		t.Fatalf("marshal market sale policy: %v", err)
	}
	for _, want := range []string{`"market_code":"AU-VIC"`, `"sell_past_date_mark":"best_before"`, `"stored":true`, `"revision":2`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("MarketSalePolicy JSON = %s, want %s", payload, want)
		}
	}
	if strings.Contains(string(payload), `"history"`) {
		t.Fatalf("MarketSalePolicy JSON = %s, want empty history omitted", payload)
	}
}

func TestPastDateMarkNoticeOmitsUnitsUnlessRequested(t *testing.T) {
	at := time.Date(2026, 11, 1, 13, 0, 0, 0, time.UTC)
	notice := warehouse.PastDateMarkNotice{Kind: warehouse_enums.InventoryDateMarkExpiry, DateMarkAt: at, Timezone: "Australia/Melbourne"}
	payload, err := json.Marshal(notice)
	if err != nil {
		t.Fatalf("marshal notice: %v", err)
	}
	if got, want := string(payload), `{"kind":"EXPIRY","date_mark_at":"2026-11-01T13:00:00Z","timezone":"Australia/Melbourne"}`; got != want {
		t.Fatalf("PastDateMarkNotice JSON = %s, want %s", got, want)
	}
	units := int64(3)
	notice.PastDateBaseUnits = &units
	payload, err = json.Marshal(notice)
	if err != nil {
		t.Fatalf("marshal notice: %v", err)
	}
	if !strings.Contains(string(payload), `"past_date_base_units":3`) {
		t.Fatalf("PastDateMarkNotice JSON = %s, want past_date_base_units", payload)
	}
}
