package orders

import "github.com/Potato-Mart/Backend-Shared-Contract/v37/pkg/contracts/common/packaging"

// OrderEditedItemSnapshot is the customer-safe requested package composition
// of one order line before or after a committed product edit. It excludes
// payment details and mutable picked, packed or substituted composition.
type OrderEditedItemSnapshot struct {
	OrderItemID          string                               `json:"order_item_id"`
	SKUCode              string                               `json:"sku_code"`
	ProductName          string                               `json:"product_name"`
	PackageOptionCode    string                               `json:"package_option_code"`
	RequestedComposition packaging.PackageCompositionSnapshot `json:"requested_composition"`
}
