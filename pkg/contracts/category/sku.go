package category

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

// SKU corresponds to one of the top-level SKU codes that
// identify a product family (e.g. A0 = 特色台灣商品/吊飾, F2 = 冷凍-肉品).
// Products reference a category by its SKU Code.
type SKU struct {
	ID         string                 `json:"id"`
	Code       string                 `json:"code"`
	Storage    enums.StorageType      `json:"storage"`
	OtherNames []common.LocalizedName `json:"other_names,omitempty"`
	Products   []string               `json:"products,omitempty"`
	SortOrder  int                    `json:"sort_order"`

	common.AuditFields
}
