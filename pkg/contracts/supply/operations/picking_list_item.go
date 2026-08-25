package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

type PickingListItem struct {
	ID                     string                               `json:"id"`
	PickingListID          string                               `json:"picking_list_id"`
	OrderItemID            string                               `json:"order_item_id"`
	SKUCode                string                               `json:"sku_code"`
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
