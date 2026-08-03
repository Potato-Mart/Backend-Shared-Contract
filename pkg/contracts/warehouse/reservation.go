package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// StockReservation is a logical hold for one product at one depot.
type StockReservation struct {
	ID                   string                               `json:"id"`
	ProductSKUCode       string                               `json:"product_sku_code"`
	DepotCode            string                               `json:"depot_code"`
	OrderNumber          string                               `json:"order_number,omitempty"`
	GroupOrderCode       string                               `json:"group_order_code,omitempty"`
	OfferID              string                               `json:"offer_id"`
	OfferRevision        int64                                `json:"offer_revision"`
	RequestedComposition common.PackageCompositionSnapshot    `json:"requested_composition"`
	ReservedComposition  common.PackageCompositionSnapshot    `json:"reserved_composition"`
	Status               warehouseenum.StockReservationStatus `json:"status"`
	Revision             int64                                `json:"revision"`
	Timezone             string                               `json:"timezone"`
	ExpiresAt            *time.Time                           `json:"expires_at,omitempty"`

	common.AuditFields
}

// StockReservationAllocation binds part of a logical reservation to exact
// inventory identity.
type StockReservationAllocation struct {
	ID                   string                            `json:"id"`
	ReservationID        string                            `json:"reservation_id"`
	BucketID             string                            `json:"bucket_id"`
	StockUnitIDs         []string                          `json:"stock_unit_ids,omitempty"`
	LotID                string                            `json:"lot_id,omitempty"`
	PackageOptionID      string                            `json:"package_option_id"`
	HandlingUnit         common.PackageHandlingUnit        `json:"handling_unit"`
	AllocatedComposition common.PackageCompositionSnapshot `json:"allocated_composition"`
	Revision             int64                             `json:"revision"`

	common.AuditFields
}

// StockStagingRecord captures the physical transfer of an allocation into an
// online-order staging location.
type StockStagingRecord struct {
	ID                  string                            `json:"id"`
	ReservationID       string                            `json:"reservation_id"`
	AllocationID        string                            `json:"allocation_id"`
	OrderNumber         string                            `json:"order_number"`
	ProductSKUCode      string                            `json:"product_sku_code"`
	LotID               string                            `json:"lot_id,omitempty"`
	PackageOptionID     string                            `json:"package_option_id"`
	SourceBucketID      string                            `json:"source_bucket_id"`
	DestinationBucketID string                            `json:"destination_bucket_id"`
	SourceLocation      StockLocationRef                  `json:"source_location"`
	DestinationLocation StockLocationRef                  `json:"destination_location"`
	StagedComposition   common.PackageCompositionSnapshot `json:"staged_composition"`
	MovementID          string                            `json:"movement_id"`
	StagedBy            string                            `json:"staged_by"`
	StagedAt            time.Time                         `json:"staged_at"`
}
