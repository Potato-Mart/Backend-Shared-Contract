package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
)

// DiscountSpec is the discount type/value pair shared by promotions and
// coupons. DiscountValue is a string so percentage ("10") and fixed
// amounts ("5.00") share one wire shape; DiscountType disambiguates.
type DiscountSpec struct {
	DiscountType  promotion_enums.DiscountType `json:"discount_type,omitempty"`
	DiscountValue string                       `json:"discount_value,omitempty"`
}

// ActiveWindow groups the activation flag with the optional start/expiry
// window of a promotion or coupon.
type ActiveWindow struct {
	StartsAt         *time.Time `json:"starts_at,omitempty"`
	ExpiresAt        *time.Time `json:"expires_at,omitempty"`
	ScheduleTimezone string     `json:"schedule_timezone"`
	IsActive         bool       `json:"is_active"`
}

// UsageLimits groups redemption caps and counters shared by promotions
// and coupons.
type UsageLimits struct {
	UsageLimit       int `json:"usage_limit,omitempty"`
	UsedCount        int `json:"used_count"`
	PerCustomerLimit int `json:"per_customer_limit,omitempty"`
}
