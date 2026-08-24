package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
)

// CollectionRef identifies a collection by its immutable business code.
// Display names and presentation slugs are resolved from the root master.
type CollectionRef struct {
	Code string `json:"code"`
}

type Collection struct {
	ID           string                       `json:"id"`
	Code         string                       `json:"code"`
	Slug         string                       `json:"slug"`
	Name         []localization.LocalizedName `json:"name"`
	Icon         *ObjectMediaRef              `json:"icon,omitempty"`
	CategoryTags []CategoryTag                `json:"category_tags"`

	audit.AuditFields
}
