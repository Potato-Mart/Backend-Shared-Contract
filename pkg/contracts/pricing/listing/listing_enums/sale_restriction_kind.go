package listing_enums

// SaleRestrictionKind names a market-specific restriction that applies to
// selling one SKU. The contract defines the vocabulary; enforcement remains
// service behaviour.
type SaleRestrictionKind string

const (
	// SaleRestrictionKindAge restricts retail purchase using saved Customers
	// DOB and a strict below/above threshold. It never restricts viewing or
	// wholesale purchase and does not require proof-based age verification.
	SaleRestrictionKindAge SaleRestrictionKind = "age"
	// SaleRestrictionKindQuantityLimit caps the quantity one buyer may
	// purchase.
	SaleRestrictionKindQuantityLimit SaleRestrictionKind = "quantity_limit"
	// SaleRestrictionKindChannelExcluded blocks one order channel.
	SaleRestrictionKindChannelExcluded SaleRestrictionKind = "channel_excluded"
	// SaleRestrictionKindDeliveryExcluded blocks the explicitly listed
	// canonical delivery methods: delivery, pickup or outsourced.
	SaleRestrictionKindDeliveryExcluded SaleRestrictionKind = "delivery_excluded"
	// SaleRestrictionKindPrescription requires an authorised prescription.
	SaleRestrictionKindPrescription SaleRestrictionKind = "prescription"
)

// IsValid reports whether k is a known SaleRestrictionKind.
func (k SaleRestrictionKind) IsValid() bool {
	switch k {
	case SaleRestrictionKindAge, SaleRestrictionKindQuantityLimit,
		SaleRestrictionKindChannelExcluded, SaleRestrictionKindDeliveryExcluded,
		SaleRestrictionKindPrescription:
		return true
	}
	return false
}

func (k SaleRestrictionKind) String() string { return string(k) }
