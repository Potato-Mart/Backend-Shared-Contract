package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

type PickingList struct {
	ID          string                  `json:"id"`
	DepotID     string                  `json:"depot_id,omitempty"`
	OrderID     string                  `json:"order_id,omitempty"`
	OrderNumber string                  `json:"order_number,omitempty"`
	Status      enums.PickingListStatus `json:"status"`
	AssignedTo  string                  `json:"assigned_to,omitempty"`
	Note        string                  `json:"note,omitempty"`
	Items       []PickingListItem       `json:"items,omitempty"`
	History     []shared.HistoryEntry   `json:"history,omitempty"`

	common.AuditFields `bson:",inline"`
}

type PickingListItem struct {
	ID               string                  `json:"id"`
	PickingListID    string                  `json:"picking_list_id"`
	ProductID        string                  `json:"product_id"`
	ProductName      string                  `json:"product_name"`
	Barcode          string                  `json:"barcode,omitempty"`
	Location         string                  `json:"location,omitempty"`
	QuantityRequired int                     `json:"quantity_required"`
	QuantityPicked   int                     `json:"quantity_picked"`
	Status           enums.PickingItemStatus `json:"status"`
	CreatedAt        time.Time               `json:"created_at"`
}
