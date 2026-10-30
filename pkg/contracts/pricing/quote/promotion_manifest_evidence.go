package quote

// PromotionManifestEvidence is quote-level, not a per-line allocation. An
// adopting service must include this complete object in its authoritative
// immutable decision and expose it only after independent capability
// negotiation. Omitted evidence retains historical wire and hash semantics.
type PromotionManifestEvidence struct {
	Version                   int32                               `json:"version"`
	Manifests                 []PromotionCandidateManifest        `json:"manifests"`
	SettlementPools           []PromotionSettlementPool           `json:"settlement_pools"`
	ApplicationRoundingProofs []PromotionApplicationRoundingProof `json:"application_rounding_proofs"`
	GroupedNetEvents          []PromotionGroupedNetEvent          `json:"grouped_net_events,omitempty"`
	StageAudits               []PromotionStageAudit               `json:"stage_audits"`
}
