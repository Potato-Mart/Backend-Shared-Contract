package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

type PickingList struct {
	ID          string                            `json:"id"`
	DepotCode   string                            `json:"depot_code"`
	OrderNumber string                            `json:"order_number"`
	Status      warehouse_enums.PickingListStatus `json:"status"`
	AssignedTo  string                            `json:"assigned_to,omitempty"`
	Note        string                            `json:"note,omitempty"`
	Items       []PickingListItem                 `json:"items,omitempty"`
	History     []security.HistoryEntry           `json:"history,omitempty"`

	audit.AuditFields
}

type PickingListItem struct {
	ID                     string                               `json:"id"`
	PickingListID          string                               `json:"picking_list_id"`
	OrderItemID            string                               `json:"order_item_id"`
	SKUID                  string                               `json:"sku_id"`
	ProductName            string                               `json:"product_name,omitempty"`
	RequestedComposition   packaging.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition   packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	PickedComposition      packaging.PackageCompositionSnapshot `json:"picked_composition"`
	PackedComposition      packaging.PackageCompositionSnapshot `json:"packed_composition"`
	SubstitutedComposition packaging.PackageCompositionSnapshot `json:"substituted_composition"`
	ReturnedComposition    packaging.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition    packaging.PackageCompositionSnapshot `json:"refunded_composition"`
	Allocations            []PickingAllocation                  `json:"allocations,omitempty"`
	Substitutions          []PackageSubstitutionSnapshot        `json:"substitutions,omitempty"`
	Status                 warehouse_enums.PickingItemStatus    `json:"status"`
	CreatedAt              time.Time                            `json:"created_at"`
}

type PickingAllocation struct {
	ReservationAllocationID string                               `json:"reservation_allocation_id"`
	SourceBucketID          string                               `json:"source_bucket_id"`
	StockUnitIDs            []string                             `json:"stock_unit_ids,omitempty"`
	SourceLocation          warehouse.StockLocationRef           `json:"source_location"`
	LotID                   string                               `json:"lot_id,omitempty"`
	PackageOptionID         string                               `json:"package_option_id"`
	AllocatedComposition    packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	PickedComposition       packaging.PackageCompositionSnapshot `json:"picked_composition"`
	ScannedBarcodes         []string                             `json:"scanned_barcodes,omitempty"`
}
