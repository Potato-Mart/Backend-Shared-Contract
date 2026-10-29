package fulfilment

import (
	"errors"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/inventory"
	"strings"
)

// OutboundContainerContent identifies inventory packed into one outbound
// container. An unidentified tracked item has no asserted physical bucket or lot;
// its allocation reference identifies only the original reservation.
type OutboundContainerContent struct {
	OrderItemID                 string                                         `json:"order_item_id"`
	SKUCode                     string                                         `json:"sku_code"`
	AllocationID                string                                         `json:"allocation_id"`
	BucketID                    string                                         `json:"bucket_id,omitempty"`
	UnidentifiedTrackedPhysical *inventory.UnidentifiedTrackedPhysicalEvidence `json:"unidentified_tracked_physical,omitempty"`
	LotID                       string                                         `json:"lot_id,omitempty"`
	PackageOption               packaging.PackageOptionRef                     `json:"package_option"`
	PackedComposition           packaging.PackageCompositionSnapshot           `json:"packed_composition"`
	Substitutions               []PackageSubstitutionSnapshot                  `json:"substitutions,omitempty"`
}

// ValidatePhysicalEvidence checks packing identity and quantity shape. Supply
// must independently verify the reservation and any asserted physical source.
func (c OutboundContainerContent) ValidatePhysicalEvidence() error {
	clean := func(s string) bool { return s != "" && strings.TrimSpace(s) == s }
	if !clean(c.OrderItemID) || !clean(c.AllocationID) || !inventory.ValidPhysicalIssueComposition(c.SKUCode, c.PackageOption, c.PackedComposition) {
		return errors.New("invalid packing identity or composition")
	}
	if c.UnidentifiedTrackedPhysical != nil {
		if c.BucketID != "" || c.LotID != "" || len(c.Substitutions) != 0 {
			return errors.New("unidentified tracked packing claims a source, lot or substitution")
		}
		return nil
	}
	if !clean(c.BucketID) || (c.LotID != "" && !clean(c.LotID)) {
		return errors.New("verified packing source is required")
	}
	return nil
}
