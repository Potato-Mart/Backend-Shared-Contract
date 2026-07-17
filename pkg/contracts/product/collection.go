package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"

type CollectionRef struct {
	ID   string                 `json:"id"`
	Slug string                 `json:"slug,omitempty"`
	Name []common.LocalizedName `json:"name"`
}

type Collection struct {
	ID           string                 `json:"id"`
	Slug         string                 `json:"slug,omitempty"`
	Name         []common.LocalizedName `json:"name"`
	CategoryTags []CategoryTag          `json:"category_tags"`

	common.AuditFields
}
