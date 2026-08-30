package shipping

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"time"
)

// DeliverySlot is one customer-selectable delivery window. ID is opaque and
// stable for the area, date, window, and schedule revision that produced it.
// Availability is a service-owned wire value such as available, limited, or
// full.
type DeliverySlot struct {
	ID           string       `json:"id"`
	StartAt      time.Time    `json:"start_at"`
	EndAt        time.Time    `json:"end_at"`
	Label        string       `json:"label,omitempty"`
	Availability string       `json:"availability"`
	Fee          *money.Money `json:"fee,omitempty"`
}
