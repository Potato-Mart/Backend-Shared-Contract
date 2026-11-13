package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
)

// InventoryLot identifies inventory received or manufactured together.
type InventoryLot struct {
	ID                  string             `json:"id"`
	SKUCode             string             `json:"sku_code"`
	SupplierLotCode     string             `json:"supplier_lot_code,omitempty"`
	ManufacturerLotCode string             `json:"manufacturer_lot_code,omitempty"`
	ReceivedAt          time.Time          `json:"received_at"`
	ManufacturedAt      *time.Time         `json:"manufactured_at,omitempty"`
	DateMark            *InventoryDateMark `json:"date_mark,omitempty"`
	// OriginLotID is set on a lot split off from a received lot at put-away,
	// so receipt, purchase-order, supplier and cost linkage can be traced to
	// the origin lot.
	OriginLotID string `json:"origin_lot_id,omitempty"`
	// PastDateMark warns that the lot is at or past its date mark. Supply
	// sets it on lot list reads; it is omitted for an in-date or undated lot.
	// Information only: whether the lot may be sold is the
	// sell-past-date-mark policy's answer.
	PastDateMark *PastDateMarkNotice `json:"past_date_mark,omitempty"`

	audit.AuditFields
}
