package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

type InventoryReservationChangedEvent struct {
	ReservationID        string                                 `json:"reservation_id"`
	SKUCode              string                                 `json:"sku_code"`
	DepotCode            string                                 `json:"depot_code"`
	PreviousStatus       warehouse_enums.StockReservationStatus `json:"previous_status,omitempty"`
	Status               warehouse_enums.StockReservationStatus `json:"status"`
	RequestedComposition packaging.PackageCompositionSnapshot   `json:"requested_composition"`
	ReservedComposition  packaging.PackageCompositionSnapshot   `json:"reserved_composition"`
	Revision             int64                                  `json:"revision"`
	OccurredAt           time.Time                              `json:"occurred_at"`
}
