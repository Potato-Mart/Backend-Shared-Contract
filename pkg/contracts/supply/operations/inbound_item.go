package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse"
)

type InboundItem struct {
	ID                  string                               `json:"id"`
	SKUCode             string                               `json:"sku_code"`
	ProductName         string                               `json:"product_name,omitempty"`
	ScannedBarcode      string                               `json:"scanned_barcode,omitempty"`
	LotID               string                               `json:"lot_id,omitempty"`
	PackageOptionCode   string                               `json:"package_option_code"`
	HandlingUnit        packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	StorageType         classification_enums.StorageType     `json:"storage_type"`
	ExpectedComposition packaging.PackageCompositionSnapshot `json:"expected_composition"`
	ReceivedComposition packaging.PackageCompositionSnapshot `json:"received_composition"`
	DestinationLocation warehouse.StockLocationRef           `json:"destination_location"`
	DestinationBucketID string                               `json:"destination_bucket_id,omitempty"`
	CreatedAt           time.Time                            `json:"created_at"`
}
