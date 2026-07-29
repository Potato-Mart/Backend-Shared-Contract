package analytics

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
)

// OrderItemFact is the immutable product and merchandising snapshot used by
// sales rollups. Dimension values are canonical identifiers captured at purchase
// time, so later catalogue edits cannot rewrite historical analytics.
type OrderItemFact struct {
	ProductSKUCode string       `json:"product_sku_code"`
	ProductName    string       `json:"product_name,omitempty"`
	BrandID        string       `json:"brand_id,omitempty"`
	StorageType    string       `json:"storage_type,omitempty"`
	CollectionSlug string       `json:"collection_slug,omitempty"`
	CategorySlugs  []string     `json:"category_slugs,omitempty"`
	Production     string       `json:"production,omitempty"`
	Quantity       int64        `json:"quantity"`
	Gross          common.Money `json:"gross"`
}

// RefundItemFact identifies quantities and value reversed by a completed
// line-level refund. Amount-only refunds intentionally carry no item rows.
type RefundItemFact struct {
	ProductSKUCode string       `json:"product_sku_code"`
	BrandID        string       `json:"brand_id,omitempty"`
	StorageType    string       `json:"storage_type,omitempty"`
	CollectionSlug string       `json:"collection_slug,omitempty"`
	CategorySlugs  []string     `json:"category_slugs,omitempty"`
	Production     string       `json:"production,omitempty"`
	Quantity       int64        `json:"quantity"`
	Amount         common.Money `json:"amount"`
}

// OrderFact is the immutable analytical projection of an order event.
type OrderFact struct {
	EventID              string          `json:"event_id"`
	OrderNumber          string          `json:"order_number"`
	RetailCustomerNumber string          `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string          `json:"organisation_access_id,omitempty"`
	Status               string          `json:"status"`
	Channel              string          `json:"channel,omitempty"`
	ItemCount            int             `json:"item_count"`
	Total                common.Money    `json:"total"`
	Items                []OrderItemFact `json:"items,omitempty"`
	OccurredAt           time.Time       `json:"occurred_at"`
}

// PaymentFact is the immutable analytical projection of a payment event.
type PaymentFact struct {
	EventID     string       `json:"event_id"`
	PaymentID   string       `json:"payment_id"`
	OrderNumber string       `json:"order_number"`
	Method      string       `json:"method,omitempty"`
	Status      string       `json:"status"`
	Amount      common.Money `json:"amount"`
	OccurredAt  time.Time    `json:"occurred_at"`
}

// RefundFact is the immutable analytical projection of a refund event.
type RefundFact struct {
	EventID     string           `json:"event_id"`
	RefundID    string           `json:"refund_id"`
	OrderNumber string           `json:"order_number"`
	Status      string           `json:"status"`
	Amount      common.Money     `json:"amount"`
	Items       []RefundItemFact `json:"items,omitempty"`
	OccurredAt  time.Time        `json:"occurred_at"`
}

// MetricRollup is one hourly or daily aggregate consumed by Admin reports.
type MetricRollup struct {
	Metric       string       `json:"metric"`
	Granularity  string       `json:"granularity"`
	WindowStart  time.Time    `json:"window_start"`
	WindowEnd    time.Time    `json:"window_end"`
	Dimension    string       `json:"dimension,omitempty"`
	Count        int64        `json:"count"`
	Amount       common.Money `json:"amount"`
	CalculatedAt time.Time    `json:"calculated_at"`
}
