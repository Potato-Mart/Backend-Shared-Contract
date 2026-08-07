package membership

// TierProgressReason explains why customer tier progress is unavailable.
type TierProgressReason string

const (
	TierProgressReasonNoActiveTiers         TierProgressReason = "no_active_tiers"
	TierProgressReasonManualQualification   TierProgressReason = "manual_qualification"
	TierProgressReasonUnsupportedMetric     TierProgressReason = "unsupported_metric"
	TierProgressReasonCurrencyMismatch      TierProgressReason = "currency_mismatch"
	TierProgressReasonMembershipNotAssigned TierProgressReason = "membership_not_assigned"
)

func (r TierProgressReason) IsValid() bool {
	switch r {
	case TierProgressReasonNoActiveTiers, TierProgressReasonManualQualification,
		TierProgressReasonUnsupportedMetric, TierProgressReasonCurrencyMismatch,
		TierProgressReasonMembershipNotAssigned:
		return true
	default:
		return false
	}
}

func (r TierProgressReason) String() string { return string(r) }

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
	case TierBenefitKindQualifyingSpend, TierBenefitKindPointsMultiplier,
		TierBenefitKindDiscountPercent, TierBenefitKindFreeShippingThreshold,
		TierBenefitKindBirthdayBonusPoints:
		return true
	default:
		return false
	}
}

func (k TierBenefitKind) String() string { return string(k) }
