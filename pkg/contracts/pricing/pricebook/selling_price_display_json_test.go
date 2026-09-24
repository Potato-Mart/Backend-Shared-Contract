package pricebook

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/pricing/market/market_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/pricing/pricebook/pricebook_enums"
)

func TestSellingPriceLegacyJSONShapeIsUnchanged(t *testing.T) {
	instant := time.Date(2026, 9, 8, 0, 0, 0, 0, time.UTC)
	value := SellingPrice{
		UnitPrice:        money.Money{AmountMinor: 500, Currency: "AUD"},
		CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		MarketCode:       "market_au", Channel: commerce_enums.OrderTypeOnline,
		Audience:        market_enums.PriceAudienceRetail,
		PriceVisibility: pricebook_enums.PriceVisibilityVisible,
		TaxInclusion:    pricebook_enums.PriceTaxInclusionInclusive,
		ValidFrom:       instant, AsOf: instant,
	}
	payload := assertSellingJSON(t, value, `{
		"unit_price":{"amount_minor":500,"currency":"AUD"},
		"currency_exponent":{"currency":"AUD","exponent":2},
		"market_code":"market_au","channel":"online","audience":"retail",
		"price_visibility":"visible","tax_inclusion":"tax_inclusive",
		"valid_from":"2026-09-08T00:00:00Z","as_of":"2026-09-08T00:00:00Z"
	}`)
	var decoded SellingPrice
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(value, decoded) {
		t.Fatalf("legacy SellingPrice changed after round trip: %+v", decoded)
	}
}

func TestSellingPriceDisplayExactJSONAndRoundTrip(t *testing.T) {
	from := time.Date(2026, 9, 8, 0, 0, 0, 0, time.UTC)
	until := from.Add(7 * 24 * time.Hour)
	regular := money.Money{AmountMinor: 500, Currency: "AUD"}
	comparison := money.Money{AmountMinor: 80, Currency: "AUD"}
	value := SellingPriceDisplay{
		RegularUnitPrice:   regular,
		EffectiveUnitPrice: money.Money{AmountMinor: 400, Currency: "AUD"},
		CompareAtUnitPrice: &regular,
		UnitPricePerMeasure: &SellingUnitPriceDisplay{
			NetContent: measurement.NetContent{
				NetQuantity:     measurement.Measure{Amount: 500, Unit: "g"},
				StandardMeasure: measurement.Measure{Amount: 100, Unit: "g"},
			},
			ComparisonAmount: &comparison,
		},
		PromotionDisplays: []SellingPromotionDisplay{{
			Kind:                    "weekly_promotion",
			Messages:                []localization.LocalizedText{{Language: "en", Text: "Weekly special"}, {Language: "zh-TW", Text: "本週特價"}},
			AppliedToEffectivePrice: true, ValidFrom: &from, ValidUntil: &until,
		}},
	}
	payload := assertSellingJSON(t, value, `{
		"regular_unit_price":{"amount_minor":500,"currency":"AUD"},
		"effective_unit_price":{"amount_minor":400,"currency":"AUD"},
		"compare_at_unit_price":{"amount_minor":500,"currency":"AUD"},
		"unit_price_per_measure":{
			"net_content":{"net_quantity":{"amount":500,"exponent":0,"unit":"g"},"standard_measure":{"amount":100,"exponent":0,"unit":"g"}},
			"comparison_amount":{"amount_minor":80,"currency":"AUD"},"exempt":false
		},
		"promotions":[{"kind":"weekly_promotion","messages":[{"language":"en","text":"Weekly special"},{"language":"zh-TW","text":"本週特價"}],
			"conditional":false,"applied_to_effective_price":true,"valid_from":"2026-09-08T00:00:00Z","valid_until":"2026-09-15T00:00:00Z"}]
	}`)
	var decoded SellingPriceDisplay
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(value, decoded) {
		t.Fatalf("display evidence changed after round trip: %+v", decoded)
	}
}

func TestSellingUnitPriceDisplayDistinguishesAbsentAndZeroAmounts(t *testing.T) {
	net := measurement.NetContent{
		NetQuantity:     measurement.Measure{Amount: 125, Exponent: -1, Unit: "g"},
		StandardMeasure: measurement.Measure{Amount: 100, Unit: "g"},
	}
	assertSellingJSON(t, SellingUnitPriceDisplay{
		NetContent: net, Exempt: true, ExemptionReason: "soon_expiry_markdown",
	}, `{"net_content":{"net_quantity":{"amount":125,"exponent":-1,"unit":"g"},"standard_measure":{"amount":100,"exponent":0,"unit":"g"}},"exempt":true,"exemption_reason":"soon_expiry_markdown"}`)
	zero := money.Money{AmountMinor: 0, Currency: "AUD"}
	assertSellingJSON(t, SellingUnitPriceDisplay{NetContent: net, ComparisonAmount: &zero},
		`{"net_content":{"net_quantity":{"amount":125,"exponent":-1,"unit":"g"},"standard_measure":{"amount":100,"exponent":0,"unit":"g"}},"comparison_amount":{"amount_minor":0,"currency":"AUD"},"exempt":false}`)
	assertSellingJSON(t, SellingPriceDisplay{RegularUnitPrice: zero, EffectiveUnitPrice: zero},
		`{"regular_unit_price":{"amount_minor":0,"currency":"AUD"},"effective_unit_price":{"amount_minor":0,"currency":"AUD"}}`)
	assertSellingJSON(t, SellingPriceDisplay{RegularUnitPrice: zero, EffectiveUnitPrice: zero, CompareAtUnitPrice: &zero},
		`{"regular_unit_price":{"amount_minor":0,"currency":"AUD"},"effective_unit_price":{"amount_minor":0,"currency":"AUD"},"compare_at_unit_price":{"amount_minor":0,"currency":"AUD"}}`)
}

