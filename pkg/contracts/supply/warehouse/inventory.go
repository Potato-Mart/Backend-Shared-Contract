package warehouse

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// InventoryDateMark is a timezone-qualified date mark attached to a lot.
type InventoryDateMark struct {
	Kind       InventoryDateMarkKind `json:"kind"`
	DateMarkAt time.Time             `json:"date_mark_at"`
	Timezone   string                `json:"timezone"`
}

// InventoryLot identifies inventory received or manufactured together.
type InventoryLot struct {
	ID                  string             `json:"id"`
	ProductSKUCode      string             `json:"product_sku_code"`
	SupplierLotCode     string             `json:"supplier_lot_code,omitempty"`
	ManufacturerLotCode string             `json:"manufacturer_lot_code,omitempty"`
	ReceivedAt          time.Time          `json:"received_at"`
	ManufacturedAt      *time.Time         `json:"manufactured_at,omitempty"`
	DateMark            *InventoryDateMark `json:"date_mark,omitempty"`

	common.AuditFields
}

// InventoryStockBucket is the quantity authority for one package form,
// location, lot, condition, and disposition combination. A CASE bucket
// represents intact cases and an EACH bucket represents loose base units.
type InventoryStockBucket struct {
	ID                 string                            `json:"id"`
	Location           StockLocationRef                  `json:"location"`
	ProductSKUCode     string                            `json:"product_sku_code"`
	LotID              string                            `json:"lot_id,omitempty"`
	PackageOptionID    string                            `json:"package_option_id"`
	HandlingUnit       common.PackageHandlingUnit        `json:"handling_unit"`
	Condition          InventoryCondition                `json:"condition"`
	Disposition        InventoryDisposition              `json:"disposition"`
	PackageComposition common.PackageCompositionSnapshot `json:"package_composition"`
	OnHandBaseUnits    int64                             `json:"on_hand_base_units"`
	ReservedBaseUnits  int64                             `json:"reserved_base_units"`
	// AvailableBaseUnits is a derived JSON projection for this bucket.
	AvailableBaseUnits int64     `json:"available_base_units"`
	Revision           int64     `json:"revision"`
	DepotTimezone      string    `json:"depot_timezone"`
	AsOf               time.Time `json:"as_of"`

	common.AuditFields
}

// InventoryStockUnit identifies an individually labelled or evidenced stock
// unit while its bucket remains the quantity authority.
type InventoryStockUnit struct {
	ID                 string                     `json:"id"`
	BucketID           string                     `json:"bucket_id"`
	ProductSKUCode     string                     `json:"product_sku_code"`
	LotID              string                     `json:"lot_id,omitempty"`
	PackageOptionID    string                     `json:"package_option_id"`
	HandlingUnit       common.PackageHandlingUnit `json:"handling_unit"`
	BaseUnits          int64                      `json:"base_units"`
	Condition          InventoryCondition         `json:"condition"`
	Disposition        InventoryDisposition       `json:"disposition"`
	UnitLabelCode      string                     `json:"unit_label_code,omitempty"`
	ClearanceLabelCode string                     `json:"clearance_label_code,omitempty"`
	EvidenceMediaURLs  []string                   `json:"evidence_media_urls,omitempty"`

	common.AuditFields
}

// QualityAssessment captures an observed condition/disposition decision and
// the resulting physical inventory movements.
type QualityAssessment struct {
	ID                   string                            `json:"id"`
	ProductSKUCode       string                            `json:"product_sku_code"`
	BucketID             string                            `json:"bucket_id"`
	StockUnitID          string                            `json:"stock_unit_id,omitempty"`
	AssessedComposition  common.PackageCompositionSnapshot `json:"assessed_composition"`
	PreviousCondition    InventoryCondition                `json:"previous_condition"`
	ResultCondition      InventoryCondition                `json:"result_condition"`
	PreviousDisposition  InventoryDisposition              `json:"previous_disposition"`
	ResultDisposition    InventoryDisposition              `json:"result_disposition"`
	AssessedBy           string                            `json:"assessed_by"`
	ReasonCode           string                            `json:"reason_code"`
	Note                 string                            `json:"note,omitempty"`
	EvidenceMediaURLs    []string                          `json:"evidence_media_urls,omitempty"`
	AssessedAt           time.Time                         `json:"assessed_at"`
	ResultingMovementIDs []string                          `json:"resulting_movement_ids,omitempty"`
}
