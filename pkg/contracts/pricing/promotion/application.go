package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
)

// PromotionApplication is the immutable result of applying one promotion
// relation to products in a sale order. It retains resolved inputs and
// customer-approved copy without embedding the mutable Promotion definition.
type PromotionApplication struct {
	PromotionID       string `json:"promotion_id"`
	PromotionKind     string `json:"promotion_kind"`
	PromotionRevision int64  `json:"promotion_revision"`
	RelationID        string `json:"relation_id"`

	ResolvedQualifierProductSKUCodes []string                     `json:"resolved_qualifier_product_sku_codes,omitempty"`
	ResolvedTargetProductSKUCodes    []string                     `json:"resolved_target_product_sku_codes,omitempty"`
	ResolvedTerms                    []PromotionTerm              `json:"resolved_terms,omitempty"`
	ResolvedAmounts                  []PromotionAmount            `json:"resolved_amounts,omitempty"`
	DisplayMessages                  []localization.LocalizedText `json:"display_messages,omitempty"`
	ReceiptMessages                  []localization.LocalizedText `json:"receipt_messages,omitempty"`
	AppliedAt                        time.Time                    `json:"applied_at"`
}
