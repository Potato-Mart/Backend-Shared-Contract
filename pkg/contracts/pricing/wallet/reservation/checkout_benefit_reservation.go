package reservation

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

// CheckoutBenefitReservation is the durable checkout hold for applied benefits.
type CheckoutBenefitReservation struct {
	ID          string                                        `json:"id"`
	Owner       benefit.OwnerRef                              `json:"owner"`
	OrderNumber string                                        `json:"order_number,omitempty"`
	Coupon      *CouponBenefitReservation                     `json:"coupon,omitempty"`
	Voucher     *VoucherBenefitReservation                    `json:"voucher,omitempty"`
	GiftCards   []GiftCardBenefitReservation                  `json:"gift_cards,omitempty"`
	Status      wallet_enums.CheckoutBenefitReservationStatus `json:"status"`
	ExpiresAt   time.Time                                     `json:"expires_at"`
	CommittedAt *time.Time                                    `json:"committed_at,omitempty"`
	CancelledAt *time.Time                                    `json:"cancelled_at,omitempty"`
	RefundedAt  *time.Time                                    `json:"refunded_at,omitempty"`

	audit.AuditFields
	security.DataProtectionFields
}
