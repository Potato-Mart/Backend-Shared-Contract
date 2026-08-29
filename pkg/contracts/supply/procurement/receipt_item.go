package procurement

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse"
)

type ReceiptItem struct {
	ID string `json:"id,omitempty"`
	// SKUCode is the frozen SKU code captured when the receipt line was recorded.
	SKUCode             string                               `json:"sku_code"`
	ProductName         string                               `json:"product_name,omitempty"`
	PackageOptionCode   string                               `json:"package_option_code"`
	LotID               string                               `json:"lot_id"`
	SupplierLotCode     string                               `json:"supplier_lot_code,omitempty"`
	ManufacturerLotCode string                               `json:"manufacturer_lot_code,omitempty"`
	DestinationBucketID string                               `json:"destination_bucket_id"`
	DestinationLocation warehouse.StockLocationRef           `json:"destination_location"`
	DateMark            *warehouse.InventoryDateMark         `json:"date_mark,omitempty"`
	OrderedComposition  packaging.PackageCompositionSnapshot `json:"ordered_composition"`
	ReceivedComposition packaging.PackageCompositionSnapshot `json:"received_composition"`
	RejectedComposition packaging.PackageCompositionSnapshot `json:"rejected_composition"`
	Note                string                               `json:"note,omitempty"`
}
