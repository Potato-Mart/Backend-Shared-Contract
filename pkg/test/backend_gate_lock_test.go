package pkg_test

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"

	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/apiresponse/apiresponse_enums"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging/packaging_enums"
)

// TestV27BackendGateModelSurface locks the reusable model primitives needed by
// the v27 stock, geography, package-pricing, and availability gates. HTTP DTOs, stock
// commands, resolution rules, and error envelopes remain service-owned.
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
		"ProductSKUCode":        "product_sku_code",
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
		"ProductSKUCode": "product_sku_code",
		"Depots":         "depots,omitempty",
		"Revision":       "revision",
		"IsOutOfStock":   "is_out_of_stock",
		"AsOf":           "as_of",
	})
	assertJSONFields(t, reflect.TypeOf(product.ProductCommerce{}), map[string]string{
		"Status":        "status,omitempty",
		"Selling":       "selling,omitempty",
		"FirstListedAt": "first_listed_at,omitempty",
		"Packages":      "packages,omitempty",
	})
	assertJSONFields(t, reflect.TypeOf(product.ProductPackageCommerce{}), map[string]string{
		"PackageOptionID":       "package_option_id",
		"PackagePrice":          "package_price",
		"TaxAmount":             "tax_amount",
		"StockState":            "stock_state,omitempty",
		"PromotionApplications": "promotion_applications,omitempty",
		"AsOf":                  "as_of",
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
