package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/wallet/wallet_enums"
)

// RewardRedemptionSnapshot records a catalog reward applied to an order.
type RewardRedemptionSnapshot struct {
	RewardRedemptionID string                  `json:"reward_redemption_id"`
	RewardCode         string                  `json:"reward_code"`
	CustomerNumber     string                  `json:"customer_number"`
	RewardType         wallet_enums.RewardType `json:"reward_type"`
	PointsSpent        int                     `json:"points_spent"`
	DiscountAmount     *money.Money            `json:"discount_amount,omitempty"`
	SKUCode            string                  `json:"sku_code,omitempty"`
	VoucherCode        string                  `json:"voucher_code,omitempty"`
	OccurredAt         *time.Time              `json:"occurred_at,omitempty"`
}
