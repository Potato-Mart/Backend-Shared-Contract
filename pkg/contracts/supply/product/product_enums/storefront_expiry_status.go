package product_enums

// StorefrontExpiryStatus is the broad customer-facing expiry state. AlertLevel
// below carries the existing 30-day/7-day warning or critical classification.
type StorefrontExpiryStatus string

const (
	StorefrontExpiryStatusNotApplicable StorefrontExpiryStatus = "not_applicable"
	StorefrontExpiryStatusSoonExpiry    StorefrontExpiryStatus = "soon_expiry"
	StorefrontExpiryStatusExpired       StorefrontExpiryStatus = "expired"
)

func (s StorefrontExpiryStatus) IsValid() bool {
	switch s {
	case StorefrontExpiryStatusNotApplicable, StorefrontExpiryStatusSoonExpiry, StorefrontExpiryStatusExpired:
		return true
	}
	return false
}

func (s StorefrontExpiryStatus) String() string { return string(s) }
