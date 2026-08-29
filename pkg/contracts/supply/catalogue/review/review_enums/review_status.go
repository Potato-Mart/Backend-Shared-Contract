package review_enums

// ReviewStatus is the current moderation and publication state of a review.
// A modification is represented by a new revision and history entry, rather
// than by a separate status value.
type ReviewStatus string

const (
	ReviewStatusDraft           ReviewStatus = "draft"
	ReviewStatusPendingApproval ReviewStatus = "pending_approval"
	ReviewStatusPublished       ReviewStatus = "published"
	ReviewStatusRejected        ReviewStatus = "rejected"
	ReviewStatusSuppressed      ReviewStatus = "suppressed"
)

func (s ReviewStatus) String() string { return string(s) }

func (s ReviewStatus) IsValid() bool {
	switch s {
	case ReviewStatusDraft, ReviewStatusPendingApproval, ReviewStatusPublished,
		ReviewStatusRejected, ReviewStatusSuppressed:
		return true
	default:
		return false
	}
}
