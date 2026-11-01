package payments

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// RefundAlternativeLine proposes one alternative product for an item on
// the refund request. Stock is not reserved by the proposal.
type RefundAlternativeLine struct {
	OriginalLineRef    string      `json:"original_line_ref"`
	OriginalSKUCode    string      `json:"original_sku_code,omitempty"`
	AlternativeSKUCode string      `json:"alternative_sku_code"`
	AlternativeName    string      `json:"alternative_name,omitempty"`
	Quantity           int         `json:"quantity"`
	UnitPrice          money.Money `json:"unit_price"`
}
