package payment_enums

// MerchantProfileStatus is the lifecycle state of an effective-dated merchant
// legal profile.
type MerchantProfileStatus string

const (
	MerchantProfileStatusDraft    MerchantProfileStatus = "draft"
	MerchantProfileStatusActive   MerchantProfileStatus = "active"
	MerchantProfileStatusInactive MerchantProfileStatus = "inactive"
)

// IsValid reports whether s is a known MerchantProfileStatus.
func (s MerchantProfileStatus) IsValid() bool {
	switch s {
	case MerchantProfileStatusDraft, MerchantProfileStatusActive, MerchantProfileStatusInactive:
		return true
	}
	return false
}

func (s MerchantProfileStatus) String() string { return string(s) }
