package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// InventoryStockBucket is the quantity authority for one package form,
// location, lot, condition, and disposition combination. A CASE bucket
// represents intact cases and an EACH bucket represents loose base units.
type InventoryStockBucket struct {
	ID                 string                               `json:"id"`
	Location           StockLocationRef                     `json:"location"`
	SKUCode            string                               `json:"sku_code"`
	LotID              string                               `json:"lot_id,omitempty"`
	PackageOptionCode  string                               `json:"package_option_code"`
	HandlingUnit       packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	Condition          warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition        warehouse_enums.InventoryDisposition `json:"disposition"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	OnHandBaseUnits    int64                                `json:"on_hand_base_units"`
	ReservedBaseUnits  int64                                `json:"reserved_base_units"`
	// AvailableBaseUnits is a derived JSON projection for this bucket.
	AvailableBaseUnits int64     `json:"available_base_units"`
	Revision           int64     `json:"revision"`
	DepotTimezone      string    `json:"depot_timezone"`
	AsOf               time.Time `json:"as_of"`

	audit.AuditFields
}
