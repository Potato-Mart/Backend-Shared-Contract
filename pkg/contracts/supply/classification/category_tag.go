package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/localization"
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
