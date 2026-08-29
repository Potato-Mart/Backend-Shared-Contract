package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"

// FulfilmentCompletedEvent is emitted on the fulfilment-events topic when
// fulfilment finishes end-to-end for an order.
type FulfilmentCompletedEvent struct {
	OrderNumber string `json:"order_number"`
	Note        string `json:"note,omitempty"`
	// MarketCode, DepotCode, and CountryCode are the denormalized geography
	// the event belongs to. Empty values provide no geographic evidence; a
	// consumer that persists a geographically scoped record must fail closed
	// rather than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}