func TestConditionalSellingPromotionsCarryConditionsWithoutImplyingSavings(t *testing.T) {
	for _, kind := range []string{"bogo", "group_order", "volume_discount", "future_configured_promotion"} {
		t.Run(kind, func(t *testing.T) {
			regular := money.Money{AmountMinor: 500, Currency: "AUD"}
			value := SellingPriceDisplay{
				RegularUnitPrice: regular, EffectiveUnitPrice: regular,
				PromotionDisplays: []SellingPromotionDisplay{{
					Kind: kind, Conditional: true,
					Conditions: []localization.LocalizedText{{Language: "en", Text: "Qualifying purchases only"}},
				}},
			}
			assertSellingJSON(t, value, `{
				"regular_unit_price":{"amount_minor":500,"currency":"AUD"},
				"effective_unit_price":{"amount_minor":500,"currency":"AUD"},
				"promotions":[{"kind":"`+kind+`","conditions":[{"language":"en","text":"Qualifying purchases only"}],"conditional":true,"applied_to_effective_price":false}]
			}`)
		})
	}
}

func TestSellingPriceDisplayRetainsIndependentChannelAndAudience(t *testing.T) {
	for _, tc := range []struct {
		channel  commerce_enums.OrderType
		audience market_enums.PriceAudience
	}{
		{commerce_enums.OrderTypeOnline, market_enums.PriceAudienceRetail},
		{commerce_enums.OrderTypePOS, market_enums.PriceAudienceRetail},
		{commerce_enums.OrderTypeB2B, market_enums.PriceAudienceWholesale},
		{commerce_enums.OrderTypePOS, market_enums.PriceAudienceWholesale},
	} {
		t.Run(string(tc.channel)+"/"+string(tc.audience), func(t *testing.T) {
			amount := money.Money{AmountMinor: 400, Currency: "AUD"}
			value := SellingPrice{
				UnitPrice: amount, Channel: tc.channel, Audience: tc.audience,
				Display: &SellingPriceDisplay{RegularUnitPrice: amount, EffectiveUnitPrice: amount},
			}
			payload, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var decoded SellingPrice
			if err := json.Unmarshal(payload, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(value, decoded) {
				t.Fatalf("channel/audience or display evidence changed: %+v", decoded)
			}
		})
	}
}

func TestSellingPriceOfferCarriesPackageMemberAndConditionalDisplayEvidence(t *testing.T) {
	from := time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC)
	until := from.Add(7 * 24 * time.Hour)
	compareAt := money.Money{AmountMinor: 3600, Currency: "AUD"}
	value := SellingPrice{
		UnitPrice:        money.Money{AmountMinor: 300, Currency: "AUD"},
		CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		MarketCode:       "AU-VIC", Channel: commerce_enums.OrderTypePOS,
		Audience: market_enums.PriceAudienceRetail, MembershipTierKey: "gold",
		PriceVisibility: pricebook_enums.PriceVisibilityVisible,
		TaxInclusion:    pricebook_enums.PriceTaxInclusionInclusive,
		ValidFrom:       from, AsOf: from,
		Display: &SellingPriceDisplay{
			RegularUnitPrice:   money.Money{AmountMinor: 300, Currency: "AUD"},
			EffectiveUnitPrice: money.Money{AmountMinor: 300, Currency: "AUD"},
			Offers: []SellingPriceOffer{{
				PackageOptionCode: "CASE_12", MembershipTierKey: "gold", BaseUnits: 12,
				RegularAmount:     money.Money{AmountMinor: 3600, Currency: "AUD"},
				EffectiveAmount:   money.Money{AmountMinor: 2999, Currency: "AUD"},
				CompareAtAmount:   &compareAt,
				Messages:          []localization.LocalizedText{{Language: "en", Text: "Gold member case price"}},
				Conditions:        []localization.LocalizedText{{Language: "en", Text: "Gold membership required"}},
				Conditional:       true,
				PromotionDisplays: []SellingPromotionDisplay{{Kind: "member_package_price", Conditional: true}},
				ValidFrom:         from, ValidUntil: &until,
			}},
		},
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal selling price offer: %v", err)
	}
	for _, want := range []string{
		`"membership_tier_key":"gold"`, `"offers":[`, `"package_option_code":"CASE_12"`,
		`"base_units":12`, `"regular_amount":{"amount_minor":3600,"currency":"AUD"}`,
		`"effective_amount":{"amount_minor":2999,"currency":"AUD"}`,
		`"compare_at_amount":{"amount_minor":3600,"currency":"AUD"}`,
		`"conditional":true`, `"valid_until":"2026-09-27T00:00:00Z"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("SellingPrice offer JSON = %s, want %s", payload, want)
		}
	}

	var decoded SellingPrice
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal selling price offer: %v", err)
	}
	if !reflect.DeepEqual(value, decoded) {
		t.Fatalf("selling price offer changed after round trip: %+v", decoded)
	}
}

func assertSellingJSON(t *testing.T, value any, want string) []byte {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	var gotShape, wantShape any
	if err := json.Unmarshal(payload, &gotShape); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(want), &wantShape); err != nil {
		t.Fatalf("invalid expected JSON: %v", err)
	}
	if !reflect.DeepEqual(gotShape, wantShape) {
		t.Fatalf("%T JSON = %s, want %s", value, payload, want)
	}
	return payload
}
