package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/payments/payment/payment_enums"
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
	PlacedAt           time.Time                           `json:"placed_at"`
	UpdatedAt          time.Time                           `json:"updated_at"`
	FulfilmentLocation shipping.FulfilmentLocationSnapshot `json:"fulfilment_location"`
	Total              money.Money                         `json:"total"`
	ItemCount          int                                 `json:"item_count"`
	Items              []OrderLineSummary                  `json:"items,omitempty"`
	TrackingNumber     string                              `json:"tracking_number,omitempty"`
	TrackingURL        string                              `json:"tracking_url,omitempty"`
}
