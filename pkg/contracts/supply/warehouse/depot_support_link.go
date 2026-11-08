package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
)

// DepotSupportLink marks one depot as a supporting depot of one shop depot
// inside one commercial market. A supporting depot has no customers of its
// own: its sellable stock is offered through the shops it supports, and stock
// moves into the shop's receiving location for the stock's storage type when
// a shop sale needs it; stock of a storage type with no receiving location is
// never moved. A
// shop may have any number of supporting depots and a supporting depot may
// support several shops, but both depots must service MarketCode. Priority is
// ascending: the lowest value is used first. Supply owns validation, ordering
// and the stock movements the link enables.
type DepotSupportLink struct {
	ID                 string                          `json:"id"`
	ShopDepotCode      string                          `json:"shop_depot_code"`
	SupportDepotCode   string                          `json:"support_depot_code"`
	MarketCode         string                          `json:"market_code"`
	Priority           int                             `json:"priority"`
	ReceivingLocations []DepotSupportReceivingLocation `json:"receiving_locations"`
	IsActive           bool                            `json:"is_active"`
	EffectiveFrom      time.Time                       `json:"effective_from"`
	EffectiveTo        *time.Time                      `json:"effective_to,omitempty"`

	audit.AuditFields
}
