package operations

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
)

// OutboundContainerContent identifies inventory packed into one outbound
// container.
type OutboundContainerContent struct {
	OrderItemID       string                               `json:"order_item_id"`
	SKUCode           string                               `json:"sku_code"`
	AllocationID      string                               `json:"allocation_id"`
	BucketID          string                               `json:"bucket_id"`
	LotID             string                               `json:"lot_id,omitempty"`
	PackageOptionCode string                               `json:"package_option_code"`
	PackedComposition packaging.PackageCompositionSnapshot `json:"packed_composition"`
	Substitutions     []PackageSubstitutionSnapshot        `json:"substitutions,omitempty"`
}
