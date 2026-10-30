package quote

// PromotionExactMinor is an exact rational amount in minor units of the
// enclosing run's currency. Denominator is positive and the fraction is reduced.
// It is not itself a rounded payable amount.
type PromotionExactMinor struct {
	Numerator   int64 `json:"numerator"`
	Denominator int64 `json:"denominator"`
}
