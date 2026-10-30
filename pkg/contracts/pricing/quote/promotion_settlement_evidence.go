package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// PromotionSettlementPool freezes the complete original-line peer population
// for one actual transition. Purpose separates settled-price and contribution
// pools; they are never interchangeable. Peers include non-target, zero and
// locked units as well as targets. The total and eligibility are checked
// against authenticated original source and all preceding stages; consumers
// verify the supported rounding mode, floors, awards and tie ordering without
// expanding compressed runs. Pool and predecessor identities are bound to the
// same immutable root decision as the candidate manifest. Retained amendments
// authenticate the original pool and clip inherited awards without rerounding.
// Incomplete peer population, prior lock lineage or source facts fail closed.
type PromotionSettlementPool struct {
	ID                string                            `json:"id"`
	Purpose           string                            `json:"purpose"`
	RootLineID        string                            `json:"root_line_id"`
	StageOrdinal      int32                             `json:"stage_ordinal"`
	SubstageOrdinal   int32                             `json:"substage_ordinal"`
	PreviousPoolID    string                            `json:"previous_pool_id,omitempty"`
	OriginalBaseUnits int64                             `json:"original_base_units"`
	BookTaxInclusion  pricebook_enums.PriceTaxInclusion `json:"book_tax_inclusion"`
	Total             RoundingEvidence                  `json:"total"`
	TieBreakRule      string                            `json:"tie_break_rule"`
	Peers             []PromotionSettlementPoolPeer     `json:"peers"`
}

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

// PromotionSourceBeforeEvidence binds a first monetary stage to the actual
// approved source and the absence of an earlier stage on the same original
// units. BookUnit and NetUnit are separately authenticated exact minor-unit
// rationals. The actual scorer-derived net candidate amount and independently
// settled book-after price must each be evidenced; their ratio is not claimed
// as an execution step. No equality between book-basis line discount and net
// application discount is implied. Missing net-candidate, book-settlement,
// source, rounding or prior-stage evidence fails closed even if totals happen
// to reconcile.
type PromotionSourceBeforeEvidence struct {
	Version                    int32               `json:"version"`
	SourceEvidenceID           string              `json:"source_evidence_id"`
	EarlierStageAuditID        string              `json:"earlier_stage_audit_id"`
	BookUnit                   PromotionExactMinor `json:"book_unit"`
	NetUnit                    PromotionExactMinor `json:"net_unit"`
	NetCandidateEvidenceID     string              `json:"net_candidate_evidence_id"`
	BookSettlementPoolID       string              `json:"book_settlement_pool_id"`
	ApplicationRoundingProofID string              `json:"application_rounding_proof_id"`
}

// PromotionApplicationRoundingProof identifies the complete actual candidate
// fixed-evaluator event and its per-line book/net reductions. It includes every peer
// needed by AllocationRule, including zero contributions. For the bounded
// candidate_net_book_after_half_up_v1 mode, the original net candidate effect
// and recorded net reduction are authenticated independently from the
// exact book-after consideration and its actual half-up settlement. A
// rounded-book-discount conversion is never presented as an executed net
// rounding step. Optional rounding evidence is used only where the scorer
// actually performed that rounding.
// A consumer must recognize the declared allocation rule and prove its full
// peer set and rounding; an unimplemented rule is unavailable, not estimated.
// ID is computed before the containing quote revision and binds every field
// except ID; the root Pricing decision and owner lookup bind this proof to the
// finalized quote and manifest without a fingerprint cycle.
type PromotionApplicationRoundingProof struct {
	ID                string                          `json:"id"`
	ManifestID        string                          `json:"manifest_id"`
	PromotionID       string                          `json:"promotion_id"`
	PromotionRevision int64                           `json:"promotion_revision"`
	CandidateID       string                          `json:"candidate_id"`
	StageOrdinal      int32                           `json:"stage_ordinal"`
	AllocationRule    string                          `json:"allocation_rule"`
	FirstGroup        int64                           `json:"first_group"`
	GroupCount        int64                           `json:"group_count"`
	Peers             []PromotionApplicationPeerProof `json:"peers"`
}

// PromotionApplicationPeerProof uses aggregate amounts for this root ordinal
// interval, not a per-unit amount. BookAfterRounding is the actual scorer
// event and may reference a larger complete population via its pool evidence.
type PromotionApplicationPeerProof struct {
	RootLineID              string              `json:"root_line_id"`
	RequirementIndex        int32               `json:"requirement_index"`
	FirstUnit               int64               `json:"first_unit"`
	BaseUnits               int64               `json:"base_units"`
	ExactBookReduction      PromotionExactMinor `json:"exact_book_reduction"`
	ExactBookAfter          PromotionExactMinor `json:"exact_book_after"`
	BookAfterRounding       RoundingEvidence    `json:"book_after_rounding"`
	ExactNetCandidateEffect PromotionExactMinor `json:"exact_net_candidate_effect"`
	RecordedNetReduction    money.Money         `json:"recorded_net_reduction"`
	NetRounding             *RoundingEvidence   `json:"net_rounding,omitempty"`
}

// PromotionStageAudit freezes the complete earlier stage population for a
// root ordinal interval. It includes all earlier monetary applications and
// exclusive/custom paths, even when they did not contribute a reduction.
// Empty Entries proves absence only when the Pricing owner authenticates the
// complete root-line decision and the audit's source/evidence binding. An
// omitted or partial audit never proves absence. IDs exclude the containing
// quote revision/fingerprint and are bound by the finalized root decision.
type PromotionStageAudit struct {
	ID               string                     `json:"id"`
	SourceEvidenceID string                     `json:"source_evidence_id"`
	RootLineID       string                     `json:"root_line_id"`
	FirstUnit        int64                      `json:"first_unit"`
	BaseUnits        int64                      `json:"base_units"`
	BeforeStage      int32                      `json:"before_stage"`
	BeforeSubstage   int32                      `json:"before_substage"`
	Entries          []PromotionStageAuditEntry `json:"entries"`
}

type PromotionStageAuditEntry struct {
	StageOrdinal      int32  `json:"stage_ordinal"`
	SubstageOrdinal   int32  `json:"substage_ordinal"`
	Kind              string `json:"kind"`
	Role              string `json:"role"`
	PromotionID       string `json:"promotion_id,omitempty"`
	PromotionRevision int64  `json:"promotion_revision,omitempty"`
	RuleID            string `json:"rule_id"`
	FirstUnit         int64  `json:"first_unit"`
	BaseUnits         int64  `json:"base_units"`
}
