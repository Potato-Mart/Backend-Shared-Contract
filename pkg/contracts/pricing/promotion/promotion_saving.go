package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// PromotionSaving attributes one discount amount to the promotion revision that
// produced it. Names repeat approved promotion copy and are omitted when the
// promotion has none. LineID is omitted for an order-level saving.
type PromotionSaving struct {
	PromotionID       string                       `json:"promotion_id"`
	PromotionRevision int64                        `json:"promotion_revision"`
	PromotionKind     string                       `json:"promotion_kind"`
	Names             []localization.LocalizedName `json:"names,omitempty"`
	LineID            string                       `json:"line_id,omitempty"`
	DiscountAmount    money.Money                  `json:"discount_amount"`
}
