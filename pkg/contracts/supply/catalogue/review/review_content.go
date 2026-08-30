package review

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
)

// ReviewContent is the customer-submitted material for a review. Score is a
// required integer in the inclusive range 1 through 5; services enforce that
// rule before accepting or publishing a revision.
type ReviewContent struct {
	Body   []localization.LocalizedText `json:"body"`
	Score  int                          `json:"score"`
	Images []security.ObjectMedia       `json:"images,omitempty"`
	Video  *security.ObjectMedia        `json:"video,omitempty"`
}
