package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"

// Brand is the canonical localized brand master used by product catalogues.
// Aliases carry alternate localized spellings and search labels without
// replacing the canonical localized Name values.
type Brand struct {
	ID      string                 `json:"id"`
	Slug    string                 `json:"slug"`
	Name    []common.LocalizedName `json:"name"`
	Aliases []common.LocalizedName `json:"aliases,omitempty"`

	common.AuditFields
}

// BrandRef is the stable, lightweight brand identity embedded in product
// records and snapshots. Brand master audit and alias data remain on Brand.
type BrandRef struct {
	ID   string                 `json:"id"`
	Slug string                 `json:"slug"`
	Name []common.LocalizedName `json:"name"`
}
