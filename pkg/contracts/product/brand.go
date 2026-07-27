package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/product"
)

// Brand is the canonical localized brand master used by product catalogues.
// BrandKey is the required lowercase, URL-safe canonical key and is immutable.
type Brand struct {
	ID       string                 `json:"id"`
	BrandKey string                 `json:"brand_key"`
	Slug     string                 `json:"slug"`
	Name     []common.LocalizedName `json:"name"`

	common.AuditFields
}

// BrandRef is the stable, lightweight brand identity embedded in product
// records and snapshots. Brand master audit data remains on Brand.
// BrandKey has the same immutable lowercase URL-safe semantics as Brand.
type BrandRef struct {
	BrandKey string                 `json:"brand_key"`
	Slug     string                 `json:"slug"`
	Name     []common.LocalizedName `json:"name"`
}

// BrandSummary is the customer-safe brand catalogue projection. BrandKey is
// the stable navigation/filter key; Names contains localized display labels.
// Featured is the explicit merchandising control. Audience makes
// ActiveProductCount unambiguous between retail and wholesale catalogues.
type BrandSummary struct {
	BrandKey           string                    `json:"brand_key"`
	Names              []common.LocalizedName    `json:"names"`
	Audience           productenum.PriceAudience `json:"audience,omitempty"`
	LogoURL            string                    `json:"logo_url,omitempty"`
	Featured           bool                      `json:"featured"`
	ActiveProductCount int64                     `json:"active_product_count"`
}
