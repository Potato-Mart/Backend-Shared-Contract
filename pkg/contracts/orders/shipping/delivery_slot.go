package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/supply/courier/courier_enums"
)

// DeliverySlot is one customer-selectable delivery window. ID is opaque and
// stable for the area, date, window, and schedule revision that produced it.
// Availability is a service-owned wire value such as available, limited, or
// full. Source distinguishes configured windows from provider offers. ExpiresAt
// is offer expiry, not a deadline for fulfilling an already accepted selection.
type DeliverySlot struct {
	ID                string                           `json:"id"`
	StartAt           time.Time                        `json:"start_at"`
	EndAt             time.Time                        `json:"end_at"`
	Label             string                           `json:"label,omitempty"`
	Availability      string                           `json:"availability"`
	Fee               *money.Money                     `json:"fee,omitempty"`
	DeliveryCompany   *courier.DeliveryCompanyRef      `json:"delivery_company,omitempty"`
	Source            courier_enums.DeliverySlotSource `json:"source,omitempty"`
	ScheduleCode      string                           `json:"schedule_code,omitempty"`
	UnavailableReason string                           `json:"unavailable_reason,omitempty"`
	ExpiresAt         *time.Time                       `json:"expires_at,omitempty"`
}
