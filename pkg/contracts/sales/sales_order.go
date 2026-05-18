package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

type Order struct {
	ID                string                   `json:"id"`
	OrderNumber       string                   `json:"order_number"`
	Channel           enums.OrderType          `json:"channel"`
	Status            enums.SalesOrderStatus   `json:"status"`
	PaymentStatus     enums.PaymentStatus      `json:"payment_status"`
	FulfillmentStatus enums.FulfillmentStatus  `json:"fulfillment_status"`
	Customer          Customer                 `json:"customer"`
	Items             []OrderItem              `json:"items"`
	ShippingContact   *common.Recipient        `json:"shipping_contact"`
	ShippingAddress   *common.Address          `json:"shipping_address,omitempty"`
	BillingContact    *common.Recipient        `json:"billing_contact,omitempty"`
	BillingAddress    *common.Address          `json:"billing_address,omitempty"`
	ShippingMethod    string                   `json:"shipping_method,omitempty"`
	ShippingZoneID    string                   `json:"shipping_zone_id,omitempty"`
	ShippingRateID    string                   `json:"shipping_rate_id,omitempty"`
	ShippingPackages  []common.PhysicalPackage `json:"shipping_packages,omitempty"`
	Subtotal          common.Money             `json:"subtotal"`
	DiscountAmount    common.Money             `json:"discount_amount"`
	ShippingAmount    common.Money             `json:"shipping_amount"`
	TaxAmount         common.Money             `json:"tax_amount"`
	// TipAmount and SurchargeAmount surface card-terminal-applied
	// extras (for example, Adyen Terminal API tip/surcharge result
	// fields). They are part of the customer-paid total and must appear
	// on tax invoices.
	// Cashout is intentionally not on Order - it is a parallel cash
	// withdrawal that does not change the order's payable total - and
	// lives only on Payment / TerminalTransaction.
	TipAmount         common.Money       `json:"tip_amount,omitempty"`
	SurchargeAmount   common.Money       `json:"surcharge_amount,omitempty"`
	Total             common.Money       `json:"total"`
	CouponCode        string             `json:"coupon_code,omitempty"`
	AppliedPromotions []AppliedPromotion `json:"applied_promotions,omitempty"`
	TrackingNumber    string             `json:"tracking_number,omitempty"`
	TrackingURL       string             `json:"tracking_url,omitempty"`
	CustomerNote      string             `json:"customer_note,omitempty"`
	InternalNote      string             `json:"internal_note,omitempty"`
	Tags              []string           `json:"tags,omitempty"`
	ConfirmedAt       *time.Time         `json:"confirmed_at,omitempty"`
	CancelledAt       *time.Time         `json:"cancelled_at,omitempty"`
	CompletedAt       *time.Time         `json:"completed_at,omitempty"`
	ShippedAt         *time.Time         `json:"shipped_at,omitempty"`
	DeliveredAt       *time.Time         `json:"delivered_at,omitempty"`
	PickingPrintedAt  *time.Time         `json:"picking_printed_at,omitempty"`
	PackedAt          *time.Time         `json:"packed_at,omitempty"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

type Customer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type OrderItem struct {
	ID             string          `json:"id,omitempty"`
	Product        ProductSnapshot `json:"product"`
	VariantTitle   string          `json:"variant_title,omitempty"`
	UnitPrice      common.Money    `json:"unit_price"`
	Quantity       int             `json:"quantity"`
	DiscountAmount common.Money    `json:"discount_amount"`
	Total          common.Money    `json:"total"`
	CartonQty      int             `json:"carton_qty,omitempty"`
	CartonSize     int             `json:"carton_size,omitempty"`
	Properties     common.Metadata `json:"properties,omitempty"`
}

type ProductSnapshot struct {
	ID         string                 `json:"id,omitempty"`
	SKU        string                 `json:"sku,omitempty"`
	Name       string                 `json:"name"`
	OtherNames []common.LocalizedName `json:"other_names,omitempty"`
	Brand      string                 `json:"brand,omitempty"`
	ImageURL   string                 `json:"image_url,omitempty"`
	Storage    string                 `json:"storage,omitempty"`
}

type AppliedPromotion struct {
	ID             string             `json:"id,omitempty"`
	Name           string             `json:"name,omitempty"`
	DiscountType   enums.DiscountType `json:"discount_type,omitempty"`
	DiscountValue  string             `json:"discount_value,omitempty"`
	DiscountAmount *common.Money      `json:"discount_amount,omitempty"`
}
