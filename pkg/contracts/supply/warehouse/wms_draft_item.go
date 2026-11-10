package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
)

// WMSDraftItem captures an observed package scan and canonical inventory refs.
type WMSDraftItem struct {
	ID               string                               `json:"id"`
	SKUCode          string                               `json:"sku_code"`
	ProductName      string                               `json:"product_name,omitempty"`
	ScannedBarcode   string                               `json:"scanned_barcode,omitempty"`
	PackageOption    packaging.PackageOptionRef           `json:"package_option"`
	HandlingUnit     packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	LotID            string                               `json:"lot_id,omitempty"`
	Location         StockLocationRef                     `json:"location"`
	Composition      packaging.PackageCompositionSnapshot `json:"composition"`
	ObservedDateMark *InventoryDateMark                   `json:"observed_date_mark,omitempty"`
	// ExpiryMonth is the operator-entered expiry month (YYYY-MM) used to pick
	// a lot. It is distinct from ObservedDateMark.
	ExpiryMonth string `json:"expiry_month,omitempty"`
	// Move is the stock move this draft item submits, kept so a retried
	// submit repeats the same request.
	Move *WMSDraftItemMove `json:"move,omitempty"`
}
