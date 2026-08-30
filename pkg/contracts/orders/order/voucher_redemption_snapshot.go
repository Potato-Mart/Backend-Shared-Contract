package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// VoucherRedemptionSnapshot records the single voucher applied to an order.
type VoucherRedemptionSnapshot struct {
	VoucherCode   string      `json:"voucher_code"`
	AppliedAmount money.Money `json:"applied_amount"`
	ReservationID string      `json:"reservation_id,omitempty"`
	OccurredAt    *time.Time  `json:"occurred_at,omitempty"`
}
