package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

func TestApprovedStorefrontProductJSONShape(t *testing.T) {
	p := wholesale.ApprovedStorefrontProduct{
		ID:                "prd_1",
		SKUCode:           "SKU-1",
		Name:              "Bulk Potatoes",
		Description:       []common.LocalizedDescription{{Language: "en", Description: "Foodservice carton"}},
		Brand:             []common.LocalizedName{{Language: "en", Name: "Potato Mart"}},
		Supplier:          "SUP-1",
		BrandKey:          "potato-mart",
		Storage:           enums.StorageDry,
		DisplayStatus:     "active",
		CoverURL:          "https://example.com/potatoes.jpg",
		CategoryKey:       "produce",
		CategoryPath:      []string{"produce", "potatoes"},
		Price:             common.Money{AmountMinor: 2500, Currency: "AUD"},
		StockAvailable:    true,
		AvailabilityState: "available",
	}

	raw, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	for _, key := range []string{
		"id",
		"sku_code",
		"name",
		"description",
		"brand",
		"supplier",
		"brand_key",
		"storage",
		"display_status",
		"cover_url",
		"category_key",
		"category_path",
		"price",
		"stock_available",
		"availability_state",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}

	for _, forbidden := range []string{"pricing", "online", "original", "cost", "current_stock", "history"} {
		if _, ok := got[forbidden]; ok {
			t.Fatalf("approved wholesale storefront product leaked %q in %s", forbidden, raw)
		}
	}
}
