package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
)

type CollectionRef struct {
	ID   string                       `json:"id"`
	Slug string                       `json:"slug,omitempty"`
	Name []localization.LocalizedName `json:"name"`
}

type Collection struct {
	ID           string                       `json:"id"`
	Slug         string                       `json:"slug,omitempty"`
	Name         []localization.LocalizedName `json:"name"`
	Icon         *security.ObjectMedia        `json:"icon,omitempty"`
	CategoryTags []CategoryTag                `json:"category_tags"`

	audit.AuditFields
}
