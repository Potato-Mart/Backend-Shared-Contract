package retail

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// RetailCustomerCommerceProfile groups aggregated commerce statistics. Values
// are computed by sync jobs and must never be manually edited.
type RetailCustomerCommerceProfile struct {
	TotalOrders       int         `json:"total_orders"`
	TotalUnits        int         `json:"total_units,omitempty"`
	TotalSpend        money.Money `json:"total_spend"`
	AverageOrderValue money.Money `json:"average_order_value"`
	FirstOrderAt      *time.Time  `json:"first_order_at,omitempty"`
	LastOrderAt       *time.Time  `json:"last_order_at,omitempty"`
	Provinces         []string    `json:"provinces,omitempty"`
	Suburbs           []string    `json:"suburbs,omitempty"`
	SyncedAt          *time.Time  `json:"synced_at,omitempty"`
}
