package quote

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// CartManualDiscountEvidence is the Pricing-owned evidence of one cart-level
// manual discount. Pricing is the only price authority: it applies the
// discount to merchandise after both promotion stages and after the coupon,
// never to shipping, and allocates AppliedAmount and AppliedTax onto the line
// snapshots (PriceSnapshot.ManualDiscount). Transaction owners freeze this
// evidence; they never compute or re-derive the discounted amount.
//
// Exactly one requested value is present for Mode: PercentBasisPoints for
// percent (10000 is 100%), RequestedAmount for amount, and TargetTotal for
// set_total. All amounts are gross, tax-inclusive merchandise values.
type CartManualDiscountEvidence struct {
	Mode               quote_enums.ManualDiscountMode `json:"mode"`
	PercentBasisPoints *int64                         `json:"percent_basis_points,omitempty"`
	RequestedAmount    *money.Money                   `json:"requested_amount,omitempty"`
	TargetTotal        *money.Money                   `json:"target_total,omitempty"`
	// PreDiscountTotal is the gross merchandise total after promotions and
	// the coupon and before this discount. It is the basis the discount was
	// valued on; shipping is never part of it.
	PreDiscountTotal money.Money `json:"pre_discount_total"`
	// AppliedAmount is the gross amount actually taken off and AppliedTax its
	// tax portion. They equal the sums of the line allocations.
	AppliedAmount money.Money `json:"applied_amount"`
	AppliedTax    money.Money `json:"applied_tax"`
	// CouponCode and CouponAmount record the coupon that was valued on the
	// pre-discount basis before this discount applied, when there was one.
	CouponCode   string       `json:"coupon_code,omitempty"`
	CouponAmount *money.Money `json:"coupon_amount,omitempty"`
	// ReasonCode is deliberately a free-form, service-validated string, not
	// an enum, because POS discount reasons are owner-configured.
	// ReasonCode categorizes the discount; Reason is the optional human
	// explanation. ActorUserID is the staff member who set the discount.
	ReasonCode  string    `json:"reason_code"`
	Reason      string    `json:"reason,omitempty"`
	ActorUserID string    `json:"actor_user_id"`
	AppliedAt   time.Time `json:"applied_at"`
}
