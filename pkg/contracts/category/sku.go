package category

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/warehouse"
)

// SKU corresponds to one of the top-level SKU codes that
// identify a product family (e.g. A0 = ç‰¹è‰²å°ç£å•†å“/åŠé£¾, F2 = å†·å‡-è‚‰å“).
// Products reference a category by its SKU Code.
type SKU struct {
	ID          string                    `json:"id"`
	Code        string                    `json:"code"`
	Storage     warehouseenum.StorageType `json:"storage"`
	PrimaryName common.LocalizedName      `json:"primary_name"`
	OtherNames  []common.LocalizedName    `json:"other_names,omitempty"`
	Products    []product.Snapshot        `json:"products,omitempty"`
	SortOrder   int                       `json:"sort_order"`

	common.AuditFields
}
