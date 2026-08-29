package supply

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	warehouse "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse"
)

type StockStagingChangedEvent struct {
	StagingRecordID     string                               `json:"staging_record_id"`
	ReservationID       string                               `json:"reservation_id"`
	AllocationID        string                               `json:"allocation_id"`
	OrderNumber         string                               `json:"order_number"`
	SKUCode             string                               `json:"sku_code"`
	SourceLocation      warehouse.StockLocationRef           `json:"source_location"`
	DestinationLocation warehouse.StockLocationRef           `json:"destination_location"`
	StagedComposition   packaging.PackageCompositionSnapshot `json:"staged_composition"`
	MovementID          string                               `json:"movement_id"`
	Revision            int64                                `json:"revision"`
	OccurredAt          time.Time                            `json:"occurred_at"`
}
