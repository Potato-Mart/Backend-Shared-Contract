package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// DeliverySlot is one customer-selectable delivery window. ID is opaque and
// stable for the area, date, window, and schedule revision that produced it.
// Availability is a service-owned wire value such as available, limited, or
// full. Source distinguishes configured windows from provider offers. ExpiresAt
// is offer expiry, not a deadline for fulfilling an already accepted selection.
// TimeBand classifies the window by local time of day in the slot's timezone.
// RequiresTimeAcknowledgement is true when the customer is asked to confirm a
// late-night or early-morning window; Orders decides whether a missing
// acknowledgement blocks acceptance.
type DeliverySlot struct {
	ID                          string                           `json:"id"`
	StartAt                     time.Time                        `json:"start_at"`
	EndAt                       time.Time                        `json:"end_at"`
	Label                       string                           `json:"label,omitempty"`
	Availability                string                           `json:"availability"`
	Fee                         *money.Money                     `json:"fee,omitempty"`
	DeliveryCompany             *courier.DeliveryCompanyRef      `json:"delivery_company,omitempty"`
	Source                      courier_enums.DeliverySlotSource `json:"source,omitempty"`
	ScheduleCode                string                           `json:"schedule_code,omitempty"`
	UnavailableReason           string                           `json:"unavailable_reason,omitempty"`
	ExpiresAt                   *time.Time                       `json:"expires_at,omitempty"`
	TimeBand                    shipping_enums.DeliveryTimeBand  `json:"time_band,omitempty"`
	RequiresTimeAcknowledgement bool                             `json:"requires_time_acknowledgement,omitempty"`
}
