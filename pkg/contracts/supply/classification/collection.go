package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/localization"
)

type Collection struct {
	ID           string                       `json:"id"`
	Code         string                       `json:"code"`
	Slug         string                       `json:"slug"`
	Name         []localization.LocalizedName `json:"name"`
	Icon         *ObjectMediaRef              `json:"icon,omitempty"`
	CategoryTags []CategoryTag                `json:"category_tags"`

	audit.AuditFields
}
