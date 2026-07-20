package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/product"
)

// Brand is the canonical localized brand master used by product catalogues.
// Aliases carry alternate localized spellings and search labels without
// replacing the canonical localized Name values.
// BrandKey is the lowercase, URL-safe canonical key and is immutable after the
// owning catalogue assigns it. It remains optional while existing masters are
// backfilled.
type Brand struct {
	ID       string                 `json:"id"`
	BrandKey string                 `json:"brand_key,omitempty"`
	Slug     string                 `json:"slug"`
	Name     []common.LocalizedName `json:"name"`
	Aliases  []common.LocalizedName `json:"aliases,omitempty"`

	common.AuditFields
}

// BrandRef is the stable, lightweight brand identity embedded in product
// records and snapshots. Brand master audit and alias data remain on Brand.
// BrandKey has the same immutable lowercase URL-safe semantics as Brand and is
// optional so references written before backfill keep their JSON shape.
type BrandRef struct {
	ID       string                 `json:"id"`
	BrandKey string                 `json:"brand_key,omitempty"`
	Slug     string                 `json:"slug"`
	Name     []common.LocalizedName `json:"name"`
}

// BrandSummary is the customer-safe brand catalogue projection. BrandKey is
// the stable navigation/filter key; Names contains localized display labels.
// Featured and SortOrder are explicit merchandising controls. Audience makes
// ActiveProductCount unambiguous between retail and wholesale catalogues; an
// omitted audience retains the legacy retail interpretation. LogoURL is
// omitted when no approved HTTPS asset is configured.
type BrandSummary struct {
	BrandKey           string                    `json:"brand_key"`
	Names              []common.LocalizedName    `json:"names"`
	Audience           productenum.PriceAudience `json:"audience,omitempty"`
	LogoURL            string                    `json:"logo_url,omitempty"`
	Featured           bool                      `json:"featured"`
	SortOrder          int                       `json:"sort_order"`
	ActiveProductCount int64                     `json:"active_product_count"`
}
