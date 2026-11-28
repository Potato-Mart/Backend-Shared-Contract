package ads_enums

// AdReviewStatus is the platform advertising-policy verdict on an object,
// normalised across Meta and Google. It is independent of AdEntityStatus: an
// approved ad can still be paused, and a rejected ad never delivers.
type AdReviewStatus string

const (
	// AdReviewStatusPending is awaiting the platform review.
	AdReviewStatusPending AdReviewStatus = "pending"
	// AdReviewStatusApproved passed the platform review.
	AdReviewStatusApproved AdReviewStatus = "approved"
	// AdReviewStatusRejected was refused by the platform. The reason codes
	// travel with the record that carries this status.
	AdReviewStatusRejected AdReviewStatus = "rejected"
)

// IsValid reports whether s is a known AdReviewStatus.
func (s AdReviewStatus) IsValid() bool {
	switch s {
	case AdReviewStatusPending, AdReviewStatusApproved, AdReviewStatusRejected:
		return true
	}
	return false
}

func (s AdReviewStatus) String() string { return string(s) }
