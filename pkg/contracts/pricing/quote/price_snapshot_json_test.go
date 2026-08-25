package quote

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/quote/quote_enums"
)

func TestTaxSnapshotKeepsTheRateExactAsAFraction(t *testing.T) {
	// A tax-inclusive Australian consideration extracts tax with the exact
	// fraction 1/11; a tax-exclusive one adds it with 1/10. Neither may be
	// stored as a decimal approximation.
	for _, tc := range []struct {
		basis       pricebook_enums.PriceTaxInclusion
		source      quote_enums.TaxCalculationSource
		numerator   int64
		denominator int64
	}{
		{pricebook_enums.PriceTaxInclusionInclusive, quote_enums.TaxCalculationSourceInclusiveExtraction, 1, 11},
		{pricebook_enums.PriceTaxInclusionExclusive, quote_enums.TaxCalculationSourceExclusiveAddition, 1, 10},
	} {
		payload, err := json.Marshal(TaxSnapshot{
			TaxCategoryCode: "tax_au_gst", TaxRuleID: "rule_au_gst", TaxRuleRevision: 1,
			InclusionBasis: tc.basis, RateNumerator: tc.numerator, RateDenominator: tc.denominator,
			TaxableBase:       money.Money{AmountMinor: 1100, Currency: "AUD"},
			AllocatedTax:      money.Money{AmountMinor: 100, Currency: "AUD"},
			CalculationSource: tc.source,
			RoundingMethod:    quote_enums.TaxRoundingMethodSumExactThenRound,
		})
		if err != nil {
			t.Fatalf("marshal tax snapshot: %v", err)
		}
		var decoded TaxSnapshot
		if err := json.Unmarshal(payload, &decoded); err != nil {
			t.Fatalf("unmarshal tax snapshot: %v", err)
		}
		if decoded.RateNumerator != tc.numerator || decoded.RateDenominator != tc.denominator {
			t.Fatalf("exact rate changed: %+v", decoded)
		}
		if decoded.RoundingMethod != quote_enums.TaxRoundingMethodSumExactThenRound {
			t.Fatalf("rounding method changed: %+v", decoded)
		}
		if strings.Contains(string(payload), "0.0909") || strings.Contains(string(payload), "0.09") {
			t.Fatalf("superseded decimal tax factor must not appear: %s", payload)
		}
	}
}

