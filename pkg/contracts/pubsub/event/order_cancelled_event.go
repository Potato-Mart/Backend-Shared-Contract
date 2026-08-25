package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"time"
)

// OrderCancelledEvent is emitted on the order-events topic when an order is
// cancelled before fulfilment completes.
type OrderCancelledEvent struct {
	OrderID              string `json:"order_id"`
	OrderNumber          string `json:"order_number"`
	RetailCustomerNumber string `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string `json:"organisation_access_id,omitempty"`
	CancelledBy          string `json:"cancelled_by,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CancelledAt time.Time             `json:"cancelled_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}
