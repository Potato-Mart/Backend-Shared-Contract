package pricebook

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

func TestSellingPromotionTermsOmitAbsentFields(t *testing.T) {
	assertSellingJSON(t, SellingPromotionTerms{}, `{}`)
	assertSellingJSON(t, SellingPromotionGiftTerms{}, `{}`)
	assertSellingJSON(t, SellingPromotionDisplay{Kind: "bogo"},
		`{"kind":"bogo","conditional":false,"applied_to_effective_price":false}`)
}

func TestSellingPromotionTermsKeepZeroValuesOnPointers(t *testing.T) {
	zero := int64(0)
	assertSellingJSON(t, SellingPromotionTerms{DiscountBasisPoints: &zero, AddonPackagePrice: &money.Money{Currency: "AUD"}},
		`{"addon_package_price":{"amount_minor":0,"currency":"AUD"},"discount_basis_points":0}`)
	assertSellingJSON(t, SellingPromotionGiftTerms{GiftBaseUnits: &zero, MinimumOrderAmount: &money.Money{Currency: "AUD"}},
		`{"gift_base_units":0,"minimum_order_amount":{"amount_minor":0,"currency":"AUD"}}`)
}

func TestSellingPromotionTermsCarryBuyOneGetOne(t *testing.T) {
	qualifier, gift, perAward := int64(1), int64(1), int64(2)
	assertSellingJSON(t, SellingPromotionDisplay{Kind: "bogo", Conditional: true, Terms: &SellingPromotionTerms{QualifierBaseUnits: &qualifier, GiftBaseUnits: &gift, BaseUnitsPerAward: &perAward}},
		`{"kind":"bogo","conditional":true,"applied_to_effective_price":false,"terms":{"qualifier_base_units":1,"gift_base_units":1,"base_units_per_award":2}}`)
}

func TestSellingPromotionTermsRoundTripEveryField(t *testing.T) {
	number := func(value int64) *int64 { return &value }
	amount := func(minor int64) *money.Money { return &money.Money{AmountMinor: minor, Currency: "AUD"} }
	value := SellingPromotionDisplay{
		Kind: "spend_gift",
		Terms: &SellingPromotionTerms{
			QualifierBaseUnits:             number(2),
			GiftBaseUnits:                  number(1),
			BaseUnitsPerAward:              number(3),
			MinimumOrderAmount:             amount(5000),
			AddonPackagePrice:              amount(100),
			BundlePackagePrice:             amount(900),
			DiscountBasisPoints:            number(1000),
			DiscountAmount:                 amount(200),
			FixedPackagePrice:              amount(700),
			MaximumApplicationsPerCustomer: number(4),
			GroupPackageCount:              number(6),
			GiftTiers: []SellingPromotionGiftTerms{
				{GiftBaseUnits: number(1), MinimumOrderAmount: amount(5000)},
				{QualifierBaseUnits: number(3), GiftBaseUnits: number(2), MinimumOrderAmount: amount(10000)},
			},
		},
	}
	payload := assertSellingJSON(t, value, `{
		"kind":"spend_gift","conditional":false,"applied_to_effective_price":false,
		"terms":{
			"qualifier_base_units":2,"gift_base_units":1,"base_units_per_award":3,
			"minimum_order_amount":{"amount_minor":5000,"currency":"AUD"},
			"addon_package_price":{"amount_minor":100,"currency":"AUD"},
			"bundle_package_price":{"amount_minor":900,"currency":"AUD"},
			"discount_basis_points":1000,
			"discount_amount":{"amount_minor":200,"currency":"AUD"},
			"fixed_package_price":{"amount_minor":700,"currency":"AUD"},
			"maximum_applications_per_customer":4,
			"group_package_count":6,
			"gift_tiers":[
				{"gift_base_units":1,"minimum_order_amount":{"amount_minor":5000,"currency":"AUD"}},
				{"qualifier_base_units":3,"gift_base_units":2,"minimum_order_amount":{"amount_minor":10000,"currency":"AUD"}}
			]
		}
	}`)
	var decoded SellingPromotionDisplay
	if err := json.Unmarshal(payload, &decoded); err != nil || !reflect.DeepEqual(value, decoded) {
		t.Fatalf("round trip: %+v %v", decoded, err)
	}
}

func TestSellingPromotionTermsCarryNoExecutableRuleFields(t *testing.T) {
	number := func(value int64) *int64 { return &value }
	amount := &money.Money{AmountMinor: 1, Currency: "AUD"}
	payload, err := json.Marshal(SellingPromotionTerms{
		QualifierBaseUnits: number(1), GiftBaseUnits: number(1), BaseUnitsPerAward: number(2),
		MinimumOrderAmount: amount, AddonPackagePrice: amount, BundlePackagePrice: amount,
		DiscountBasisPoints: number(1), DiscountAmount: amount, FixedPackagePrice: amount,
		MaximumApplicationsPerCustomer: number(1), GroupPackageCount: number(1),
		GiftTiers: []SellingPromotionGiftTerms{{GiftBaseUnits: number(1)}},
	})
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded) != 12 {
		t.Fatalf("terms expose %d fields, want the 12 reviewed ones: %s", len(decoded), payload)
	}
	for _, forbidden := range []string{"rule_id", "rule_ids", "priority", "stacking", "stackable", "usage_count"} {
		if _, found := decoded[forbidden]; found {
			t.Fatalf("terms expose %s", forbidden)
		}
	}
}
