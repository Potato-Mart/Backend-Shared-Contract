package supply

import (
	"errors"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pubsub/supply/supply_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/inventory"
	"math"
	"strings"
)

// InventorySaleCommittedEvent records one physical issue against a reservation.
// Exactly one physical evidence variant is required. Unknown returns always
// have unresolved cost; known-lot and unallocated no-lot stock are valued
// only when their actual movement cost is verified.
// Unidentified tracked issues omit bucket and revision and have unresolved cost.
// Reservation shortages are releases, not sale or deficit movements.
type InventorySaleCommittedEvent struct {
	MovementID                  string                                         `json:"movement_id"`
	OrderNumber                 string                                         `json:"order_number"`
	DepotCode                   string                                         `json:"depot_code"`
	ReservationID               string                                         `json:"reservation_id"`
	AllocationID                string                                         `json:"allocation_id"`
	BucketID                    string                                         `json:"bucket_id,omitempty"`
	SKUCode                     string                                         `json:"sku_code"`
	PackageOption               packaging.PackageOptionRef                     `json:"package_option"`
	CommittedComposition        packaging.PackageCompositionSnapshot           `json:"committed_composition"`
	InventoryRevision           int64                                          `json:"inventory_revision,omitempty"`
	CostStatus                  supply_enums.InventorySaleCostStatus           `json:"cost_status"`
	KnownLotPhysical            *inventory.KnownLotPhysicalEvidence            `json:"known_lot_physical,omitempty"`
	UnknownIdentityPhysical     *inventory.UnknownReturnPhysicalEvidence       `json:"unknown_identity_physical,omitempty"`
	UnallocatedNoLotPhysical    *inventory.UnallocatedNoLotPhysicalEvidence    `json:"unallocated_no_lot_physical,omitempty"`
	UnidentifiedTrackedPhysical *inventory.UnidentifiedTrackedPhysicalEvidence `json:"unidentified_tracked_physical,omitempty"`
}

// ValidatePhysicalEvidence checks the v7 wire invariants. It does not replace
// Supply's independent proof of settlement, source, quantity or valuation.
func (e InventorySaleCommittedEvent) ValidatePhysicalEvidence() error {
	clean := func(s string) bool { return s != "" && strings.TrimSpace(s) == s }
	if !clean(e.MovementID) || !clean(e.OrderNumber) || !clean(e.DepotCode) || !clean(e.ReservationID) || !clean(e.AllocationID) || !inventory.ValidPhysicalIssueComposition(e.SKUCode, e.PackageOption, e.CommittedComposition) || !e.CostStatus.IsValid() {
		return errors.New("invalid sale identity or composition")
	}
	count := 0
	if e.KnownLotPhysical != nil {
		count++
	}
	if e.UnknownIdentityPhysical != nil {
		count++
	}
	if e.UnallocatedNoLotPhysical != nil {
		count++
	}
	if e.UnidentifiedTrackedPhysical != nil {
		count++
	}
	if count != 1 {
		return errors.New("exactly one physical evidence variant is required")
	}
	if e.UnidentifiedTrackedPhysical != nil {
		if e.BucketID != "" || e.InventoryRevision != 0 || e.CostStatus != supply_enums.InventorySaleCostStatusUnresolved {
			return errors.New("unidentified tracked issue claims source or cost")
		}
		return nil
	}
	if !clean(e.BucketID) || e.InventoryRevision <= 0 {
		return errors.New("verified physical source and revision are required")
	}
	if e.KnownLotPhysical != nil {
		if !clean(e.KnownLotPhysical.LotID) {
			return errors.New("known lot is required")
		}
		seen := map[string]bool{}
		for _, id := range e.KnownLotPhysical.StockUnitIDs {
			if !clean(id) || seen[id] {
				return errors.New("invalid unit identity")
			}
			seen[id] = true
		}
	}
	if e.UnallocatedNoLotPhysical != nil && !clean(e.UnallocatedNoLotPhysical.SourceBucketID) {
		return errors.New("original no-lot bucket is required")
	}
	if e.UnknownIdentityPhysical != nil {
		if e.CostStatus != supply_enums.InventorySaleCostStatusUnresolved || len(e.UnknownIdentityPhysical.ReturnSources) == 0 {
			return errors.New("unknown return requires source evidence and unresolved cost")
		}
		var total int64
		seen := map[string]bool{}
		for _, source := range e.UnknownIdentityPhysical.ReturnSources {
			if !clean(source.ReceivingID) || seen[source.ReceivingID] || source.BaseUnits <= 0 || total > math.MaxInt64-source.BaseUnits {
				return errors.New("invalid return source")
			}
			seen[source.ReceivingID] = true
			total += source.BaseUnits
		}
		if total != e.CommittedComposition.TotalBaseUnits {
			return errors.New("return source quantity mismatch")
		}
	}
	return nil
}
