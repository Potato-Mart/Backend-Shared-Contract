package orders

import "github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/packaging"

// OrderEditedItemSnapshot is the customer-safe requested package composition
// of one order line before or after a committed product edit. It excludes
// payment details and mutable picked, packed or substituted composition.
type OrderEditedItemSnapshot struct {
	OrderItemID          string                               `json:"order_item_id"`
	SKUCode              string                               `json:"sku_code"`
	ProductName          string                               `json:"product_name"`
	PackageOption        packaging.PackageOptionRef           `json:"package_option"`
	RequestedComposition packaging.PackageCompositionSnapshot `json:"requested_composition"`
}
