package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

// InventoryCategoryTagEvidence records a location-qualified operational tag
// applied to inventory. Soon-expiry and damaged stock are represented by the
// category tag and inventory evidence here, never by a promotion kind.
type InventoryCategoryTagEvidence struct {
	SKUCode           string                               `json:"sku_code"`
	PackageOptionCode string                               `json:"package_option_code"`
	CategoryTag       classification.CategoryTagRef        `json:"category_tag"`
	StockLocation     warehouse.StockLocationRef           `json:"stock_location"`
	Condition         warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition       warehouse_enums.InventoryDisposition `json:"disposition"`
	DateMark          *warehouse.InventoryDateMark         `json:"date_mark,omitempty"`
	AsOf              time.Time                            `json:"as_of"`
}
