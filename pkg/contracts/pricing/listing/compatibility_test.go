package listing_test

import (
	"encoding/json"
	"reflect"
	"testing"

	pricing "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/listing"
	pe "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/listing/listing_enums"
	supply "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/supply/catalogue/listing"
	se "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/supply/catalogue/listing/listing_enums"
)

func assertShape(t *testing.T, a, b reflect.Type) {
	t.Helper()
	if a == b {
		return
	}
	if a.Kind() != b.Kind() {
		t.Fatalf("kind mismatch: %v / %v", a, b)
	}
	switch a.Kind() {
	case reflect.Struct:
		if a.NumField() != b.NumField() {
			t.Fatalf("field count mismatch: %v / %v", a, b)
		}
		for i := 0; i < a.NumField(); i++ {
			x, y := a.Field(i), b.Field(i)
			if x.Name != y.Name || x.Tag != y.Tag || x.Anonymous != y.Anonymous {
				t.Fatalf("field mismatch: %+v / %+v", x, y)
			}
			assertShape(t, x.Type, y.Type)
		}
	case reflect.Slice, reflect.Pointer:
		assertShape(t, a.Elem(), b.Elem())
	default:
		if a.Name() != b.Name() {
			t.Fatalf("type mismatch: %v / %v", a, b)
		}
	}
}

func TestSupplyPricingListingWireCompatibility(t *testing.T) {
	assertShape(t, reflect.TypeOf(pricing.MarketListing{}), reflect.TypeOf(supply.MarketListing{}))
	fixture := []byte(`{"id":"listing_1","market_code":"AU","country_code":"AU","sku_code":"sku_1","status":"active","tax_category_code":"GST","restrictions":[{"kind":"quantity_limit","value":4,"note":"limit"}],"expiry_lead_days_override":0,"unit_pricing_required":true,"available_from":"2026-10-01T00:00:00Z","available_until":"2026-11-01T00:00:00Z","revision":7}`)
	for _, data := range [][]byte{fixture, []byte(`{"market_code":"AU","sku_code":"sku_1","status":"draft"}`)} {
		var current pricing.MarketListing
		var legacy supply.MarketListing
		if err := json.Unmarshal(data, &current); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(data, &legacy); err != nil {
			t.Fatal(err)
		}
		x, err := json.Marshal(current)
		if err != nil {
			t.Fatal(err)
		}
		y, err := json.Marshal(legacy)
		if err != nil {
			t.Fatal(err)
		}
		if string(x) != string(y) {
			t.Fatalf("wire mismatch: %s / %s", x, y)
		}
	}
}

func TestSupplyPricingListingEnumCompatibility(t *testing.T) {
	for _, pair := range [][2]string{
		{string(pe.MarketListingStatusDraft), string(se.MarketListingStatusDraft)},
		{string(pe.MarketListingStatusComingSoon), string(se.MarketListingStatusComingSoon)},
		{string(pe.MarketListingStatusActive), string(se.MarketListingStatusActive)},
		{string(pe.MarketListingStatusSuspended), string(se.MarketListingStatusSuspended)},
		{string(pe.MarketListingStatusUnavailable), string(se.MarketListingStatusUnavailable)},
		{string(pe.MarketListingStatusDelisted), string(se.MarketListingStatusDelisted)},
	} {
		if pair[0] != pair[1] || !pe.MarketListingStatus(pair[0]).IsValid() || !se.MarketListingStatus(pair[1]).IsValid() {
			t.Fatalf("status mismatch: %v", pair)
		}
	}
	for _, pair := range [][2]string{
		{string(pe.SaleRestrictionKindAgeVerification), string(se.SaleRestrictionKindAgeVerification)},
		{string(pe.SaleRestrictionKindQuantityLimit), string(se.SaleRestrictionKindQuantityLimit)},
		{string(pe.SaleRestrictionKindChannelExcluded), string(se.SaleRestrictionKindChannelExcluded)},
		{string(pe.SaleRestrictionKindDeliveryExcluded), string(se.SaleRestrictionKindDeliveryExcluded)},
		{string(pe.SaleRestrictionKindPrescription), string(se.SaleRestrictionKindPrescription)},
	} {
		if pair[0] != pair[1] || !pe.SaleRestrictionKind(pair[0]).IsValid() || !se.SaleRestrictionKind(pair[1]).IsValid() {
			t.Fatalf("restriction mismatch: %v", pair)
		}
	}
	for _, unknown := range []string{"", "future_value"} {
		if pe.MarketListingStatus(unknown).IsValid() || pe.SaleRestrictionKind(unknown).IsValid() {
			t.Fatalf("unknown enum accepted: %q", unknown)
		}
	}
}
