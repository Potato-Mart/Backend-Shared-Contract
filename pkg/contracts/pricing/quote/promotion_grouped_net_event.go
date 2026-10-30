package quote

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
