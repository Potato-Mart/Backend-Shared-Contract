package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
)

// FulfilmentDeliveredEvent is emitted on the fulfilment-events topic when a
// shipment is confirmed delivered.
type FulfilmentDeliveredEvent struct {
	OrderNumber string `json:"order_number"`
	ShipmentID  string `json:"shipment_id,omitempty"`
	Note        string `json:"note,omitempty"`
	// MarketCode, DepotCode, and CountryCode are the denormalized geography
	// the event belongs to. Empty values provide no geographic evidence; a
	// consumer that persists a geographically scoped record must fail closed
	// rather than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
