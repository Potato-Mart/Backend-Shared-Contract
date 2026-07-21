package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/payment"
)

// PaymentCapturedEvent is emitted on the payment-events topic when a payment
// is captured (online provider webhook, terminal settlement or manual
// record). AggregateID is the order number.
type PaymentCapturedEvent struct {
	PaymentID         string                    `json:"payment_id"`
	OrderID           string                    `json:"order_id,omitempty"`
	OrderNumber       string                    `json:"order_number"`
	Method            paymentenum.PaymentMethod `json:"method,omitempty"`
	Amount            common.Money              `json:"amount"`
	ProviderSessionID string                    `json:"provider_session_id,omitempty"`
	CapturedAt        time.Time                 `json:"captured_at"`
	RequestID         string                    `json:"request_id,omitempty"`
}

// PaymentFailedEvent is emitted on the payment-events topic when a payment
// attempt terminally fails (provider decline, validation failure, expiry).
type PaymentFailedEvent struct {
	PaymentID         string                    `json:"payment_id,omitempty"`
	OrderID           string                    `json:"order_id,omitempty"`
	OrderNumber       string                    `json:"order_number"`
	Method            paymentenum.PaymentMethod `json:"method,omitempty"`
	Amount            common.Money              `json:"amount,omitempty"`
	ProviderSessionID string                    `json:"provider_session_id,omitempty"`
	Reason            string                    `json:"reason,omitempty"`
	FailedAt          time.Time                 `json:"failed_at"`
	RequestID         string                    `json:"request_id,omitempty"`
}

// RefundRequestedEvent is emitted on the refund-events topic when a refund is
// raised against an order's captured payments.
type RefundRequestedEvent struct {
	RefundID    string       `json:"refund_id"`
	OrderID     string       `json:"order_id,omitempty"`
	OrderNumber string       `json:"order_number"`
	Amount      common.Money `json:"amount"`
	RequestedBy string       `json:"requested_by,omitempty"`
	RequestedAt time.Time    `json:"requested_at"`
	Reason      string       `json:"reason,omitempty"`
	RequestID   string       `json:"request_id,omitempty"`
}

// RefundCompletedEvent is emitted on the refund-events topic when a refund
// settles. Consumers reverse benefits/points and update order payment state.
// The optional restoration fields carry everything the pricing owner needs to
// reverse checkout benefits and points without a synchronous read-back.
type RefundCompletedEvent struct {
	RefundID             string        `json:"refund_id"`
	OrderID              string        `json:"order_id,omitempty"`
	OrderNumber          string        `json:"order_number"`
	PaymentID            string        `json:"payment_id,omitempty"`
	Amount               common.Money  `json:"amount"`
	BenefitReservationID string        `json:"benefit_reservation_id,omitempty"`
	GiftCardRefundAmount *common.Money `json:"gift_card_refund_amount,omitempty"`
	PointsToRestore      int           `json:"points_to_restore,omitempty"`
	MembershipAccountID  string        `json:"membership_account_id,omitempty"`
	FullOrderRefund      bool          `json:"full_order_refund,omitempty"`
	CompletedAt          time.Time     `json:"completed_at"`
	RequestID            string        `json:"request_id,omitempty"`
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

// RefundFailedEvent is emitted on the refund-events topic when a refund
// attempt terminally fails.
type RefundFailedEvent struct {
	RefundID    string       `json:"refund_id"`
	OrderID     string       `json:"order_id,omitempty"`
	OrderNumber string       `json:"order_number"`
	Amount      common.Money `json:"amount,omitempty"`
	Reason      string       `json:"reason,omitempty"`
	FailedAt    time.Time    `json:"failed_at"`
	RequestID   string       `json:"request_id,omitempty"`
}
