package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// InventoryStockUnit identifies an individually labelled or evidenced stock
// unit while its bucket remains the quantity authority.
type InventoryStockUnit struct {
	ID                 string                               `json:"id"`
	BucketID           string                               `json:"bucket_id"`
	SKUCode            string                               `json:"sku_code"`
	LotID              string                               `json:"lot_id,omitempty"`
	PackageOptionCode  string                               `json:"package_option_code"`
	HandlingUnit       packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	BaseUnits          int64                                `json:"base_units"`
	Condition          warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition        warehouse_enums.InventoryDisposition `json:"disposition"`
	UnitLabelCode      string                               `json:"unit_label_code,omitempty"`
	ClearanceLabelCode string                               `json:"clearance_label_code,omitempty"`
	EvidenceMediaURLs  []string                             `json:"evidence_media_urls,omitempty"`

	audit.AuditFields
}
