package category

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// SKU corresponds to one of the top-level SKU codes that
// identify a product family (e.g. A0 = 特色台灣商品/吊飾, F2 = 冷凍-肉品).
// Products reference a category by its SKU Code.
type SKU struct {
	ID          string                 `json:"id"`
	Code        string                 `json:"code"`
	Storage     enums.StorageType      `json:"storage"`
	PrimaryName common.LocalizedName   `json:"primary_name"`
	OtherNames  []common.LocalizedName `json:"other_names,omitempty"`
	Products    []product.Snapshot     `json:"products,omitempty"`
	SortOrder   int                    `json:"sort_order"`

	common.AuditFields `bson:",inline"`
}
