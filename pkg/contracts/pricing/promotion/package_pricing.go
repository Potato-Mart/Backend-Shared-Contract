package promotion

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackagePricing is the revisioned package price accepted for a product at a
// location. It carries the inventory evidence needed by an accepted order line
// and the ordered promotion applications resolved for that price.
type PackagePricing struct {
	ID                string `json:"id"`
	Revision          int64  `json:"revision"`
	InventoryRevision int64  `json:"inventory_revision"`

	ProductSKUCode  string      `json:"product_sku_code"`
	PackageOptionID string      `json:"package_option_id"`
	PackagePrice    money.Money `json:"package_price"`
	TaxAmount       money.Money `json:"tax_amount"`

	ValidFrom         time.Time                   `json:"valid_from"`
	ValidTo           *time.Time                  `json:"valid_to,omitempty"`
	Timezone          string                      `json:"timezone"`
	GeographicContext geography.GeographicContext `json:"geographic_context"`

	StockLocation         warehouse.StockLocationRef           `json:"stock_location"`
	SourceBucketID        string                               `json:"source_bucket_id,omitempty"`
	SourceStockUnitID     string                               `json:"source_stock_unit_id,omitempty"`
	AvailablePackageCount int64                                `json:"available_package_count"`
	AvailableBaseUnits    int64                                `json:"available_base_units"`
	Condition             warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition           warehouse_enums.InventoryDisposition `json:"disposition"`
	DateMark              *warehouse.InventoryDateMark         `json:"date_mark,omitempty"`

	PromotionApplications []PromotionApplication `json:"promotion_applications"`
	CapturedAt            time.Time              `json:"captured_at"`
}
