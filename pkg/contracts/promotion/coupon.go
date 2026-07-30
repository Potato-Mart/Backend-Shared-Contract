package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/shared"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/promotion"
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

	AppliesTo       promotionenum.CouponAppliesTo `json:"applies_to"`
	ProductSKUCodes []string                      `json:"product_sku_codes,omitempty"`
	CategoryTags    []product.CategoryTag         `json:"category_tags,omitempty"`
	History         []shared.HistoryEntry         `json:"history,omitempty"`

	common.AuditFields
}

// CouponAssignment is an owner-specific issuance of a Coupon. Owner supports
// both retail customers and wholesale organisations. The Coupon aggregate
// remains the source of discount math and eligibility.
type CouponAssignment struct {
	ID                  string                     `json:"id"`
	CouponID            string                     `json:"coupon_id"`
	CouponCode          string                     `json:"coupon_code"`
	Owner               benefit.OwnerRef           `json:"owner"`
	Source              promotionenum.CouponSource `json:"source"`
	Status              string                     `json:"status"`
	ExpiresAt           *time.Time                 `json:"expires_at,omitempty"`
	RedeemedAt          *time.Time                 `json:"redeemed_at,omitempty"`
	RedeemedOrderNumber string                     `json:"redeemed_order_number,omitempty"`
	VoidedAt            *time.Time                 `json:"voided_at,omitempty"`
	Note                string                     `json:"note,omitempty"`
	History             []shared.HistoryEntry      `json:"history,omitempty"`
	CreatedAt           time.Time                  `json:"created_at"`
}

// CouponUsageRecord is Pricing's durable idempotent redemption result.
type CouponUsageRecord struct {
	ID                  string            `json:"id"`
	CouponCode          string            `json:"coupon_code"`
	Owner               *benefit.OwnerRef `json:"owner,omitempty"`
	RedeemedOrderNumber string            `json:"redeemed_order_number"`
	DiscountAmount      common.Money      `json:"discount_amount"`
	RedeemedAt          time.Time         `json:"redeemed_at"`
	RefundID            string            `json:"refund_id,omitempty"`
	RefundedAt          *time.Time        `json:"refunded_at,omitempty"`
	CreatedAt           time.Time         `json:"created_at"`
}
