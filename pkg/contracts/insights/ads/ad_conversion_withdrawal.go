package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// AdConversionWithdrawal is the reversal evidence the ads module needs when an
// order it already reported as a purchase is cancelled or refunded, so it can
// withdraw or restate that conversion on the platform.
//
// Full distinguishes the two cases a platform treats differently: a full
// reversal retracts the conversion entirely, while a partial one leaves it in
// place at a lower value. WithdrawnRevenue is how much attributed revenue is
// being taken back, and RemainingRevenue is what is still attributable
// afterwards; both are on the same basis as the paid order's revenue, net of
// tax and discount and excluding shipping, so the restated value can be sent
// without recomputing it. For a full reversal RemainingRevenue is absent
// rather than a zero amount. An absent field is missing evidence, never a
// zero: a consumer that cannot establish the amounts must leave the platform
// conversion alone instead of withdrawing a guess.
type AdConversionWithdrawal struct {
	Full             bool         `json:"full,omitempty"`
	WithdrawnRevenue *money.Money `json:"withdrawn_revenue,omitempty"`
	RemainingRevenue *money.Money `json:"remaining_revenue,omitempty"`
}
