package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

type PickingList struct {
	ID          string                          `json:"id"`
	DepotCode   string                          `json:"depot_code"`
	OrderNumber string                          `json:"order_number"`
	Status      warehouseenum.PickingListStatus `json:"status"`
	AssignedTo  string                          `json:"assigned_to,omitempty"`
	Note        string                          `json:"note,omitempty"`
	Items       []PickingListItem               `json:"items,omitempty"`
	History     []shared.HistoryEntry           `json:"history,omitempty"`

	common.AuditFields
}

type PickingListItem struct {
	ID                     string                            `json:"id"`
	PickingListID          string                            `json:"picking_list_id"`
	OrderItemID            string                            `json:"order_item_id"`
	ProductSKUCode         string                            `json:"product_sku_code"`
	ProductName            string                            `json:"product_name,omitempty"`
	RequestedComposition   common.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition   common.PackageCompositionSnapshot `json:"allocated_composition"`
	PickedComposition      common.PackageCompositionSnapshot `json:"picked_composition"`
	PackedComposition      common.PackageCompositionSnapshot `json:"packed_composition"`
	SubstitutedComposition common.PackageCompositionSnapshot `json:"substituted_composition"`
	ReturnedComposition    common.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition    common.PackageCompositionSnapshot `json:"refunded_composition"`
	Allocations            []PickingAllocation               `json:"allocations,omitempty"`
	Substitutions          []PackageSubstitutionSnapshot     `json:"substitutions,omitempty"`
	Status                 warehouseenum.PickingItemStatus   `json:"status"`
	CreatedAt              time.Time                         `json:"created_at"`
}

type PickingAllocation struct {
	ReservationAllocationID string                            `json:"reservation_allocation_id"`
	SourceBucketID          string                            `json:"source_bucket_id"`
	StockUnitIDs            []string                          `json:"stock_unit_ids,omitempty"`
	SourceLocation          StockLocationRef                  `json:"source_location"`
	LotID                   string                            `json:"lot_id,omitempty"`
	PackageOptionID         string                            `json:"package_option_id"`
	AllocatedComposition    common.PackageCompositionSnapshot `json:"allocated_composition"`
	PickedComposition       common.PackageCompositionSnapshot `json:"picked_composition"`
	ScannedBarcodes         []string                          `json:"scanned_barcodes,omitempty"`
}
