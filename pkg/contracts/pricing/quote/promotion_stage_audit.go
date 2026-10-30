package quote

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
