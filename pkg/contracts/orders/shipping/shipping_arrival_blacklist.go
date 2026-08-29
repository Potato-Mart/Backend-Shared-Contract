package shipping

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/temporal"
)

// ShippingArrivalBlacklist represents a local calendar date on which warehouse
// arrivals are unavailable.
type ShippingArrivalBlacklist struct {
	ID          string        `json:"id"`
	BlockedDate temporal.Date `json:"blocked_date"`
	Timezone    string        `json:"timezone"`
	Reason      string        `json:"reason,omitempty"`

	audit.AuditFields
}
