package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	ID            string             `json:"id"`
	Code          string             `json:"code"`
	Description   string             `json:"description,omitempty"`
	DiscountType  enums.DiscountType `json:"discount_type"`
	DiscountValue string             `json:"discount_value"`

	MinOrderAmount    *common.Money `json:"min_order_amount,omitempty"`
	MaxDiscountAmount *common.Money `json:"max_discount_amount,omitempty"`
	UsageLimit        int           `json:"usage_limit,omitempty"`
	UsedCount         int           `json:"used_count"`
	PerCustomerLimit  int           `json:"per_customer_limit,omitempty"`

	StartsAt  *time.Time `json:"starts_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	IsActive  bool       `json:"is_active"`

	AppliesTo   enums.CouponAppliesTo `json:"applies_to"`
	ProductIDs  []string              `json:"product_ids,omitempty"`
	CategoryIDs []string              `json:"category_ids,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomerCoupon is an assignment of a coupon to a specific customer,
// created by the system (RFM comeback, birthday, referral, etc.) or manually.
type CustomerCoupon struct {
	ID                string             `json:"id"`
	CustomerProfileID string             `json:"customer_profile_id"`
	CouponID          string             `json:"coupon_id"`
	Source            enums.CouponSource `json:"source"`
	ExpiresAt         *time.Time         `json:"expires_at,omitempty"`
	RedeemedAt        *time.Time         `json:"redeemed_at,omitempty"`
	RedeemedOrderID   string             `json:"redeemed_order_id,omitempty"`
	Note              string             `json:"note,omitempty"`
	CreatedAt         time.Time          `json:"created_at"`
}
