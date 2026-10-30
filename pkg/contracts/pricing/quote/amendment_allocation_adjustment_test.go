package quote

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

func TestAmendmentAllocationAdjustmentJSON(t *testing.T) {
	amount := func(v int64) money.Money { return money.Money{Currency: "AUD", AmountMinor: v} }
	legacy := PriceSnapshot{LineID: "line"}
	raw, err := json.Marshal(legacy)
	if err != nil || strings.Contains(string(raw), `"amendment_allocation_adjustment"`) || strings.Contains(string(raw), `"unit_allocations"`) {
		t.Fatalf("legacy encoding changed: %s %v", raw, err)
	}
	legacy.SourceAllocationVersion = 2
	legacy.SourceAllocations = []PriceSourceAllocation{{FirstUnit: 1, BaseUnits: 1, UnitAllocationVersion: 1, UnitAllocations: []PriceUnitAllocation{{FirstUnit: 1, BaseUnits: 1, UnitListAmount: amount(100), UnitDiscountAmount: amount(0), UnitConsideration: amount(100), UnitTaxableBase: amount(91), UnitTaxAmount: amount(9), UnitLineTotal: amount(100)}}}}
	legacy.SourceAllocations[0].UnitAllocations[0].UnitConsiderationRounding = RoundingEvidence{ExactNumerator: 201, ExactDenominator: 2, RoundedAmount: amount(100), TieBreakKey: "line:1"}
	legacy.SourceAllocations[0].UnitAllocations[0].UnitTaxRounding = RoundingEvidence{ExactNumerator: 201, ExactDenominator: 22, RoundedAmount: amount(9), TieBreakKey: "line:1"}
	legacy.SourceAllocations[0].PromotionAllocations = []PricePromotionAllocation{{PromotionID: "promo", ReductionAmount: amount(1), ReductionRounding: &RoundingEvidence{ExactNumerator: 1, ExactDenominator: 2, RoundedAmount: amount(1), TieBreakKey: "line:1"}}}
	legacy.AmendmentAllocationAdjustment = &AmendmentAllocationAdjustment{Version: 1, SourceVersionID: "v", SourceItemID: "i", SourceComponentID: "c", SourceQuoteID: "q", SourceQuoteRevision: "revision", SourceQuoteFingerprint: "fingerprint", SourceLineID: "source", PackageCount: 2, UnitsPerPackage: 1, FirstPackageOrdinal: 1, LastPackageOrdinal: 1, FirstWithinPackageBaseUnitOrdinal: 1, LastWithinPackageBaseUnitOrdinal: 1, ListAmountDelta: amount(390), DiscountAmountDelta: amount(-10), ConsiderationDelta: amount(400), TaxableBaseDelta: amount(364), TaxAmountDelta: amount(36), LineTotalDelta: amount(400), TaxAdjustments: []AmendmentTaxAdjustment{{TaxableBaseDelta: amount(364), TaxAmountDelta: amount(36), TargetBaseRounding: RoundingEvidence{ExactNumerator: 455, ExactDenominator: 1, RoundedAmount: amount(455), TieBreakKey: "tax-group"}, TargetTaxRounding: RoundingEvidence{ExactNumerator: 45, ExactDenominator: 1, RoundedAmount: amount(45), TieBreakKey: "tax-group"}}}}
	raw, err = json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	var got PriceSnapshot
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.AmendmentAllocationAdjustment == nil || got.AmendmentAllocationAdjustment.LineTotalDelta.AmountMinor != 400 || got.AmendmentAllocationAdjustment.DiscountAmountDelta.AmountMinor != -10 || len(got.AmendmentAllocationAdjustment.TaxAdjustments) != 1 || got.AmendmentAllocationAdjustment.TaxAdjustments[0].TargetTaxRounding.TieBreakKey != "tax-group" || len(got.SourceAllocations) != 1 || len(got.SourceAllocations[0].UnitAllocations) != 1 || got.SourceAllocations[0].UnitAllocations[0].UnitLineTotal.AmountMinor != 100 || got.SourceAllocations[0].UnitAllocations[0].UnitTaxRounding.ExactDenominator != 22 || len(got.SourceAllocations[0].PromotionAllocations) != 1 || got.SourceAllocations[0].PromotionAllocations[0].ReductionRounding == nil || got.SourceAllocations[0].PromotionAllocations[0].ReductionRounding.ExactDenominator != 2 {
		t.Fatalf("unit or adjustment evidence lost: %+v", got)
	}
}
