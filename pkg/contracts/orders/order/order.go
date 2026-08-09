package order

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/device"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
)

// Buyer describes who is buying, independently of Channel. POS is a
// channel, not a buyer type — see sales.BuyerContext. Optional pointer
// so it is omitted entirely when unset.
type Order struct {
	ID                string                        `json:"id"`
	OrderNumber       string                        `json:"order_number"`
	Channel           commerce_enums.OrderType      `json:"channel"`
	Status            order_enums.SalesOrderStatus  `json:"status"`
	PaymentStatus     payment_enums.PaymentStatus   `json:"payment_status"`
	PaymentMethod     payment_enums.PaymentMethod   `json:"payment_method"`
	FulfillmentStatus order_enums.FulfillmentStatus `json:"fulfillment_status"`
	Customer          party.PartyRef                `json:"customer"`

	Buyer                *BuyerContext               `json:"buyer,omitempty"`
	GeographicContext    geography.GeographicContext `json:"geographic_context"`
	GroupOrder           *GroupOrderContext          `json:"group_order,omitempty"`
	GroupOrderFulfilment *GroupOrderFulfilmentPlan   `json:"group_order_fulfilment,omitempty"`
	Items                []OrderItem                 `json:"items"`
	SourceDevice         SourceDevice                `json:"source_device,omitempty"`

	// ── Shipping & billing ────────────────────────────────────────────
	Shipping         party.ContactAddress            `json:"shipping"`
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

	CouponCode           string                           `json:"coupon_code,omitempty"`
	AppliedPromotions    []AppliedPromotion               `json:"applied_promotions,omitempty"`
	PointRedemption      *PointRedemptionSnapshot         `json:"point_redemption,omitempty"`
	RewardRedemptions    []RewardRedemptionSnapshot       `json:"reward_redemptions,omitempty"`
	VoucherRedemption    *VoucherRedemptionSnapshot       `json:"voucher_redemption,omitempty"`
	GiftCardRedemptions  []GiftCardRedemptionSnapshot     `json:"gift_card_redemptions,omitempty"`
	TrackingNumber       string                           `json:"tracking_number,omitempty"`
	TrackingURL          string                           `json:"tracking_url,omitempty"`
	Packing              *OrderPackingProgress            `json:"packing,omitempty"`
	PackingRevision      int64                            `json:"packing_revision,omitempty"`
	FulfillmentReadiness order_enums.FulfillmentReadiness `json:"fulfillment_readiness"`
	CustomerNote         string                           `json:"customer_note,omitempty"`
	InternalNote         string                           `json:"internal_note,omitempty"`
	Tags                 []string                         `json:"tags,omitempty"`

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

// OrderPackingProgress is the order-owned packing state shown to both staff
// and customer order views. It replaces admin-facing dependence on a separate
// packing-session aggregate for current packing progress.
type OrderPackingProgress struct {
	Status        order_enums.FulfillmentStatus     `json:"status,omitempty"`
	Operator      string                            `json:"operator,omitempty"`
	Lines         []warehouse.PackingLine           `json:"lines,omitempty"`
	Containers    []warehouse.OutboundContainerPlan `json:"containers,omitempty"`
	Damages       []warehouse.PackingDamage         `json:"damages,omitempty"`
	Discrepancies []warehouse.PackingDiscrepancy    `json:"discrepancies,omitempty"`
	StartedAt     *time.Time                        `json:"started_at,omitempty"`
	UpdatedAt     *time.Time                        `json:"updated_at,omitempty"`
	PackedAt      *time.Time                        `json:"packed_at,omitempty"`
	FulfilledAt   *time.Time                        `json:"fulfilled_at,omitempty"`
}

type OrderItem struct {
	ID                   string                                  `json:"id"`
	Product              product.Snapshot                        `json:"product"`
	VariantTitle         string                                  `json:"variant_title,omitempty"`
	Components           []PricedPackageComponent                `json:"components"`
	TotalBaseUnits       int64                                   `json:"total_base_units"`
	Pricing              *PricingContext                         `json:"pricing,omitempty"`
	SubstitutionPolicy   LooseSubstitutionPolicySnapshot         `json:"substitution_policy"`
	RequestedComposition packaging.PackageCompositionSnapshot    `json:"requested_composition"`
	AllocatedComposition packaging.PackageCompositionSnapshot    `json:"allocated_composition"`
	PickedComposition    packaging.PackageCompositionSnapshot    `json:"picked_composition"`
	PackedComposition    packaging.PackageCompositionSnapshot    `json:"packed_composition"`
	ReturnedComposition  packaging.PackageCompositionSnapshot    `json:"returned_composition"`
	RefundedComposition  packaging.PackageCompositionSnapshot    `json:"refunded_composition"`
	Substitutions        []warehouse.PackageSubstitutionSnapshot `json:"substitutions,omitempty"`
	DiscountAmount       money.Money                             `json:"discount_amount"`
	Total                money.Money                             `json:"total"`
	Preorder             *PreorderItemState                      `json:"preorder,omitempty"`
}

type AppliedPromotion struct {
	ID                string                       `json:"id,omitempty"`
	Name              string                       `json:"name,omitempty"`
	DiscountType      promotion_enums.DiscountType `json:"discount_type,omitempty"`
	DiscountValue     string                       `json:"discount_value,omitempty"`
	DiscountAmount    *money.Money                 `json:"discount_amount,omitempty"`
	GeographicContext geography.GeographicContext  `json:"geographic_context"`
}

// PointRedemptionSnapshot records a points discount applied to an order without
// overloading coupon or promotion fields.
type PointRedemptionSnapshot struct {
	CustomerNumber string      `json:"customer_number"`
	ReservationID  string      `json:"reservation_id,omitempty"`
	LedgerEntryID  string      `json:"ledger_entry_id,omitempty"`
	Points         int         `json:"points"`
	DiscountAmount money.Money `json:"discount_amount"`
	OccurredAt     *time.Time  `json:"occurred_at,omitempty"`
}

// RewardRedemptionSnapshot records a catalog reward applied to an order.
type RewardRedemptionSnapshot struct {
	RewardRedemptionID string                                `json:"reward_redemption_id"`
	RewardCode         string                                `json:"reward_code"`
	CustomerNumber     string                                `json:"customer_number"`
	RewardType         membership_enums.MembershipRewardType `json:"reward_type"`
	PointsSpent        int                                   `json:"points_spent"`
	DiscountAmount     *money.Money                          `json:"discount_amount,omitempty"`
	ProductSKUCode     string                                `json:"product_sku_code,omitempty"`
	VoucherCode        string                                `json:"voucher_code,omitempty"`
	OccurredAt         *time.Time                            `json:"occurred_at,omitempty"`
}

// VoucherRedemptionSnapshot records the single voucher applied to an order.
type VoucherRedemptionSnapshot struct {
	VoucherCode   string      `json:"voucher_code"`
	AppliedAmount money.Money `json:"applied_amount"`
	ReservationID string      `json:"reservation_id,omitempty"`
	OccurredAt    *time.Time  `json:"occurred_at,omitempty"`
}

// GiftCardRedemptionSnapshot records one ordered gift-card allocation applied
// to an order. WalletTransactionID links the snapshot and completed gift-card
// payment to the committed wallet ledger entry.
type GiftCardRedemptionSnapshot struct {
	GiftCardCode        string      `json:"gift_card_code"`
	AppliedAmount       money.Money `json:"applied_amount"`
	ReservationID       string      `json:"reservation_id,omitempty"`
	WalletTransactionID string      `json:"wallet_transaction_id,omitempty"`
	OccurredAt          *time.Time  `json:"occurred_at,omitempty"`
}

// POSAttribution carries first-class in-store sale attribution (store, event,
// register, device, operator, shift, platform, form factor) on the order's
// source device.
type POSAttribution struct {
	StoreID        string `json:"store_id,omitempty"`
	EventID        string `json:"event_id,omitempty"`
	RegisterID     string `json:"register_id,omitempty"`
	ShiftID        string `json:"shift_id,omitempty"`
	OperatorUserID string `json:"operator_user_id,omitempty"`
	Platform       string `json:"platform,omitempty"`
	FormFactor     string `json:"form_factor,omitempty"`
}

type SourceDevice struct {
	Type    order_enums.OrderSourceDeviceType `json:"type,omitempty"`
	LocalID string                            `json:"local_id,omitempty"`
	Name    string                            `json:"name,omitempty"`

	// POS carries first-class in-store attribution when the order originates
	// at a point of sale.
	POS *POSAttribution `json:"pos,omitempty"`

	// Metadata stores source-specific details that should not become first-class
	// contract fields yet, for example app_version, terminal_id, store_id,
	// operator_id, forwarded_for, device_model, or network_interface.
	Metadata metadata.Metadata `json:"metadata,omitempty"`

	// DeviceRecord carries shared fingerprint/request attributes such as
	// device_key, ip_address, user_agent, os, and browser.
	device.DeviceRecord
}
