package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

type Order struct {
	ID                string                  `json:"id"`
	OrderNumber       string                  `json:"order_number"`
	Channel           enums.OrderType         `json:"channel"`
	Status            enums.SalesOrderStatus  `json:"status"`
	PaymentStatus     enums.PaymentStatus     `json:"payment_status"`
	PaymentMethod     enums.PaymentMethod     `json:"payment_method"`
	FulfillmentStatus enums.FulfillmentStatus `json:"fulfillment_status"`
	Customer          common.PartyRef         `json:"customer"`
	Items             []OrderItem             `json:"items"`
	SourceDevice      SourceDevice            `json:"source_device,omitempty"`

	// ── Shipping & billing ────────────────────────────────────────────
	Shipping         common.ContactAddress    `json:"shipping"`
	Billing          *common.ContactAddress   `json:"billing,omitempty"`
	ShippingMethod   enums.ShippingRateName   `json:"shipping_method,omitempty"`
	ShippingZoneID   string                   `json:"shipping_zone_id,omitempty"`
	ShippingRateID   string                   `json:"shipping_rate_id,omitempty"`
	ShippingPackages []common.PhysicalPackage `json:"shipping_packages,omitempty"`

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

	CouponCode        string             `json:"coupon_code,omitempty"`
	AppliedPromotions []AppliedPromotion `json:"applied_promotions,omitempty"`
	TrackingNumber    string             `json:"tracking_number,omitempty"`
	TrackingURL       string             `json:"tracking_url,omitempty"`
	CustomerNote      string             `json:"customer_note,omitempty"`
	InternalNote      string             `json:"internal_note,omitempty"`
	Tags              []string           `json:"tags,omitempty"`

	// ── Lifecycle timestamps ──────────────────────────────────────────
	ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
	CancelledAt      *time.Time `json:"cancelled_at,omitempty"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
	ShippedAt        *time.Time `json:"shipped_at,omitempty"`
	DeliveredAt      *time.Time `json:"delivered_at,omitempty"`
	PickingPrintedAt *time.Time `json:"picking_printed_at,omitempty"`
	PackedAt         *time.Time `json:"packed_at,omitempty"`

	common.AuditFields `bson:",inline"`
}

type OrderItem struct {
	ID             string           `json:"id,omitempty"`
	Product        product.Snapshot `json:"product"`
	VariantTitle   string           `json:"variant_title,omitempty"`
	UnitPrice      common.Money     `json:"unit_price"`
	Quantity       int              `json:"quantity"`
	DiscountAmount common.Money     `json:"discount_amount"`
	Total          common.Money     `json:"total"`
	CartonQty      int              `json:"carton_qty,omitempty"`
	CartonSize     int              `json:"carton_size,omitempty"`
	Properties     common.Metadata  `json:"properties,omitempty"`
}

type AppliedPromotion struct {
	ID             string             `json:"id,omitempty"`
	Name           string             `json:"name,omitempty"`
	DiscountType   enums.DiscountType `json:"discount_type,omitempty"`
	DiscountValue  string             `json:"discount_value,omitempty"`
	DiscountAmount *common.Money      `json:"discount_amount,omitempty"`
}

type SourceDevice struct {
	Type     enums.OrderSourceDeviceType `json:"type,omitempty"`
	LocalID  string                      `json:"local_id,omitempty"`
	DeviceID string                      `json:"device_id,omitempty"`
	Name     string                      `json:"name,omitempty"`

	// Metadata stores source-specific details that should not become first-class
	// contract fields yet, for example app_version, terminal_id, store_id,
	// operator_id, forwarded_for, device_model, or network_interface.
	Metadata common.Metadata `json:"metadata,omitempty"`

	// DeviceRecord carries shared fingerprint/request attributes such as
	// device_key, ip_address, user_agent, os, and browser.
	shared.DeviceRecord `bson:",inline"`
}
