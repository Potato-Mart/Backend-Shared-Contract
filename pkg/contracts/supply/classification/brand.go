package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"
)

// Brand is the canonical localized brand master used by product catalogues.
type Brand struct {
	ID   string                       `json:"id"`
	Code string                       `json:"code"`
	Slug string                       `json:"slug"`
	Name []localization.LocalizedName `json:"name"`
	Logo *security.ObjectMedia        `json:"logo,omitempty"`
}

// BrandRef is the stable display identity embedded in product records and
// snapshots. Code is the immutable canonical brand business key; mutable
// slugs and database identifiers are deliberately excluded.
type BrandRef struct {
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name,omitempty"`
	Logo *security.ObjectMedia        `json:"logo,omitempty"`
}
