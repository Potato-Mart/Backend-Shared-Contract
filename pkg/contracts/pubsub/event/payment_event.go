package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/payments/payment/payment_enums"
)

// PaymentCapturedEvent is emitted on the payment-events topic when a payment
// is captured (online provider webhook, terminal settlement or manual
// record). AggregateID is the order number.
type PaymentCapturedEvent struct {
	PaymentID         string                      `json:"payment_id"`
	OrderID           string                      `json:"order_id,omitempty"`
	OrderNumber       string                      `json:"order_number"`
	Method            payment_enums.PaymentMethod `json:"method,omitempty"`
	Amount            money.Money                 `json:"amount"`
	ProviderSessionID string                      `json:"provider_session_id,omitempty"`
	CapturedAt        time.Time                   `json:"captured_at"`
	RequestID         string                      `json:"request_id,omitempty"`
}

// PaymentFailedEvent is emitted on the payment-events topic when a payment
// attempt terminally fails (provider decline, validation failure, expiry).
type PaymentFailedEvent struct {
	PaymentID            string                      `json:"payment_id,omitempty"`
	OrderID              string                      `json:"order_id,omitempty"`
	OrderNumber          string                      `json:"order_number"`
	Method               payment_enums.PaymentMethod `json:"method,omitempty"`
	Amount               money.Money                 `json:"amount,omitempty"`
	ProviderSessionID    string                      `json:"provider_session_id,omitempty"`
	RetailCustomerNumber string                      `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                      `json:"organisation_access_id,omitempty"`
	Reason               string                      `json:"reason,omitempty"`
	FailedAt             time.Time                   `json:"failed_at"`
	RequestID            string                      `json:"request_id,omitempty"`
}

// ReceiptGeneratedEvent is emitted on the payment-events topic when a POS
// receipt snapshot is written or its revision advances. AggregateID carries
// the order number.
type ReceiptGeneratedEvent struct {
	OrderNumber          string                     `json:"order_number"`
	RetailCustomerNumber string                     `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                     `json:"organisation_access_id,omitempty"`
	DocumentKind         payment_enums.DocumentKind `json:"document_kind"`
	Revision             int64                      `json:"revision"`
	IssuedAt             time.Time                  `json:"issued_at"`
	RequestID            string                     `json:"request_id,omitempty"`
}

// InvoiceIssuedEvent is emitted on the payment-events topic when an invoice
// is issued for an order. AggregateID is the order number.
type InvoiceIssuedEvent struct {
	InvoiceNumber        string    `json:"invoice_number"`
	OrderNumber          string    `json:"order_number"`
	RetailCustomerNumber string    `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string    `json:"organisation_access_id,omitempty"`
	IssuedAt             time.Time `json:"issued_at"`
	RequestID            string    `json:"request_id,omitempty"`
}

// InvoiceDeliveryRequestedEvent requests an idempotent redelivery of an
// already-issued invoice through the customer-notification consumer.
type InvoiceDeliveryRequestedEvent struct {
	DeliveryID           string    `json:"delivery_id"`
	InvoiceNumber        string    `json:"invoice_number"`
	OrderNumber          string    `json:"order_number"`
	RetailCustomerNumber string    `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string    `json:"organisation_access_id,omitempty"`
	RequestedAt          time.Time `json:"requested_at"`
	RequestID            string    `json:"request_id,omitempty"`
}
