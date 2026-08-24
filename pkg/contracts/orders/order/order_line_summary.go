package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/product"
)

// OrderLineSummary is a customer-facing package-aware order-line snapshot with
// direct frozen product facts, rather than an embedded catalogue projection.
type OrderLineSummary struct {
	// SKUCode is the frozen SKU code captured when the line was priced.
	SKUCode               string                               `json:"sku_code"`
	ProductName           string                               `json:"product_name"`
	ProductImage          *security.ObjectMedia                `json:"product_image,omitempty"`
	ProductPackageOption  product.ProductPackageOption         `json:"product_package_option"`
	CapturedAt            time.Time                            `json:"captured_at"`
	PackageCount          int64                                `json:"package_count"`
	TotalBaseUnits        int64                                `json:"total_base_units"`
	PackagePrice          money.Money                          `json:"package_price"`
	Subtotal              money.Money                          `json:"subtotal"`
	TaxAmount             money.Money                          `json:"tax_amount"`
	DiscountAmount        money.Money                          `json:"discount_amount"`
	PromotionApplications []promotion.PromotionApplication     `json:"promotion_applications"`
	PackedComposition     packaging.PackageCompositionSnapshot `json:"packed_composition"`
	ReturnedComposition   packaging.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition   packaging.PackageCompositionSnapshot `json:"refunded_composition"`
	Total                 money.Money                          `json:"total"`
}
