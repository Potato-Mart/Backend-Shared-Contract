package supply

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply/supply_enums"
)

// InventoryLotReceivedEvent reports stock arriving in an inventory lot.
// Cause says why the stock arrived; an absent cause comes from a publisher
// that predates it and means a receipt. OriginLotID names the lot the stock
// came from, for a lot_split or lot_reassignment cause.
type InventoryLotReceivedEvent struct {
	LotID               string                                 `json:"lot_id"`
	SKUCode             string                                 `json:"sku_code"`
	DepotCode           string                                 `json:"depot_code"`
	DestinationBucketID string                                 `json:"destination_bucket_id"`
	ReceivedComposition packaging.PackageCompositionSnapshot   `json:"received_composition"`
	MovementID          string                                 `json:"movement_id"`
	LotRevision         int64                                  `json:"lot_revision"`
	ReceivedAt          time.Time                              `json:"received_at"`
	Cause               supply_enums.InventoryLotReceivedCause `json:"cause,omitempty"`
	OriginLotID         string                                 `json:"origin_lot_id,omitempty"`
}
