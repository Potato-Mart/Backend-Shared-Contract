package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v47/pkg/contracts/common/packaging"

// InventoryPOSSaleCommittedEvent records one physical or deficit movement from
// an unreserved POS sale. Deficit movements remain pending cost reconciliation.
type InventoryPOSSaleCommittedEvent struct {
	SaleID                string                               `json:"sale_id"`
	OrderNumber           string                               `json:"order_number"`
	FulfillmentGeneration int64                                `json:"fulfillment_generation"`
	OrderItemID           string                               `json:"order_item_id"`
	MovementID            string                               `json:"movement_id"`
	DepotCode             string                               `json:"depot_code"`
	LocationCode          string                               `json:"location_code"`
	BucketID              string                               `json:"bucket_id"`
	SKUCode               string                               `json:"sku_code"`
	LotID                 string                               `json:"lot_id,omitempty"`
	PackageOption         packaging.PackageOptionRef           `json:"package_option"`
	CommittedComposition  packaging.PackageCompositionSnapshot `json:"committed_composition"`
	BaseUnits             int64                                `json:"base_units"`
	IsDeficit             bool                                 `json:"is_deficit"`
	InventoryRevision     int64                                `json:"inventory_revision"`
}
