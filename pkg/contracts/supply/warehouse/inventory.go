package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/warehouse/warehouse_enums"
)

// InventoryDateMark is a timezone-qualified date mark attached to a lot.
type InventoryDateMark struct {
	Kind       warehouse_enums.InventoryDateMarkKind `json:"kind"`
	DateMarkAt time.Time                             `json:"date_mark_at"`
	Timezone   string                                `json:"timezone"`
}

// InventoryLot identifies inventory received or manufactured together.
type InventoryLot struct {
	ID                  string             `json:"id"`
	SKUCode             string             `json:"sku_code"`
	SupplierLotCode     string             `json:"supplier_lot_code,omitempty"`
	ManufacturerLotCode string             `json:"manufacturer_lot_code,omitempty"`
	ReceivedAt          time.Time          `json:"received_at"`
	ManufacturedAt      *time.Time         `json:"manufactured_at,omitempty"`
	DateMark            *InventoryDateMark `json:"date_mark,omitempty"`

	audit.AuditFields
}

// InventoryStockBucket is the quantity authority for one package form,
// location, lot, condition, and disposition combination. A CASE bucket
// represents intact cases and an EACH bucket represents loose base units.
type InventoryStockBucket struct {
	ID                 string                               `json:"id"`
	Location           StockLocationRef                     `json:"location"`
	SKUCode            string                               `json:"sku_code"`
	LotID              string                               `json:"lot_id,omitempty"`
	PackageOptionCode  string                               `json:"package_option_code"`
	HandlingUnit       packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	Condition          warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition        warehouse_enums.InventoryDisposition `json:"disposition"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	OnHandBaseUnits    int64                                `json:"on_hand_base_units"`
	ReservedBaseUnits  int64                                `json:"reserved_base_units"`
	// AvailableBaseUnits is a derived JSON projection for this bucket.
	AvailableBaseUnits int64     `json:"available_base_units"`
	Revision           int64     `json:"revision"`
	DepotTimezone      string    `json:"depot_timezone"`
	AsOf               time.Time `json:"as_of"`

	audit.AuditFields
}

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

// QualityAssessment captures an observed condition/disposition decision and
// the resulting physical inventory movements.
type QualityAssessment struct {
	ID                   string                               `json:"id"`
	SKUCode              string                               `json:"sku_code"`
	BucketID             string                               `json:"bucket_id"`
	StockUnitID          string                               `json:"stock_unit_id,omitempty"`
	AssessedComposition  packaging.PackageCompositionSnapshot `json:"assessed_composition"`
	PreviousCondition    warehouse_enums.InventoryCondition   `json:"previous_condition"`
	ResultCondition      warehouse_enums.InventoryCondition   `json:"result_condition"`
	PreviousDisposition  warehouse_enums.InventoryDisposition `json:"previous_disposition"`
	ResultDisposition    warehouse_enums.InventoryDisposition `json:"result_disposition"`
	AssessedBy           string                               `json:"assessed_by"`
	ReasonCode           string                               `json:"reason_code"`
	Note                 string                               `json:"note,omitempty"`
	EvidenceMediaURLs    []string                             `json:"evidence_media_urls,omitempty"`
	AssessedAt           time.Time                            `json:"assessed_at"`
	ResultingMovementIDs []string                             `json:"resulting_movement_ids,omitempty"`
}
