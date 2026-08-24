package pricebook_enums

type PriceEndingPolicy string

const (
	PriceEndingPolicyNone      PriceEndingPolicy = "none"
	PriceEndingPolicyCharmNine PriceEndingPolicy = "charm_nine"
)

func (p PriceEndingPolicy) IsValid() bool {
	return p == PriceEndingPolicyNone || p == PriceEndingPolicyCharmNine
}
func (p PriceEndingPolicy) String() string { return string(p) }
