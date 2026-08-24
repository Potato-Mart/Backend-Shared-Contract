package membership_enums

// TierBenefitKind classifies a typed membership tier benefit.
type TierBenefitKind string

const (
	TierBenefitKindQualifyingSpend       TierBenefitKind = "qualifying_spend"
	TierBenefitKindPointsMultiplier      TierBenefitKind = "points_multiplier"
	TierBenefitKindDiscountPercent       TierBenefitKind = "discount_percent"
	TierBenefitKindFreeShippingThreshold TierBenefitKind = "free_shipping_threshold"
	TierBenefitKindBirthdayBonusPoints   TierBenefitKind = "birthday_bonus_points"
)

func (k TierBenefitKind) IsValid() bool {
	switch k {
	case TierBenefitKindQualifyingSpend, TierBenefitKindPointsMultiplier, TierBenefitKindDiscountPercent, TierBenefitKindFreeShippingThreshold, TierBenefitKindBirthdayBonusPoints:
		return true
	}
	return false
}
func (k TierBenefitKind) String() string { return string(k) }
