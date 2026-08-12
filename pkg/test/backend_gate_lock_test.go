package pkg_test

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"

	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/apiresponse/apiresponse_enums"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/market"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/pricebook"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/quote"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/cost"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/listing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/purchase"
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
	assertJSONFields(t, reflect.TypeOf(quote.PriceSnapshot{}), map[string]string{
		"QuoteID":            "quote_id",
		"SKUID":              "sku_id",
		"MarketID":           "market_id",
		"PriceBookID":        "price_book_id",
		"PriceBookRevision":  "price_book_revision",
		"PriceEntryID":       "price_entry_id",
		"PriceEntryRevision": "price_entry_revision",
		"Currency":           "currency",
		"CurrencyExponent":   "currency_exponent",
		"ListUnitPrice":      "list_unit_price",
		"DiscountAmount":     "discount_amount",
		"TaxableBase":        "taxable_base",
		"TaxAmount":          "tax_amount",
		"LineTotal":          "line_total",
		"Tax":                "tax",
		"AppliedRules":       "applied_rules,omitempty",
		"Rounding":           "rounding",
		"UnitPrice":          "unit_price,omitempty",
		"CustomOverride":     "custom_override,omitempty",
		"Eligibility":        "eligibility",
		"Fingerprint":        "fingerprint",
		"CapturedAt":         "captured_at",
	})
	assertJSONFields(t, reflect.TypeOf(quote.TaxSnapshot{}), map[string]string{
		"TaxCategoryID":     "tax_category_id",
		"TaxRuleID":         "tax_rule_id",
		"TaxRuleRevision":   "tax_rule_revision",
		"InclusionBasis":    "inclusion_basis",
		"RateNumerator":     "rate_numerator",
		"RateDenominator":   "rate_denominator",
		"TaxableBase":       "taxable_base",
		"AllocatedTax":      "allocated_tax",
		"CalculationSource": "calculation_source",
		"RoundingMethod":    "rounding_method",
	})
	assertJSONFields(t, reflect.TypeOf(listing.SaleEligibilitySnapshot{}), map[string]string{
		"MarketID":           "market_id",
		"SKUID":              "sku_id",
		"ListingID":          "listing_id",
		"ListingRevision":    "listing_revision",
		"TaxCategoryID":      "tax_category_id",
		"StockLocation":      "stock_location",
		"Condition":          "condition",
		"Disposition":        "disposition",
		"DateMark":           "date_mark,omitempty",
		"DamageApproval":     "damage_approval,omitempty",
		"AvailableBaseUnits": "available_base_units",
		"InventoryRevision":  "inventory_revision",
		"ValidityToken":      "validity_token",
		"CapturedAt":         "captured_at",
	})
	assertJSONFields(t, reflect.TypeOf(payment.MerchantLegalSnapshot{}), map[string]string{
		"ProfileID":             "profile_id",
		"ProfileRevision":       "profile_revision",
		"MarketID":              "market_id",
		"LegalName":             "legal_name",
		"BusinessNumberScheme":  "business_number_scheme",
		"BusinessNumber":        "business_number",
		"TaxRegistrationStatus": "tax_registration_status",
		"Address":               "address",
		"FrozenAt":              "frozen_at",
	})
	assertJSONFields(t, reflect.TypeOf(pos.CashRoundingSnapshot{}), map[string]string{
		"ExactAmountDue":     "exact_amount_due",
		"RoundingAdjustment": "rounding_adjustment",
		"CashAmountDue":      "cash_amount_due",
		"IncrementMinor":     "increment_minor",
		"Mode":               "mode",
	})
	assertJSONFields(t, reflect.TypeOf(cost.BaseAcquisitionCost{}), map[string]string{
		"SKUID":    "sku_id",
		"Currency": "currency",
		"Amount":   "amount",
		"Revision": "revision",
	})
	assertJSONFields(t, reflect.TypeOf(cost.DepotCarryingCost{}), map[string]string{
		"SKUID":                  "sku_id",
		"DepotCode":              "depot_code",
		"TotalCarryingCostMinor": "total_carrying_cost_minor",
		"BaseUnitQuantity":       "base_unit_quantity",
		"CurrentAverageUnitCost": "current_average_unit_cost,omitempty",
	})
	assertJSONFields(t, reflect.TypeOf(purchase.SupplierInvoiceLine{}), map[string]string{
		"SKUID":         "sku_id",
		"SKUCode":       "sku_code",
		"TaxTreatment":  "tax_treatment",
		"PriceBasis":    "price_basis",
		"DeclaredTax":   "declared_tax,omitempty",
		"CalculatedTax": "calculated_tax,omitempty",
		"TaxSource":     "tax_source",
		"InputTaxClaim": "input_tax_claim",
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
