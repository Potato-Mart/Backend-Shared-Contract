package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/shared"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/promotion"
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

// CouponAssignment is a customer-specific issuance of a Coupon. It is owned by
// the Coupon aggregate; the coupon definition remains the source of discount
// math and eligibility.
type CouponAssignment struct {
	ID                  string                     `json:"id"`
	CouponID            string                     `json:"coupon_id"`
	CouponCode          string                     `json:"coupon_code"`
	CustomerNumber      string                     `json:"customer_number"`
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

// CouponAssignmentSummary is embedded in coupon detail responses so clients can
// show assignment state without fetching the full assignment list.
type CouponAssignmentSummary struct {
	TotalAssignments    int `json:"total_assignments"`
	ActiveAssignments   int `json:"active_assignments"`
	RedeemedAssignments int `json:"redeemed_assignments"`
	VoidedAssignments   int `json:"voided_assignments"`
	ExpiredAssignments  int `json:"expired_assignments"`
}

// CouponDetail expands a Coupon with assignment and usage summaries.
type CouponDetail struct {
	Coupon
	AssignmentSummary CouponAssignmentSummary `json:"assignment_summary"`
}

// CouponIssueSpec describes assignment issuance under a Coupon.
type CouponIssueSpec struct {
	CustomerNumbers []string                   `json:"customer_numbers,omitempty"`
	Source          promotionenum.CouponSource `json:"source,omitempty"`
	ExpiresAt       *time.Time                 `json:"expires_at,omitempty"`
	Note            string                     `json:"note,omitempty"`
}

// CouponRecipientPreview reports the backend-resolved recipients for an issue
// request. Future criteria-based issuance can extend Criteria without changing
// assignment records.
type CouponRecipientPreview struct {
	CustomerNumbers []string `json:"customer_numbers"`
	Count           int      `json:"count"`
}
