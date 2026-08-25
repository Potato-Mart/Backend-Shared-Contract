package operations

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
)

type PickingList struct {
	ID          string                            `json:"id"`
	DepotCode   string                            `json:"depot_code"`
	MarketCode  string                            `json:"market_code,omitempty"`
	OrderNumber string                            `json:"order_number"`
	Status      warehouse_enums.PickingListStatus `json:"status"`
	AssignedTo  string                            `json:"assigned_to,omitempty"`
	Note        string                            `json:"note,omitempty"`
	Items       []PickingListItem                 `json:"items,omitempty"`
	History     []security.HistoryEntry           `json:"history,omitempty"`

	audit.AuditFields
}
