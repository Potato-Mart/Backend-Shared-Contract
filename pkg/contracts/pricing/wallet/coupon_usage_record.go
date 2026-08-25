package wallet

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/benefit"
	"time"
)

// CouponUsageRecord is a durable idempotent coupon redemption result.
type CouponUsageRecord struct {
	ID                  string                      `json:"id"`
	CouponCode          string                      `json:"coupon_code"`
	Owner               *benefit.OwnerRef           `json:"owner,omitempty"`
	RedeemedOrderNumber string                      `json:"redeemed_order_number"`
	DiscountAmount      money.Money                 `json:"discount_amount"`
	GeographicContext   geography.GeographicContext `json:"geographic_context"`
	RedeemedAt          time.Time                   `json:"redeemed_at"`
	RefundID            string                      `json:"refund_id,omitempty"`
	RefundedAt          *time.Time                  `json:"refunded_at,omitempty"`
	CreatedAt           time.Time                   `json:"created_at"`
}
