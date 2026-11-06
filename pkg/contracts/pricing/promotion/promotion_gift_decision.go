package promotion

// PromotionGiftDecision tells a till how to settle a gift-type relation.
// RequiresStaffDecision is true when staff must accept or decline the gift
// before checkout; MaxGiftBaseUnits is the total gift base units the relation
// allows across its GiftOptions.
type PromotionGiftDecision struct {
	RequiresStaffDecision bool  `json:"requires_staff_decision"`
	MaxGiftBaseUnits      int64 `json:"max_gift_base_units"`
}
