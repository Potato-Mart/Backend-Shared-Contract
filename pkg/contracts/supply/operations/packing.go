package operations

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackageSubstitutionSnapshot records the exact loose-item replacement evidence
// for one requested sealed-case substitution.
type PackageSubstitutionSnapshot struct {
	ID                             string    `json:"id"`
	RequestedCasePackageOptionID   string    `json:"requested_case_package_option_id"`
	RequestedCaseCount             int64     `json:"requested_case_count"`
	RequestedUnitsPerCase          int64     `json:"requested_units_per_case"`
	FulfilledSealedCaseCount       int64     `json:"fulfilled_sealed_case_count"`
	ReplacementEachPackageOptionID string    `json:"replacement_each_package_option_id"`
	ReplacementBaseUnits           int64     `json:"replacement_base_units"`
	LotID                          string    `json:"lot_id"`
	SourceBucketID                 string    `json:"source_bucket_id"`
	StockUnitIDs                   []string  `json:"stock_unit_ids,omitempty"`
	ReasonCode                     string    `json:"reason_code"`
	Operator                       string    `json:"operator"`
	CapturedAt                     time.Time `json:"captured_at"`
}

// PackingLine carries package compositions through fulfilment and returns.
type PackingLine struct {
	ID                     string                               `json:"id"`
	OrderItemID            string                               `json:"order_item_id"`
	SKUID                  string                               `json:"sku_id"`
	ProductName            string                               `json:"product_name,omitempty"`
	RequestedComposition   packaging.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition   packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	PickedComposition      packaging.PackageCompositionSnapshot `json:"picked_composition"`
	PackedComposition      packaging.PackageCompositionSnapshot `json:"packed_composition"`
	SubstitutedComposition packaging.PackageCompositionSnapshot `json:"substituted_composition"`
	ReturnedComposition    packaging.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition    packaging.PackageCompositionSnapshot `json:"refunded_composition"`
	Substitutions          []PackageSubstitutionSnapshot        `json:"substitutions,omitempty"`
}

// PackingDamage links packing evidence to canonical inventory assessment and
// movement identities.
type PackingDamage struct {
	ID                   string                                `json:"id"`
	SKUID                string                                `json:"sku_id"`
	SourceBucketID       string                                `json:"source_bucket_id"`
	StockUnitID          string                                `json:"stock_unit_id,omitempty"`
	QualityAssessmentID  string                                `json:"quality_assessment_id"`
	AffectedComposition  packaging.PackageCompositionSnapshot  `json:"affected_composition"`
	Handling             warehouse_enums.PackingDamageHandling `json:"handling"`
	Note                 string                                `json:"note,omitempty"`
	ResultingMovementIDs []string                              `json:"resulting_movement_ids,omitempty"`
	CreatedAt            time.Time                             `json:"created_at"`
	CreatedBy            string                                `json:"created_by,omitempty"`
}

// OutboundContainerPlan describes one outbound shipping container and its
// package-aware contents.
type OutboundContainerPlan struct {
	ID                string                      `json:"id"`
	ContainerCode     string                      `json:"container_code"`
	StorageType       warehouse_enums.StorageType `json:"storage_type"`
	Contents          []OutboundContainerContent  `json:"contents,omitempty"`
	IsManuallyPlanned bool                        `json:"is_manually_planned"`
	UpdatedAt         time.Time                   `json:"updated_at"`
}

// OutboundContainerContent identifies inventory packed into one outbound
// container.
type OutboundContainerContent struct {
	OrderItemID       string                               `json:"order_item_id"`
	SKUID             string                               `json:"sku_id"`
	AllocationID      string                               `json:"allocation_id"`
	BucketID          string                               `json:"bucket_id"`
	LotID             string                               `json:"lot_id,omitempty"`
	PackageOptionID   string                               `json:"package_option_id"`
	PackedComposition packaging.PackageCompositionSnapshot `json:"packed_composition"`
	Substitutions     []PackageSubstitutionSnapshot        `json:"substitutions,omitempty"`
}
