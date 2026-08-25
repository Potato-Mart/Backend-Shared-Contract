package review

import "time"

// PublishedReview is the immutable customer-safe snapshot currently visible
// to customers. When a customer edits a published review, this snapshot stays
// visible until the new revision is approved; rejected edits do not replace it.
type PublishedReview struct {
	ID                  string                 `json:"id"`
	Subject             PublishedReviewSubject `json:"subject"`
	ReviewerDisplayName string                 `json:"reviewer_display_name"`
	SubmissionChannel   string                 `json:"submission_channel"`
	Content             ReviewContent          `json:"content"`
	Revision            int64                  `json:"revision"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	PublishedAt         time.Time              `json:"published_at"`
}
