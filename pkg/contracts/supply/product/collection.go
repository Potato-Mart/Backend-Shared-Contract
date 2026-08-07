package product

import common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"

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
