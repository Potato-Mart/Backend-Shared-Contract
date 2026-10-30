package possettlement

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"

// Line preserves one original priced package grain. Reserved lines carry the
// complete authenticated native allocation; direct merchandise has none and
// cannot prescribe a bucket, lot or physical unit to consume.
type Line struct {
	OrderItemID   string                               `json:"order_item_id"`
	QuoteLineID   string                               `json:"quote_line_id"`
	SKUCode       string                               `json:"sku_code"`
	PackageOption packaging.PackageOptionRef           `json:"package_option"`
	Composition   packaging.PackageCompositionSnapshot `json:"composition"`
	SourceMode    string                               `json:"source_mode"`
	Reservations  []ReservationAllocation              `json:"reservations,omitempty"`
}
