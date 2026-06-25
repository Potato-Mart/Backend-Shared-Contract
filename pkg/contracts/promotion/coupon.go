package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	ID           string `json:"id"`
	Code         string `json:"code"`
	Description  string `json:"description,omitempty"`
	DiscountSpec `bson:",inline"`

	MinOrderAmount    *common.Money `json:"min_order_amount,omitempty"`
	MaxDiscountAmount *common.Money `json:"max_discount_amount,omitempty"`
	UsageLimits       `bson:",inline"`
	ActiveWindow      `bson:",inline"`

	AppliesTo   enums.CouponAppliesTo `json:"applies_to"`
	ProductIDs  []string              `json:"product_ids,omitempty"`
	CategoryIDs []string              `json:"category_ids,omitempty"`
	History     []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields `bson:",inline"`
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
