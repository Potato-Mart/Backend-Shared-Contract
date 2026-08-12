package listing_enums

// MarketListingStatus records whether a SKU may be sold in one market.
// Commercial availability is a listing fact and is independent of the physical
// inventory a depot happens to hold.
type MarketListingStatus string

const (
	// MarketListingStatusDraft is a listing being prepared. Price drafts may
	// be stored against it, but no price entry may be approved or quoted.
	MarketListingStatusDraft MarketListingStatus = "draft"
	// MarketListingStatusComingSoon is a listing published for display
	// before its availability window opens.
	MarketListingStatusComingSoon MarketListingStatus = "coming_soon"
	// MarketListingStatusActive is a listing that may be quoted and sold.
	MarketListingStatusActive MarketListingStatus = "active"
	// MarketListingStatusSuspended is a listing temporarily withheld from
	// sale in this market.
	MarketListingStatusSuspended MarketListingStatus = "suspended"
	// MarketListingStatusUnavailable is a listing that is not sold in this
	// market.
	MarketListingStatusUnavailable MarketListingStatus = "unavailable"
	// MarketListingStatusDelisted is a listing permanently withdrawn from
	// this market and retained for history.
	MarketListingStatusDelisted MarketListingStatus = "delisted"
)

// IsValid reports whether s is a known MarketListingStatus.
func (s MarketListingStatus) IsValid() bool {
	switch s {
	case MarketListingStatusDraft, MarketListingStatusComingSoon, MarketListingStatusActive,
		MarketListingStatusSuspended, MarketListingStatusUnavailable, MarketListingStatusDelisted:
		return true
	}
	return false
}

func (s MarketListingStatus) String() string { return string(s) }

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
