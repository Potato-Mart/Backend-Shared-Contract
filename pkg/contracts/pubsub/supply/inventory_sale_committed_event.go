package supply

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply/supply_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/inventory"
)

// InventorySaleCommittedEvent records one physical issue against a reservation.
// Exactly one physical evidence variant is required. Unknown returns always
// have unresolved cost; known-lot and unallocated no-lot stock are valued
// only when their actual movement cost is verified.
// Reservation shortages are releases, not sale or deficit movements.
type InventorySaleCommittedEvent struct {
	MovementID               string                                      `json:"movement_id"`
	OrderNumber              string                                      `json:"order_number"`
	DepotCode                string                                      `json:"depot_code"`
	ReservationID            string                                      `json:"reservation_id"`
	AllocationID             string                                      `json:"allocation_id"`
	BucketID                 string                                      `json:"bucket_id"`
	SKUCode                  string                                      `json:"sku_code"`
	PackageOption            packaging.PackageOptionRef                  `json:"package_option"`
	CommittedComposition     packaging.PackageCompositionSnapshot        `json:"committed_composition"`
	InventoryRevision        int64                                       `json:"inventory_revision"`
	CostStatus               supply_enums.InventorySaleCostStatus        `json:"cost_status"`
	KnownLotPhysical         *inventory.KnownLotPhysicalEvidence         `json:"known_lot_physical,omitempty"`
	UnknownIdentityPhysical  *inventory.UnknownReturnPhysicalEvidence    `json:"unknown_identity_physical,omitempty"`
	UnallocatedNoLotPhysical *inventory.UnallocatedNoLotPhysicalEvidence `json:"unallocated_no_lot_physical,omitempty"`
}
