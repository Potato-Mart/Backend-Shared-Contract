package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

type PickingList struct {
	ID        string `json:"id"`
	DepotCode string `json:"depot_code,omitempty"`
	// Deprecated: use DepotCode.
	DepotID     string `json:"depot_id,omitempty"`
	OrderNumber string `json:"order_number,omitempty"`
	// Deprecated: use OrderNumber.
	OrderID    string                  `json:"order_id,omitempty"`
	Status     enums.PickingListStatus `json:"status"`
	AssignedTo string                  `json:"assigned_to,omitempty"`
	Note       string                  `json:"note,omitempty"`
	Items      []PickingListItem       `json:"items,omitempty"`
	History    []shared.HistoryEntry   `json:"history,omitempty"`

	common.AuditFields `bson:",inline"`
}

type PickingListItem struct {
	ID             string `json:"id"`
	PickingListID  string `json:"picking_list_id"`
	ProductSKUCode string `json:"product_sku_code"`
	// Deprecated: use ProductSKUCode.
	ProductID        string                  `json:"product_id,omitempty"`
	ProductName      string                  `json:"product_name"`
	Barcode          string                  `json:"barcode,omitempty"`
	Location         string                  `json:"location,omitempty"`
	QuantityRequired int                     `json:"quantity_required"`
	QuantityPicked   int                     `json:"quantity_picked"`
	Status           enums.PickingItemStatus `json:"status"`
	CreatedAt        time.Time               `json:"created_at"`
}
