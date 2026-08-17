package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging/packaging_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

type InboundReceipt struct {
	ID           string                               `json:"id"`
	DepotCode    string                               `json:"depot_code"`
	Reference    string                               `json:"reference,omitempty"`
	SupplierCode string                               `json:"supplier_code,omitempty"`
	ETA          *time.Time                           `json:"eta,omitempty"`
	Operator     string                               `json:"operator,omitempty"`
	Status       warehouse_enums.InboundReceiptStatus `json:"status"`
	Items        []InboundItem                        `json:"items"`
	Note         string                               `json:"note,omitempty"`
	ConfirmedAt  *time.Time                           `json:"confirmed_at,omitempty"`
	History      []security.HistoryEntry              `json:"history,omitempty"`

	audit.AuditFields
}

type InboundItem struct {
	ID                  string                               `json:"id"`
	SKUID               string                               `json:"sku_id"`
	ProductName         string                               `json:"product_name,omitempty"`
	ScannedBarcode      string                               `json:"scanned_barcode,omitempty"`
	LotID               string                               `json:"lot_id,omitempty"`
	PackageOptionID     string                               `json:"package_option_id"`
	HandlingUnit        packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	StorageType         warehouse_enums.StorageType          `json:"storage_type"`
	ExpectedComposition packaging.PackageCompositionSnapshot `json:"expected_composition"`
	ReceivedComposition packaging.PackageCompositionSnapshot `json:"received_composition"`
	DestinationLocation warehouse.StockLocationRef           `json:"destination_location"`
	DestinationBucketID string                               `json:"destination_bucket_id,omitempty"`
	CreatedAt           time.Time                            `json:"created_at"`
}
