package warehouse_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
)

func TestShopRefillLevelCarriesMinMaxInBaseUnits(t *testing.T) {
	value := warehouse.ShopRefillLevel{
		ID: "refill_1", ShopDepotCode: "A7", SKUCode: "SKU-APPLE-JUICE-1L",
		MinBaseUnits: 6, MaxBaseUnits: 24, IsActive: true,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal shop refill level: %v", err)
	}
	for _, want := range []string{
		`"shop_depot_code":"A7"`, `"sku_code":"SKU-APPLE-JUICE-1L"`,
		`"min_base_units":6`, `"max_base_units":24`, `"is_active":true`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("ShopRefillLevel JSON = %s, want %s", payload, want)
		}
	}
	for _, forbidden := range []string{`"support_depot_code"`, `"case_count"`, `"currency"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("ShopRefillLevel leaked %s: %s", forbidden, payload)
		}
	}

	var decoded warehouse.ShopRefillLevel
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal shop refill level: %v", err)
	}
	if decoded != value {
		t.Fatalf("shop refill level did not round-trip: %+v", decoded)
	}
}
