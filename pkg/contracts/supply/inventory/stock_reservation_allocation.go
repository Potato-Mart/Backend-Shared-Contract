package inventory

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
)

// StockReservationAllocation binds part of a logical reservation to owner-verified
// physical stock. Exactly one identity variant is required. No-lot
// allocation evidence does not assert a lot, unit identity, verified cost or
// that its reserved units will still be physically available at picking time.
type StockReservationAllocation struct {
	ID                       string                               `json:"id"`
	ReservationID            string                               `json:"reservation_id"`
	BucketID                 string                               `json:"bucket_id"`
	KnownLotPhysical         *KnownLotPhysicalEvidence            `json:"known_lot_physical,omitempty"`
	UnknownIdentityPhysical  *UnknownReturnPhysicalEvidence       `json:"unknown_identity_physical,omitempty"`
	UnallocatedNoLotPhysical *UnallocatedNoLotPhysicalEvidence    `json:"unallocated_no_lot_physical,omitempty"`
	PackageOption            packaging.PackageOptionRef           `json:"package_option"`
	HandlingUnit             packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	AllocatedComposition     packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	Revision                 int64                                `json:"revision"`

	audit.AuditFields
}
