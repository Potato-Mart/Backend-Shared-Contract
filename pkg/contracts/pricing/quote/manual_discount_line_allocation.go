package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// ManualDiscountLineAllocation is the share of a cart-level manual discount
// Pricing allocated to one priced line: the gross amount and its tax.
type ManualDiscountLineAllocation struct {
	AllocatedGross money.Money `json:"allocated_gross"`
	AllocatedTax   money.Money `json:"allocated_tax"`
}
