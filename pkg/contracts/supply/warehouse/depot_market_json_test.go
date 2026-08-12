package warehouse_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse"
)

func TestDepotMarketAssociatesPhysicalSitesWithCommercialMarkets(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	value := warehouse.DepotMarket{
		ID: "depot_market_1", DepotCode: "AU-VIC-MEL-DC-01", MarketID: "market_au",
		IsActive: true, EffectiveFrom: effectiveFrom,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal depot market: %v", err)
	}
	for _, want := range []string{
		`"depot_code":"AU-VIC-MEL-DC-01"`, `"market_id":"market_au"`,
		`"is_active":true`, `"effective_from":"2026-08-12T00:00:00Z"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("DepotMarket JSON = %s, want %s", payload, want)
		}
	}
	if strings.Contains(string(payload), `"effective_to"`) {
		t.Fatalf("an open-ended association must omit effective_to: %s", payload)
	}
	// A depot services a market commercially; it never carries the market's
	// country, currency, or price configuration.
	for _, forbidden := range []string{`"country_code"`, `"currency"`, `"price_book_id"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("DepotMarket leaked commercial configuration %s: %s", forbidden, payload)
		}
	}

	var decoded warehouse.DepotMarket
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal depot market: %v", err)
	}
	if decoded.DepotCode != value.DepotCode || decoded.MarketID != value.MarketID || !decoded.IsActive {
		t.Fatalf("depot market did not round-trip: %+v", decoded)
	}

	// One depot may service several markets, so the association is many to
	// many and Depot.Country is never assumed to equal Market.Country.
	closedAt := effectiveFrom.AddDate(1, 0, 0)
	second := warehouse.DepotMarket{
		ID: "depot_market_2", DepotCode: "AU-VIC-MEL-DC-01", MarketID: "market_nz",
		IsActive: false, EffectiveFrom: effectiveFrom, EffectiveTo: &closedAt,
	}
	both, err := json.Marshal([]warehouse.DepotMarket{value, second})
	if err != nil {
		t.Fatalf("marshal depot market set: %v", err)
	}
	if strings.Count(string(both), `"depot_code":"AU-VIC-MEL-DC-01"`) != 2 {
		t.Fatalf("one depot must be able to service several markets: %s", both)
	}
	if !strings.Contains(string(both), `"effective_to":"2027-08-12T00:00:00Z"`) {
		t.Fatalf("a closed association must record effective_to: %s", both)
	}
}
