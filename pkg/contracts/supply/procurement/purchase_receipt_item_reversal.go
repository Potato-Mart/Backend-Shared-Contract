package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
)

// PurchaseReceiptItemReversal records one reversal of part or all of a
// received purchase-receipt line. Composition is what this reversal took back
// from the line. ReversalID is shared by every line reversed in the same
// request, and Reason is that request's reason.
type PurchaseReceiptItemReversal struct {
	ReversalID  string                               `json:"reversal_id"`
	Composition packaging.PackageCompositionSnapshot `json:"composition"`
	Reason      string                               `json:"reason"`
	ReversedAt  time.Time                            `json:"reversed_at"`
	ReversedBy  string                               `json:"reversed_by"`
}
