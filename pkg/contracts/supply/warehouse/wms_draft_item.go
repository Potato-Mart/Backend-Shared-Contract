package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging/packaging_enums"
)

// WMSDraftItem captures an observed package scan and canonical inventory refs.
type WMSDraftItem struct {
	ID                string                               `json:"id"`
	SKUCode           string                               `json:"sku_code"`
	ProductName       string                               `json:"product_name,omitempty"`
	ScannedBarcode    string                               `json:"scanned_barcode,omitempty"`
	PackageOptionCode string                               `json:"package_option_code"`
	HandlingUnit      packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	LotID             string                               `json:"lot_id,omitempty"`
	Location          StockLocationRef                     `json:"location"`
	Composition       packaging.PackageCompositionSnapshot `json:"composition"`
	ObservedDateMark  *InventoryDateMark                   `json:"observed_date_mark,omitempty"`
}
