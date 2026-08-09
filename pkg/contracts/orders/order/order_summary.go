package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"
)

// OrderSummary is a slim, customer-facing projection of an Order for "my orders"
// views (e.g. a bounded recent-orders strip on a customer profile). It is a read
// projection — deliberately carries no audit/actor fields — and is built from
// Order in pkg/logic. Reads and links use OrderNumber, the human business key.
type OrderSummary struct {
	OrderNumber       string                        `json:"order_number"`
	Status            order_enums.SalesOrderStatus  `json:"status"`
	PaymentStatus     payment_enums.PaymentStatus   `json:"payment_status"`
	FulfillmentStatus order_enums.FulfillmentStatus `json:"fulfillment_status"`
	Channel           commerce_enums.OrderType      `json:"channel,omitempty"`
	// PlacedAt is the customer-meaningful order time (Order.ConfirmedAt,
	// falling back to Order.CreatedAt — Order has no separate order_date).
	PlacedAt       time.Time             `json:"placed_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	Shipping       *party.ContactAddress `json:"shipping,omitempty"`
	Total          money.Money           `json:"total"`
	ItemCount      int                   `json:"item_count"`
	Items          []OrderLineSummary    `json:"items,omitempty"`
	TrackingNumber string                `json:"tracking_number,omitempty"`
	TrackingURL    string                `json:"tracking_url,omitempty"`
}

// OrderLineSummary is a customer-facing package-aware order-line snapshot with
// direct frozen product facts, rather than an embedded catalogue projection.
type OrderLineSummary struct {
	ProductSKUCode       string                               `json:"product_sku_code"`
	ProductName          string                               `json:"product_name"`
	ProductImage         *security.ObjectMedia                `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption         `json:"product_package_option"`
	CapturedAt           time.Time                            `json:"captured_at"`
	Components           []PricedPackageComponent             `json:"components"`
	TotalBaseUnits       int64                                `json:"total_base_units"`
	PackedComposition    packaging.PackageCompositionSnapshot `json:"packed_composition"`
	ReturnedComposition  packaging.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition  packaging.PackageCompositionSnapshot `json:"refunded_composition"`
	Total                money.Money                          `json:"total"`
}
