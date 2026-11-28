package orders

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
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
	// Tenders lists, for a POS split-tender order, every captured part of
	// its tender plan in collection (sequence) order. Method stays the
	// order's primary external tender (Order.PaymentMethod). Gift-card and
	// points redemptions are not listed. Omitted when the order was not
	// split.
	Tenders []OrderPaidTender `json:"tenders,omitempty"`

	// AdAttribution is the order's read-only paid-advertising origin, copied
	// from the order. It is absent for an order that did not come from an ad,
	// which is not the same as an unknown origin. It carries no personal data.
	AdAttribution *order.AdAttribution `json:"ad_attribution,omitempty"`
	// Revenue is the attributed revenue of this order: the merchandise value
	// net of tax and net of discount, excluding shipping. It is deliberately
	// not AmountPaid, which is settlement evidence, and not Subtotal, which is
	// before discount. The same figure is the purchase value reported back to
	// an advertising platform, so every consumer uses one number. Absence is
	// missing evidence, never a zero revenue: a consumer that needs it must
	// skip the order rather than infer it from the other amounts.
	Revenue *money.Money `json:"revenue,omitempty"`
	// IsFirstPaidOrder is true when this is the customer's first paid order,
	// which is what makes it a new customer for advertising reporting. False
	// and absent are not distinguishable on the wire, so a consumer that
	// counts new customers must treat absence as "not established" rather
	// than as a returning customer.
	IsFirstPaidOrder bool `json:"is_first_paid_order,omitempty"`
	// SourcePlatform is the client the order was placed from: the website or
	// one of the apps. It decides which device rules a server-side conversion
	// must honour.
	SourcePlatform ads_enums.AdSourcePlatform `json:"source_platform,omitempty"`
	// HashedEmail, HashedPhone and HashedExternalID are normalised SHA-256
	// digests, and nothing else. A raw email address, phone number, name or
	// address must never appear in this event, and a consumer that finds one
	// must reject the event rather than forward it. The digests are present
	// only when the customer granted ad_user_data consent, and only for a
	// source platform whose device rules permit it; absence is a withheld
	// consent, so a consumer must send no contact data at all for that order
	// rather than falling back to another identifier. They exist for
	// server-side advertising conversions and must never be copied into a
	// read API, a report or a durable projection.
	HashedEmail      string `json:"hashed_email,omitempty"`
	HashedPhone      string `json:"hashed_phone,omitempty"`
	HashedExternalID string `json:"hashed_external_id,omitempty"`
}
