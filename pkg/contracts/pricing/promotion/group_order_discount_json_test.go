package promotion

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography"

	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography/geography_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/pricing/promotion/promotion_enums"
)

// A fixed-amount group discount must cross the wire as money.Money minor units,
// never as a stringified major-unit value like DiscountSpec.DiscountValue.
func TestGroupOrderDiscountFixedAmountUsesMinorUnitMoney(t *testing.T) {
	body, err := json.Marshal(GroupOrderDiscountApplication{
		ID:                        "goda_1",
		GroupOrderCode:            "GO-2601010001",
		WholesaleOrganisationCode: "WO-1",
		State:                     promotion_enums.GroupOrderDiscountStateApproved,
		ApprovedPromotionID:       "prm_1",
		Proposal: &GroupOrderDiscountProposal{
			DiscountType: promotion_enums.DiscountTypeFixedAmount,
			Amount:       &money.Money{AmountMinor: 500, Currency: "AUD"},
		},
	})
	if err != nil {
		t.Fatalf("marshal application: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal application JSON: %v", err)
	}
	proposal, ok := got["proposal"].(map[string]any)
	if !ok {
		t.Fatalf("proposal = %#v, want object (%s)", got["proposal"], body)
	}
	amount, ok := proposal["amount"].(map[string]any)
	if !ok {
		t.Fatalf("proposal.amount = %#v, want money object (%s)", proposal["amount"], body)
	}
	if amount["amount_minor"].(float64) != 500 || amount["currency"] != "AUD" {
		t.Fatalf("proposal.amount = %#v, want 500 AUD minor units (%s)", amount, body)
	}
	if got["state"] != "approved" {
		t.Fatalf("state = %v, want approved (%s)", got["state"], body)
	}
}

func TestGroupOrderDiscountDecisionFreezesGeographicContext(t *testing.T) {
	payload, err := json.Marshal(GroupOrderDiscountDecision{
		QuoteKey: "quote_1", GroupOrderCode: "GO-2601010001", Applied: true,
		ApprovedPromotionID: "prm_1",
		DiscountAmount:      money.Money{AmountMinor: 500, Currency: "AUD"},
		GeographicContext: geography.GeographicContext{
			Source:      geography_enums.GeographicContextSourceWholesaleOrganisationProfile,
			CountryCode: "AU", SubdivisionCode: "AU-VIC",
			MatchedTargetKind: geography_enums.GeographicTargetSubdivision,
			MatchedTargetCode: "AU-VIC", ScopeRevision: 3, RuleRevision: 7,
			EvaluationTimezone: "Australia/Melbourne",
		},
	})
	if err != nil {
		t.Fatalf("marshal group-order discount decision: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal group-order discount decision: %v", err)
	}
	context, ok := got["geographic_context"].(map[string]any)
	if !ok || context["source"] != "WHOLESALE_ORGANISATION_PROFILE" || context["matched_target_code"] != "AU-VIC" {
		t.Fatalf("group-order discount geographic context mismatch: %s", payload)
	}
}

// A percentage group discount carries integer basis points, not a float/string.
func TestGroupOrderDiscountPercentUsesBasisPoints(t *testing.T) {
	body, err := json.Marshal(GroupOrderDiscountProposal{
		DiscountType:       promotion_enums.DiscountTypePercentage,
		PercentBasisPoints: 1000,
	})
	if err != nil {
		t.Fatalf("marshal proposal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal proposal JSON: %v", err)
	}
	if got["percent_basis_points"].(float64) != 1000 {
		t.Fatalf("percent_basis_points = %#v, want 1000 (%s)", got["percent_basis_points"], body)
	}
	if _, present := got["amount"]; present {
		t.Fatalf("amount should be omitted for a percentage proposal (%s)", body)
	}
}
