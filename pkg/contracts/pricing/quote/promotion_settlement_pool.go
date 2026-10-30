package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"

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
