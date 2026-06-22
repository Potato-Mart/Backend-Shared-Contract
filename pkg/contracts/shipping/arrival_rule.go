package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/common"
)

// ShippingArrivalRule defines a time window within the week that maps to
// a specific warehouse arrival day. Rules replace the hard-coded
// Monday-14:00 cut-off logic and are evaluated in sort_order priority.
//
// Days of week: 0=Sunday … 6=Saturday.
// Times are in "HH:MM" 24-hour format (Melbourne local time).
type ShippingArrivalRule struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	FromDOW    int              `json:"from_dow"` // 0-6
	FromTime   common.TimeOfDay `json:"from_time"`
	ToDOW      int              `json:"to_dow"` // 0-6
	ToTime     common.TimeOfDay `json:"to_time"`
	ArrivalDOW int              `json:"arrival_dow"` // 0-6
	WeekOffset int              `json:"week_offset"` // 0 = this week, 1 = next week
	IsActive   bool             `json:"is_active"`
	SortOrder  int              `json:"sort_order"`

	common.AuditFields `bson:",inline"`
}

// ShippingArrivalBlacklist lists calendar dates on which warehouse arrivals
// are not permitted (public holidays, depot closures, etc.).
// When a computed arrival date falls on a blocked date it rolls forward
// one day at a time until a non-blocked date is found.
type ShippingArrivalBlacklist struct {
	ID          string      `json:"id"`
	BlockedDate common.Date `json:"blocked_date"`
	Reason      string      `json:"reason,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
}
