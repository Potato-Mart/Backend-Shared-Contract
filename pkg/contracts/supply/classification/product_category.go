package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

// ProductCategory corresponds to one of the top-level catalogue category
// codes that identify a product family (e.g. A0 = 特色台灣商品/吊飾,
// F2 = 冷凍-肉品). Products reference a category by its Code.
type ProductCategory struct {
	ID          string                       `json:"id"`
	Code        string                       `json:"code"`
	StorageType warehouse_enums.StorageType  `json:"storage_type"`
	PrimaryName localization.LocalizedName   `json:"primary_name"`
	OtherNames  []localization.LocalizedName `json:"other_names,omitempty"`
	audit.AuditFields
}
