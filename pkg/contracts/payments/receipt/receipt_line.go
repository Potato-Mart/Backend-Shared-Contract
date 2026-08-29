package receipt

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/product"
	"time"
)

// ReceiptLine is the customer-safe, immutable product and monetary evidence
// recorded for one receipt line. It never exposes the inventory evidence in
// accepted package pricing.
type ReceiptLine struct {
	// SKUCode is the frozen SKU code captured when the receipt was issued.
	SKUCode               string                           `json:"sku_code"`
	ProductName           string                           `json:"product_name"`
	ProductImage          *security.ObjectMedia            `json:"product_image,omitempty"`
	ProductPackageOption  product.ProductPackageOption     `json:"product_package_option"`
	CapturedAt            time.Time                        `json:"captured_at"`
	PackageCount          int64                            `json:"package_count"`
	TotalBaseUnits        int64                            `json:"total_base_units"`
	PackagePrice          money.Money                      `json:"package_price"`
	Subtotal              money.Money                      `json:"subtotal"`
	TaxAmount             money.Money                      `json:"tax_amount"`
	DiscountAmount        money.Money                      `json:"discount_amount"`
	Total                 money.Money                      `json:"total"`
	PromotionApplications []promotion.PromotionApplication `json:"promotion_applications"`
}
