package wallet

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

// CouponBenefitReservation is the coupon capacity held by one checkout.
type CouponBenefitReservation struct {
	CouponCode    string                                        `json:"coupon_code"`
	AppliedAmount money.Money                                   `json:"applied_amount"`
	Status        wallet_enums.CheckoutBenefitReservationStatus `json:"status"`
}
