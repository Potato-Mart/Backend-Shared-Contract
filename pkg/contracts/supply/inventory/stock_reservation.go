package inventory

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/warehouse/warehouse_enums"
)

// StockReservation is a logical hold for one product at one depot.
type StockReservation struct {
	ID             string `json:"id"`
	SKUCode        string `json:"sku_code"`
	DepotCode      string `json:"depot_code"`
	OrderNumber    string `json:"order_number,omitempty"`
	GroupOrderCode string `json:"group_order_code,omitempty"`
	MarketCode     string `json:"market_code"`
	// EligibilityToken is the validity token from the sale-eligibility
	// snapshot the reservation was taken against, and ListingRevision pins
	// the market listing that was in force. A stale token requires a
	// re-quote rather than a silent re-price.
	EligibilityToken     string                                 `json:"eligibility_token"`
	ListingRevision      int64                                  `json:"listing_revision"`
	RequestedComposition packaging.PackageCompositionSnapshot   `json:"requested_composition"`
	ReservedComposition  packaging.PackageCompositionSnapshot   `json:"reserved_composition"`
	Status               warehouse_enums.StockReservationStatus `json:"status"`
	Revision             int64                                  `json:"revision"`
	Timezone             string                                 `json:"timezone"`
	ExpiresAt            *time.Time                             `json:"expires_at,omitempty"`

	audit.AuditFields
}
