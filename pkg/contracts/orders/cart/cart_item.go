package cart

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/orders/buyer"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/orders/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/catalogue/product"
)

// CartItem is the frozen product selection and, when available, pricing context.
type CartItem struct {
	// SKUCode is the frozen SKU code captured for this selection.
	SKUCode              string                       `json:"sku_code"`
	ProductName          string                       `json:"product_name"`
	ProductImage         *security.ObjectMedia        `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption `json:"product_package_option"`
	CapturedAt           time.Time                    `json:"captured_at"`
	// UnpricedSelection is a POS selection before Pricing issues its quote.
	// Owners require either this selection or priced components, never both.
	UnpricedSelection  *CartPackageSelection                 `json:"unpriced_selection,omitempty"`
	Components         []fulfilment.PricedPackageComponent   `json:"components"`
	TotalBaseUnits     int64                                 `json:"total_base_units"`
	Pricing            *buyer.PricingContext                 `json:"pricing,omitempty"`
	SubstitutionPolicy order.LooseSubstitutionPolicySnapshot `json:"substitution_policy"`
	Preorder           *order.PreorderItemSnapshot           `json:"preorder,omitempty"`
}
