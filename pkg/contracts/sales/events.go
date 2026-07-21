package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/payment"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/sales"
)

// OrderCreatedEvent is emitted on the order-events topic when a sales order
// is first persisted. AggregateID is the order number.
type OrderCreatedEvent struct {
	OrderID                   string           `json:"order_id"`
	OrderNumber               string           `json:"order_number"`
	OrderType                 salesenum.OrderType `json:"order_type,omitempty"`
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
	OrderID     string                      `json:"order_id"`
	OrderNumber string                      `json:"order_number"`
	PaymentID   string                      `json:"payment_id,omitempty"`
	Method      paymentenum.PaymentMethod   `json:"method,omitempty"`
	AmountPaid  common.Money                `json:"amount_paid"`
	PaidAt      time.Time                   `json:"paid_at"`
	RequestID   string                      `json:"request_id,omitempty"`
}

// OrderStatusChangedEvent is emitted on the order-events topic for every
// order lifecycle transition (processing, picking, packed, shipped, …).
type OrderStatusChangedEvent struct {
	OrderID        string                     `json:"order_id"`
	OrderNumber    string                     `json:"order_number"`
	PreviousStatus salesenum.SalesOrderStatus `json:"previous_status,omitempty"`
	Status         salesenum.SalesOrderStatus `json:"status"`
	ChangedBy      string                     `json:"changed_by,omitempty"`
	ChangedAt      time.Time                  `json:"changed_at"`
	Reason         string                     `json:"reason,omitempty"`
	RequestID      string                     `json:"request_id,omitempty"`
}

// OrderCancelledEvent is emitted on the order-events topic when an order is
// cancelled before fulfilment completes.
type OrderCancelledEvent struct {
	OrderID     string    `json:"order_id"`
	OrderNumber string    `json:"order_number"`
	CancelledBy string    `json:"cancelled_by,omitempty"`
	CancelledAt time.Time `json:"cancelled_at"`
	Reason      string    `json:"reason,omitempty"`
	RequestID   string    `json:"request_id,omitempty"`
}
