package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"

type Collection struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	CategoryTags []CategoryTag `json:"category_tags"`

	common.AuditFields
}
