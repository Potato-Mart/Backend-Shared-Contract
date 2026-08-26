package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"
)

type InventoryLotReceivedEvent struct {
	LotID               string                               `json:"lot_id"`
	SKUCode             string                               `json:"sku_code"`
	DepotCode           string                               `json:"depot_code"`
	DestinationBucketID string                               `json:"destination_bucket_id"`
	ReceivedComposition packaging.PackageCompositionSnapshot `json:"received_composition"`
	MovementID          string                               `json:"movement_id"`
	LotRevision         int64                                `json:"lot_revision"`
	ReceivedAt          time.Time                            `json:"received_at"`
	OccurredAt          time.Time                            `json:"occurred_at"`
}
