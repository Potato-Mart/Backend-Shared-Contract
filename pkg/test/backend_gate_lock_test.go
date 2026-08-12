package pkg_test

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"

	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/apiresponse/apiresponse_enums"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/market"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/pricebook"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/listing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging/packaging_enums"
)

// TestV27BackendGateModelSurface locks the reusable model primitives needed by
// the v27 stock, geography, market-pricing, listing, and availability gates.
// HTTP DTOs, stock commands, resolution rules, and error envelopes remain
// service-owned.
func TestV27BackendGateModelSurface(t *testing.T) {
	assertJSONFields(t, reflect.TypeOf(geography.Address{}), map[string]string{
		"Locality":           "locality",
		"AdministrativeArea": "administrative_area,omitempty",
		"PostalCode":         "postal_code",
		"Country":            "country",
	})
	assertJSONFields(t, reflect.TypeOf(geography.GeographicContext{}), map[string]string{
		"DepotCode":          "depot_code,omitempty",
		"ScopeRevision":      "scope_revision",
		"RuleRevision":       "rule_revision",
		"EvaluationTimezone": "evaluation_timezone",
	})
	assertJSONFields(t, reflect.TypeOf(promotion.PackagePricing{}), map[string]string{
		"ID":                    "id",
		"SKUID":                 "sku_id",
		"PackageOptionID":       "package_option_id",
		"PackagePrice":          "package_price",
		"StockLocation":         "stock_location",
		"AvailablePackageCount": "available_package_count",
		"Revision":              "revision",
		"InventoryRevision":     "inventory_revision",
		"GeographicContext":     "geographic_context",
		"PromotionApplications": "promotion_applications",
		"CapturedAt":            "captured_at",
	})
	assertJSONFields(t, reflect.TypeOf(operations.ProductStockSummary{}), map[string]string{
		"SKUID":        "sku_id",
		"Depots":       "depots,omitempty",
		"Revision":     "revision",
		"IsOutOfStock": "is_out_of_stock",
		"AsOf":         "as_of",
	})
	assertJSONFields(t, reflect.TypeOf(product.SKU{}), map[string]string{
		"ID":                 "id",
		"ProductID":          "product_id",
		"Code":               "code",
		"PackageOptions":     "package_options",
		"BarcodeAssignments": "barcode_assignments,omitempty",
		"NetContent":         "net_content,omitempty",
		"StorageType":        "storage_type,omitempty",
		"Status":             "status",
	})
	assertJSONFields(t, reflect.TypeOf(market.Market{}), map[string]string{
		"ID":               "id",
		"Code":             "code",
		"CountryCode":      "country_code",
		"DefaultCurrency":  "default_currency",
		"CurrencyExponent": "currency_exponent",
		"Status":           "status",
		"Revision":         "revision",
	})
	assertJSONFields(t, reflect.TypeOf(pricebook.PriceBook{}), map[string]string{
		"MarketID":         "market_id",
		"Currency":         "currency",
		"CurrencyExponent": "currency_exponent",
		"Channel":          "channel",
		"Audience":         "audience",
		"TaxInclusion":     "tax_inclusion",
		"PriceEnding":      "price_ending",
		"Status":           "status",
	})
	assertJSONFields(t, reflect.TypeOf(pricebook.PriceEntry{}), map[string]string{
		"PriceBookID": "price_book_id",
		"SKUID":       "sku_id",
		"Amount":      "amount",
		"Status":      "status",
		"Derivation":  "derivation",
		"Approval":    "approval,omitempty",
		"Revision":    "revision",
	})
	assertJSONFields(t, reflect.TypeOf(pricebook.PriceBookAssignment{}), map[string]string{
		"MarketID":             "market_id",
		"PriceBookID":          "price_book_id",
		"Kind":                 "kind",
		"OrganisationCategory": "organisation_category,omitempty",
		"OrganisationCode":     "organisation_code,omitempty",
	})
	assertJSONFields(t, reflect.TypeOf(listing.MarketListing{}), map[string]string{
		"MarketID":               "market_id",
		"SKUID":                  "sku_id",
		"Status":                 "status",
		"TaxCategoryID":          "tax_category_id",
		"ExpiryLeadDaysOverride": "expiry_lead_days_override,omitempty",
		"UnitPricingRequired":    "unit_pricing_required",
		"Revision":               "revision",
	})
	assertJSONFields(t, reflect.TypeOf(warehouse.DepotMarket{}), map[string]string{
		"DepotCode": "depot_code",
		"MarketID":  "market_id",
		"IsActive":  "is_active",
	})
	assertJSONFields(t, reflect.TypeOf(sales.BuyerContext{}), map[string]string{
		"Type":                 "type,omitempty",
		"RetailCustomerNumber": "retail_customer_number,omitempty",
		"FulfilmentIntent":     "fulfilment_intent,omitempty",
	})

	if !packaging_enums.PackageHandlingUnitEach.IsValid() || packaging_enums.PackageHandlingUnitEach.String() != "EACH" {
		t.Fatal("retail package pricing requires the canonical EACH handling unit")
	}
	for _, code := range []apiresponse_enums.Code{
		apiresponse_enums.CodeInsufficientStock,
		apiresponse_enums.CodeServiceUnavailable,
		apiresponse_enums.CodeConflict,
	} {
		if !code.IsValid() {
			t.Fatalf("backend gate response code %q must remain valid", code)
		}
	}
}

func assertJSONFields(t *testing.T, model reflect.Type, expected map[string]string) {
	t.Helper()
	for name, wantTag := range expected {
		field, ok := model.FieldByName(name)
		if !ok {
			t.Errorf("%s is missing required field %s", model, name)
			continue
		}
		if got := field.Tag.Get("json"); got != wantTag {
			t.Errorf("%s.%s JSON tag = %q, want %q", model, name, got, wantTag)
		}
	}
}
