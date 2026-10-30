package quote

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
