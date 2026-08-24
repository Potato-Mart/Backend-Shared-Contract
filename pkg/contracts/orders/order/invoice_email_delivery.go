package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/payments/payment/payment_enums"
)

// InvoiceEmailDelivery records one customer-requested invoice email
// redelivery and its lifecycle state.
type InvoiceEmailDelivery struct {
	InvoiceNumber string                            `json:"invoice_number"`
	DeliveryID    string                            `json:"delivery_id"`
	Status        payment_enums.InvoiceResendStatus `json:"status"`
	RecipientHint string                            `json:"recipient_hint,omitempty"`
	RequestedAt   time.Time                         `json:"requested_at"`
}
