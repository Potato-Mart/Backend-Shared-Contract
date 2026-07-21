package product

import (
	"time"

	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/product"
)

// ProductChangedEvent is emitted on the catalog-events topic when a product
// master changes. Change is a discriminator: created, updated, archived,
// deleted. AggregateID is the SKU code.
type ProductChangedEvent struct {
	SKUCode    string                    `json:"sku_code"`
	Name       string                    `json:"name,omitempty"`
	Status     productenum.ProductStatus `json:"status,omitempty"`
	Change     string                    `json:"change"`
	OccurredAt time.Time                 `json:"occurred_at"`
	RequestID  string                    `json:"request_id,omitempty"`
}

// BrandChangedEvent is emitted on the catalog-events topic when a brand
// master changes.
type BrandChangedEvent struct {
	BrandID    string    `json:"brand_id"`
	BrandKey   string    `json:"brand_key,omitempty"`
	Slug       string    `json:"slug,omitempty"`
	Change     string    `json:"change"`
	OccurredAt time.Time `json:"occurred_at"`
	RequestID  string    `json:"request_id,omitempty"`
}

// CategoryChangedEvent is emitted on the catalog-events topic when a
// collection or category tag changes.
type CategoryChangedEvent struct {
	CollectionSlug string    `json:"collection_slug,omitempty"`
	CategorySlug   string    `json:"category_slug,omitempty"`
	Change         string    `json:"change"`
	OccurredAt     time.Time `json:"occurred_at"`
	RequestID      string    `json:"request_id,omitempty"`
}

// ProductSalesRollup is the product-stats topic payload: per-SKU sales
// statistics published on an interval by the sales-order owner and
// materialised by the catalog owner. AggregateID is the SKU code.
type ProductSalesRollup struct {
	ProductSKUCode string           `json:"product_sku_code"`
	Last7Days      SalesWindowStats `json:"last_7_days"`
	Last30Days     SalesWindowStats `json:"last_30_days"`
	Last90Days     SalesWindowStats `json:"last_90_days"`
	Lifetime       SalesTotals      `json:"lifetime"`
	LastSoldAt     *time.Time       `json:"last_sold_at,omitempty"`
	AsOf           time.Time        `json:"as_of"`
	Timezone       string           `json:"timezone,omitempty"`
}
