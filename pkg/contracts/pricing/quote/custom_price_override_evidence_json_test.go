package quote

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/pricing/quote/quote_enums"
)

func TestCustomPriceOverrideJSONPreservesPriceBases(t *testing.T) {
	value := customPriceFixture()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"actor_user_id":"cashier_1","reason":"Approved clearance for damaged outer packaging","source_approved_price":{"amount_minor":800,"currency":"AUD"},"source_price_basis":"base_unit","source_tax_inclusion":"tax_exclusive","override_gross_amount":{"amount_minor":500,"currency":"AUD"},"override_price_basis":"package","cost_comparison":"unavailable","below_cost_warning":false,"overridden_at":"2026-09-08T00:00:00Z"}`
	if string(payload) != want {
		t.Fatalf("custom override JSON = %s, want %s", payload, want)
	}
	var decoded CustomPriceOverrideEvidence
	if err := json.Unmarshal([]byte(want), &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(value, decoded) {
		t.Fatalf("custom override evidence changed: %+v", decoded)
	}
}

func TestCustomPriceOverrideStructuredReasonAndDefinitionRoundTrip(t *testing.T) {
	for _, reason := range []quote_enums.CustomPriceReason{
		quote_enums.CustomPriceReasonQuickSale,
		quote_enums.CustomPriceReasonSoonExpiry,
		quote_enums.CustomPriceReasonDamaged,
	} {
		t.Run(string(reason), func(t *testing.T) {
			revision := int64(3)
			evidence := customPriceFixture()
			evidence.ReasonCode = reason
			evidence.CustomPriceCode = "CLEARANCE_AU_001"
			evidence.CustomPriceRevision = &revision
			payload, err := json.Marshal(evidence)
			if err != nil {
				t.Fatal(err)
			}
			want := `{"actor_user_id":"cashier_1","reason":"Approved clearance for damaged outer packaging","reason_code":"` + string(reason) + `","custom_price_code":"CLEARANCE_AU_001","custom_price_revision":3,"source_approved_price":{"amount_minor":800,"currency":"AUD"},"source_price_basis":"base_unit","source_tax_inclusion":"tax_exclusive","override_gross_amount":{"amount_minor":500,"currency":"AUD"},"override_price_basis":"package","cost_comparison":"unavailable","below_cost_warning":false,"overridden_at":"2026-09-08T00:00:00Z"}`
			if string(payload) != want {
				t.Fatalf("structured custom override JSON = %s, want %s", payload, want)
			}
			snapshot := PriceSnapshot{CustomOverride: &evidence}
			frozen, err := json.Marshal(snapshot)
			if err != nil {
				t.Fatal(err)
			}
			var decoded PriceSnapshot
			if err := json.Unmarshal(frozen, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(snapshot, decoded) {
				t.Fatalf("frozen snapshot changed custom evidence: %+v", decoded.CustomOverride)
			}
		})
	}
}

func TestCustomPriceRevisionDistinguishesMissingFromZero(t *testing.T) {
	// Zero remains explicit model data; acceptance and revision validation
	// belong to Pricing, not to JSON marshaling in Shared Contract.
	zero := int64(0)
	value := customPriceFixture()
	value.CustomPriceRevision = &zero
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(payload), `"custom_price_revision":0`) {
		t.Fatalf("explicit zero revision was omitted: %s", payload)
	}
}

func TestHistoricalCustomPriceBasesRemainUnknown(t *testing.T) {
	raw := `{"actor_user_id":"staff","reason":"recorded reason","source_approved_price":{"amount_minor":800,"currency":"AUD"},"override_gross_amount":{"amount_minor":500,"currency":"AUD"}}`
	var got CustomPriceOverrideEvidence
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatal(err)
	}
	if got.SourcePriceBasis != "" || got.OverridePriceBasis != "" || got.SourceTaxInclusion != "" {
		t.Fatalf("missing historical bases must stay unknown: %+v", got)
	}
}

func TestCustomPriceOverridePreservesIndependentAmountBases(t *testing.T) {
	for _, tc := range []struct {
		name      string
		source    quote_enums.PriceAmountBasis
		override  quote_enums.PriceAmountBasis
		sourceTax pricebook_enums.PriceTaxInclusion
	}{
		{"manual package price with exclusive base price", quote_enums.PriceAmountBasisBaseUnit, quote_enums.PriceAmountBasisPackage, pricebook_enums.PriceTaxInclusionExclusive},
		{"approved base-unit price with inclusive package price", quote_enums.PriceAmountBasisPackage, quote_enums.PriceAmountBasisBaseUnit, pricebook_enums.PriceTaxInclusionInclusive},
	} {
		t.Run(tc.name, func(t *testing.T) {
			want := customPriceFixture()
			want.SourcePriceBasis, want.OverridePriceBasis, want.SourceTaxInclusion = tc.source, tc.override, tc.sourceTax
			raw, err := json.Marshal(want)
			if err != nil {
				t.Fatal(err)
			}
			var got CustomPriceOverrideEvidence
			if err := json.Unmarshal(raw, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("price basis changed: %#v", got)
			}
		})
	}
}

func customPriceFixture() CustomPriceOverrideEvidence {
	return CustomPriceOverrideEvidence{
		ActorUserID: "cashier_1", Reason: "Approved clearance for damaged outer packaging",
		SourceApprovedPrice: money.Money{AmountMinor: 800, Currency: "AUD"},
		SourcePriceBasis:    quote_enums.PriceAmountBasisBaseUnit,
		SourceTaxInclusion:  pricebook_enums.PriceTaxInclusionExclusive,
		OverrideGrossAmount: money.Money{AmountMinor: 500, Currency: "AUD"},
		OverridePriceBasis:  quote_enums.PriceAmountBasisPackage,
		CostComparison:      quote_enums.CostComparisonUnavailable,
		OverriddenAt:        time.Date(2026, 9, 8, 0, 0, 0, 0, time.UTC),
	}
}
