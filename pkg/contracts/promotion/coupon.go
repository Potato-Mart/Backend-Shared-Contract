package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	DiscountSpec

	MinOrderAmount    *common.Money `json:"min_order_amount,omitempty"`
	MaxDiscountAmount *common.Money `json:"max_discount_amount,omitempty"`
	UsageLimits
	ActiveWindow

	AppliesTo       enums.CouponAppliesTo `json:"applies_to"`
	ProductSKUCodes []string              `json:"product_sku_codes,omitempty"`
	CategoryTags    []product.CategoryTag `json:"category_tags,omitempty"`
	History         []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// CustomerCoupon is an assignment of a coupon to a specific customer,
// created by the system (RFM comeback, birthday, referral, etc.) or manually.
type CustomerCoupon struct {
	ID                  string             `json:"id"`
	CustomerNumber      string             `json:"customer_number"`
	CouponCode          string             `json:"coupon_code"`
	Source              enums.CouponSource `json:"source"`
	ExpiresAt           *time.Time         `json:"expires_at,omitempty"`
	RedeemedAt          *time.Time         `json:"redeemed_at,omitempty"`
	RedeemedOrderNumber string             `json:"redeemed_order_number,omitempty"`
	Note                string             `json:"note,omitempty"`
	CreatedAt           time.Time          `json:"created_at"`
}
