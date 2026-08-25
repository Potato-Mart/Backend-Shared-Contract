package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
)

// DepotMarket associates one physical depot with one commercial market. A
// depot may service several markets and a market may be serviced by several
// depots, so Depot.Country is never assumed to equal Market.Country.
type DepotMarket struct {
	ID            string     `json:"id"`
	DepotCode     string     `json:"depot_code"`
	MarketCode    string     `json:"market_code"`
	IsActive      bool       `json:"is_active"`
	EffectiveFrom time.Time  `json:"effective_from"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`

	audit.AuditFields
}
