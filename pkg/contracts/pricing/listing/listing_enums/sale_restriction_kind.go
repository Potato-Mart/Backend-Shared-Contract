package listing_enums

// SaleRestrictionKind names a market-specific restriction that applies to
// selling one SKU. The contract defines the vocabulary; enforcement remains
// service behaviour.
type SaleRestrictionKind string

const (
	// SaleRestrictionKindAgeVerification requires proof of age at sale.
	SaleRestrictionKindAgeVerification SaleRestrictionKind = "age_verification"
	// SaleRestrictionKindQuantityLimit caps the quantity one buyer may
	// purchase.
	SaleRestrictionKindQuantityLimit SaleRestrictionKind = "quantity_limit"
	// SaleRestrictionKindChannelExcluded blocks one order channel.
	SaleRestrictionKindChannelExcluded SaleRestrictionKind = "channel_excluded"
	// SaleRestrictionKindDeliveryExcluded blocks delivery fulfilment.
	SaleRestrictionKindDeliveryExcluded SaleRestrictionKind = "delivery_excluded"
	// SaleRestrictionKindPrescription requires an authorised prescription.
	SaleRestrictionKindPrescription SaleRestrictionKind = "prescription"
)

// IsValid reports whether k is a known SaleRestrictionKind.
func (k SaleRestrictionKind) IsValid() bool {
	switch k {
	case SaleRestrictionKindAgeVerification, SaleRestrictionKindQuantityLimit,
		SaleRestrictionKindChannelExcluded, SaleRestrictionKindDeliveryExcluded,
		SaleRestrictionKindPrescription:
		return true
	}
	return false
}

func (k SaleRestrictionKind) String() string { return string(k) }
