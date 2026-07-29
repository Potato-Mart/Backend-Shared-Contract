package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/analytics"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/payment"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/sales"
)

// OrderCreatedEvent is emitted on the order-events topic when a sales order
// is first persisted. AggregateID is the order number.
type OrderCreatedEvent struct {
	OrderID                   string              `json:"order_id"`
	OrderNumber               string              `json:"order_number"`
	OrderType                 salesenum.OrderType `json:"order_type,omitempty"`
	BuyerUserID               string              `json:"buyer_user_id,omitempty"`
	RetailCustomerNumber      string              `json:"retail_customer_number,omitempty"`
	WholesaleOrganisationCode string              `json:"wholesale_organisation_code,omitempty"`
	Total                     common.Money        `json:"total"`
	CreatedAt                 time.Time           `json:"created_at"`
	RequestID                 string              `json:"request_id,omitempty"`
}

// OrderPaidEvent is emitted on the order-events topic when an order reaches
// the paid state.
type OrderPaidEvent struct {
	OrderID              string                    `json:"order_id"`
	OrderNumber          string                    `json:"order_number"`
	PaymentID            string                    `json:"payment_id,omitempty"`
	Method               paymentenum.PaymentMethod `json:"method,omitempty"`
	Channel              salesenum.OrderType       `json:"channel,omitempty"`
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
	OrderID              string                     `json:"order_id"`
	OrderNumber          string                     `json:"order_number"`
	PreviousStatus       salesenum.SalesOrderStatus `json:"previous_status,omitempty"`
	Status               salesenum.SalesOrderStatus `json:"status"`
	RetailCustomerNumber string                     `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                     `json:"organisation_access_id,omitempty"`
	TrackingNumber       string                     `json:"tracking_number,omitempty"`
	InvoiceNumber        string                     `json:"invoice_number,omitempty"`
	ChangedBy            string                     `json:"changed_by,omitempty"`
	ChangedAt            time.Time                  `json:"changed_at"`
	Reason               string                     `json:"reason,omitempty"`
	RequestID            string                     `json:"request_id,omitempty"`
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

// PreorderAvailabilityEvent is emitted on the order-events topic when a
// preorder line becomes available to fulfil. AggregateID is the order number.
type PreorderAvailabilityEvent struct {
	OrderNumber          string     `json:"order_number"`
	RetailCustomerNumber string     `json:"retail_customer_number,omitempty"`
	ProductSKUCode       string     `json:"product_sku_code"`
	ProductName          string     `json:"product_name,omitempty"`
	Quantity             int64      `json:"quantity"`
	AvailableAt          time.Time  `json:"available_at"`
	ExpectedAvailableAt  *time.Time `json:"expected_available_at,omitempty"`
	RequestID            string     `json:"request_id,omitempty"`
}

// FulfilmentShippedEvent is emitted on the fulfilment-events topic when a
// shipment leaves the warehouse. AggregateID is the order number.
type FulfilmentShippedEvent struct {
	OrderNumber    string    `json:"order_number"`
	ShipmentID     string    `json:"shipment_id,omitempty"`
	Carrier        string    `json:"carrier,omitempty"`
	TrackingNumber string    `json:"tracking_number,omitempty"`
	TrackingURL    string    `json:"tracking_url,omitempty"`
	Note           string    `json:"note,omitempty"`
	OccurredAt     time.Time `json:"occurred_at"`
	RequestID      string    `json:"request_id,omitempty"`
}

// FulfilmentDeliveredEvent is emitted on the fulfilment-events topic when a
// shipment is confirmed delivered.
type FulfilmentDeliveredEvent struct {
	OrderNumber string    `json:"order_number"`
	ShipmentID  string    `json:"shipment_id,omitempty"`
	Note        string    `json:"note,omitempty"`
	OccurredAt  time.Time `json:"occurred_at"`
	RequestID   string    `json:"request_id,omitempty"`
}

// FulfilmentCompletedEvent is emitted on the fulfilment-events topic when
// fulfilment finishes end-to-end for an order.
type FulfilmentCompletedEvent struct {
	OrderNumber string    `json:"order_number"`
	Note        string    `json:"note,omitempty"`
	OccurredAt  time.Time `json:"occurred_at"`
	RequestID   string    `json:"request_id,omitempty"`
}

// FulfilmentTrackingEvent is emitted on the fulfilment-events topic when
// carrier tracking details are assigned or corrected.
type FulfilmentTrackingEvent struct {
	OrderNumber    string    `json:"order_number"`
	TrackingNumber string    `json:"tracking_number"`
	TrackingURL    string    `json:"tracking_url,omitempty"`
	OccurredAt     time.Time `json:"occurred_at"`
	RequestID      string    `json:"request_id,omitempty"`
}
