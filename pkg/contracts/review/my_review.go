package review

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/review/review_enums"
)

// MyReview is the owner-safe review projection. It intentionally omits the
// private customer number, moderation note, actor history, and data-protection
// administration fields.
type MyReview struct {
	ID                  string                             `json:"id"`
	Subject             ReviewSubject                      `json:"subject"`
	ReviewerDisplayName string                             `json:"reviewer_display_name"`
	SubmissionChannel   string                             `json:"submission_channel"`
	Qualification       ReviewQualification                `json:"qualification"`
	Content             ReviewContent                      `json:"content"`
	Status              review_enums.ReviewStatus          `json:"status"`
	Revision            int64                              `json:"revision"`
	LastPublished       *PublishedReview                   `json:"last_published,omitempty"`
	RejectionReason     review_enums.ReviewRejectionReason `json:"rejection_reason,omitempty"`
	CreatedAt           time.Time                          `json:"created_at"`
	UpdatedAt           time.Time                          `json:"updated_at"`
}
