package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/temporal"
)

// ShippingArrivalRule represents a weekly warehouse-arrival window.
// Days of week: 0=Sunday … 6=Saturday. Times are interpreted in Timezone.
type ShippingArrivalRule struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Timezone   string             `json:"timezone"`
	FromDOW    int                `json:"from_dow"` // 0-6
	FromTime   temporal.TimeOfDay `json:"from_time"`
	ToDOW      int                `json:"to_dow"` // 0-6
	ToTime     temporal.TimeOfDay `json:"to_time"`
	ArrivalDOW int                `json:"arrival_dow"` // 0-6
	WeekOffset int                `json:"week_offset"` // 0 = this week, 1 = next week
	IsActive   bool               `json:"is_active"`
	audit.AuditFields
}

// ShippingArrivalBlacklist represents a local calendar date on which warehouse
// arrivals are unavailable.
type ShippingArrivalBlacklist struct {
	ID          string        `json:"id"`
	BlockedDate temporal.Date `json:"blocked_date"`
	Timezone    string        `json:"timezone"`
	Reason      string        `json:"reason,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
}
