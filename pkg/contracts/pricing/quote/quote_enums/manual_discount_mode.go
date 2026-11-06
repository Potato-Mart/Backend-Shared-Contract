package quote_enums

// ManualDiscountMode is how a staff member expressed a cart-level manual
// discount. Pricing applies the discount to merchandise after promotions and
// the coupon, and allocates the applied amount onto line snapshots and tax.
// percent carries basis points, amount a gross amount off, and set_total the
// gross merchandise total the cart should reach.
type ManualDiscountMode string

const (
	ManualDiscountModePercent  ManualDiscountMode = "percent"
	ManualDiscountModeAmount   ManualDiscountMode = "amount"
	ManualDiscountModeSetTotal ManualDiscountMode = "set_total"
)

func (m ManualDiscountMode) IsValid() bool {
	switch m {
	case ManualDiscountModePercent, ManualDiscountModeAmount, ManualDiscountModeSetTotal:
		return true
	}
	return false
}

func (m ManualDiscountMode) String() string { return string(m) }
