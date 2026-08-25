package review

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/customers/review/review_enums"
)

// Review is the protected internal review record. CustomerNumber is mandatory
// for every review, including app feedback, and must never be included in a
// public review projection.
type Review struct {
	ID                  string                             `json:"id"`
	Subject             ReviewSubject                      `json:"subject"`
	CustomerNumber      string                             `json:"customer_number"`
	ReviewerDisplayName string                             `json:"reviewer_display_name"`
	SubmissionChannel   string                             `json:"submission_channel"`
	Qualification       ReviewQualification                `json:"qualification"`
	Content             ReviewContent                      `json:"content"`
	Status              review_enums.ReviewStatus          `json:"status"`
	Revision            int64                              `json:"revision"`
	LastPublished       *PublishedReview                   `json:"last_published,omitempty"`
	RejectionReason     review_enums.ReviewRejectionReason `json:"rejection_reason,omitempty"`
	ModerationNote      string                             `json:"moderation_note,omitempty"`
	History             []security.HistoryEntry            `json:"history,omitempty"`

	audit.AuditFields
	security.DataProtectionFields
}
