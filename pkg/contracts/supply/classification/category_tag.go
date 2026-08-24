package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
)

type CategoryTag struct {
	ID             string                       `json:"id"`
	Code           string                       `json:"code"`
	Slug           string                       `json:"slug"`
	Name           []localization.LocalizedName `json:"name"`
	CollectionCode string                       `json:"collection_code"`
	CollectionName []localization.LocalizedName `json:"collection_name"`

	audit.AuditFields
}

// CategoryTagRef identifies a category tag by its immutable code without
// embedding mutable classification or audit details.
type CategoryTagRef struct {
	Code string `json:"code"`
}
