package reservation

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/wallet/wallet_enums"
)

// CouponBenefitReservation is the coupon capacity held by one checkout.
type CouponBenefitReservation struct {
	CouponID      string                                        `json:"coupon_id"`
	EntitlementID string                                        `json:"entitlement_id"`
	CouponCode    string                                        `json:"coupon_code"`
	AppliedAmount money.Money                                   `json:"applied_amount"`
	Status        wallet_enums.CheckoutBenefitReservationStatus `json:"status"`
}
