package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
)

type CategoryTag struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CollectionID string `json:"collection_id"`

	common.AuditFields
}
