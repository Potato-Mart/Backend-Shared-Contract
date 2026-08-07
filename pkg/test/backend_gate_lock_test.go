package pkg_test

import (
	apiresponseenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/apiresponse"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
	"reflect"
	"testing"
)

// TestV23BackendGateModelSurface locks the reusable model primitives needed by
// the V23 stock, geography, offer, and availability gates. HTTP DTOs, stock
// commands, resolution rules, and error envelopes remain service-owned.
func TestV23BackendGateModelSurface(t *testing.T) {
	assertJSONFields(t, reflect.TypeOf(common.Address{}), map[string]string{
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
	assertJSONFields(t, reflect.TypeOf(product.SellableOfferSnapshot{}), map[string]string{
		"ID":                    "id",
		"ProductSKUCode":        "product_sku_code",
		"DepotCode":             "depot_code",
		"PackageOption":         "package_option",
		"AvailablePackageCount": "available_package_count",
		"Revision":              "revision",
		"InventoryRevision":     "inventory_revision",
		"GeographicContext":     "geographic_context",
		"CapturedAt":            "captured_at",
	})
	assertJSONFields(t, reflect.TypeOf(product.ProductStockSummary{}), map[string]string{
		"ProductSKUCode": "product_sku_code",
		"Depots":         "depots,omitempty",
		"Revision":       "revision",
		"IsOutOfStock":   "is_out_of_stock",
		"AsOf":           "as_of",
	})
	assertJSONFields(t, reflect.TypeOf(product.StorefrontCommercial{}), map[string]string{
		"Price":      "price,omitempty",
		"Package":    "package_option",
		"StockState": "stock_state",
		"Market":     "market",
		"AsOf":       "as_of",
	})
	assertJSONFields(t, reflect.TypeOf(sales.BuyerContext{}), map[string]string{
		"Type":                 "type,omitempty",
		"RetailCustomerNumber": "retail_customer_number,omitempty",
		"FulfilmentIntent":     "fulfilment_intent,omitempty",
	})

	if !common.PackageHandlingUnitEach.IsValid() || common.PackageHandlingUnitEach.String() != "EACH" {
		t.Fatal("retail offer resolution requires the canonical EACH handling unit")
	}
	for _, code := range []apiresponseenum.Code{
		apiresponseenum.CodeInsufficientStock,
		apiresponseenum.CodeServiceUnavailable,
		apiresponseenum.CodeConflict,
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
