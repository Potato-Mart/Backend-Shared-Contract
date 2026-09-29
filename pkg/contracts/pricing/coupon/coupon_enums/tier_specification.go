package coupon_enums

// TierSpecification selects all membership tiers or one tier reference.
type TierSpecification string

const (
	TierSpecificationAllTiers     TierSpecification = "ALL_TIERS"
	TierSpecificationSelectedTier TierSpecification = "SELECTED_TIER"
)

func (v TierSpecification) IsValid() bool {
	return v == TierSpecificationAllTiers || v == TierSpecificationSelectedTier
}
func (v TierSpecification) String() string { return string(v) }
