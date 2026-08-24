package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
)

// Brand is the canonical localized brand master used by product catalogues.
type Brand struct {
	ID   string                       `json:"id"`
	Code string                       `json:"code"`
	Slug string                       `json:"slug"`
	Name []localization.LocalizedName `json:"name"`
	Logo *ObjectMediaRef              `json:"logo,omitempty"`
}

// BrandRef is the stable relationship embedded in product records and
// snapshots. Display data is resolved from the brand master by its immutable
// code; mutable names, logos, slugs, and database identifiers are excluded.
type BrandRef struct {
	Code string `json:"code"`
}
