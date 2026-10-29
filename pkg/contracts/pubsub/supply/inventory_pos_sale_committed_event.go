package supply

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pubsub/supply/supply_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/supply/inventory"
)

// InventoryPOSSaleCommittedEvent records one actual movement from an unreserved
// POS sale. Exactly one evidence variant is required. Producers and consumers
// must verify variant exclusivity and all source, quantity, and cost evidence.
// Unknown-identity physical stock and deficits always have unresolved cost.
// Known-lot and unallocated no-lot physical stock are valued only when their
// actual issue cost is verified; neither variant invents missing provenance.
type InventoryPOSSaleCommittedEvent struct {
	SaleID                   string                                      `json:"sale_id"`
	OrderNumber              string                                      `json:"order_number"`
	FulfillmentGeneration    int64                                       `json:"fulfillment_generation"`
	OrderItemID              string                                      `json:"order_item_id"`
	MovementID               string                                      `json:"movement_id"`
	DepotCode                string                                      `json:"depot_code"`
	LocationCode             string                                      `json:"location_code"`
	BucketID                 string                                      `json:"bucket_id"`
	SKUCode                  string                                      `json:"sku_code"`
	PackageOption            packaging.PackageOptionRef                  `json:"package_option"`
	CommittedComposition     packaging.PackageCompositionSnapshot        `json:"committed_composition"`
	BaseUnits                int64                                       `json:"base_units"`
	InventoryRevision        int64                                       `json:"inventory_revision"`
	CostStatus               supply_enums.POSSaleCostStatus              `json:"cost_status"`
	KnownLotPhysical         *POSKnownLotPhysicalEvidence                `json:"known_lot_physical,omitempty"`
	UnknownIdentityPhysical  *POSUnknownIdentityPhysicalEvidence         `json:"unknown_identity_physical,omitempty"`
	UnallocatedNoLotPhysical *inventory.UnallocatedNoLotPhysicalEvidence `json:"unallocated_no_lot_physical,omitempty"`
	Deficit                  *POSDeficitEvidence                         `json:"deficit,omitempty"`
}
