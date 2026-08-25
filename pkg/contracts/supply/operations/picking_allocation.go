package operations

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse"
)

type PickingAllocation struct {
	ReservationAllocationID string                               `json:"reservation_allocation_id"`
	SourceBucketID          string                               `json:"source_bucket_id"`
	StockUnitIDs            []string                             `json:"stock_unit_ids,omitempty"`
	SourceLocation          warehouse.StockLocationRef           `json:"source_location"`
	LotID                   string                               `json:"lot_id,omitempty"`
	PackageOptionCode       string                               `json:"package_option_code"`
	AllocatedComposition    packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	PickedComposition       packaging.PackageCompositionSnapshot `json:"picked_composition"`
	ScannedBarcodes         []string                             `json:"scanned_barcodes,omitempty"`
}
