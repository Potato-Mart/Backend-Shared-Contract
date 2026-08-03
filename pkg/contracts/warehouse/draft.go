package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// WMSDraft is an uncommitted package-aware warehouse operation.
type WMSDraft struct {
	ID               string                            `json:"id"`
	Type             warehouseenum.WMSDraftType        `json:"type"`
	Operator         string                            `json:"operator"`
	DepotCode        string                            `json:"depot_code"`
	Reference        string                            `json:"reference,omitempty"`
	Items            []WMSDraftItem                    `json:"items"`
	ItemCount        int64                             `json:"item_count"`
	TotalComposition common.PackageCompositionSnapshot `json:"total_composition"`
	Status           warehouseenum.WMSDraftStatus      `json:"status"`
	Note             string                            `json:"note,omitempty"`
	SubmittedAt      *time.Time                        `json:"submitted_at,omitempty"`
	History          []shared.HistoryEntry             `json:"history,omitempty"`

	common.AuditFields
}

// WMSDraftItem captures an observed package scan and canonical inventory refs.
type WMSDraftItem struct {
	ID               string                            `json:"id"`
	ProductSKUCode   string                            `json:"product_sku_code"`
	ProductName      string                            `json:"product_name,omitempty"`
	ScannedBarcode   string                            `json:"scanned_barcode,omitempty"`
	PackageOptionID  string                            `json:"package_option_id"`
	HandlingUnit     common.PackageHandlingUnit        `json:"handling_unit"`
	LotID            string                            `json:"lot_id,omitempty"`
	Location         StockLocationRef                  `json:"location"`
	Composition      common.PackageCompositionSnapshot `json:"composition"`
	ObservedDateMark *InventoryDateMark                `json:"observed_date_mark,omitempty"`
}
