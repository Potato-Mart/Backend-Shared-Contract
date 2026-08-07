package review

// ReviewModerationStatus is the customer-safe lifecycle state of review text.
type ReviewModerationStatus string

const (
	ReviewModerationStatusNotRequired ReviewModerationStatus = "not_required"
	ReviewModerationStatusPending     ReviewModerationStatus = "pending"
	ReviewModerationStatusApproved    ReviewModerationStatus = "approved"
	ReviewModerationStatusRejected    ReviewModerationStatus = "rejected"
	ReviewModerationStatusSuppressed  ReviewModerationStatus = "suppressed"
)

func (s ReviewModerationStatus) String() string { return string(s) }

func (s ReviewModerationStatus) IsValid() bool {
	switch s {
	case ReviewModerationStatusNotRequired, ReviewModerationStatusPending,
		ReviewModerationStatusApproved, ReviewModerationStatusRejected,
		ReviewModerationStatusSuppressed:
		return true
	default:
		return false
	}
}

// ReviewRejectionReason is a customer-safe explanation category. Internal
// moderation notes and actor identity are deliberately not represented here.
type ReviewRejectionReason string

const (
	ReviewRejectionReasonSpam                ReviewRejectionReason = "spam"
	ReviewRejectionReasonOffTopic            ReviewRejectionReason = "off_topic"
	ReviewRejectionReasonInappropriate       ReviewRejectionReason = "inappropriate"
	ReviewRejectionReasonPersonalInformation ReviewRejectionReason = "personal_information"
	ReviewRejectionReasonUnsupportedLanguage ReviewRejectionReason = "unsupported_language"
	ReviewRejectionReasonOther               ReviewRejectionReason = "other"
)

func (r ReviewRejectionReason) String() string { return string(r) }

func (r ReviewRejectionReason) IsValid() bool {
	switch r {
	case ReviewRejectionReasonSpam, ReviewRejectionReasonOffTopic,
		ReviewRejectionReasonInappropriate, ReviewRejectionReasonPersonalInformation,
		ReviewRejectionReasonUnsupportedLanguage, ReviewRejectionReasonOther:
		return true
	default:
		return false
	}
}

// ReviewErrorCode is a stable domain error code shared by review clients.
type ReviewErrorCode string

const (
	ReviewErrorCodeNotFound         ReviewErrorCode = "REVIEW_NOT_FOUND"
	ReviewErrorCodePurchaseRequired ReviewErrorCode = "REVIEW_PURCHASE_REQUIRED"
)

func (c ReviewErrorCode) String() string { return string(c) }

func (c ReviewErrorCode) IsValid() bool {
	switch c {
	case ReviewErrorCodeNotFound, ReviewErrorCodePurchaseRequired:
		return true
	default:
		return false
	}
}
