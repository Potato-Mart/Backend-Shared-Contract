package review

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/review/review_enums"
)

// ReviewSubject identifies the internal subject being evaluated. Reference is
// the stable order number, product SKU code, campaign code, or app-surface
// code for the declared Type.
type ReviewSubject struct {
	Type        review_enums.ReviewType      `json:"type"`
	Reference   string                       `json:"reference"`
	DisplayName []localization.LocalizedName `json:"display_name"`
}
