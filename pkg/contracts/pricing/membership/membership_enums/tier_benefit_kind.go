package membership_enums

// TierBenefitKind classifies a typed membership tier benefit.
type TierBenefitKind string

const (
	TierBenefitKindQualifyingSpend  TierBenefitKind = "qualifying_spend"
	TierBenefitKindPointsMultiplier TierBenefitKind = "points_multiplier"
)

func (k TierBenefitKind) IsValid() bool {
	switch k {
	case TierBenefitKindQualifyingSpend, TierBenefitKindPointsMultiplier:
		return true
	}
	return false
}
func (k TierBenefitKind) String() string { return string(k) }
