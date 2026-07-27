package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
)

type CategoryTag struct {
	ID             string                 `json:"id"`
	Slug           string                 `json:"slug,omitempty"`
	Name           []common.LocalizedName `json:"name"`
	CollectionID   string                 `json:"collection_id"`
	CollectionName []common.LocalizedName `json:"collection_name"`

	common.AuditFields
}
