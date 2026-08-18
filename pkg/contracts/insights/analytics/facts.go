package analytics

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/classification/classification_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/packaging"
)

// OrderItemFact is the immutable product and merchandising snapshot used by
// sales rollups. Dimension values are canonical identifiers captured at purchase
// time, so later catalogue edits cannot rewrite historical analytics.
type OrderItemFact struct {
	SKUCode string `json:"sku_code"`
	// MarketCode qualifies the fact; Insights never infers a market from
	// country or currency.
	MarketCode string `json:"market_code"`
	// CountryCode is the denormalized country the fact is attributed to,
	// carried so a country-scoped principal can be filtered by a plain
	// indexed match instead of joining every market back to its country.
	// It is absent on facts recorded before v28.0.0; an absent value means
	// "unattributed" and the consumer handles it fail-closed rather than
	// widening the reader's visibility.
	CountryCode        geography.CountryCode                `json:"country_code,omitempty"`
	ProductName        string                               `json:"product_name,omitempty"`
	BrandCode          string                               `json:"brand_code,omitempty"`
	StorageType        classification_enums.StorageType     `json:"storage_type,omitempty"`
	CollectionCode     string                               `json:"collection_code,omitempty"`
	CategoryTagCodes   []string                             `json:"category_tag_codes,omitempty"`
	Production         string                               `json:"production,omitempty"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	Gross              money.Money                          `json:"gross"`
}

// RefundItemFact identifies quantities and value reversed by a completed
// line-level refund. Amount-only refunds intentionally carry no item rows.
type RefundItemFact struct {
	SKUCode    string `json:"sku_code"`
	MarketCode string `json:"market_code"`
	// CountryCode carries the same denormalized attribution, and the same
	// fail-closed handling of an absent value, as OrderItemFact.CountryCode.
	CountryCode        geography.CountryCode                `json:"country_code,omitempty"`
	BrandCode          string                               `json:"brand_code,omitempty"`
	StorageType        classification_enums.StorageType     `json:"storage_type,omitempty"`
	CollectionCode     string                               `json:"collection_code,omitempty"`
	CategoryTagCodes   []string                             `json:"category_tag_codes,omitempty"`
	Production         string                               `json:"production,omitempty"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	Amount             money.Money                          `json:"amount"`
}

// MetricRollup is one hourly or daily aggregate consumed by Admin reports.
//
// MarketCode and CountryCode are the denormalized geography the aggregate is
// attributed to. A rollup carrying neither is platform-wide and is readable
// only at global scope; a geographically scoped principal filters on these
// two fields and a rollup that does not match is not returned. Both are
// absent on rollups computed before v28.0.0, and an absent value means
// "unattributed" rather than "everywhere".
type MetricRollup struct {
	Metric       string                `json:"metric"`
	Granularity  string                `json:"granularity"`
	WindowStart  time.Time             `json:"window_start"`
	WindowEnd    time.Time             `json:"window_end"`
	Dimension    string                `json:"dimension,omitempty"`
	MarketCode   string                `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode `json:"country_code,omitempty"`
	Count        int64                 `json:"count"`
	Amount       money.Money           `json:"amount"`
	CalculatedAt time.Time             `json:"calculated_at"`
}
