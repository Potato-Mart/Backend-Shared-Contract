package wallet

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/wallet/wallet_enums"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`

	Scope    promotion.PromotionScope    `json:"scope"`
	Period   promotion.PromotionPeriod   `json:"period"`
	Terms    []promotion.PromotionTerm   `json:"terms,omitempty"`
	Controls promotion.PromotionControls `json:"controls"`
	History  []security.HistoryEntry     `json:"history,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}

// CouponAssignment is an owner-specific issuance of a Coupon. Owner supports
// both retail customers and wholesale organisations. The Coupon aggregate
// remains the source of discount math and eligibility.
type CouponAssignment struct {
	ID                  string                    `json:"id"`
	CouponID            string                    `json:"coupon_id"`
	CouponCode          string                    `json:"coupon_code"`
	Owner               benefit.OwnerRef          `json:"owner"`
	Source              wallet_enums.CouponSource `json:"source"`
	Status              string                    `json:"status"`
	ExpiresAt           *time.Time                `json:"expires_at,omitempty"`
	RedeemedAt          *time.Time                `json:"redeemed_at,omitempty"`
	RedeemedOrderNumber string                    `json:"redeemed_order_number,omitempty"`
	VoidedAt            *time.Time                `json:"voided_at,omitempty"`
	Note                string                    `json:"note,omitempty"`
	History             []security.HistoryEntry   `json:"history,omitempty"`
	CreatedAt           time.Time                 `json:"created_at"`
}

// CouponUsageRecord is Pricing's durable idempotent redemption result.
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
