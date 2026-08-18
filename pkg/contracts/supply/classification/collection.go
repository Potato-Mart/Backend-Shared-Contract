package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"
)

type CollectionRef struct {
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name"`
}

type Collection struct {
	ID           string                       `json:"id"`
	Code         string                       `json:"code"`
	Slug         string                       `json:"slug"`
	Name         []localization.LocalizedName `json:"name"`
	Icon         *security.ObjectMedia        `json:"icon,omitempty"`
	CategoryTags []CategoryTag                `json:"category_tags"`

	audit.AuditFields
}
