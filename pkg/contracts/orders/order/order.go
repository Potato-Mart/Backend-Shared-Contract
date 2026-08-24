package order

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/shipping"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/promotion"
)

// Buyer describes who is buying, independently of Channel. POS is a
// channel, not a buyer type — see sales.BuyerContext. Optional pointer
// so it is omitted entirely when unset.
type Order struct {
	ID          string `json:"id"`
	OrderNumber string `json:"order_number"`
	// MarketCode is the immutable commercial market the order was placed
	// in. It is captured when the cart is created and never re-derived
	// from country or currency.
	MarketCode string `json:"market_code"`
	// CountryCode is the denormalized country of MarketCode, carried so a
	// country-scoped staff query is a plain indexed match.
	CountryCode       geography.CountryCode         `json:"country_code,omitempty"`
	Channel           commerce_enums.OrderType      `json:"channel"`
	Status            order_enums.SalesOrderStatus  `json:"status"`
	PaymentStatus     payment_enums.PaymentStatus   `json:"payment_status"`
	PaymentMethod     payment_enums.PaymentMethod   `json:"payment_method"`
	FulfillmentStatus order_enums.FulfillmentStatus `json:"fulfillment_status"`
	Customer          party.PartyRef                `json:"customer"`

	Buyer                *BuyerContext                       `json:"buyer,omitempty"`
	FulfilmentLocation   shipping.FulfilmentLocationSnapshot `json:"fulfilment_location"`
	GroupOrder           *GroupOrderContext                  `json:"group_order,omitempty"`
	GroupOrderFulfilment *GroupOrderFulfilmentPlan           `json:"group_order_fulfilment,omitempty"`
	Items                []OrderItem                         `json:"items"`
	SourceDevice         SourceDevice                        `json:"source_device,omitempty"`

	// ── Shipping & billing ────────────────────────────────────────────
	Billing          *party.ContactAddress           `json:"billing,omitempty"`
	ShippingMethod   shipping_enums.ShippingRateName `json:"shipping_method,omitempty"`
	ShippingZoneID   string                          `json:"shipping_zone_id,omitempty"`
	ShippingRateID   string                          `json:"shipping_rate_id,omitempty"`
	ShippingPackages []packaging.PhysicalPackage     `json:"shipping_packages,omitempty"`

	// ── Delivery scheduling ───────────────────────────────────────────
	// ExpectedDeliveryDate/Time are the promised delivery slot, set by
	// staff or checkout — distinct from the DeliveredAt lifecycle
	// timestamp recorded after the fact.
	ExpectedDeliveryDate temporal.Date                 `json:"expected_delivery_date,omitempty"`
	ExpectedDeliveryTime temporal.TimeOfDay            `json:"expected_delivery_time,omitempty"`
	DeliveryMethod       shipping_enums.DeliveryMethod `json:"delivery_method,omitempty"`
	// OutsourcedCarrier names the third-party delivery company; set only
	// when DeliveryMethod is outsourced.
	OutsourcedCarrier string `json:"outsourced_carrier,omitempty"`

	// ── Money ─────────────────────────────────────────────────────────
	Subtotal       money.Money `json:"subtotal"`
	DiscountAmount money.Money `json:"discount_amount"`
	ShippingAmount money.Money `json:"shipping_amount"`
	TaxAmount      money.Money `json:"tax_amount"`
	// TipAmount and SurchargeAmount surface card-terminal-applied
	// extras. They are part of the customer-paid total and must appear
	// on tax invoices.
	// Cashout is intentionally not on Order - it is a parallel cash
	// withdrawal that does not change the order's payable total - and
	// lives only on Payment / TerminalTransaction.
	TipAmount       money.Money `json:"tip_amount,omitempty"`
	SurchargeAmount money.Money `json:"surcharge_amount,omitempty"`
	Total           money.Money `json:"total"`

	CouponCode            string                           `json:"coupon_code,omitempty"`
	PromotionApplications []promotion.PromotionApplication `json:"promotion_applications"`
	PointRedemption       *PointRedemptionSnapshot         `json:"point_redemption,omitempty"`
	RewardRedemptions     []RewardRedemptionSnapshot       `json:"reward_redemptions,omitempty"`
	VoucherRedemption     *VoucherRedemptionSnapshot       `json:"voucher_redemption,omitempty"`
	GiftCardRedemptions   []GiftCardRedemptionSnapshot     `json:"gift_card_redemptions,omitempty"`
	TrackingNumber        string                           `json:"tracking_number,omitempty"`
	TrackingURL           string                           `json:"tracking_url,omitempty"`
	Packing               *OrderPackingProgress            `json:"packing,omitempty"`
	PackingRevision       int64                            `json:"packing_revision,omitempty"`
	FulfillmentReadiness  order_enums.FulfillmentReadiness `json:"fulfillment_readiness"`
	CustomerNote          string                           `json:"customer_note,omitempty"`
	InternalNote          string                           `json:"internal_note,omitempty"`
	Tags                  []string                         `json:"tags,omitempty"`

	// ── Lifecycle timestamps ──────────────────────────────────────────
	ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
	PaidAt           *time.Time `json:"paid_at,omitempty"`
	CancelledAt      *time.Time `json:"cancelled_at,omitempty"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
	ShippedAt        *time.Time `json:"shipped_at,omitempty"`
	DeliveredAt      *time.Time `json:"delivered_at,omitempty"`
	PickingPrintedAt *time.Time `json:"picking_printed_at,omitempty"`
	PackedAt         *time.Time `json:"packed_at,omitempty"`

	History []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}
