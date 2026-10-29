package selection

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/catalogue/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/catalogue/selection/selection_enums"
)

// ProductSelection is a curated, ordered list of products referenced by
// sku_code. It is independent of collections and category tags; membership is
// carried by ProductSelectionItem records.
//
// Code is immutable once assigned. Slug is the public address and is editable
// only while the selection is in draft status. Revision is advanced by every
// successful write, including a membership replacement. Soft deletion uses the
// audit DeletedAt/DeletedBy fields and is separate from archiving.
type ProductSelection struct {
	ID           string                                 `json:"id"`
	Code         string                                 `json:"code"`
	Slug         string                                 `json:"slug"`
	Kind         selection_enums.ProductSelectionKind   `json:"kind"`
	Status       selection_enums.ProductSelectionStatus `json:"status"`
	Name         []localization.LocalizedName           `json:"name"`
	Description  []localization.LocalizedDescription    `json:"description,omitempty"`
	Icon         *classification.ObjectMediaRef         `json:"icon,omitempty"`
	Cover        *classification.ObjectMediaRef         `json:"cover,omitempty"`
	DisplayOrder int                                    `json:"display_order"`
	Revision     int64                                  `json:"revision"`

	audit.AuditFields
}
