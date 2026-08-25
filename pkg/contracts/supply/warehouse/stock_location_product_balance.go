package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
)

// StockLocationProductBalance is a revisioned quantity projection for one
// stock-location assignment.
type StockLocationProductBalance struct {
	AssignmentID       string                               `json:"assignment_id"`
	DepotCode          string                               `json:"depot_code"`
	LocationCode       string                               `json:"location_code"`
	SKUCode            string                               `json:"sku_code"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	OnHandBaseUnits    int64                                `json:"on_hand_base_units"`
	ReservedBaseUnits  int64                                `json:"reserved_base_units"`
	AvailableBaseUnits int64                                `json:"available_base_units"`
	IsOutOfStock       bool                                 `json:"is_out_of_stock"`
	Revision           int64                                `json:"revision"`
	LastRestockedAt    *time.Time                           `json:"last_restocked_at,omitempty"`
	DepotTimezone      string                               `json:"depot_timezone"`
	AsOf               time.Time                            `json:"as_of"`
}
