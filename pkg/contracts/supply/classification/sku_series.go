package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/classification/classification_enums"
)

// SKUSeries is the catalogue family identified by the immutable two-character
// prefix used by every SKU in the series (for example A0 or F2).
type SKUSeries struct {
	ID          string                           `json:"id"`
	Code        string                           `json:"code"`
	StorageType classification_enums.StorageType `json:"storage_type"`
	PrimaryName localization.LocalizedName       `json:"primary_name"`
	OtherNames  []localization.LocalizedName     `json:"other_names,omitempty"`
	audit.AuditFields
}
