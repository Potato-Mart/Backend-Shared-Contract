package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

type Order struct {
	ID                string                  `json:"id"`
	OrderNumber       string                  `json:"order_number"`
	Channel           enums.OrderType         `json:"type"`
	Status            enums.SalesOrderStatus  `json:"status"`
	PaymentStatus     enums.PaymentStatus     `json:"payment_status"`
	FulfillmentStatus enums.FulfillmentStatus `json:"fulfillment_status"`
	Customer          Customer                `json:"customers"`
	Items             []OrderItem             `json:"items"`
	ShippingContact   *common.Recipient       `json:"shipping_contact"`
	ShippingAddress   *common.Address         `json:"shipping_address,omitempty"`
	BillingContact    *common.Recipient       `json:"billing_contact,omitempty"`
	BillingAddress    *common.Address         `json:"billing_address,omitempty"`
	ShippingMethod    string                  `json:"shipping_method,omitempty"`
	ShippingZoneID    string                  `json:"shipping_zone_id,omitempty"`
	ShippingRateID    string                  `json:"shipping_rate_id,omitempty"`
	Subtotal          float64                 `json:"subtotal"`
	DiscountAmount    float64                 `json:"discount_amount"`
	ShippingAmount    float64                 `json:"shipping_amount"`
	TaxAmount         float64                 `json:"tax_amount"`
	Total             float64                 `json:"total"`
	Currency          string                  `json:"currency"`
	CouponCode        string                  `json:"coupon_code,omitempty"`
	AppliedPromotions []AppliedPromotion      `json:"applied_promotions,omitempty"`
	TrackingNumber    string                  `json:"tracking_number,omitempty"`
	TrackingURL       string                  `json:"tracking_url,omitempty"`
	CustomerNote      string                  `json:"customer_note,omitempty"`
	InternalNote      string                  `json:"internal_note,omitempty"`
	Tags              []string                `json:"tags,omitempty"`
	ConfirmedAt       *time.Time              `json:"confirmed_at,omitempty"`
	CancelledAt       *time.Time              `json:"cancelled_at,omitempty"`
	CompletedAt       *time.Time              `json:"completed_at,omitempty"`
	ShippedAt         *time.Time              `json:"shipped_at,omitempty"`
	DeliveredAt       *time.Time              `json:"delivered_at,omitempty"`
	PickingPrintedAt  *time.Time              `json:"picking_printed_at,omitempty"`
	PackedAt          *time.Time              `json:"packed_at,omitempty"`
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"updated_at"`
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
	UnitPrice      float64         `json:"unit_price"`
	Quantity       int             `json:"quantity"`
	DiscountAmount float64         `json:"discount_amount"`
	Total          float64         `json:"total"`
	CartonQty      int             `json:"carton_qty,omitempty"`
	CartonSize     int             `json:"carton_size,omitempty"`
	Properties     common.Metadata `json:"properties,omitempty"`
}

type ProductSnapshot struct {
	ID         string   `json:"id,omitempty"`
	SKU        string   `json:"sku,omitempty"`
	Name       string   `json:"name"`
	OtherNames []string `json:"other_names,omitempty"`
	Brand      string   `json:"brand,omitempty"`
	ImageURL   string   `json:"image_url,omitempty"`
	Storage    string   `json:"storage,omitempty"`
}

type AppliedPromotion struct {
	ID             string             `json:"id,omitempty"`
	Name           string             `json:"name,omitempty"`
	DiscountType   enums.DiscountType `json:"discount_type,omitempty"`
	DiscountValue  float64            `json:"discount_value,omitempty"`
	DiscountAmount float64            `json:"discount_amount,omitempty"`
}
