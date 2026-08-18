package promotion

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
)

func TestPromotionApplicationsFreezeVisibleQualifierTargetRelationships(t *testing.T) {
	appliedAt := time.Date(2026, 8, 9, 1, 2, 3, 0, time.UTC)
	basisPoints := int64(1_000)
	applications := []PromotionApplication{
		{PromotionID: "prm_addon", PromotionKind: "addon_purchase", PromotionRevision: 2, RelationID: "rel_addon", ResolvedQualifierSKUCodes: []string{"POTATO-A"}, ResolvedTargetSKUCodes: []string{"POTATO-B"}, ResolvedTerms: []PromotionTerm{{Key: "addon_price", MoneyValue: &money.Money{AmountMinor: 300, Currency: "AUD"}}}, AppliedAt: appliedAt},
		{PromotionID: "prm_bogo", PromotionKind: "bogo", PromotionRevision: 3, RelationID: "rel_bogo", ResolvedQualifierSKUCodes: []string{"POTATO-C"}, ResolvedTargetSKUCodes: []string{"POTATO-C"}, ResolvedTerms: []PromotionTerm{{Key: "discount_basis_points", BasisPointsValue: &basisPoints}}, AppliedAt: appliedAt},
		{PromotionID: "prm_bundle", PromotionKind: "bundle", PromotionRevision: 4, RelationID: "rel_bundle", ResolvedQualifierSKUCodes: []string{"POTATO-D", "POTATO-E"}, ResolvedTargetSKUCodes: []string{"POTATO-D", "POTATO-E"}, ResolvedAmounts: []PromotionAmount{{Key: "bundle_total", Amount: money.Money{AmountMinor: 1200, Currency: "AUD"}}}, DisplayMessages: []localization.LocalizedText{{Language: "en", Text: "Bundle applied"}}, ReceiptMessages: []localization.LocalizedText{{Language: "en", Text: "Potato bundle"}}, AppliedAt: appliedAt},
	}

	body, err := json.Marshal(applications)
	if err != nil {
		t.Fatalf("marshal applications: %v", err)
	}
	var got []PromotionApplication
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal applications: %v", err)
	}
	if len(got) != 3 || got[0].RelationID != "rel_addon" || got[0].ResolvedQualifierSKUCodes[0] != "POTATO-A" || got[0].ResolvedTargetSKUCodes[0] != "POTATO-B" {
		t.Fatalf("add-on qualifier-to-target relation changed: %+v", got)
	}
	if got[1].PromotionKind != "bogo" || got[1].ResolvedTerms[0].BasisPointsValue == nil || *got[1].ResolvedTerms[0].BasisPointsValue != 1_000 {
		t.Fatalf("BOGO relation changed: %+v", got[1])
	}
	if got[2].PromotionKind != "bundle" || len(got[2].ResolvedQualifierSKUCodes) != 2 || got[2].ReceiptMessages[0].Text != "Potato bundle" {
		t.Fatalf("bundle relation or approved receipt content changed: %+v", got[2])
	}
}
