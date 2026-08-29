package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/analytics"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// OrderPaidEvent is emitted on the order-events topic when an order reaches
// the paid state.
//
// Subtotal, DiscountAmount, and Tags are qualification evidence. An empty
// Subtotal or DiscountAmount currency and a nil Tags slice mean "no evidence",
// never a zero subtotal, a zero discount, or an untagged order. Consumers must
// skip qualification for such an event rather than infer a value from AmountPaid.
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
	// MarketCode, DepotCode, and CountryCode are the denormalized geography
	// the event belongs to. Empty values provide no geographic evidence; a
	// consumer that persists a geographically scoped record must fail closed
	// rather than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	PaidAt      time.Time             `json:"paid_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
