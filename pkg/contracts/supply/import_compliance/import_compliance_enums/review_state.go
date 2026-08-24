package import_compliance_enums

type ReviewState string

const (
	ReviewStateDraft    ReviewState = "draft"
	ReviewStateInReview ReviewState = "in_review"
	ReviewStateApproved ReviewState = "approved"
	ReviewStateRejected ReviewState = "rejected"
	ReviewStateArchived ReviewState = "archived"
)

func (s ReviewState) IsValid() bool {
	switch s {
	case ReviewStateDraft, ReviewStateInReview, ReviewStateApproved, ReviewStateRejected, ReviewStateArchived:
		return true
	}
	return false
}

func (s ReviewState) String() string { return string(s) }
