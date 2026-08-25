package analytics

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification/classification_enums"
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
	// indexed match instead of joining every market back to its country. An
	// empty value provides no geographic evidence and must be handled
	// fail-closed rather than widening the reader's visibility.
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
