package product_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/product"
	geographyenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/geography"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

func TestStorefrontProductJSONUsesPackageOffersAndAvailability(t *testing.T) {
	now := time.Date(2026, 8, 4, 1, 2, 3, 0, time.UTC)
	discount := 20
	option := product.ProductPackageOptionSnapshot{
		ID: "pkg_case_12", Code: "CASE-12", ProductSKUCode: "A00001",
		HandlingUnit: common.PackageHandlingUnitCase, UnitsPerPackage: 12,
		EffectiveFrom: now, CapturedAt: now,
	}
	projection := product.StorefrontProduct{
		SKUCode: "A00001", CategorySKUCode: "CAT-A", Name: "Product",
		BrandRef: &product.BrandRef{
			ID: "64c13ab08edf48a008793ca1", Slug: "happy-potato",
			Name: []common.LocalizedName{{Language: "en", Name: "Happy Potato"}},
		},
		PackageOptions: []product.ProductPackageOptionSnapshot{option},
		BarcodeAssignments: []product.ProductBarcodeAssignmentSnapshot{
			{ID: "barcode_1", ProductSKUCode: "A00001", PackageOptionID: "pkg_case_12", Value: "930000000001", Format: productenum.BarcodeFormatEAN13, IsPrimary: true, EffectiveFrom: now, CapturedAt: now},
		},
		Offers: []product.SellableOfferSnapshot{
			{
				ID: "offer_1", ProductSKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01",
				PackageOption: option, AvailablePackageCount: 2, AvailableBaseUnits: 24,
				Condition: warehouseenum.InventoryConditionGood, Disposition: warehouseenum.InventoryDispositionStandardSellable,
				Revision: 4, InventoryRevision: 9,
				PackagePrice: common.Money{AmountMinor: 800, Currency: "AUD"}, TaxAmount: common.Money{AmountMinor: 73, Currency: "AUD"},
				ValidFrom: now, Timezone: "Etc/UTC", CapturedAt: now,
				GeographicContext: common.GeographicContext{Source: geographyenum.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU", ScopeRevision: 2, RuleRevision: 4, EvaluationTimezone: "Australia/Melbourne"},
			},
		},
		Availability: &product.ProductStockSummary{
			ProductSKUCode: "A00001",
			AllDepots:      product.ProductStockQuantitySnapshot{AvailableBaseUnits: 24, SellableBaseUnits: 24},
			Revision:       9, Timezone: "Australia/Melbourne", AsOf: now,
		},
		Audience:          productenum.PriceAudienceRetail,
		StorefrontDisplay: product.StorefrontDisplay{},
		PromotionBadge: &product.StorefrontPromotionBadge{
			PromotionID: "promo_1", SeriesKey: "potato-august", DiscountPercent: &discount,
			ScheduleTimezone: "Etc/UTC",
			GeographicContext: common.GeographicContext{
				Source: geographyenum.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU",
				MatchedTargetKind: geographyenum.GeographicTargetCountry, MatchedTargetCode: "AU",
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
	for _, key := range []string{"sku_code", "category_sku_code", "package_options", "barcode_assignments", "offers", "availability", "audience"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("storefront product JSON missing %s: %s", key, payload)
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
