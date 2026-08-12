package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/insights/analytics"

	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/commerce/commerce_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/payment/payment_enums"
)

// OrderCreatedEvent is emitted on the order-events topic when a sales order
// is first recorded. AggregateID is the order number.
type OrderCreatedEvent struct {
	OrderID                   string                   `json:"order_id"`
	OrderNumber               string                   `json:"order_number"`
	OrderType                 commerce_enums.OrderType `json:"order_type,omitempty"`
	BuyerUserID               string                   `json:"buyer_user_id,omitempty"`
	RetailCustomerNumber      string                   `json:"retail_customer_number,omitempty"`
	WholesaleOrganisationCode string                   `json:"wholesale_organisation_code,omitempty"`
	Total                     money.Money              `json:"total"`
	CreatedAt                 time.Time                `json:"created_at"`
	RequestID                 string                   `json:"request_id,omitempty"`
}

// OrderPaidEvent is emitted on the order-events topic when an order reaches
// the paid state.
type OrderPaidEvent struct {
	OrderID              string                      `json:"order_id"`
	OrderNumber          string                      `json:"order_number"`
	PaymentID            string                      `json:"payment_id,omitempty"`
	Method               payment_enums.PaymentMethod `json:"method,omitempty"`
	Channel              commerce_enums.OrderType    `json:"channel,omitempty"`
	AmountPaid           money.Money                 `json:"amount_paid"`
	Items                []analytics.OrderItemFact   `json:"items,omitempty"`
	RetailCustomerNumber string                      `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                      `json:"organisation_access_id,omitempty"`
	PaidAt               time.Time                   `json:"paid_at"`
	RequestID            string                      `json:"request_id,omitempty"`
}

// CheckoutCompensationRequestedEvent restores Pricing-owned value after a
// checkout saga has committed points or benefits but cannot complete. It is
// replay-safe by CompensationID and carries no customer PII.
type CheckoutCompensationRequestedEvent struct {
	CompensationID       string      `json:"compensation_id"`
	OrderNumber          string      `json:"order_number"`
	RetailCustomerNumber string      `json:"retail_customer_number"`
	BenefitReservationID string      `json:"benefit_reservation_id,omitempty"`
	GiftCardAmount       money.Money `json:"gift_card_amount"`
	PointsToRestore      int         `json:"points_to_restore,omitempty"`
	FullOrder            bool        `json:"full_order"`
	RequestedAt          time.Time   `json:"requested_at"`
	RequestID            string      `json:"request_id,omitempty"`
}

// OrderStatusChangedEvent is emitted on the order-events topic for every
// order lifecycle transition (processing, picking, packed, shipped, …).
// Buyer identity, tracking, and invoice references are optional enrichment
// so notification consumers need no synchronous read-back.
type OrderStatusChangedEvent struct {
	OrderID              string                       `json:"order_id"`
	OrderNumber          string                       `json:"order_number"`
	PreviousStatus       order_enums.SalesOrderStatus `json:"previous_status,omitempty"`
	Status               order_enums.SalesOrderStatus `json:"status"`
	RetailCustomerNumber string                       `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                       `json:"organisation_access_id,omitempty"`
	TrackingNumber       string                       `json:"tracking_number,omitempty"`
	InvoiceNumber        string                       `json:"invoice_number,omitempty"`
	ChangedBy            string                       `json:"changed_by,omitempty"`
	ChangedAt            time.Time                    `json:"changed_at"`
	Reason               string                       `json:"reason,omitempty"`
	RequestID            string                       `json:"request_id,omitempty"`
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
