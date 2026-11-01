package promotion_enums

// PromotionProgressStatus states how far a cart is from one promotion.
type PromotionProgressStatus string

const (
	// PromotionProgressStatusApplied means pricing used the promotion.
	PromotionProgressStatusApplied PromotionProgressStatus = "applied"
	// PromotionProgressStatusNear means one quantity or spend shortfall remains.
	PromotionProgressStatusNear PromotionProgressStatus = "near"
	// PromotionProgressStatusOpen means the cart is eligible but has no progress.
	PromotionProgressStatusOpen PromotionProgressStatus = "open"
)

// IsValid reports whether s is a supported promotion progress status.
func (s PromotionProgressStatus) IsValid() bool {
	switch s {
	case PromotionProgressStatusApplied, PromotionProgressStatusNear, PromotionProgressStatusOpen:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s PromotionProgressStatus) String() string { return string(s) }
