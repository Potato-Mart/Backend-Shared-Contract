package sales

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// OrderItemPromotionAllocation identifies the packages and discount allocated
// to a paid item by one applied promotion, in the item's currency.
type OrderItemPromotionAllocation struct {
	PromotionID         string      `json:"promotion_id"`
	AppliedPackageCount int64       `json:"applied_package_count"`
	DiscountAmount      money.Money `json:"discount_amount"`
}
