package coupon_enums

// TierMatch describes matching against a service-owned tier order.
type TierMatch string

const (
	TierMatchExact     TierMatch = "EXACT"
	TierMatchAtOrAbove TierMatch = "AT_OR_ABOVE"
)

func (v TierMatch) IsValid() bool  { return v == TierMatchExact || v == TierMatchAtOrAbove }
func (v TierMatch) String() string { return string(v) }
