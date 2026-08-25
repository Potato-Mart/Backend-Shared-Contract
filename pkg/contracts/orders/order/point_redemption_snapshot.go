package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
)

// PointRedemptionSnapshot records a points discount applied to an order without
// overloading coupon or promotion fields.
type PointRedemptionSnapshot struct {
	CustomerNumber string      `json:"customer_number"`
	ReservationID  string      `json:"reservation_id,omitempty"`
	LedgerEntryID  string      `json:"ledger_entry_id,omitempty"`
	Points         int         `json:"points"`
	DiscountAmount money.Money `json:"discount_amount"`
	OccurredAt     *time.Time  `json:"occurred_at,omitempty"`
}
