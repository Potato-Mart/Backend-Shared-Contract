package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/packaging"
)

// PromotionTargetPackageAllocation freezes one order item's purchased package
// quantity and discount within a promotion application. OrderItemID and
// PackageOption identify the priced requested component; the canonical
// purchased quantity remains on the order item's requested composition and
// priced component snapshot. GroupOrdinal is one-based for complete repeated
// groups and omitted for an application without grouping.
type PromotionTargetPackageAllocation struct {
	OrderItemID         string                     `json:"order_item_id"`
	SKUCode             string                     `json:"sku_code"`
	PackageOption       packaging.PackageOptionRef `json:"package_option"`
	AppliedPackageCount int64                      `json:"applied_package_count"`
	GroupOrdinal        int64                      `json:"group_ordinal,omitempty"`
	DiscountAmount      money.Money                `json:"discount_amount"`
}
