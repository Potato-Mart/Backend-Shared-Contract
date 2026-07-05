package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"

type CollectionRef struct {
	ID   string                 `json:"id"`
	Name []common.LocalizedName `json:"name"`
}

type Collection struct {
	ID           string                 `json:"id"`
	Name         []common.LocalizedName `json:"name"`
	CategoryTags []CategoryTag          `json:"category_tags"`

	common.AuditFields
}
