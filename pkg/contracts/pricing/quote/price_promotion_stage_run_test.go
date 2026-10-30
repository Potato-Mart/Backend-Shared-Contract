package quote

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
)

func TestPromotionStageEvidenceIsAdditiveAndRoundTrips(t *testing.T) {
	legacy, err := json.Marshal(PriceSnapshot{})
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{`"promotion_stage_version"`, `"promotion_stages"`, `"admin_discount_witness"`} {
		if strings.Contains(string(legacy), field) {
			t.Fatalf("legacy snapshot gained %s", field)
		}
	}
	amount := money.Money{Currency: "AUD", AmountMinor: 401}
	eligible := false
	lockStage := int32(1)
	snapshot := PriceSnapshot{PromotionStageVersion: 1, AdminDiscountWitness: &MixedAdminLineWitness{Version: 1, LineID: "added"}, PromotionStages: []PricePromotionStageRun{{
		PromotionID: "offer", PromotionRevision: 2, StageOrdinal: 2, Role: "target", FirstUnit: 1, BaseUnits: 2,
		PackageOption: packaging.PackageOptionRef{SKUCode: "SKU", Code: "EACH", Version: 1}, UnitsPerPackage: 1, FirstPackageOrdinal: 1, FirstWithinPackageUnit: 1,
		Currency: "AUD", CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2}, EligibleUnitBasis: PromotionExactMinor{Numerator: 1001, Denominator: 2}, BeforeUnit: PromotionExactMinor{Numerator: 1001, Denominator: 2}, ReductionUnit: PromotionExactMinor{Numerator: 1001, Denominator: 10}, AfterUnit: PromotionExactMinor{Numerator: 2002, Denominator: 5}, SettledAfterUnit: &amount,
		GroupSpan: &PromotionGroupSpan{CandidateID: "candidate", FirstGroup: 1, GroupCount: 2, UnitsPerGroup: 1, FirstGroupUnitOffset: 1}, ContributionRoundingPool: &PromotionStageRoundingPool{ID: "peers", RootLineID: "line", TieBreakRule: "line_unit_ordinal_v1"}, SettledAwardEligible: &eligible, SettledLockStage: &lockStage, UnitContributionRounding: &RoundingEvidence{TieBreakKey: "00000000000000000005"}, Lineage: &PromotionStageLineage{RootQuoteID: "quote", RootQuoteFingerprint: "fingerprint", RootLineID: "line", RootFirstUnit: 5},
	}}}
	raw, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	var got PriceSnapshot
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.PromotionStageVersion != 1 || got.AdminDiscountWitness == nil || len(got.PromotionStages) != 1 || got.PromotionStages[0].GroupSpan.CandidateID != "candidate" || got.PromotionStages[0].Lineage.RootFirstUnit != 5 || got.PromotionStages[0].SettledAfterUnit.AmountMinor != 401 || got.PromotionStages[0].SettledAwardEligible == nil || *got.PromotionStages[0].SettledAwardEligible || *got.PromotionStages[0].SettledLockStage != 1 {
		t.Fatalf("evidence lost: %+v", got)
	}
}

func TestMixedAdminPredecessorEvidenceIsAdditive(t *testing.T) {
	legacy, err := json.Marshal(MixedAdminDiscountWitness{Version: 1})
	if err != nil || strings.Contains(string(legacy), "predecessor_evidence_source") || strings.Contains(string(legacy), "predecessor_administrative_discount_hash") {
		t.Fatalf("legacy witness changed: %s %v", legacy, err)
	}
	in := MixedAdminDiscountWitness{Version: 2, PredecessorEvidenceSource: "pricing_lookup", PredecessorAdministrativeDiscountHash: "hash"}
	raw, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out MixedAdminDiscountWitness
	if err := json.Unmarshal(raw, &out); err != nil || out.PredecessorEvidenceSource != in.PredecessorEvidenceSource || out.PredecessorAdministrativeDiscountHash != in.PredecessorAdministrativeDiscountHash {
		t.Fatalf("predecessor evidence lost: %+v %v", out, err)
	}
}
