package category

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// SKU corresponds to one of the top-level SKU codes that
// identify a product family (e.g. A0 = 特色台灣商品/吊飾, F2 = 冷凍-肉品).
// Products reference a category by its SKU Code.
type SKU struct {
	ID          string                    `json:"id"`
	Code        string                    `json:"code"`
	StorageType warehouseenum.StorageType `json:"storage_type"`
	PrimaryName common.LocalizedName      `json:"primary_name"`
	OtherNames  []common.LocalizedName    `json:"other_names,omitempty"`
	common.AuditFields
}
