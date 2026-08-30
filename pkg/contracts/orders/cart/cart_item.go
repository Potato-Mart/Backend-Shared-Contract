package cart

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/buyer"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/product"
)

// CartItem is the frozen product and pricing context for one cart line.
type CartItem struct {
	// SKUCode is the frozen SKU code captured when the line was priced.
	SKUCode              string                                `json:"sku_code"`
	ProductName          string                                `json:"product_name"`
	ProductImage         *security.ObjectMedia                 `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption          `json:"product_package_option"`
	CapturedAt           time.Time                             `json:"captured_at"`
	Components           []fulfilment.PricedPackageComponent   `json:"components"`
	TotalBaseUnits       int64                                 `json:"total_base_units"`
	Pricing              *buyer.PricingContext                 `json:"pricing,omitempty"`
	SubstitutionPolicy   order.LooseSubstitutionPolicySnapshot `json:"substitution_policy"`
	Preorder             *order.PreorderItemSnapshot           `json:"preorder,omitempty"`
}
