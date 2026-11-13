package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement/purchase_enums"
)

// PurchaseOrderItemReplacement records one replacement of part of a
// purchase-order line's outstanding quantity by another line.
// ReplacedComposition is the quantity of the original line that was replaced,
// and ReplacementItemID is the purchase-order line that now carries it.
type PurchaseOrderItemReplacement struct {
	ID                  string                                            `json:"id"`
	ReplacementItemID   string                                            `json:"replacement_item_id"`
	ReplacedComposition packaging.PackageCompositionSnapshot              `json:"replaced_composition"`
	ReasonCode          purchase_enums.PurchaseOrderItemReplacementReason `json:"reason_code"`
	Note                string                                            `json:"note,omitempty"`
	ReplacedBy          string                                            `json:"replaced_by"`
	ReplacedAt          time.Time                                         `json:"replaced_at"`
}