func TestRoundingEvidenceRecordsTheExactValueAndTieBreak(t *testing.T) {
	payload, err := json.Marshal(RoundingEvidence{
		Mode:        quote_enums.RoundingModeLargestRemainder,
		PriceEnding: pricebook_enums.PriceEndingPolicyNone,
		Exponent:    2,
		// 100/11 exact minor units becomes 9 with a remainder that the
		// document allocates by largest remainder, tie-broken on line ID.
		ExactNumerator: 100, ExactDenominator: 11,
		RoundedAmount:         money.Money{AmountMinor: 9, Currency: "AUD"},
		RemainderRank:         1,
		RemainderMinorApplied: 1,
		TieBreakKey:           "line-1",
	})
	if err != nil {
		t.Fatalf("marshal rounding evidence: %v", err)
	}
	for _, want := range []string{
		`"mode":"largest_remainder"`, `"exact_numerator":100`, `"exact_denominator":11`,
		`"remainder_rank":1`, `"remainder_minor_applied":1`, `"tie_break_key":"line-1"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("RoundingEvidence JSON = %s, want %s", payload, want)
		}
	}
}

func TestAppliedPriceRuleKeepsExclusiveFactorsExact(t *testing.T) {
	appliedAt := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	// Group 15/20 and the fixed damage tier mappings are exact fractions,
	// never one minus a damage percentage.
	for kind, factor := range map[string][2]int64{
		"group_15":       {85, 100},
		"group_20":       {80, 100},
		"expiry_20":      {80, 100},
		"damage_tier_30": {8, 10},
		"damage_tier_50": {5, 10},
		"damage_tier_80": {3, 10},
	} {
		payload, err := json.Marshal(AppliedPriceRule{
			Kind: kind, Exclusive: true,
			FactorNumerator: factor[0], FactorDenominator: factor[1],
			AmountBefore: money.Money{AmountMinor: 1000, Currency: "AUD"},
			AmountAfter:  money.Money{AmountMinor: factor[0] * 1000 / factor[1], Currency: "AUD"},
			AppliedAt:    appliedAt,
		})
		if err != nil {
			t.Fatalf("marshal applied rule: %v", err)
		}
		var decoded AppliedPriceRule
		if err := json.Unmarshal(payload, &decoded); err != nil {
			t.Fatalf("unmarshal applied rule: %v", err)
		}
		if decoded.Kind != kind || !decoded.Exclusive || decoded.FactorNumerator != factor[0] || decoded.FactorDenominator != factor[1] {
			t.Fatalf("exclusive rule evidence changed for %s: %+v", kind, decoded)
		}
	}

	bogo, err := json.Marshal(AppliedPriceRule{
		Kind: "expiry_bogo", Exclusive: true, ChargeableBaseUnits: 1,
		AmountBefore: money.Money{AmountMinor: 400, Currency: "AUD"},
		AmountAfter:  money.Money{AmountMinor: 200, Currency: "AUD"},
		AppliedAt:    appliedAt,
	})
	if err != nil {
		t.Fatalf("marshal bogo rule: %v", err)
	}
	if !strings.Contains(string(bogo), `"chargeable_base_units":1`) {
		t.Fatalf("a buy-one-get-one rule must record the chargeable quantity: %s", bogo)
	}
}

func TestCustomPriceOverrideEvidenceFreezesActorReasonAndCostComparison(t *testing.T) {
	overriddenAt := time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC)
	cost := money.Money{AmountMinor: 550, Currency: "AUD"}
	payload, err := json.Marshal(CustomPriceOverrideEvidence{
		ActorUserID:         "cashier_1",
		Reason:              "manager approved clearance",
		SourceApprovedPrice: money.Money{AmountMinor: 800, Currency: "AUD"},
		OverrideGrossAmount: money.Money{AmountMinor: 500, Currency: "AUD"},
		CostComparison:      quote_enums.CostComparisonBelowCost,
		ComparedCost:        &cost,
		BelowCostWarning:    true,
		OverriddenAt:        overriddenAt,
	})
	if err != nil {
		t.Fatalf("marshal custom override: %v", err)
	}
	for _, want := range []string{
		`"actor_user_id":"cashier_1"`, `"reason":"manager approved clearance"`,
		`"source_approved_price"`, `"override_gross_amount"`,
		`"cost_comparison":"below_cost"`, `"below_cost_warning":true`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("CustomPriceOverrideEvidence JSON = %s, want %s", payload, want)
		}
	}

	unavailable, err := json.Marshal(CustomPriceOverrideEvidence{
		ActorUserID: "cashier_1", Reason: "damaged pack",
		SourceApprovedPrice: money.Money{AmountMinor: 800, Currency: "AUD"},
		OverrideGrossAmount: money.Money{AmountMinor: 500, Currency: "AUD"},
		CostComparison:      quote_enums.CostComparisonUnavailable,
		OverriddenAt:        overriddenAt,
	})
	if err != nil {
		t.Fatalf("marshal unavailable-cost override: %v", err)
	}
	if strings.Contains(string(unavailable), `"compared_cost"`) {
		t.Fatalf("an unavailable cost comparison must omit compared_cost: %s", unavailable)
	}
}

func TestUnitPriceEvidenceCarriesComparisonAmountOrExemption(t *testing.T) {
	payload, err := json.Marshal(UnitPriceEvidence{
		NetContent: measurement.NetContent{
			NetQuantity:     measurement.Measure{Amount: 500, Unit: "g"},
			StandardMeasure: measurement.Measure{Amount: 100, Unit: "g"},
		},
		ComparisonAmount: money.Money{AmountMinor: 64, Currency: "AUD"},
	})
	if err != nil {
		t.Fatalf("marshal unit price evidence: %v", err)
	}
	if !strings.Contains(string(payload), `"comparison_amount":{"amount_minor":64,"currency":"AUD"}`) {
		t.Fatalf("UnitPriceEvidence JSON = %s", payload)
	}
	if strings.Contains(string(payload), `"exemption_reason"`) {
		t.Fatalf("a non-exempt listing must omit the exemption reason: %s", payload)
	}

	exempt, err := json.Marshal(UnitPriceEvidence{
		NetContent: measurement.NetContent{
			NetQuantity:     measurement.Measure{Amount: 500, Unit: "g"},
			StandardMeasure: measurement.Measure{Amount: 100, Unit: "g"},
		},
		ComparisonAmount: money.Money{AmountMinor: 51, Currency: "AUD"},
		Exempt:           true,
		ExemptionReason:  "soon_expiry_markdown",
	})
	if err != nil {
		t.Fatalf("marshal exempt unit price evidence: %v", err)
	}
	if !strings.Contains(string(exempt), `"exemption_reason":"soon_expiry_markdown"`) {
		t.Fatalf("an exempt markdown must record its reason: %s", exempt)
	}
}
