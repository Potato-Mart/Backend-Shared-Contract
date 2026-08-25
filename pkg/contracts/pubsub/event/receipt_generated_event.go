package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// ReceiptGeneratedEvent is emitted on the payment-events topic when a POS
// receipt snapshot is written or its revision advances. AggregateID carries
// the order number.
type ReceiptGeneratedEvent struct {
	OrderNumber          string                     `json:"order_number"`
	RetailCustomerNumber string                     `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                     `json:"organisation_access_id,omitempty"`
	DocumentKind         payment_enums.DocumentKind `json:"document_kind"`
	// MarketCode, DepotCode, and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	Revision    int64                 `json:"revision"`
	IssuedAt    time.Time             `json:"issued_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
