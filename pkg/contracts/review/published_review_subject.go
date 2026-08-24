package review

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/review/review_enums"
)

// PublishedReviewSubject is a customer-safe review subject. Services must
// omit Reference for order reviews, because an order number is never public.
type PublishedReviewSubject struct {
	Type        review_enums.ReviewType      `json:"type"`
	Reference   string                       `json:"reference,omitempty"`
	DisplayName []localization.LocalizedName `json:"display_name"`
}
