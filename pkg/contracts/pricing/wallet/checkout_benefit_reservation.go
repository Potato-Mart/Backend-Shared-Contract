package wallet

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
	"time"
)

// CheckoutBenefitReservation is the durable idempotency-keyed checkout hold.
type CheckoutBenefitReservation struct {
	ID             string                                        `json:"id"`
	IdempotencyKey string                                        `json:"idempotency_key"`
	Owner          benefit.OwnerRef                              `json:"owner"`
	OrderNumber    string                                        `json:"order_number,omitempty"`
	Coupon         *CouponBenefitReservation                     `json:"coupon,omitempty"`
	Voucher        *VoucherBenefitReservation                    `json:"voucher,omitempty"`
	GiftCards      []GiftCardBenefitReservation                  `json:"gift_cards,omitempty"`
	Status         wallet_enums.CheckoutBenefitReservationStatus `json:"status"`
	ExpiresAt      time.Time                                     `json:"expires_at"`
	CommittedAt    *time.Time                                    `json:"committed_at,omitempty"`
	CancelledAt    *time.Time                                    `json:"cancelled_at,omitempty"`
	RefundedAt     *time.Time                                    `json:"refunded_at,omitempty"`
	CreatedAt      time.Time                                     `json:"created_at"`
	UpdatedAt      time.Time                                     `json:"updated_at"`
}
