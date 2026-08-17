package analytics

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

// OrderItemFact is the immutable product and merchandising snapshot used by
// sales rollups. Dimension values are canonical identifiers captured at purchase
// time, so later catalogue edits cannot rewrite historical analytics.
type OrderItemFact struct {
	SKUID string `json:"sku_id"`
	// MarketID qualifies the fact; Insights never infers a market from
	// country or currency.
	MarketID           string                               `json:"market_id"`
	ProductName        string                               `json:"product_name,omitempty"`
	BrandID            string                               `json:"brand_id,omitempty"`
	StorageType        warehouse_enums.StorageType          `json:"storage_type,omitempty"`
	CollectionSlug     string                               `json:"collection_slug,omitempty"`
	CategorySlugs      []string                             `json:"category_slugs,omitempty"`
	Production         string                               `json:"production,omitempty"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	Gross              money.Money                          `json:"gross"`
}

// RefundItemFact identifies quantities and value reversed by a completed
// line-level refund. Amount-only refunds intentionally carry no item rows.
type RefundItemFact struct {
	SKUID              string                               `json:"sku_id"`
	MarketID           string                               `json:"market_id"`
	BrandID            string                               `json:"brand_id,omitempty"`
	StorageType        warehouse_enums.StorageType          `json:"storage_type,omitempty"`
	CollectionSlug     string                               `json:"collection_slug,omitempty"`
	CategorySlugs      []string                             `json:"category_slugs,omitempty"`
	Production         string                               `json:"production,omitempty"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	Amount             money.Money                          `json:"amount"`
}

// MetricRollup is one hourly or daily aggregate consumed by Admin reports.
type MetricRollup struct {
	Metric       string      `json:"metric"`
	Granularity  string      `json:"granularity"`
	WindowStart  time.Time   `json:"window_start"`
	WindowEnd    time.Time   `json:"window_end"`
	Dimension    string      `json:"dimension,omitempty"`
	Count        int64       `json:"count"`
	Amount       money.Money `json:"amount"`
	CalculatedAt time.Time   `json:"calculated_at"`
}
