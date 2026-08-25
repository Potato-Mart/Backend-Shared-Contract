package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
)

// SupplierAvailableProduct describes one supplier's terms and identity for a
// Potato Mart SKU. Unobserved commercial terms remain absent.
type SupplierAvailableProduct struct {
	SKUCode                 string                       `json:"sku_code"`
	SupplierSKUCode         string                       `json:"supplier_sku_code,omitempty"`
	ProductNames            []localization.LocalizedName `json:"product_names"`
	OfferedPrice            *money.Money                 `json:"offered_price,omitempty"`
	MinimumPurchaseQuantity int64                        `json:"minimum_purchase_quantity,omitempty"`
	Manufacturing           *ProductManufacturing        `json:"manufacturing,omitempty"`
}
