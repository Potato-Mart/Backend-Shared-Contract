package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

// InventoryDateMark is a timezone-qualified date mark attached to a lot.
type InventoryDateMark struct {
	Kind       warehouse_enums.InventoryDateMarkKind `json:"kind"`
	DateMarkAt time.Time                             `json:"date_mark_at"`
	Timezone   string                                `json:"timezone"`
}
