package product

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

type CategoryTag struct {
	ID             string                 `json:"id"`
	Slug           string                 `json:"slug,omitempty"`
	Name           []common.LocalizedName `json:"name"`
	CollectionID   string                 `json:"collection_id"`
	CollectionName []common.LocalizedName `json:"collection_name"`

	common.AuditFields
}
