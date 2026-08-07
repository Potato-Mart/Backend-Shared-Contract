package order

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/payments/payment"
)

// OrderSummary is a slim, customer-facing projection of an Order for "my orders"
// views (e.g. a bounded recent-orders strip on a customer profile). It is a read
// projection — deliberately carries no audit/actor fields — and is built from
// Order in pkg/logic. Reads and links use OrderNumber, the human business key.
type OrderSummary struct {
	OrderNumber       string                    `json:"order_number"`
	Status            SalesOrderStatus          `json:"status"`
	PaymentStatus     paymentenum.PaymentStatus `json:"payment_status"`
	FulfillmentStatus FulfillmentStatus         `json:"fulfillment_status"`
	Channel           common.OrderType          `json:"channel,omitempty"`
	// PlacedAt is the customer-meaningful order time (Order.ConfirmedAt,
	// falling back to Order.CreatedAt — Order has no separate order_date).
	PlacedAt       time.Time              `json:"placed_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
	Shipping       *common.ContactAddress `json:"shipping,omitempty"`
	Total          common.Money           `json:"total"`
	ItemCount      int                    `json:"item_count"`
	Items          []OrderLineSummary     `json:"items,omitempty"`
	TrackingNumber string                 `json:"tracking_number,omitempty"`
	TrackingURL    string                 `json:"tracking_url,omitempty"`
}

// OrderLineSummary is a customer-facing package-aware order-line snapshot.
type OrderLineSummary struct {
	SKUCode             string                            `json:"sku_code"`
	Name                string                            `json:"name"`
	ImageURL            string                            `json:"image_url,omitempty"`
	Components          []PricedPackageComponent          `json:"components"`
	TotalBaseUnits      int64                             `json:"total_base_units"`
	PackedComposition   common.PackageCompositionSnapshot `json:"packed_composition"`
	ReturnedComposition common.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition common.PackageCompositionSnapshot `json:"refunded_composition"`
	Total               common.Money                      `json:"total"`
}
