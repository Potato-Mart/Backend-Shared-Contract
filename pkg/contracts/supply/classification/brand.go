package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
)

// Brand is the canonical localized brand master used by product catalogues.
type Brand struct {
	ID   string                       `json:"id"`
	Code string                       `json:"code"`
	Slug string                       `json:"slug"`
	Name []localization.LocalizedName `json:"name"`
	Logo *ObjectMediaRef              `json:"logo,omitempty"`
}
