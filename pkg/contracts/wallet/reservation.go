package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/contracts/membership"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/wallet"
)

// CouponBenefitReservation is the coupon capacity held by one checkout.
type CouponBenefitReservation struct {
	CouponCode    string                                      `json:"coupon_code"`
	AppliedAmount common.Money                                `json:"applied_amount"`
	Status        walletenum.CheckoutBenefitReservationStatus `json:"status"`
}

// VoucherBenefitReservation is the exclusive voucher hold owned by one
// checkout.
type VoucherBenefitReservation struct {
	VoucherCode   string                                      `json:"voucher_code"`
	AppliedAmount common.Money                                `json:"applied_amount"`
	Status        walletenum.CheckoutBenefitReservationStatus `json:"status"`
}

// GiftCardBenefitReservation is one ordered gift-card balance allocation.
// RefundedAmount supports deterministic, idempotent partial refunds without
// losing the original application amount or priority.
type GiftCardBenefitReservation struct {
	GiftCardCode   string                                      `json:"gift_card_code"`
	AppliedAmount  common.Money                                `json:"applied_amount"`
	RefundedAmount common.Money                                `json:"refunded_amount"`
	Status         walletenum.CheckoutBenefitReservationStatus `json:"status"`
}

// CheckoutBenefitReservation is the durable, idempotency-keyed record for the
// coupon-or-voucher and ordered gift-card value held by a checkout. The owning
// service sets ExpiresAt (normally 30 minutes after creation) and releases
// reservations both opportunistically and through its expiry sweep.
type CheckoutBenefitReservation struct {
	ID             string                                      `json:"id"`
	IdempotencyKey string                                      `json:"idempotency_key"`
	Owner          membership.MembershipOwnerRef               `json:"owner"`
	OrderNumber    string                                      `json:"order_number,omitempty"`
	Coupon         *CouponBenefitReservation                   `json:"coupon,omitempty"`
	Voucher        *VoucherBenefitReservation                  `json:"voucher,omitempty"`
	GiftCards      []GiftCardBenefitReservation                `json:"gift_cards,omitempty"`
	Status         walletenum.CheckoutBenefitReservationStatus `json:"status"`
	ExpiresAt      time.Time                                   `json:"expires_at"`
	CommittedAt    *time.Time                                  `json:"committed_at,omitempty"`
	CancelledAt    *time.Time                                  `json:"cancelled_at,omitempty"`
	RefundedAt     *time.Time                                  `json:"refunded_at,omitempty"`
	CreatedAt      time.Time                                   `json:"created_at"`
	UpdatedAt      time.Time                                   `json:"updated_at"`
}
