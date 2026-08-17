package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

// WMSDraft is an uncommitted package-aware warehouse operation.
type WMSDraft struct {
	ID               string                               `json:"id"`
	Type             warehouse_enums.WMSDraftType         `json:"type"`
	Operator         string                               `json:"operator"`
	DepotCode        string                               `json:"depot_code"`
	Reference        string                               `json:"reference,omitempty"`
	Items            []WMSDraftItem                       `json:"items"`
	ItemCount        int64                                `json:"item_count"`
	TotalComposition packaging.PackageCompositionSnapshot `json:"total_composition"`
	Status           warehouse_enums.WMSDraftStatus       `json:"status"`
	Note             string                               `json:"note,omitempty"`
	SubmittedAt      *time.Time                           `json:"submitted_at,omitempty"`
	History          []security.HistoryEntry              `json:"history,omitempty"`

	audit.AuditFields
}

// WMSDraftItem captures an observed package scan and canonical inventory refs.
type WMSDraftItem struct {
	ID               string                               `json:"id"`
	SKUID            string                               `json:"sku_id"`
	ProductName      string                               `json:"product_name,omitempty"`
	ScannedBarcode   string                               `json:"scanned_barcode,omitempty"`
	PackageOptionID  string                               `json:"package_option_id"`
	HandlingUnit     packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	LotID            string                               `json:"lot_id,omitempty"`
	Location         StockLocationRef                     `json:"location"`
	Composition      packaging.PackageCompositionSnapshot `json:"composition"`
	ObservedDateMark *InventoryDateMark                   `json:"observed_date_mark,omitempty"`
}
