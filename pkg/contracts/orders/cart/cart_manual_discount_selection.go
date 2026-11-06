package cart

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// CartManualDiscountSelection is the staff-entered cart-level manual
// discount request Orders keeps on a cart and forwards to Pricing. It is a
// request, not a price: Pricing values it after promotions and the coupon,
// on merchandise only, and returns quote.CartManualDiscountEvidence.
//
// Exactly one requested value is present for Mode: PercentBasisPoints for
// percent (10000 is 100%), RequestedAmount for amount, and TargetTotal for
// set_total.
// Amounts are gross, tax-inclusive merchandise values.
type CartManualDiscountSelection struct {
	Mode               quote_enums.ManualDiscountMode `json:"mode"`
	PercentBasisPoints *int64                         `json:"percent_basis_points,omitempty"`
	RequestedAmount    *money.Money                   `json:"requested_amount,omitempty"`
	TargetTotal        *money.Money                   `json:"target_total,omitempty"`
	// ReasonCode is a deliberately free-form, service-validated code, not an
	// enum: POS reasons are owner-configured and change without a release.
	ReasonCode string `json:"reason_code"`
	Reason     string `json:"reason,omitempty"`
	// ActorUserID is the staff member who set the discount; Orders checks
	// their authority before keeping or forwarding the selection.
	ActorUserID string    `json:"actor_user_id"`
	SetAt       time.Time `json:"set_at"`
}
