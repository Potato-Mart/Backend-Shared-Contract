package product

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
