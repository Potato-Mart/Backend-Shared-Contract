package warehouse_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
)

func TestDepotSupportLinkNamesShopSupportMarketAndReceivingLocation(t *testing.T) {
	effectiveFrom := time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC)
	value := warehouse.DepotSupportLink{
		ID: "support_link_1", ShopDepotCode: "A7", SupportDepotCode: "D3", MarketCode: "AU-VIC",
		Priority: 1, IsActive: true, EffectiveFrom: effectiveFrom,
		ReceivingLocations: []warehouse.DepotSupportReceivingLocation{
			{StorageType: classification_enums.StorageAmbient, LocationCode: "A7-UNSLOTTED-AMBIENT"},
			{StorageType: classification_enums.StorageFrozen, LocationCode: "A7-UNSLOTTED-FROZEN"},
		},
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal depot support link: %v", err)
	}
	for _, want := range []string{
		`"shop_depot_code":"A7"`, `"support_depot_code":"D3"`, `"market_code":"AU-VIC"`,
		`"priority":1`, `"is_active":true`,
		`"receiving_locations":[{"storage_type":"AMBIENT","location_code":"A7-UNSLOTTED-AMBIENT"},{"storage_type":"FROZEN","location_code":"A7-UNSLOTTED-FROZEN"}]`,
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
	if !reflect.DeepEqual(decoded, value) {
		t.Fatalf("depot support link did not round-trip: %+v", decoded)
	}

	// A supporting depot may support several shops and a shop may have
	// several supporting depots.
	closedAt := effectiveFrom.AddDate(0, 6, 0)
	second := warehouse.DepotSupportLink{
		ID: "support_link_2", ShopDepotCode: "A8", SupportDepotCode: "D3", MarketCode: "AU-VIC",
		Priority: 2, EffectiveFrom: effectiveFrom, EffectiveTo: &closedAt,
		ReceivingLocations: []warehouse.DepotSupportReceivingLocation{
			{StorageType: classification_enums.StorageAmbient, LocationCode: "A8-UNSLOTTED-AMBIENT"},
		},
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
