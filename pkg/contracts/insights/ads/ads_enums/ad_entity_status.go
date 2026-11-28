package ads_enums

// AdEntityStatus is the delivery state of a campaign, ad set or ad as the
// platform reports it, normalised across Meta and Google. It is the state the
// owner asked for, not the platform policy verdict, which AdReviewStatus
// carries separately.
type AdEntityStatus string

const (
	// AdEntityStatusActive is enabled and eligible to spend.
	AdEntityStatusActive AdEntityStatus = "active"
	// AdEntityStatusPaused is switched off and not spending.
	AdEntityStatusPaused AdEntityStatus = "paused"
	// AdEntityStatusArchived is kept for reporting only and cannot resume
	// delivery.
	AdEntityStatusArchived AdEntityStatus = "archived"
	// AdEntityStatusDeleted was removed on the platform. Historical spend
	// stays attributable to it.
	AdEntityStatusDeleted AdEntityStatus = "deleted"
)

// IsValid reports whether s is a known AdEntityStatus.
func (s AdEntityStatus) IsValid() bool {
	switch s {
	case AdEntityStatusActive, AdEntityStatusPaused, AdEntityStatusArchived, AdEntityStatusDeleted:
		return true
	}
	return false
}

func (s AdEntityStatus) String() string { return string(s) }
