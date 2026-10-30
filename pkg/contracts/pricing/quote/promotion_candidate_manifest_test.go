package quote

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPromotionManifestLineageIsAdditive(t *testing.T) {
	legacy, err := json.Marshal(PricePromotionStageRun{Lineage: &PromotionStageLineage{RootQuoteID: "quote", RootQuoteFingerprint: "fingerprint", RootLineID: "line", RootFirstUnit: 1}})
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"manifest_id", "requirement_index", "root_quote_key", "root_quote_revision", "root_manifest_id"} {
		if strings.Contains(string(legacy), `"`+name+`"`) {
			t.Fatalf("legacy wire gained %s", name)
		}
	}
	value := PricePromotionStageRun{ManifestID: "manifest", RequirementIndex: 2, Lineage: &PromotionStageLineage{RootQuoteKey: "key", RootQuoteRevision: "revision", RootManifestID: "manifest", RootQuoteID: "quote", RootQuoteFingerprint: "fingerprint", RootLineID: "line", RootFirstUnit: 3}}
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var got PricePromotionStageRun
	if err := json.Unmarshal(raw, &got); err != nil || got.ManifestID != value.ManifestID || got.RequirementIndex != 2 || got.Lineage == nil || got.Lineage.RootQuoteKey != "key" || got.Lineage.RootQuoteRevision != "revision" || got.Lineage.RootManifestID != "manifest" {
		t.Fatalf("manifest lineage lost: %+v %v", got, err)
	}
}

func TestPromotionManifestPreservesPeerAndAuditIdentity(t *testing.T) {
	proof := PromotionManifestEvidence{Version: 1, Manifests: []PromotionCandidateManifest{{Version: 1, ManifestID: "manifest", PromotionID: "p", CompetitionClass: "product_competing", CandidateID: "candidate", RequirementCount: 1, Requirements: []PromotionCandidateRequirement{{Index: 1, Role: "qualifier", WholePackage: true, UnitsPerGroup: 2, Peers: []PromotionCandidatePeer{{RootLineID: "line", TaxRuleID: "gst", TaxRuleRevision: 2, TaxRateNumerator: 1, TaxRateDenominator: 11, Spans: []PromotionCandidatePeerSpan{{FirstGroup: 1, GroupCount: 1, FirstUnit: 3}}}}}}}}, StageAudits: []PromotionStageAudit{{ID: "audit", RootLineID: "line", FirstUnit: 3, BaseUnits: 2, Entries: []PromotionStageAuditEntry{}}}}
	raw, err := json.Marshal(proof)
	if err != nil {
		t.Fatal(err)
	}
	var got PromotionManifestEvidence
	if err := json.Unmarshal(raw, &got); err != nil || len(got.Manifests) != 1 || len(got.Manifests[0].Requirements) != 1 || !got.Manifests[0].Requirements[0].WholePackage || got.Manifests[0].Requirements[0].Peers[0].TaxRuleRevision != 2 || len(got.StageAudits) != 1 || got.StageAudits[0].Entries == nil {
		t.Fatalf("manifest evidence lost: %+v %v", got, err)
	}
}
