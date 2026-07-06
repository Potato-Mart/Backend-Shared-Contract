package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
)

type CategoryTag struct {
	ID             string                 `json:"id"`
	Name           []common.LocalizedName `json:"name"`
	CollectionID   string                 `json:"collection_id"`
	CollectionName []common.LocalizedName `json:"collection_name"`

	common.AuditFields
}
