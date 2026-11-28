package ads_enums

// AdDraftStatus is the state of an ad built in our admin, from the platform
// review through to delivery. Every draft is created paused on the platform,
// so approved means approved and not yet delivering.
type AdDraftStatus string

const (
	// AdDraftStatusInReview was submitted and is awaiting the platform
	// review.
	AdDraftStatusInReview AdDraftStatus = "in_review"
	// AdDraftStatusApproved passed review and is paused on the platform.
	AdDraftStatusApproved AdDraftStatus = "approved"
	// AdDraftStatusRejected was refused by the platform. The draft carries
	// the reason codes.
	AdDraftStatusRejected AdDraftStatus = "rejected"
	// AdDraftStatusActive is approved and delivering after someone resumed
	// it through the delivery-change flow.
	AdDraftStatusActive AdDraftStatus = "active"
)

// IsValid reports whether s is a known AdDraftStatus.
func (s AdDraftStatus) IsValid() bool {
	switch s {
	case AdDraftStatusInReview, AdDraftStatusApproved, AdDraftStatusRejected, AdDraftStatusActive:
		return true
	}
	return false
}

func (s AdDraftStatus) String() string { return string(s) }
