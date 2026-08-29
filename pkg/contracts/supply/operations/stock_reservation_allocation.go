package operations

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
)

// StockReservationAllocation binds part of a logical reservation to exact
// inventory identity.
type StockReservationAllocation struct {
	ID                   string                               `json:"id"`
	ReservationID        string                               `json:"reservation_id"`
	BucketID             string                               `json:"bucket_id"`
	StockUnitIDs         []string                             `json:"stock_unit_ids,omitempty"`
	LotID                string                               `json:"lot_id,omitempty"`
	PackageOptionCode    string                               `json:"package_option_code"`
	HandlingUnit         packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	AllocatedComposition packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	Revision             int64                                `json:"revision"`

	audit.AuditFields
}
