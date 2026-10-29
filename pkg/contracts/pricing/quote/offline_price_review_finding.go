package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/money"

// OfflinePriceReviewFinding records a pricing issue found while reconciling an
// offline sale. A promotion may be unsupported before a comparable amount is
// known; nil amounts and difference mean unknown, not zero. The owning service
// decides when a finding is raised and how it is resolved.
type OfflinePriceReviewFinding struct {
	ClientSaleID   string       `json:"client_sale_id"`
	LineID         string       `json:"line_id"`
	LineNumber     int64        `json:"line_number"`
	SKUCode        string       `json:"sku_code"`
	PromotionID    string       `json:"promotion_id,omitempty"`
	Reason         string       `json:"reason"`
	ChargedAmount  *money.Money `json:"charged_amount,omitempty"`
	ExpectedAmount *money.Money `json:"expected_amount,omitempty"`
	Difference     *money.Money `json:"difference,omitempty"`
}
