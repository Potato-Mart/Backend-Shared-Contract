package purchase

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/purchase/purchase_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse"
)

type Receipt struct {
	ID           string                             `json:"id"`
	OrderNumber  string                             `json:"order_number"`
	DepotCode    string                             `json:"depot_code,omitempty"`
	Reference    string                             `json:"reference,omitempty"`
	SupplierCode string                             `json:"supplier_code,omitempty"`
	Operator     string                             `json:"operator,omitempty"`
	Status       purchase_enums.PurchaseOrderStatus `json:"status"`
	ReceivedAt   *time.Time                         `json:"received_at,omitempty"`
	ConfirmedAt  *time.Time                         `json:"confirmed_at,omitempty"`
	Note         string                             `json:"note,omitempty"`
	Items        []ReceiptItem                      `json:"items"`
	History      []security.HistoryEntry            `json:"history,omitempty"`

	audit.AuditFields
}

type ReceiptItem struct {
	ID                  string                               `json:"id,omitempty"`
	ProductSKUCode      string                               `json:"product_sku_code"`
	ProductName         string                               `json:"product_name,omitempty"`
	PackageOptionID     string                               `json:"package_option_id"`
	LotID               string                               `json:"lot_id"`
	DestinationBucketID string                               `json:"destination_bucket_id"`
	DestinationLocation warehouse.StockLocationRef           `json:"destination_location"`
	DateMark            *warehouse.InventoryDateMark         `json:"date_mark,omitempty"`
	OrderedComposition  packaging.PackageCompositionSnapshot `json:"ordered_composition"`
	ReceivedComposition packaging.PackageCompositionSnapshot `json:"received_composition"`
	RejectedComposition packaging.PackageCompositionSnapshot `json:"rejected_composition"`
	Note                string                               `json:"note,omitempty"`
}
