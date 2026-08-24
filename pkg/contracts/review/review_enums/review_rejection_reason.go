package review_enums

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
