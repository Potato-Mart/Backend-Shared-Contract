package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// PromotionSettlementPoolPeer is a homogeneous original ordinal run. When
// eligible, FirstRank is the first unit's zero-based rank and later units have
// consecutive ranks and the negotiated ordinal tie rule. Split runs whenever
// exact values, awards, eligibility, prior locks or rank progression differ.
// Ineligible units have no rank/tie and must be justified by a zero exact
// value or a separately authenticated prior stage/substage lock.
type PromotionSettlementPoolPeer struct {
	FirstUnit           int64               `json:"first_unit"`
	BaseUnits           int64               `json:"base_units"`
	ExactUnit           PromotionExactMinor `json:"exact_unit"`
	SettledUnit         money.Money         `json:"settled_unit"`
	AwardEligible       bool                `json:"award_eligible"`
	Awarded             bool                `json:"awarded"`
	FirstRank           *int64              `json:"first_rank,omitempty"`
	FirstTieBreakKey    string              `json:"first_tie_break_key,omitempty"`
	LockStageOrdinal    *int32              `json:"lock_stage_ordinal,omitempty"`
	LockSubstageOrdinal *int32              `json:"lock_substage_ordinal,omitempty"`
	LockProofID         string              `json:"lock_proof_id,omitempty"`
}
