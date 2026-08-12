package promotion_test

import (
	"encoding/json"
	"testing"
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestPackagePricingJSONPreservesExplicitEmptyPromotionApplications(t *testing.T) {
	now := time.Date(2026, 8, 7, 4, 5, 6, 0, time.UTC)
	value := promotion.PackagePricing{
		ID: "pricing_1", Revision: 4, InventoryRevision: 11,
		SKUID: "A00001", PackageOptionID: "pkg_case_6",
		PackagePrice: money.Money{AmountMinor: 1200, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 109, Currency: "AUD"},
		ValidFrom: now, Timezone: "Australia/Melbourne",
		GeographicContext: geography.GeographicContext{Source: geography_enums.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU", ScopeRevision: 3, RuleRevision: 4, EvaluationTimezone: "Australia/Melbourne"},
		StockLocation:     warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"},
		SourceBucketID:    "bucket_1", SourceStockUnitID: "stock_unit_1",
		AvailablePackageCount: 7, AvailableBaseUnits: 42,
		Condition: warehouse_enums.InventoryConditionGood, Disposition: warehouse_enums.InventoryDispositionStandardSellable,
		PromotionApplications: []promotion.PromotionApplication{},
		CapturedAt:            now,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal package pricing: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal package pricing: %v", err)
	}
	for _, key := range []string{"id", "revision", "inventory_revision", "sku_id", "package_option_id", "package_price", "tax_amount", "valid_from", "timezone", "geographic_context", "stock_location", "available_package_count", "available_base_units", "condition", "disposition", "promotion_applications", "captured_at"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("package pricing JSON missing %q: %s", key, payload)
		}
	}
	applications, ok := shape["promotion_applications"].([]any)
	if !ok || len(applications) != 0 {
		t.Fatalf("zero promotion applications must remain an explicit empty array: %s", payload)
	}
	for _, removed := range []string{"package_option", "product_snapshot", "offer", "discounts"} {
		if _, ok := shape[removed]; ok {
			t.Fatalf("package pricing retained removed field %q: %s", removed, payload)
		}
	}
}
