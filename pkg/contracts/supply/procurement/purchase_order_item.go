package procurement

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/catalogue/product"
)

// PurchaseOrderItem is one frozen supplier order line.
type PurchaseOrderItem struct {
	ID string `json:"id,omitempty"`
	// SKUCode is the frozen SKU code captured when the purchase line was raised.
	SKUCode              string                       `json:"sku_code"`
	ProductName          string                       `json:"product_name"`
	ProductImage         *security.ObjectMedia        `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption `json:"product_package_option"`
	CapturedAt           time.Time                    `json:"captured_at"`
	PackageOption        packaging.PackageOptionRef   `json:"package_option"`
	// UnitCost is the supplier cost per base unit. Nil means this basis is
	// absent, not a free unit or a rounded selected-package price. An explicit
	// zero Money is evidence of a zero-cost base unit.
	UnitCost *money.Money `json:"unit_cost,omitempty"`
	// SelectedPackageCost is the exact supplier cost of ONE package identified
	// by the frozen, versioned PackageOption. Supply validates that the ordered
	// composition uses that selected package, that currency/counts are valid,
	// and that UnitCost and SelectedPackageCost are mutually exclusive bases.
	// Exact package cost and line totals never require division into minor
	// units per base unit. Nil does not authorize deriving a package price.
	SelectedPackageCost *money.Money                         `json:"selected_package_cost,omitempty"`
	OrderedComposition  packaging.PackageCompositionSnapshot `json:"ordered_composition"`
	ReceivedComposition packaging.PackageCompositionSnapshot `json:"received_composition"`
	RejectedComposition packaging.PackageCompositionSnapshot `json:"rejected_composition"`
	LineTotal           money.Money                          `json:"line_total"`
	// NetGoodsAmount is the validated, frozen total for
	// OrderedComposition.TotalBaseUnits after goods discounts and excluding
	// all tax, freight, and duty. Supply owns confirmation and provisional
	// valuation; absent evidence does not imply LineTotal has this basis.
	NetGoodsAmount *money.Money `json:"net_goods_amount,omitempty"`
	Note           string       `json:"note,omitempty"`
}
