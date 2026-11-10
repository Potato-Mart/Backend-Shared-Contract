package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// WMSDraftItemMove tracks the stock move a WMS draft item submits, so a
// retried submit repeats the same request under the same idempotency code.
type WMSDraftItemMove struct {
	// Key is the fixed idempotency code assigned on first submit.
	Key  string                               `json:"key"`
	Step warehouse_enums.WMSDraftItemMoveStep `json:"step"`
	// Request is the body sent to the stock transfer or package conversion
	// route. It is stored opaquely; the contract does not describe its keys.
	Request metadata.Metadata `json:"request,omitempty"`
}
