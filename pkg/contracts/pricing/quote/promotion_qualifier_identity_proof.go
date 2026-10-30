package quote

// PromotionQualifierIdentityProof records an actual no-change qualifier
// transition where Pricing performed no settlement on the qualifier line.
// It must not be used as a fabricated rounding pool. The owner authenticates
// the full original source, complete earlier-stage observation, actual qualifier
// event, candidate requirement and root ordinal/group spans against the root
// quote. The consumer verifies whole-package boundaries, full peer coverage,
// disjoint competing units, equal original/before/after exact unit values and
// absence of prior monetary stages or locks. Retained slices additionally
// require authenticated predecessor lineage. Usage and eligibility evidence
// remain independent. In the supported `no_earlier_events_v1` mode the
// Pricing owner authenticates its captured observation that no other selected
// candidate or prior/same-stage mechanic event used the line, and the root
// snapshot has no applied rules. This is not a separate audit record. Unknown
// modes fail closed. Identity is opaque and computed before quote revision.
type PromotionQualifierIdentityProof struct {
	ID                string              `json:"id"`
	SourceEvidenceID  string              `json:"source_evidence_id"`
	PriorStageMode    string              `json:"prior_stage_mode"`
	ManifestID        string              `json:"manifest_id"`
	PromotionID       string              `json:"promotion_id"`
	PromotionRevision int64               `json:"promotion_revision"`
	StageOrdinal      int32               `json:"stage_ordinal"`
	SubstageOrdinal   int32               `json:"substage_ordinal"`
	RequirementIndex  int32               `json:"requirement_index"`
	FirstGroup        int64               `json:"first_group"`
	GroupCount        int64               `json:"group_count"`
	RootLineID        string              `json:"root_line_id"`
	FirstUnit         int64               `json:"first_unit"`
	BaseUnits         int64               `json:"base_units"`
	BeforeUnit        PromotionExactMinor `json:"before_unit"`
	AfterUnit         PromotionExactMinor `json:"after_unit"`
}
