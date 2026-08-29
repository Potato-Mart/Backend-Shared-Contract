package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
)

// StockLocationAssignment links one product SKU to one depot-qualified
// location independently of its current quantity.
type StockLocationAssignment struct {
	ID                       string `json:"id"`
	DepotCode                string `json:"depot_code"`
	LocationCode             string `json:"location_code"`
	SKUCode                  string `json:"sku_code"`
	ElectronicShelfLabelCode string `json:"electronic_shelf_label_code,omitempty"`
	IsActive                 bool   `json:"is_active"`

	audit.AuditFields
}
