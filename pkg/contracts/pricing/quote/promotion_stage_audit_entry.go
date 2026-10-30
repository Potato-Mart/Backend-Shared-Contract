package quote

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
