package purchase

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/warehouse"
)

type Receipt struct {
	ID           string                  `json:"id"`
	OrderNumber  string                  `json:"order_number"`
	DepotCode    string                  `json:"depot_code,omitempty"`
	Reference    string                  `json:"reference,omitempty"`
	SupplierCode string                  `json:"supplier_code,omitempty"`
	Operator     string                  `json:"operator,omitempty"`
	Status       PurchaseOrderStatus     `json:"status"`
	ReceivedAt   *time.Time              `json:"received_at,omitempty"`
	ConfirmedAt  *time.Time              `json:"confirmed_at,omitempty"`
	Note         string                  `json:"note,omitempty"`
	Items        []ReceiptItem           `json:"items"`
	History      []security.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

type ReceiptItem struct {
	ID                  string                            `json:"id,omitempty"`
	ProductSKUCode      string                            `json:"product_sku_code"`
	ProductName         string                            `json:"product_name,omitempty"`
	PackageOptionID     string                            `json:"package_option_id"`
	LotID               string                            `json:"lot_id"`
	DestinationBucketID string                            `json:"destination_bucket_id"`
	DestinationLocation warehouse.StockLocationRef        `json:"destination_location"`
	DateMark            *warehouse.InventoryDateMark      `json:"date_mark,omitempty"`
	OrderedComposition  common.PackageCompositionSnapshot `json:"ordered_composition"`
	ReceivedComposition common.PackageCompositionSnapshot `json:"received_composition"`
	RejectedComposition common.PackageCompositionSnapshot `json:"rejected_composition"`
	Note                string                            `json:"note,omitempty"`
}
