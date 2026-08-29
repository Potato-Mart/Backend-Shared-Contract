package analytics

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/classification/classification_enums"
)

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
