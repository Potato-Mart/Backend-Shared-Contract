package orders

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/insights/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// OrderPaidEvent is emitted on the order-events topic when an order reaches
// the paid state.
//
// Subtotal, DiscountAmount, and Tags are qualification evidence. An empty
// Subtotal or DiscountAmount currency and a nil Tags slice mean "no evidence",
// never a zero subtotal, a zero discount, or an untagged order. Without a
// verified frozen quote, consumers must skip qualification for such an event
// rather than infer a value from AmountPaid. A verified quote supplies its own
// authoritative qualification evidence.
type OrderPaidEvent struct {
	DeferredPayment *order.DeferredPaymentAuthorization `json:"deferred_payment,omitempty"`
	OrderID         string                              `json:"order_id"`
	OrderNumber     string                              `json:"order_number"`
	PaymentID       string                              `json:"payment_id,omitempty"`
	Method          payment_enums.PaymentMethod         `json:"method,omitempty"`
	Channel         commerce_enums.OrderType            `json:"channel,omitempty"`
	// QuoteKey and QuoteFingerprint identify Pricing's accepted frozen checkout
	// quote. QuoteFingerprint carries its opaque revision exactly, without parsing,
	// hashing again or numeric conversion. Both are absent for legacy events.
	// Pricing verifies the pair against protected accepted-quote evidence; absence
	// or mismatch provides no opaque quote authority. Numeric QuoteRevision is
	// retained for released compatibility and omitted when only opaque evidence
	// exists. It is not a substitute for QuoteFingerprint; producers must never
	// fabricate it from the fingerprint. If both are present, the owning service
	// must verify they refer to the same quote, not silently choose conflicting
	// evidence. AmountPaid is settlement evidence, never qualifying goods value.
	QuoteKey         string      `json:"quote_key,omitempty"`
	QuoteFingerprint string      `json:"quote_fingerprint,omitempty"`
	QuoteRevision    int64       `json:"quote_revision,omitempty"`
	AmountPaid       money.Money `json:"amount_paid"`

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
