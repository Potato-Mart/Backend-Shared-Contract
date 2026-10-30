package quote

// PromotionCandidateManifest freezes the complete original candidate role
// vector for a single actual stage. It is carried in separately negotiated
// quote-level evidence and bound to the immutable Pricing decision; it is not
// inferred from the runs a consumer happened to receive. Version 1 requires
// contiguous one-based requirement indices, complete group coverage for each
// declared peer and bounded counts. Each stage run names ManifestID and its
// requirement index. A candidate with different peers per repetition must be
// represented by distinct manifests, never silently regrouped.
// MinimumRepetitions zero means no extra minimum beyond one selected group;
// MaximumRepetitions is a strict positive upper bound. Requirement UnitsPerGroup
// is the sum of that requirement's peer multiplicities, not a per-line guess.
// Version 1 permits competition_class "product_competing" for grouped BOGO
// and bundle candidates only. Other mechanics or unknown classes require a
// separately specified proof and fail closed.
//
// ManifestID is an opaque Pricing-owned identifier, authenticated against the
// complete immutable root manifest through the owner. Consumers must not
// recompute it from this wire struct or assume it is a cross-service canonical
// digest. Its owner preimage excludes the finalized quote revision/fingerprint:
// the quote decision binds the manifest, so including either creates a cycle.
// EligibilityEvidenceID identifies a separately frozen Pricing-owned decision
// bound to the root quote, candidate and original normalized request, including
// buyer, channel, platform, geography, line/package quantities, POS selection,
// evaluation time and the immutable definition revisions. The owner must
// authenticate the approved original source, tax and exact unit net for every
// original cart component used by the decision, including nonparticipating
// lines used by a cart-wide threshold. An excluded exclusive-path line must
// instead have its exclusion independently bound to the original request,
// actual pre-promotion exclusive rule and immutable root snapshot; it cannot
// participate in the candidate. Missing input or exclusion proof fails closed.
// The owner must authenticate this record and the decision; an unrelated quote digest or
// request fingerprint alone is not that proof. UsageEvidenceID identifies
// separately frozen authoritative usage evidence. Both IDs are computed before
// and independently of the containing quote revision/fingerprint; the final
// quote binds them. Atomic reservation remains a
// separate final race guard. A manifest alone does not prove audience,
// channel, geography, time, promotion availability or a usage reservation.
// Missing or unverifiable facts fail closed. This evidence does not alter the
// existing line-level proportional refund policy.
type PromotionCandidateManifest struct {
	Version               int32                           `json:"version"`
	ManifestID            string                          `json:"manifest_id"`
	PromotionID           string                          `json:"promotion_id"`
	PromotionKind         string                          `json:"promotion_kind"`
	CompetitionClass      string                          `json:"competition_class"`
	PromotionRevision     int64                           `json:"promotion_revision"`
	ScopeRevision         int64                           `json:"scope_revision"`
	RuleRevision          int64                           `json:"rule_revision"`
	PublicationRevision   int64                           `json:"publication_revision"`
	CandidateID           string                          `json:"candidate_id"`
	StageOrdinal          int32                           `json:"stage_ordinal"`
	FirstGroup            int64                           `json:"first_group"`
	GroupCount            int64                           `json:"group_count"`
	MinimumRepetitions    int64                           `json:"minimum_repetitions"`
	MaximumRepetitions    int64                           `json:"maximum_repetitions"`
	RequirementCount      int32                           `json:"requirement_count"`
	Requirements          []PromotionCandidateRequirement `json:"requirements"`
	EligibilityEvidenceID string                          `json:"eligibility_evidence_id"`
	UsageEvidenceID       string                          `json:"usage_evidence_id"`
}
