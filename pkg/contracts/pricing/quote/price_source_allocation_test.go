package quote

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
)

func TestPriceSnapshotSourceAllocationsJSON(t *testing.T) {
	amount := func(v int64) money.Money { return money.Money{Currency: "AUD", AmountMinor: v} }
	snapshot := PriceSnapshot{LineID: "line-a", BaseUnits: 2}
	legacy, err := json.Marshal(snapshot)
	if err != nil || strings.Contains(string(legacy), `"source_allocations"`) || strings.Contains(string(legacy), `"source_allocation_version"`) {
		t.Fatalf("legacy encoding changed: %s %v", legacy, err)
	}
	snapshot.SourceAllocationVersion = 1
	snapshot.SourceAllocations = []PriceSourceAllocation{
		{FirstUnit: 1, BaseUnits: 1, PriceBookCode: "pos", PriceBookRevision: 1, PriceEntryID: "pos-entry", PriceEntryRevision: 1, ListUnitPrice: amount(1000), ListAmount: amount(1000), DiscountAmount: amount(0), Consideration: amount(1000), Tax: TaxSnapshot{InclusionBasis: pricebook_enums.PriceTaxInclusionExclusive, AllocatedTax: amount(100)}, TaxRounding: RoundingEvidence{ExactNumerator: 100, ExactDenominator: 1, RoundedAmount: amount(100), TieBreakKey: "line-a:1"}, LineTotal: amount(1100)},
		{FirstUnit: 2, BaseUnits: 1, PriceBookCode: "online", PriceBookRevision: 2, PriceEntryID: "online-entry", PriceEntryRevision: 3, ListUnitPrice: amount(1100), ListAmount: amount(1100), DiscountAmount: amount(220), Consideration: amount(880), Tax: TaxSnapshot{InclusionBasis: pricebook_enums.PriceTaxInclusionInclusive, AllocatedTax: amount(80)}, TaxRounding: RoundingEvidence{ExactNumerator: 80, ExactDenominator: 1, RoundedAmount: amount(80), TieBreakKey: "line-a:2"}, LineTotal: amount(880), PromotionAllocations: []PricePromotionAllocation{{PromotionID: "p", PromotionRevision: 4, Role: "target", FirstUnit: 2, BaseUnits: 1, ReductionAmount: amount(220)}}},
	}
	raw, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	var decoded PriceSnapshot
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.SourceAllocationVersion != 1 || len(decoded.SourceAllocations) != 2 || decoded.SourceAllocations[0].FirstUnit != 1 || decoded.SourceAllocations[1].Tax.InclusionBasis != pricebook_enums.PriceTaxInclusionInclusive || decoded.SourceAllocations[1].LineTotal.AmountMinor != 880 || decoded.SourceAllocations[1].TaxRounding.TieBreakKey != "line-a:2" || len(decoded.SourceAllocations[1].PromotionAllocations) != 1 || decoded.SourceAllocations[1].PromotionAllocations[0].PromotionRevision != 4 {
		t.Fatalf("source evidence was not preserved: %+v", decoded.SourceAllocations)
	}
}
