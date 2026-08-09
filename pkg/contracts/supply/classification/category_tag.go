package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
)

type CategoryTag struct {
	ID             string                       `json:"id"`
	Slug           string                       `json:"slug,omitempty"`
	Name           []localization.LocalizedName `json:"name"`
	CollectionID   string                       `json:"collection_id"`
	CollectionName []localization.LocalizedName `json:"collection_name"`

	audit.AuditFields
}

// CategoryTagRef identifies a category tag without embedding mutable
// classification or audit details. It is used by canonical product records and
// other membership-style models.
type CategoryTagRef struct {
	ID   string                       `json:"id"`
	Slug string                       `json:"slug,omitempty"`
	Name []localization.LocalizedName `json:"name,omitempty"`
}
