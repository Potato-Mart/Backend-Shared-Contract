package purchase

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/product"
)

type OrderItem struct {
	ID string `json:"id,omitempty"`
	// SKUCode is the frozen SKU code captured when the purchase line was raised.
	SKUCode              string                               `json:"sku_code"`
	ProductName          string                               `json:"product_name"`
	ProductImage         *security.ObjectMedia                `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption         `json:"product_package_option"`
	CapturedAt           time.Time                            `json:"captured_at"`
	PackageOptionCode    string                               `json:"package_option_code"`
	UnitCost             money.Money                          `json:"unit_cost"`
	OrderedComposition   packaging.PackageCompositionSnapshot `json:"ordered_composition"`
	ReceivedComposition  packaging.PackageCompositionSnapshot `json:"received_composition"`
	RejectedComposition  packaging.PackageCompositionSnapshot `json:"rejected_composition"`
	LineTotal            money.Money                          `json:"line_total"`
	Note                 string                               `json:"note,omitempty"`
}
