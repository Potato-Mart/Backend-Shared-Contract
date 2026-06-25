package enums

// BuyerType is the shared, channel-independent classification of who is
// buying. It is deliberately distinct from the order channel
// (enums.OrderType): POS is a channel, never a buyer type. A wholesale
// organisation buying in the physical shop is BuyerTypeWholesaleOrganisation
// on OrderTypePOS, while a walk-in customer is BuyerTypeGuestRetail on the
// same channel.
type BuyerType string

const (
	// BuyerTypeGuestRetail is an unauthenticated retail buyer, e.g. a
	// walk-in POS customer or an anonymous web shopper.
	BuyerTypeGuestRetail BuyerType = "guest_retail"
	// BuyerTypeRetailCustomer is an identified retail customer/member.
	BuyerTypeRetailCustomer BuyerType = "retail_customer"
	// BuyerTypeWholesaleOrganisation is a wholesale organisation buyer,
	// regardless of the channel it transacts through.
	BuyerTypeWholesaleOrganisation BuyerType = "wholesale_organisation"
)

// IsValid reports whether b is a known BuyerType.
func (b BuyerType) IsValid() bool {
	switch b {
	case BuyerTypeGuestRetail, BuyerTypeRetailCustomer, BuyerTypeWholesaleOrganisation:
		return true
	}
	return false
}

func (b BuyerType) String() string { return string(b) }
