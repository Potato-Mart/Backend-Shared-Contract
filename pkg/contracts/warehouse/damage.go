package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

type DamageReport struct {
	ID          string                `json:"id"`
	ProductID   string                `json:"product_id"`
	ProductName string                `json:"product_name,omitempty"`
	DamagedQty  int                   `json:"damaged_qty"`
	LossValue   *common.Money         `json:"loss_value"`
	Stage       enums.DamageStage     `json:"stage"`                  // e.g., Inbound, Picking, Packing, Storage
	ReferenceID string                `json:"reference_id,omitempty"` // ID of the InboundReceipt, PickingList, or OutboundShipment
	Note        string                `json:"note,omitempty"`
	ReportedBy  string                `json:"reported_by"`
	ReportedAt  time.Time             `json:"reported_at"`
	History     []shared.HistoryEntry `json:"history,omitempty"`
}
