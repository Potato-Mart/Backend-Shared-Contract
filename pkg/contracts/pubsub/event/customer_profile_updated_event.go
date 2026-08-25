package event

import (
	"time"
)

// CustomerProfileUpdatedEvent is emitted on the customer-events topic when a
// retail customer profile changes. UpdatedFields carries field names only,
// never values.
type CustomerProfileUpdatedEvent struct {
	CustomerNumber string    `json:"customer_number"`
	UserID         string    `json:"user_id,omitempty"`
	UpdatedFields  []string  `json:"updated_fields,omitempty"`
	UpdatedBy      string    `json:"updated_by,omitempty"`
	UpdatedAt      time.Time `json:"updated_at"`
	RequestID      string    `json:"request_id,omitempty"`
}
