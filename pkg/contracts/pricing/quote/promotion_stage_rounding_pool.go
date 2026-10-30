package quote

// PromotionStageRoundingPool binds the common exact total, rounded total,
// rounding mode and exponent of the complete original-line peer population.
// Every member repeats the same pool proof, ID and policy; consumers must
// authenticate all peers and recompute their ranks, awards and stable tie keys.
// Version 1 permits only TieBreakRule "line_unit_ordinal_v1": a pool covers
// one original line and ranks exact fractional remainders descending, then
// eligible original unit ordinals ascending. Ineligible peers still contribute
// their floors to the pool total, but receive no award, rank or tie key.
// The tie key for eligible peers is the root unit ordinal
// rendered as a zero-padded 20-digit decimal string. RootLineID binds the line. A
// different peer order cannot be labelled with this rule and must be rejected
// until another negotiated version defines it. Pools of settled book prices
// and of promotion contributions are distinct and cannot be substituted.
type PromotionStageRoundingPool struct {
	ID           string           `json:"id"`
	Total        RoundingEvidence `json:"total"`
	TieBreakRule string           `json:"tie_break_rule"`
	RootLineID   string           `json:"root_line_id"`
}
