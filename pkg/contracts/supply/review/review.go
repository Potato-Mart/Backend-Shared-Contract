package review

import (
	"time"
)

// RatingDistributionBucket records the number of ratings for one score.
type RatingDistributionBucket struct {
	Score int   `json:"score"`
	Count int64 `json:"count"`
}

// RatingSummary is the public aggregate for one product. Distribution always
// contains exactly five entries ordered by Score from 1 through 5.
type RatingSummary struct {
	AverageScore             float64                     `json:"average_score"`
	RatingCount              int64                       `json:"rating_count"`
	PublishedTextReviewCount int64                       `json:"published_text_review_count"`
	Distribution             [5]RatingDistributionBucket `json:"distribution"`
}

// ProductReview is a public, customer-safe review projection. Title and Body
// contain only approved text; customer and moderation identity never appear.
type ProductReview struct {
	ID               string    `json:"id"`
	ProductSKUCode   string    `json:"product_sku_code"`
	Score            int       `json:"score"`
	Title            string    `json:"title,omitempty"`
	Body             string    `json:"body,omitempty"`
	Locale           string    `json:"locale,omitempty"`
	VerifiedPurchase bool      `json:"verified_purchase"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// MyProductReview is the customer-owned projection of a review. It exposes the
// customer's original text and customer-safe moderation outcome without any
// customer, account, or user identifier.
type MyProductReview struct {
	ID               string                 `json:"id"`
	ProductSKUCode   string                 `json:"product_sku_code"`
	Score            int                    `json:"score"`
	Title            string                 `json:"title,omitempty"`
	Body             string                 `json:"body,omitempty"`
	OriginalTitle    string                 `json:"original_title,omitempty"`
	OriginalBody     string                 `json:"original_body,omitempty"`
	Locale           string                 `json:"locale,omitempty"`
	VerifiedPurchase bool                   `json:"verified_purchase"`
	ModerationStatus ReviewModerationStatus `json:"moderation_status"`
	RejectionReason  ReviewRejectionReason  `json:"rejection_reason,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
}

// ProductReviewModeration is the PII-free admin moderation projection. Its
// internal note and moderation timestamp are intentionally absent from both
// ProductReview and MyProductReview.
type ProductReviewModeration struct {
	ID               string                 `json:"id"`
	ProductSKUCode   string                 `json:"product_sku_code"`
	Score            int                    `json:"score"`
	Title            string                 `json:"title,omitempty"`
	Body             string                 `json:"body,omitempty"`
	OriginalTitle    string                 `json:"original_title,omitempty"`
	OriginalBody     string                 `json:"original_body,omitempty"`
	Locale           string                 `json:"locale,omitempty"`
	VerifiedPurchase bool                   `json:"verified_purchase"`
	ModerationStatus ReviewModerationStatus `json:"moderation_status"`
	RejectionReason  ReviewRejectionReason  `json:"rejection_reason,omitempty"`
	ModerationNote   string                 `json:"moderation_note,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
	ModeratedAt      *time.Time             `json:"moderated_at,omitempty"`
}
