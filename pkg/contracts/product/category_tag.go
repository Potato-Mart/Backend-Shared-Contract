package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
)

type CategoryTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	common.AuditFields `bson:",inline"`
}
