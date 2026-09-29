package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v38/pkg/contracts/common/packaging"

type InventorySaleCommittedEvent struct {
	MovementID           string                               `json:"movement_id"`
	OrderNumber          string                               `json:"order_number"`
	DepotCode            string                               `json:"depot_code"`
	ReservationID        string                               `json:"reservation_id"`
	AllocationID         string                               `json:"allocation_id"`
	BucketID             string                               `json:"bucket_id"`
	SKUCode              string                               `json:"sku_code"`
	LotID                string                               `json:"lot_id,omitempty"`
	PackageOption        packaging.PackageOptionRef           `json:"package_option"`
	CommittedComposition packaging.PackageCompositionSnapshot `json:"committed_composition"`
	InventoryRevision    int64                                `json:"inventory_revision"`
}
