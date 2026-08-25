package shipping

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/warehouse/warehouse_enums"
	"time"
)

type OutboundShipment struct {
	ID        string `json:"id"`
	DepotCode string `json:"depot_code"`
	// MarketCode and CountryCode are the denormalized market and country of
	// the dispatching depot, carried so a geographically scoped staff query
	// is a plain indexed match.
	MarketCode    string                                 `json:"market_code,omitempty"`
	CountryCode   geography.CountryCode                  `json:"country_code,omitempty"`
	PickingListID string                                 `json:"picking_list_id,omitempty"`
	OrderNumber   string                                 `json:"order_number"`
	CustomerName  string                                 `json:"customer_name,omitempty"`
	Address       *geography.Address                     `json:"address,omitempty"`
	Operator      string                                 `json:"operator"`
	Status        warehouse_enums.OutboundShipmentStatus `json:"status"`
	Containers    []operations.OutboundContainerPlan     `json:"containers,omitempty"`
	// Carrier is the optional delivery-company code recorded by Supply
	// (for example detrack, bcrc, aupost).
	Carrier        string                  `json:"carrier,omitempty"`
	TrackingNumber string                  `json:"tracking_number,omitempty"`
	Note           string                  `json:"note,omitempty"`
	DispatchedAt   *time.Time              `json:"dispatched_at,omitempty"`
	DeliveredAt    *time.Time              `json:"delivered_at,omitempty"`
	History        []security.HistoryEntry `json:"history,omitempty"`
	CreatedAt      time.Time               `json:"created_at"`
}
