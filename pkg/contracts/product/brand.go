package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
)

// Brand is the canonical localized brand master used by product catalogues.
type Brand struct {
	ID      string                 `json:"id"`
	Slug    string                 `json:"slug"`
	Name    []common.LocalizedName `json:"name"`
	LogoURL string                 `json:"logo_url,omitempty"`
}

// BrandRef is the stable display identity embedded in product records and
// snapshots. ID is the matching canonical brand master identifier.
type BrandRef struct {
	ID      string                 `json:"id"`
	Slug    string                 `json:"slug"`
	Name    []common.LocalizedName `json:"name"`
	LogoURL string                 `json:"logo_url,omitempty"`
}
