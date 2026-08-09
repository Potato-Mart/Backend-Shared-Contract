package promotion

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
)

// A wholesale group-order manager applies for (or selects) a per-group discount
// from the wholesale storefront; a staff approver issues it by publishing the
// backing promotion. Money is always money.Money
// (minor units); the percentage case is carried as integer basis points so no
// stringified major-unit value ever crosses the wire (unlike DiscountSpec).
// GroupOrderDiscountProposal is a manager's requested discount when applying for
// a brand-new benefit (rather than selecting an existing promotion).
type GroupOrderDiscountProposal struct {
	DiscountType promotion_enums.DiscountType `json:"discount_type"`
	// PercentBasisPoints is set when DiscountType is percentage: 1000 = 10.00%.
	PercentBasisPoints int `json:"percent_basis_points,omitempty"`
	// Amount is set when DiscountType is fixed_amount; minor units.
	Amount *money.Money `json:"amount,omitempty"`
	// MaxDiscount optionally caps a percentage discount; minor units.
	MaxDiscount *money.Money `json:"max_discount,omitempty"`
}

// GroupOrderDiscountApplication is the shared application record owned by
// Pricing: its lifecycle state and, once approved, the backing promotion ID
// Orders must apply. Exactly one of SelectedPromotionID (an
// existing eligible promotion) or Proposal (a newly requested benefit) is set.
// The wire request/decision envelopes are owned by Pricing.
type GroupOrderDiscountApplication struct {
	ID                        string                                  `json:"id"`
	GroupOrderCode            string                                  `json:"group_order_code"`
	WholesaleOrganisationCode string                                  `json:"wholesale_organisation_code"`
	OrganisationAccessID      string                                  `json:"organisation_access_id,omitempty"`
	State                     promotion_enums.GroupOrderDiscountState `json:"state"`
	SelectedPromotionID       string                                  `json:"selected_promotion_id,omitempty"`
	Proposal                  *GroupOrderDiscountProposal             `json:"proposal,omitempty"`
	// ApprovedPromotionID is the published promotion Orders applies at
	// pricing/submit; set only when State is approved.
	ApprovedPromotionID string     `json:"approved_promotion_id,omitempty"`
	RequestedBy         string     `json:"requested_by,omitempty"`
	DecidedBy           string     `json:"decided_by,omitempty"`
	DecisionReason      string     `json:"decision_reason,omitempty"`
	RequestedAt         *time.Time `json:"requested_at,omitempty"`
	DecidedAt           *time.Time `json:"decided_at,omitempty"`
}

type GroupOrderDiscountDecision struct {
	QuoteKey            string                           `json:"quote_key"`
	GroupOrderCode      string                           `json:"group_order_code"`
	Applied             bool                             `json:"applied"`
	ApplicationID       string                           `json:"application_id,omitempty"`
	ApprovedPromotionID string                           `json:"approved_promotion_id,omitempty"`
	DiscountAmount      money.Money                      `json:"discount_amount"`
	GeographicContext   geography.GeographicContext      `json:"geographic_context"`
	Lines               []GroupOrderDiscountDecisionLine `json:"lines,omitempty"`
	RuleVersion         string                           `json:"rule_version,omitempty"`
	EvaluatedAt         time.Time                        `json:"evaluated_at"`
	Replayed            bool                             `json:"replayed,omitempty"`
}

type GroupOrderDiscountDecisionLine struct {
	OrderItemID    string      `json:"order_item_id"`
	ProductSKUCode string      `json:"product_sku_code"`
	DiscountAmount money.Money `json:"discount_amount"`
}
