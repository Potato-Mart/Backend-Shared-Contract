package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/shared"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/membership"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/payment"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/promotion"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/sales"
	shippingenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/shipping"
)

type Order struct {
	ID                string                      `json:"id"`
	OrderNumber       string                      `json:"order_number"`
	Channel           salesenum.OrderType         `json:"channel"`
	Status            salesenum.SalesOrderStatus  `json:"status"`
	PaymentStatus     paymentenum.PaymentStatus   `json:"payment_status"`
	PaymentMethod     paymentenum.PaymentMethod   `json:"payment_method"`
	FulfillmentStatus salesenum.FulfillmentStatus `json:"fulfillment_status"`
	Customer          common.PartyRef             `json:"customer"`
	// Buyer describes who is buying, independently of Channel. POS is a
	// channel, not a buyer type — see sales.BuyerContext. Optional pointer
	// so it is omitted entirely when unset.
	Buyer        *BuyerContext `json:"buyer,omitempty"`
	Items        []OrderItem   `json:"items"`
	SourceDevice SourceDevice  `json:"source_device,omitempty"`

	// ── Shipping & billing ────────────────────────────────────────────
	Shipping         common.ContactAddress         `json:"shipping"`
	Billing          *common.ContactAddress        `json:"billing,omitempty"`
	ShippingMethod   shippingenum.ShippingRateName `json:"shipping_method,omitempty"`
	ShippingZoneID   string                        `json:"shipping_zone_id,omitempty"`
	ShippingRateID   string                        `json:"shipping_rate_id,omitempty"`
	ShippingPackages []common.PhysicalPackage      `json:"shipping_packages,omitempty"`

	// ── Delivery scheduling ───────────────────────────────────────────
	// ExpectedDeliveryDate/Time are the promised delivery slot, set by
	// staff or checkout — distinct from the DeliveredAt lifecycle
	// timestamp recorded after the fact.
	ExpectedDeliveryDate common.Date                 `json:"expected_delivery_date,omitempty"`
	ExpectedDeliveryTime common.TimeOfDay            `json:"expected_delivery_time,omitempty"`
	DeliveryMethod       shippingenum.DeliveryMethod `json:"delivery_method,omitempty"`
	// OutsourcedCarrier names the third-party delivery company; set only
	// when DeliveryMethod is outsourced.
	OutsourcedCarrier string                      `json:"outsourced_carrier,omitempty"`
	DeliveryRegion    shippingenum.DeliveryRegion `json:"delivery_region,omitempty"`

	// ── Money ─────────────────────────────────────────────────────────
	Subtotal       common.Money `json:"subtotal"`
	DiscountAmount common.Money `json:"discount_amount"`
	ShippingAmount common.Money `json:"shipping_amount"`
	TaxAmount      common.Money `json:"tax_amount"`
	// TipAmount and SurchargeAmount surface card-terminal-applied
	// extras. They are part of the customer-paid total and must appear
	// on tax invoices.
	// Cashout is intentionally not on Order - it is a parallel cash
	// withdrawal that does not change the order's payable total - and
	// lives only on Payment / TerminalTransaction.
	TipAmount       common.Money `json:"tip_amount,omitempty"`
	SurchargeAmount common.Money `json:"surcharge_amount,omitempty"`
	Total           common.Money `json:"total"`

	CouponCode          string                       `json:"coupon_code,omitempty"`
	AppliedPromotions   []AppliedPromotion           `json:"applied_promotions,omitempty"`
	PointRedemption     *PointRedemptionSnapshot     `json:"point_redemption,omitempty"`
	RewardRedemptions   []RewardRedemptionSnapshot   `json:"reward_redemptions,omitempty"`
	GiftCardRedemptions []GiftCardRedemptionSnapshot `json:"gift_card_redemptions,omitempty"`
	TrackingNumber      string                       `json:"tracking_number,omitempty"`
	TrackingURL         string                       `json:"tracking_url,omitempty"`
	CustomerNote        string                       `json:"customer_note,omitempty"`
	InternalNote        string                       `json:"internal_note,omitempty"`
	Tags                []string                     `json:"tags,omitempty"`

	// ── Lifecycle timestamps ──────────────────────────────────────────
	ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
	CancelledAt      *time.Time `json:"cancelled_at,omitempty"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
	ShippedAt        *time.Time `json:"shipped_at,omitempty"`
	DeliveredAt      *time.Time `json:"delivered_at,omitempty"`
	PickingPrintedAt *time.Time `json:"picking_printed_at,omitempty"`
	PackedAt         *time.Time `json:"packed_at,omitempty"`

	History []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

type OrderItem struct {
	ID           string           `json:"id,omitempty"`
	Product      product.Snapshot `json:"product"`
	VariantTitle string           `json:"variant_title,omitempty"`
	UnitPrice    common.Money     `json:"unit_price"`
	// Pricing is the commercial pricing context under which UnitPrice was
	// set (retail vs wholesale audience, visibility). Optional pointer so it
	// is omitted entirely when unset.
	Pricing        *PricingContext `json:"pricing,omitempty"`
	Quantity       int             `json:"quantity"`
	DiscountAmount common.Money    `json:"discount_amount"`
	Total          common.Money    `json:"total"`
	CartonQty      int             `json:"carton_qty,omitempty"`
	CartonSize     int             `json:"carton_size,omitempty"`
	Properties     common.Metadata `json:"properties,omitempty"`
}

type AppliedPromotion struct {
	ID             string                     `json:"id,omitempty"`
	Name           string                     `json:"name,omitempty"`
	DiscountType   promotionenum.DiscountType `json:"discount_type,omitempty"`
	DiscountValue  string                     `json:"discount_value,omitempty"`
	DiscountAmount *common.Money              `json:"discount_amount,omitempty"`
}

// PointRedemptionSnapshot records a points discount applied to an order without
// overloading coupon or promotion fields.
type PointRedemptionSnapshot struct {
	MembershipAccountID  string                             `json:"membership_account_id"`
	OwnerType            membershipenum.MembershipOwnerType `json:"owner_type"`
	OwnerID              string                             `json:"owner_id"`
	OrganisationAccessID string                             `json:"organisation_access_id,omitempty"`
	ReservationID        string                             `json:"reservation_id,omitempty"`
	LedgerEntryID        string                             `json:"ledger_entry_id,omitempty"`
	Points               int                                `json:"points"`
	DiscountAmount       common.Money                       `json:"discount_amount"`
}

// RewardRedemptionSnapshot records a catalog reward applied to an order.
type RewardRedemptionSnapshot struct {
	RewardRedemptionID  string                              `json:"reward_redemption_id"`
	RewardCode          string                              `json:"reward_code"`
	MembershipAccountID string                              `json:"membership_account_id"`
	RewardType          membershipenum.MembershipRewardType `json:"reward_type"`
	PointsSpent         int                                 `json:"points_spent"`
	DiscountAmount      *common.Money                       `json:"discount_amount,omitempty"`
	ProductSKUCode      string                              `json:"product_sku_code,omitempty"`
	VoucherCode         string                              `json:"voucher_code,omitempty"`
}

// GiftCardRedemptionSnapshot records a gift-card amount applied to an order,
// alongside PointRedemptionSnapshot and RewardRedemptionSnapshot. The gift card
// is referenced by its code; its re-spendable balance ledger lives in
// pkg/contracts/wallet.
type GiftCardRedemptionSnapshot struct {
	GiftCardCode string       `json:"gift_card_code"`
	Amount       common.Money `json:"amount"`
	OrderNumber  string       `json:"order_number,omitempty"`
}

type SourceDevice struct {
	Type    salesenum.OrderSourceDeviceType `json:"type,omitempty"`
	LocalID string                          `json:"local_id,omitempty"`
	Name    string                          `json:"name,omitempty"`

	// Metadata stores source-specific details that should not become first-class
	// contract fields yet, for example app_version, terminal_id, store_id,
	// operator_id, forwarded_for, device_model, or network_interface.
	Metadata common.Metadata `json:"metadata,omitempty"`

	// DeviceRecord carries shared fingerprint/request attributes such as
	// device_key, ip_address, user_agent, os, and browser.
	shared.DeviceRecord
}
