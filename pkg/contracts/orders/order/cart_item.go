package order

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/product"
)

// CartItem is the frozen product and pricing context for one cart line.
type CartItem struct {
	// SKUCode is the frozen SKU code captured when the line was priced.
	SKUCode              string                          `json:"sku_code"`
	ProductName          string                          `json:"product_name"`
	ProductImage         *security.ObjectMedia           `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption    `json:"product_package_option"`
	CapturedAt           time.Time                       `json:"captured_at"`
	Components           []PricedPackageComponent        `json:"components"`
	TotalBaseUnits       int64                           `json:"total_base_units"`
	Pricing              *PricingContext                 `json:"pricing,omitempty"`
	SubstitutionPolicy   LooseSubstitutionPolicySnapshot `json:"substitution_policy"`
	Preorder             *PreorderItemSnapshot           `json:"preorder,omitempty"`
}
