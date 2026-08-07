package event

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/insights/analytics"
	order "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/orders/order"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/payment"
	"time"
)

// OrderCreatedEvent is emitted on the order-events topic when a sales order
// is first recorded. AggregateID is the order number.
type OrderCreatedEvent struct {
	OrderID                   string           `json:"order_id"`
	OrderNumber               string           `json:"order_number"`
	OrderType                 common.OrderType `json:"order_type,omitempty"`
	BuyerUserID               string           `json:"buyer_user_id,omitempty"`
	RetailCustomerNumber      string           `json:"retail_customer_number,omitempty"`
	WholesaleOrganisationCode string           `json:"wholesale_organisation_code,omitempty"`
	Total                     common.Money     `json:"total"`
	CreatedAt                 time.Time        `json:"created_at"`
	RequestID                 string           `json:"request_id,omitempty"`
}

// OrderPaidEvent is emitted on the order-events topic when an order reaches
// the paid state.
type OrderPaidEvent struct {
	OrderID              string                    `json:"order_id"`
	OrderNumber          string                    `json:"order_number"`
	PaymentID            string                    `json:"payment_id,omitempty"`
	Method               paymentenum.PaymentMethod `json:"method,omitempty"`
	Channel              common.OrderType          `json:"channel,omitempty"`
	AmountPaid           common.Money              `json:"amount_paid"`
	Items                []analytics.OrderItemFact `json:"items,omitempty"`
	RetailCustomerNumber string                    `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                    `json:"organisation_access_id,omitempty"`
	PaidAt               time.Time                 `json:"paid_at"`
	RequestID            string                    `json:"request_id,omitempty"`
}

// CheckoutCompensationRequestedEvent restores Pricing-owned value after a
// checkout saga has committed points or benefits but cannot complete. It is
// replay-safe by CompensationID and carries no customer PII.
type CheckoutCompensationRequestedEvent struct {
	CompensationID       string       `json:"compensation_id"`
	OrderNumber          string       `json:"order_number"`
	RetailCustomerNumber string       `json:"retail_customer_number"`
	BenefitReservationID string       `json:"benefit_reservation_id,omitempty"`
	GiftCardAmount       common.Money `json:"gift_card_amount"`
	PointsToRestore      int          `json:"points_to_restore,omitempty"`
	FullOrder            bool         `json:"full_order"`
	RequestedAt          time.Time    `json:"requested_at"`
	RequestID            string       `json:"request_id,omitempty"`
}

// OrderStatusChangedEvent is emitted on the order-events topic for every
// order lifecycle transition (processing, picking, packed, shipped, …).
// Buyer identity, tracking, and invoice references are optional enrichment
// so notification consumers need no synchronous read-back.
type OrderStatusChangedEvent struct {
	OrderID              string                 `json:"order_id"`
	OrderNumber          string                 `json:"order_number"`
	PreviousStatus       order.SalesOrderStatus `json:"previous_status,omitempty"`
	Status               order.SalesOrderStatus `json:"status"`
	RetailCustomerNumber string                 `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                 `json:"organisation_access_id,omitempty"`
	TrackingNumber       string                 `json:"tracking_number,omitempty"`
	InvoiceNumber        string                 `json:"invoice_number,omitempty"`
	ChangedBy            string                 `json:"changed_by,omitempty"`
	ChangedAt            time.Time              `json:"changed_at"`
	Reason               string                 `json:"reason,omitempty"`
	RequestID            string                 `json:"request_id,omitempty"`
}

// OrderCancelledEvent is emitted on the order-events topic when an order is
// cancelled before fulfilment completes.
type OrderCancelledEvent struct {
	OrderID              string    `json:"order_id"`
	OrderNumber          string    `json:"order_number"`
	RetailCustomerNumber string    `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string    `json:"organisation_access_id,omitempty"`
	CancelledBy          string    `json:"cancelled_by,omitempty"`
	CancelledAt          time.Time `json:"cancelled_at"`
	Reason               string    `json:"reason,omitempty"`
	RequestID            string    `json:"request_id,omitempty"`
}
