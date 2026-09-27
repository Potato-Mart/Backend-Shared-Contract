package promotion

import "github.com/Potato-Mart/Backend-Shared-Contract/v37/pkg/contracts/common/money"

// PromotionTargetPackageAllocation freezes one order item's purchased package
// quantity and discount within a promotion application. OrderItemID and
// PackageOptionCode identify the priced requested component; the canonical
// purchased quantity remains on the order item's requested composition and
// priced component snapshot. GroupOrdinal is one-based for complete repeated
// groups and omitted for an application without grouping.
type PromotionTargetPackageAllocation struct {
	OrderItemID         string      `json:"order_item_id"`
	SKUCode             string      `json:"sku_code"`
	PackageOptionCode   string      `json:"package_option_code"`
	AppliedPackageCount int64       `json:"applied_package_count"`
	GroupOrdinal        int64       `json:"group_ordinal,omitempty"`
	DiscountAmount      money.Money `json:"discount_amount"`
}
