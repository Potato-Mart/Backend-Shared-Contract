package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

// StockReservation is a logical hold for one product at one depot.
type StockReservation struct {
	ID                   string                                 `json:"id"`
	ProductSKUCode       string                                 `json:"product_sku_code"`
	DepotCode            string                                 `json:"depot_code"`
	OrderNumber          string                                 `json:"order_number,omitempty"`
	GroupOrderCode       string                                 `json:"group_order_code,omitempty"`
	OfferID              string                                 `json:"offer_id"`
	OfferRevision        int64                                  `json:"offer_revision"`
	RequestedComposition packaging.PackageCompositionSnapshot   `json:"requested_composition"`
	ReservedComposition  packaging.PackageCompositionSnapshot   `json:"reserved_composition"`
	Status               warehouse_enums.StockReservationStatus `json:"status"`
	Revision             int64                                  `json:"revision"`
	Timezone             string                                 `json:"timezone"`
	ExpiresAt            *time.Time                             `json:"expires_at,omitempty"`

	audit.AuditFields
}

// StockReservationAllocation binds part of a logical reservation to exact
// inventory identity.
type StockReservationAllocation struct {
	ID                   string                               `json:"id"`
	ReservationID        string                               `json:"reservation_id"`
	BucketID             string                               `json:"bucket_id"`
	StockUnitIDs         []string                             `json:"stock_unit_ids,omitempty"`
	LotID                string                               `json:"lot_id,omitempty"`
	PackageOptionID      string                               `json:"package_option_id"`
	HandlingUnit         packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	AllocatedComposition packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	Revision             int64                                `json:"revision"`

	audit.AuditFields
}

// StockStagingRecord captures the physical transfer of an allocation into an
// online-order staging location.
type StockStagingRecord struct {
	ID                  string                               `json:"id"`
	ReservationID       string                               `json:"reservation_id"`
	AllocationID        string                               `json:"allocation_id"`
	OrderNumber         string                               `json:"order_number"`
	ProductSKUCode      string                               `json:"product_sku_code"`
	LotID               string                               `json:"lot_id,omitempty"`
	PackageOptionID     string                               `json:"package_option_id"`
	SourceBucketID      string                               `json:"source_bucket_id"`
	DestinationBucketID string                               `json:"destination_bucket_id"`
	SourceLocation      StockLocationRef                     `json:"source_location"`
	DestinationLocation StockLocationRef                     `json:"destination_location"`
	StagedComposition   packaging.PackageCompositionSnapshot `json:"staged_composition"`
	MovementID          string                               `json:"movement_id"`
	StagedBy            string                               `json:"staged_by"`
	StagedAt            time.Time                            `json:"staged_at"`
}
