package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"time"
)

// InvoiceIssuedEvent is emitted on the payment-events topic when an invoice
// is issued for an order. AggregateID is the order number.
type InvoiceIssuedEvent struct {
	InvoiceNumber        string `json:"invoice_number"`
	OrderNumber          string `json:"order_number"`
	RetailCustomerNumber string `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string `json:"organisation_access_id,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	IssuedAt    time.Time             `json:"issued_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
