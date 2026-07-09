package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/promotion"
)

// Group-order discount endpoints. Provider: Backend-Management. Added v13.1.0.
//
// A wholesale group-order manager applies for (or selects) a per-group discount
// from the wholesale storefront; a staff approver issues it by publishing the
// backing promotion. Backend-Commerce reads the approved promotion at
// pricing/submit time via the internal endpoint. Money is always common.Money
// (minor units); the percentage case is carried as integer basis points so no
// stringified major-unit value ever crosses the wire (unlike DiscountSpec).
const (
	// PathGroupOrderDiscountInternal is the S2S read Commerce uses to resolve an
	// application's approved promotion for a group order (scope promotion:grant).
	PathGroupOrderDiscountInternal = "/v1/internal/promotions/group-order-discount"
)

// GroupOrderDiscountProposal is a manager's requested discount when applying for
// a brand-new benefit (rather than selecting an existing promotion).
type GroupOrderDiscountProposal struct {
	DiscountType promotionenum.DiscountType `json:"discount_type"`
	// PercentBasisPoints is set when DiscountType is percentage: 1000 = 10.00%.
	PercentBasisPoints int `json:"percent_basis_points,omitempty"`
	// Amount is set when DiscountType is fixed_amount; minor units.
	Amount *common.Money `json:"amount,omitempty"`
	// MaxDiscount optionally caps a percentage discount; minor units.
	MaxDiscount *common.Money `json:"max_discount,omitempty"`
}

// GroupOrderDiscountApplication is the shared application record produced by
// Backend-Management and read by Backend-Commerce (over the internal endpoint)
// and the admin console: its lifecycle state and, once approved, the backing
// promotion id Commerce must apply. Exactly one of SelectedPromotionID (an
// existing eligible promotion) or Proposal (a newly requested benefit) is set.
// The wire request/decision envelopes are owned by Backend-Management.
type GroupOrderDiscountApplication struct {
	ID                        string                                `json:"id"`
	GroupOrderCode            string                                `json:"group_order_code"`
	WholesaleOrganisationCode string                                `json:"wholesale_organisation_code"`
	OrganisationAccessID      string                                `json:"organisation_access_id,omitempty"`
	State                     promotionenum.GroupOrderDiscountState `json:"state"`
	SelectedPromotionID       string                                `json:"selected_promotion_id,omitempty"`
	Proposal                  *GroupOrderDiscountProposal           `json:"proposal,omitempty"`
	// ApprovedPromotionID is the published promotion Commerce applies at
	// pricing/submit; set only when State is approved.
	ApprovedPromotionID string     `json:"approved_promotion_id,omitempty"`
	RequestedBy         string     `json:"requested_by,omitempty"`
	DecidedBy           string     `json:"decided_by,omitempty"`
	DecisionReason      string     `json:"decision_reason,omitempty"`
	RequestedAt         *time.Time `json:"requested_at,omitempty"`
	DecidedAt           *time.Time `json:"decided_at,omitempty"`
}
