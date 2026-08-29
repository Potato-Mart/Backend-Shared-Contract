package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/operations"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// StockLocationAvailabilityChangedEvent represents a customer-accessible
// standard-location zero crossing. AssignmentID is its ordering key.
type StockLocationAvailabilityChangedEvent struct {
	AssignmentID             string                                     `json:"assignment_id"`
	DepotCode                string                                     `json:"depot_code"`
	LocationCode             string                                     `json:"location_code"`
	SKUCode                  string                                     `json:"sku_code"`
	AvailableBeforeBaseUnits int64                                      `json:"available_before_base_units"`
	AvailableAfterBaseUnits  int64                                      `json:"available_after_base_units"`
	Direction                warehouse_enums.StockAvailabilityDirection `json:"direction"`
	ElectronicShelfLabelCode string                                     `json:"electronic_shelf_label_code,omitempty"`
	Cause                    operations.InventoryCauseRef               `json:"cause"`
	Revision                 int64                                      `json:"revision"`
	OccurredAt               time.Time                                  `json:"occurred_at"`
	AsOf                     time.Time                                  `json:"as_of"`
}
