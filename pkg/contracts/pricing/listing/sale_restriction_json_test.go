package listing

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/listing/listing_enums"
)

func TestRetailAgeRestrictionCarriesThresholdWithoutBuyerOrProofData(t *testing.T) {
	for _, comparison := range []listing_enums.AgeComparison{listing_enums.AgeComparisonBelow, listing_enums.AgeComparisonAbove} {
		for _, years := range []int32{0, 18, 65} {
			value := SaleRestriction{Kind: listing_enums.SaleRestrictionKindAge, AgeYears: &years, AgeComparison: comparison}
			payload, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var got SaleRestriction
			if err := json.Unmarshal(payload, &got); err != nil {
				t.Fatal(err)
			}
			if got.Kind != value.Kind || got.AgeYears == nil || *got.AgeYears != years || got.AgeComparison != comparison {
				t.Fatalf("age policy lost: %s", payload)
			}
			var shape map[string]json.RawMessage
			if err := json.Unmarshal(payload, &shape); err != nil {
				t.Fatal(err)
			}
			if len(shape) != 3 {
				t.Fatalf("age policy must carry only kind/threshold/comparison, got %s", payload)
			}
			for _, forbidden := range []string{"date_of_birth", "customer_id", "proof", "verified_age", "visibility", "buyer_types"} {
				if strings.Contains(string(payload), forbidden) {
					t.Fatalf("buyer data or configurable applicability leaked: %s", payload)
				}
			}
		}
	}
	// The closed allowed-group vocabulary has no separate equality operator.
	// Retail applicability and completed-age evaluation are owner
	// invariants; this model deliberately contains no DOB evaluator or audience.
	for _, invalid := range []listing_enums.AgeComparison{"", "equal", "below_or_equal", "above_or_equal", "future"} {
		if invalid.IsValid() {
			t.Fatalf("unknown allowed-group comparison accepted: %q", invalid)
		}
	}
	if listing_enums.SaleRestrictionKind("age_verification").IsValid() {
		t.Fatal("proof-verification kind must be absent")
	}
}

func TestExcludedDeliveryMethodsRoundTripCanonicalValues(t *testing.T) {
	methods := []shipping_enums.DeliveryMethod{shipping_enums.DeliveryMethodDelivery, shipping_enums.DeliveryMethodPickup, shipping_enums.DeliveryMethodOutsourced}
	for _, method := range methods {
		if !method.IsValid() {
			t.Fatalf("invalid canonical method: %s", method)
		}
	}
	value := SaleRestriction{Kind: listing_enums.SaleRestrictionKindDeliveryExcluded, ExcludedDeliveryMethods: methods}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if string(payload) != `{"kind":"delivery_excluded","excluded_delivery_methods":["delivery","pickup","outsourced"]}` {
		t.Fatalf("delivery wire = %s", payload)
	}
	var got SaleRestriction
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, value) {
		t.Fatalf("delivery policy changed: %+v", got)
	}
}

func TestOtherRestrictionKindsRetainWireAndOmitAgeDeliveryFields(t *testing.T) {
	for _, value := range []SaleRestriction{
		{Kind: listing_enums.SaleRestrictionKindQuantityLimit, Value: 4, Note: "limit"},
		{Kind: listing_enums.SaleRestrictionKindChannelExcluded, Channels: []commerce_enums.OrderType{commerce_enums.OrderTypeOnline}},
		{Kind: listing_enums.SaleRestrictionKindPrescription, Note: "prescription"},
	} {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		for _, absent := range []string{"age_years", "age_comparison", "excluded_delivery_methods"} {
			if strings.Contains(string(payload), absent) {
				t.Fatalf("unrelated payload leaked: %s", payload)
			}
		}
		var got SaleRestriction
		if err := json.Unmarshal(payload, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, value) {
			t.Fatalf("restriction changed: %+v", got)
		}
	}
}
