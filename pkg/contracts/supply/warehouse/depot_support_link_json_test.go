package warehouse_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
)

func TestDepotSupportLinkNamesShopSupportMarketAndReceivingLocation(t *testing.T) {
	effectiveFrom := time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC)
	value := warehouse.DepotSupportLink{
		ID: "support_link_1", ShopDepotCode: "A7", SupportDepotCode: "D3", MarketCode: "AU-VIC",
		Priority: 1, ReceivingLocationCode: "A7-BACK-01", IsActive: true, EffectiveFrom: effectiveFrom,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal depot support link: %v", err)
	}
	for _, want := range []string{
		`"shop_depot_code":"A7"`, `"support_depot_code":"D3"`, `"market_code":"AU-VIC"`,
		`"priority":1`, `"receiving_location_code":"A7-BACK-01"`, `"is_active":true`,
		`"effective_from":"2026-11-09T00:00:00Z"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("DepotSupportLink JSON = %s, want %s", payload, want)
		}
	}
	if strings.Contains(string(payload), `"effective_to"`) {
		t.Fatalf("an open-ended link must omit effective_to: %s", payload)
	}
	for _, forbidden := range []string{`"on_hand_base_units"`, `"currency"`, `"country_code"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("DepotSupportLink leaked stock or commercial data %s: %s", forbidden, payload)
		}
	}

	var decoded warehouse.DepotSupportLink
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal depot support link: %v", err)
	}
	if decoded != value {
		t.Fatalf("depot support link did not round-trip: %+v", decoded)
	}

	// A supporting depot may support several shops and a shop may have
	// several supporting depots.
	closedAt := effectiveFrom.AddDate(0, 6, 0)
	second := warehouse.DepotSupportLink{
		ID: "support_link_2", ShopDepotCode: "A8", SupportDepotCode: "D3", MarketCode: "AU-VIC",
		Priority: 2, ReceivingLocationCode: "A8-BACK-01", EffectiveFrom: effectiveFrom, EffectiveTo: &closedAt,
	}
	both, err := json.Marshal([]warehouse.DepotSupportLink{value, second})
	if err != nil {
		t.Fatalf("marshal depot support link set: %v", err)
	}
	if strings.Count(string(both), `"support_depot_code":"D3"`) != 2 {
		t.Fatalf("one supporting depot must be able to support several shops: %s", both)
	}
	if !strings.Contains(string(both), `"effective_to":"2027-05-09T00:00:00Z"`) {
		t.Fatalf("a closed link must record effective_to: %s", both)
	}
}
