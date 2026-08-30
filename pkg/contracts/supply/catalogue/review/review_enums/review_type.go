package review_enums

// ReviewType identifies the subject that a review evaluates.
type ReviewType string

const (
	ReviewTypeOrder       ReviewType = "order"
	ReviewTypeProduct     ReviewType = "product"
	ReviewTypeCampaign    ReviewType = "campaign"
	ReviewTypeAppFeedback ReviewType = "app_feedback"
)

func (t ReviewType) String() string { return string(t) }

func (t ReviewType) IsValid() bool {
	switch t {
	case ReviewTypeOrder, ReviewTypeProduct, ReviewTypeCampaign, ReviewTypeAppFeedback:
		return true
	default:
		return false
	}
}
