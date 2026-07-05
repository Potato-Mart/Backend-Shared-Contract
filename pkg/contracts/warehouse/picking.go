package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

type PickingList struct {
	ID          string                  `json:"id"`
	DepotCode   string                  `json:"depot_code,omitempty"`
	OrderNumber string                  `json:"order_number,omitempty"`
	Status      enums.PickingListStatus `json:"status"`
	AssignedTo  string                  `json:"assigned_to,omitempty"`
	Note        string                  `json:"note,omitempty"`
	Items       []PickingListItem       `json:"items,omitempty"`
	History     []shared.HistoryEntry   `json:"history,omitempty"`

	common.AuditFields
}

type PickingListItem struct {
	ID               string                  `json:"id"`
	PickingListID    string                  `json:"picking_list_id"`
	ProductSKUCode   string                  `json:"product_sku_code"`
	ProductName      string                  `json:"product_name"`
	Barcode          string                  `json:"barcode,omitempty"`
	Location         string                  `json:"location,omitempty"`
	QuantityRequired int                     `json:"quantity_required"`
	QuantityPicked   int                     `json:"quantity_picked"`
	Status           enums.PickingItemStatus `json:"status"`
	CreatedAt        time.Time               `json:"created_at"`
}
