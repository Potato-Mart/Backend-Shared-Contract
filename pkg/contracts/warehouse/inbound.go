package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

type InboundReceipt struct {
	ID           string                             `json:"id"`
	DepotCode    string                             `json:"depot_code"`
	Reference    string                             `json:"reference,omitempty"`
	SupplierCode string                             `json:"supplier_code,omitempty"`
	ETA          *time.Time                         `json:"eta,omitempty"`
	Operator     string                             `json:"operator,omitempty"`
	Status       warehouseenum.InboundReceiptStatus `json:"status"`
	Items        []InboundItem                      `json:"items"`
	Note         string                             `json:"note,omitempty"`
	ConfirmedAt  *time.Time                         `json:"confirmed_at,omitempty"`
	History      []shared.HistoryEntry              `json:"history,omitempty"`

	common.AuditFields
}

type InboundItem struct {
	ID                  string                            `json:"id"`
	ProductSKUCode      string                            `json:"product_sku_code"`
	ProductName         string                            `json:"product_name,omitempty"`
	ScannedBarcode      string                            `json:"scanned_barcode,omitempty"`
	LotID               string                            `json:"lot_id,omitempty"`
	PackageOptionID     string                            `json:"package_option_id"`
	HandlingUnit        common.PackageHandlingUnit        `json:"handling_unit"`
	StorageType         warehouseenum.StorageType         `json:"storage_type"`
	ExpectedComposition common.PackageCompositionSnapshot `json:"expected_composition"`
	ReceivedComposition common.PackageCompositionSnapshot `json:"received_composition"`
	DestinationLocation StockLocationRef                  `json:"destination_location"`
	DestinationBucketID string                            `json:"destination_bucket_id,omitempty"`
	CreatedAt           time.Time                         `json:"created_at"`
}
