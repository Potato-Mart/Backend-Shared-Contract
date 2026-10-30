package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// PromotionGroupedNetEvent freezes the actual grouped scorer operation, not a
// fabricated PromotionApplication or per-package allocation. Version 1 modes
// are bogo_floor_v1 and bundle_floor_largest_remainder_v1. The owner binds
// promotion/definition, candidate, manifest, original request and actual fixed
// target events to the immutable root decision. Consumers verify all required
// target and qualifier peers, original source facts, integer floor operations,
// group count, and the full allocation population. Bundle price and currency
// must match the authenticated authored relation. Different candidate variants
// or unavailable group/event attribution fail closed; no per-group split is
// claimed from an aggregate event. An inclusive scorer's converted reduction
// and an evaluator's conversion of the whole after-net amount are distinct
// actual events; neither is inferred from the other, particularly when the
// authored extraction and addition rates are not inverses. Selection and
// settled amounts must be reconciled or fail closed. Event ID is opaque and excludes the finalized
// quote revision; the finalized root decision binds the actual record.
type PromotionGroupedNetEvent struct {
	Version           int32                          `json:"version"`
	ID                string                         `json:"id"`
	ManifestID        string                         `json:"manifest_id"`
	CandidateID       string                         `json:"candidate_id"`
	PromotionID       string                         `json:"promotion_id"`
	PromotionRevision int64                          `json:"promotion_revision"`
	RelationID        string                         `json:"relation_id"`
	Mode              string                         `json:"mode"`
	GroupCount        int64                          `json:"group_count"`
	BundleUnitPrice   *money.Money                   `json:"bundle_unit_price,omitempty"`
	TotalReduction    money.Money                    `json:"total_reduction"`
	Peers             []PromotionGroupedNetEventPeer `json:"peers"`
}

// PromotionGroupedNetEventPeer records the actual original rounded line net,
// chosen units and floor result. FloorNumerator is the exact nonnegative
// product as a canonical decimal integer string, because two int64 inputs may
// exceed int64; FloorDenominator is original base units. For bundle allocation
// every peer in the actual allocator population is present, including zero
// shares, in original request order. Rank/award/tie are present only when the
// allocator actually ran. BOGO floor-only events do not claim such an event.
type PromotionGroupedNetEventPeer struct {
	RootLineID            string      `json:"root_line_id"`
	RequestOrder          int32       `json:"request_order"`
	OriginalLineNet       money.Money `json:"original_line_net"`
	OriginalBaseUnits     int64       `json:"original_base_units"`
	SelectedBaseUnits     int64       `json:"selected_base_units"`
	FloorNumerator        string      `json:"floor_numerator"`
	FloorDenominator      int64       `json:"floor_denominator"`
	FlooredSelectedNet    money.Money `json:"floored_selected_net"`
	RecordedNetReduction  money.Money `json:"recorded_net_reduction"`
	AllocationRank        *int64      `json:"allocation_rank,omitempty"`
	AllocationAwarded     *bool       `json:"allocation_awarded,omitempty"`
	AllocationTieBreakKey string      `json:"allocation_tie_break_key,omitempty"`
}
