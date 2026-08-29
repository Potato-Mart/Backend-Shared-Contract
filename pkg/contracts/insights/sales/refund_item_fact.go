package sales

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"

// RefundItemFact identifies quantities and value reversed by a completed
// line-level refund. Amount-only refunds intentionally carry no item rows.
type RefundItemFact struct {
	ProductFactDimensions
	Amount money.Money `json:"amount"`
}
