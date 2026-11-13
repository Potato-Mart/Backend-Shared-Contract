package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PastDateMarkNotice tells a client that the stock offered, or about to be
// sold, includes units at or past their date mark. It is information for a
// warning and never a refusal. Kind, DateMarkAt and Timezone are those of the
// earliest such date mark. PastDateBaseUnits is present only where a quantity
// was requested: how many of the requested base units come from past-date
// stock. The notice is omitted when no past-date stock is involved.
type PastDateMarkNotice struct {
	Kind              warehouse_enums.InventoryDateMarkKind `json:"kind"`
	DateMarkAt        time.Time                             `json:"date_mark_at"`
	Timezone          string                                `json:"timezone"`
	PastDateBaseUnits *int64                                `json:"past_date_base_units,omitempty"`
}
