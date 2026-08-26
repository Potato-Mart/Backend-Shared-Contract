package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse"
)

// StockStagingRecord captures the physical transfer of an allocation into an
// online-order staging location.
type StockStagingRecord struct {
	ID                  string                               `json:"id"`
	ReservationID       string                               `json:"reservation_id"`
	AllocationID        string                               `json:"allocation_id"`
	OrderNumber         string                               `json:"order_number"`
	SKUCode             string                               `json:"sku_code"`
	LotID               string                               `json:"lot_id,omitempty"`
	PackageOptionCode   string                               `json:"package_option_code"`
	SourceBucketID      string                               `json:"source_bucket_id"`
	DestinationBucketID string                               `json:"destination_bucket_id"`
	SourceLocation      warehouse.StockLocationRef           `json:"source_location"`
	DestinationLocation warehouse.StockLocationRef           `json:"destination_location"`
	StagedComposition   packaging.PackageCompositionSnapshot `json:"staged_composition"`
	MovementID          string                               `json:"movement_id"`
	StagedBy            string                               `json:"staged_by"`
	StagedAt            time.Time                            `json:"staged_at"`
}
