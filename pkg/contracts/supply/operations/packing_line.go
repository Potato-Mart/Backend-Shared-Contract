package operations

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"
)

// PackingLine carries package compositions through fulfilment and returns.
type PackingLine struct {
	ID                     string                               `json:"id"`
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
	Substitutions          []PackageSubstitutionSnapshot        `json:"substitutions,omitempty"`
}
