package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/inventory"

// POSUnknownIdentityPhysicalEvidence identifies physically present stock from
// recorded returns without claiming a lot, unit identity or valuation. Sources
// are unique within the movement and their quantities sum to event base_units;
// Supply verifies aggregate source consumption across movements.
type POSUnknownIdentityPhysicalEvidence struct {
	ReturnSources []inventory.UnknownReturnSource `json:"return_sources"`
}
