package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/payment"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/sales"
)

// OrderSummary is a slim, customer-facing projection of an Order for "my orders"
// views (e.g. a bounded recent-orders strip on a customer profile). It is a read
// projection â€” deliberately carries no audit/actor fields â€” and is built from
// Order in pkg/logic. Reads and links use OrderNumber, the human business key.
type OrderSummary struct {
	OrderNumber       string                      `json:"order_number"`
	Status            salesenum.SalesOrderStatus  `json:"status"`
	PaymentStatus     paymentenum.PaymentStatus   `json:"payment_status"`
	FulfillmentStatus salesenum.FulfillmentStatus `json:"fulfillment_status"`
	Channel           salesenum.OrderType         `json:"channel,omitempty"`
	// PlacedAt is the customer-meaningful order time (Order.ConfirmedAt,
	// falling back to Order.CreatedAt â€” Order has no separate order_date).
	PlacedAt       time.Time              `json:"placed_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
	Shipping       *common.ContactAddress `json:"shipping,omitempty"`
	Total          common.Money           `json:"total"`
	ItemCount      int                    `json:"item_count"`
	Items          []OrderLineSummary     `json:"items,omitempty"`
	TrackingNumber string                 `json:"tracking_number,omitempty"`
	TrackingURL    string                 `json:"tracking_url,omitempty"`
}

// OrderLineSummary is a tiny per-line snapshot: the SKU code, the product's
// primary name, and the unit price paid at order time.
type OrderLineSummary struct {
	SKUCode   string       `json:"sku_code"`
	Name      string       `json:"name"`
	ImageURL  string       `json:"image_url,omitempty"`
	Quantity  int          `json:"quantity"`
	UnitPrice common.Money `json:"unit_price"`
	Total     common.Money `json:"total"`
}
