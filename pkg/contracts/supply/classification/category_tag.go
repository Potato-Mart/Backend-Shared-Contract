package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
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

// CategoryTagRef identifies a category tag without embedding mutable
// classification or audit details. It is used by canonical product records and
// other membership-style models.
type CategoryTagRef struct {
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name,omitempty"`
}
