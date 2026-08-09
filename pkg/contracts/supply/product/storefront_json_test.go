package product_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestStorefrontProductJSONUsesPackageOffersAndAvailability(t *testing.T) {
	now := time.Date(2026, 8, 4, 1, 2, 3, 0, time.UTC)
	discount := 20
	option := product.ProductPackageOptionSnapshot{
		ID: "pkg_case_12", Code: "CASE-12", ProductSKUCode: "A00001",
		HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 12,
		EffectiveFrom: now, CapturedAt: now,
	}
	projection := product.StorefrontProduct{
		SKUCode: "A00001", CategorySKUCode: "CAT-A", Name: "Product",
		BrandRef: &product.BrandRef{
			ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato",
			Name: []localization.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		},
		PackageOptions: []product.ProductPackageOptionSnapshot{option},
		BarcodeAssignments: []product.ProductBarcodeAssignmentSnapshot{
			{ID: "barcode_1", ProductSKUCode: "A00001", PackageOptionID: "pkg_case_12", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, IsPrimary: true, EffectiveFrom: now, CapturedAt: now},
		},
		Offers: []product.SellableOfferSnapshot{
			{
				ID: "offer_1", ProductSKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01",
				PackageOption: option, AvailablePackageCount: 2, AvailableBaseUnits: 24,
				Condition: warehouse_enums.InventoryConditionGood, Disposition: warehouse_enums.InventoryDispositionStandardSellable,
				Revision: 4, InventoryRevision: 9,
				PackagePrice: money.Money{AmountMinor: 800, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 73, Currency: "AUD"},
				ValidFrom: now, Timezone: "Etc/UTC", CapturedAt: now,
				GeographicContext: geography.GeographicContext{Source: geography_enums.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU", ScopeRevision: 2, RuleRevision: 4, EvaluationTimezone: "Australia/Melbourne"},
			},
		},
		Availability: &product.ProductStockSummary{
			ProductSKUCode: "A00001",
			AllDepots:      product.ProductStockQuantitySnapshot{AvailableBaseUnits: 24, SellableBaseUnits: 24},
			Revision:       9, Timezone: "Australia/Melbourne", AsOf: now,
		},
		Commercial: &product.StorefrontCommercial{
			Price: &money.Money{AmountMinor: 800, Currency: "AUD"},
			Package: product.ProductPackageOptionSnapshot{
				ID: "pkg_each", Code: "EACH", ProductSKUCode: "A00001",
				HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1,
				EffectiveFrom: now, CapturedAt: now,
			},
			StockState: product_enums.StorefrontStockStateInStock,
			Market:     "AU",
			AsOf:       now,
		},
		Audience:          product_enums.PriceAudienceRetail,
		StorefrontDisplay: product.StorefrontDisplay{},
		PromotionBadge: &product.StorefrontPromotionBadge{
			PromotionID: "promo_1", SeriesKey: "potato-august", DiscountPercent: &discount,
			ScheduleTimezone: "Etc/UTC",
			GeographicContext: geography.GeographicContext{
				Source: geography_enums.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU",
				MatchedTargetKind: geography_enums.GeographicTargetCountry, MatchedTargetCode: "AU",
				ScopeRevision: 2, RuleRevision: 4, EvaluationTimezone: "Australia/Melbourne",
			},
		},
	}
	payload, err := json.Marshal(projection)
	if err != nil {
		t.Fatalf("marshal storefront product: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal storefront product: %v", err)
	}
	for _, key := range []string{"sku_code", "category_sku_code", "package_options", "barcode_assignments", "offers", "availability", "commercial", "audience"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("storefront product JSON missing %s: %s", key, payload)
		}
	}
	commercial := got["commercial"].(map[string]any)
	if commercial["market"] != "AU" || commercial["stock_state"] != "in_stock" || commercial["as_of"] != "2026-08-04T01:02:03Z" {
		t.Fatalf("commercial projection JSON mismatch: %s", payload)
	}
	packageOption := commercial["package_option"].(map[string]any)
	if packageOption["handling_unit"] != "EACH" || packageOption["units_per_package"] != float64(1) {
		t.Fatalf("commercial package option mismatch: %s", payload)
	}
	commercialJSON, _ := json.Marshal(product.StorefrontCommercial{
		Price: &money.Money{AmountMinor: 800, Currency: "AUD"},
		Package: product.ProductPackageOptionSnapshot{
			ID: "pkg_each", Code: "EACH", ProductSKUCode: "A00001",
			HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1,
			EffectiveFrom: now, CapturedAt: now,
		},
		StockState: product_enums.StorefrontStockStateInStock, Market: "AU", AsOf: now,
	})
	for _, forbidden := range []string{`"depot_code"`, `"lot_id"`, `"available_base_units"`, `"geographic_context"`} {
		if strings.Contains(string(commercialJSON), forbidden) {
			t.Fatalf("commercial projection leaked %s: %s", forbidden, commercialJSON)
		}
	}
	for _, removed := range []string{"sku", "current_stock", "pricing", "expiry_date", "display_status", "physical_weight", "barcode"} {
		if _, ok := got[removed]; ok {
			t.Fatalf("storefront product JSON contains removed %s: %s", removed, payload)
		}
	}
	badge := got["promotion_badge"].(map[string]any)
	if badge["series_key"] != "potato-august" || badge["schedule_timezone"] != "Etc/UTC" {
		t.Fatalf("promotion badge lacks geographic schedule identity: %s", payload)
	}
	if _, ok := badge["geographic_context"].(map[string]any); !ok {
		t.Fatalf("promotion badge lacks resolved geographic context: %s", payload)
	}
}
