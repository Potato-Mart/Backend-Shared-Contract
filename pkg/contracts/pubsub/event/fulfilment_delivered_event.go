package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
)

// FulfilmentDeliveredEvent is emitted on the fulfilment-events topic when a
// shipment is confirmed delivered.
type FulfilmentDeliveredEvent struct {
	OrderNumber string `json:"order_number"`
	ShipmentID  string `json:"shipment_id,omitempty"`
	Note        string `json:"note,omitempty"`
	// MarketCode, DepotCode, and CountryCode are the denormalized geography
	// the event belongs to. They are absent on every event published before
	// v28.0.0; a consumer that persists a geographically scoped record
	// treats an absent value as "no evidence" and fails closed rather than
	// defaulting it.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
