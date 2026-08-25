package event

import (
	"time"
)

// FulfilmentTrackingEvent is emitted on the fulfilment-events topic when
// carrier tracking details are assigned or corrected.
type FulfilmentTrackingEvent struct {
	OrderNumber    string    `json:"order_number"`
	TrackingNumber string    `json:"tracking_number"`
	TrackingURL    string    `json:"tracking_url,omitempty"`
	OccurredAt     time.Time `json:"occurred_at"`
	RequestID      string    `json:"request_id,omitempty"`
}
