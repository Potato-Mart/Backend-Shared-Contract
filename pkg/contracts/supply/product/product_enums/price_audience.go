package product_enums

// PriceAudience is the commercial audience a price applies to. It lets a
// line record whether retail or wholesale pricing was used, independently
// of the order channel or the buyer's identity.
type PriceAudience string

const (
	// PriceAudienceRetail is consumer/retail pricing.
	PriceAudienceRetail PriceAudience = "retail"
	// PriceAudienceWholesale is wholesale/trade pricing.
	PriceAudienceWholesale PriceAudience = "wholesale"
)

// IsValid reports whether a is a known PriceAudience.
func (a PriceAudience) IsValid() bool {
	switch a {
	case PriceAudienceRetail, PriceAudienceWholesale:
		return true
	}
	return false
}

func (a PriceAudience) String() string { return string(a) }
