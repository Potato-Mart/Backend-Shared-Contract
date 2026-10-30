package quote

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
