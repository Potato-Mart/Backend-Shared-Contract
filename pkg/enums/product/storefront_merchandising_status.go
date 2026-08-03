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
