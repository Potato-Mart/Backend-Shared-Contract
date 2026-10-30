package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// PricePromotionAllocation identifies the participating units and role in a
// frozen promotion application. Its range is one-based, lies within its source
// allocation, and is relative to the original requested line. GroupOrdinal is
// the existing promotion group identity where applicable. ReductionAmount is
// the allocated monetary reduction, which may be zero for a qualifier. This
// is evidence of Pricing's decision, not a new refund allocation policy.
type PricePromotionAllocation struct {
	PromotionID       string      `json:"promotion_id"`
	PromotionRevision int64       `json:"promotion_revision"`
	Role              string      `json:"role"`
	FirstUnit         int64       `json:"first_unit"`
	BaseUnits         int64       `json:"base_units"`
	GroupOrdinal      int64       `json:"group_ordinal,omitempty"`
	ReductionAmount   money.Money `json:"reduction_amount"`
	// ReductionRounding preserves the exact group reduction for v2 source
	// evidence. It is absent on historical v1 allocations.
	ReductionRounding *RoundingEvidence `json:"reduction_rounding,omitempty"`
}
