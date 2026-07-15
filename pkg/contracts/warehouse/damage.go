package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/warehouse"
)

type DamageReport struct {
	ID             string                    `json:"id"`
	ProductSKUCode string                    `json:"product_sku_code"`
	ProductName    string                    `json:"product_name,omitempty"`
	DamagedQty     int                       `json:"damaged_qty"`
	LossValue      *common.Money             `json:"loss_value"`
	Stage          warehouseenum.DamageStage `json:"stage"`                  // e.g., Inbound, Picking, Packing, Storage
	ReferenceID    string                    `json:"reference_id,omitempty"` // ID of the InboundReceipt, PickingList, or OutboundShipment
	Note           string                    `json:"note,omitempty"`
	ReportedBy     string                    `json:"reported_by"`
	ReportedAt     time.Time                 `json:"reported_at"`
	History        []shared.HistoryEntry     `json:"history,omitempty"`
}
