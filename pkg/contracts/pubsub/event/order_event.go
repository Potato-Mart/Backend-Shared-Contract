package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/insights/analytics"

	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/commerce/commerce_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/payments/payment/payment_enums"
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
	// MarketID, DepotCode, and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	Total       money.Money           `json:"total"`
	CreatedAt   time.Time             `json:"created_at"`
	RequestID   string                `json:"request_id,omitempty"`
}

// OrderPaidEvent is emitted on the order-events topic when an order reaches
// the paid state.
//
// Subtotal, DiscountAmount, and Tags are qualification evidence added in
// v27.3.0. Every event published before that release carries none of them, so a
// consumer that qualifies on them must fail closed: an empty Subtotal or
// DiscountAmount currency and a nil Tags slice mean "no evidence", never a zero
// subtotal, a zero discount, or an untagged order. Consumers must skip
// qualification for such an event rather than infer a value from AmountPaid.
type OrderPaidEvent struct {
	OrderID     string                      `json:"order_id"`
	OrderNumber string                      `json:"order_number"`
	PaymentID   string                      `json:"payment_id,omitempty"`
	Method      payment_enums.PaymentMethod `json:"method,omitempty"`
	Channel     commerce_enums.OrderType    `json:"channel,omitempty"`
	AmountPaid  money.Money                 `json:"amount_paid"`

	// Subtotal is the merchandise subtotal before any discount is applied.
	Subtotal money.Money `json:"subtotal"`
	// DiscountAmount is the order-level discount applied to that subtotal.
	DiscountAmount money.Money `json:"discount_amount"`
	// Tags are the order tags carried at the time of payment.
	Tags []string `json:"tags,omitempty"`

	Items                []analytics.OrderItemFact `json:"items,omitempty"`
	RetailCustomerNumber string                    `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                    `json:"organisation_access_id,omitempty"`
	// MarketID, DepotCode, and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	PaidAt      time.Time             `json:"paid_at"`
	RequestID   string                `json:"request_id,omitempty"`
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
	// MarketID and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	ChangedAt   time.Time             `json:"changed_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}

// OrderCancelledEvent is emitted on the order-events topic when an order is
// cancelled before fulfilment completes.
type OrderCancelledEvent struct {
	OrderID              string `json:"order_id"`
	OrderNumber          string `json:"order_number"`
	RetailCustomerNumber string `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string `json:"organisation_access_id,omitempty"`
	CancelledBy          string `json:"cancelled_by,omitempty"`
	// MarketID and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CancelledAt time.Time             `json:"cancelled_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}
