package productenum

// StorefrontPreorderStatus is the public preorder display state for storefronts.
type StorefrontPreorderStatus string

const (
	StorefrontPreorderStatusUnavailable StorefrontPreorderStatus = "unavailable"
	StorefrontPreorderStatusUpcoming    StorefrontPreorderStatus = "upcoming"
	StorefrontPreorderStatusOpen        StorefrontPreorderStatus = "open"
	StorefrontPreorderStatusClosed      StorefrontPreorderStatus = "closed"
	StorefrontPreorderStatusSoldOut     StorefrontPreorderStatus = "sold_out"
)

func (s StorefrontPreorderStatus) IsValid() bool {
	switch s {
	case StorefrontPreorderStatusUnavailable, StorefrontPreorderStatusUpcoming,
		StorefrontPreorderStatusOpen, StorefrontPreorderStatusClosed,
		StorefrontPreorderStatusSoldOut:
		return true
	}
	return false
}

func (s StorefrontPreorderStatus) String() string { return string(s) }

// StorefrontExpiryStatus is the public expiry merchandising state for
// storefronts. It avoids exposing warehouse lot or service implementation
// details.
type StorefrontExpiryStatus string

const (
	StorefrontExpiryStatusNotApplicable StorefrontExpiryStatus = "not_applicable"
	StorefrontExpiryStatusSoonExpiry    StorefrontExpiryStatus = "soon_expiry"
	StorefrontExpiryStatusExpired       StorefrontExpiryStatus = "expired"
)

func (s StorefrontExpiryStatus) IsValid() bool {
	switch s {
	case StorefrontExpiryStatusNotApplicable, StorefrontExpiryStatusSoonExpiry,
		StorefrontExpiryStatusExpired:
		return true
	}
	return false
}

func (s StorefrontExpiryStatus) String() string { return string(s) }
