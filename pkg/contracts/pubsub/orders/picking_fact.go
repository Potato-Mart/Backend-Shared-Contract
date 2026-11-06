package orders

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/customers/retail/retail_enums"
)

// PickingFact is Orders' immutable analytical fact for one order's picking in
// one fulfilment generation (analytics.picking_fact on order-events). Orders
// emits it once, when picking is final and packing starts.
//
// FactID is "picking:{order_number}:{fulfillment_generation}". Times are
// server times: PickingStartedAt is when the picking list was printed and
// PickingCompletedAt is when packing started. Units are base units. Pickers
// lists the PDA picking assignment halves, and is empty when the order was
// picked without an assignment.
type PickingFact struct {
	FactID                string                 `json:"fact_id"`
	OrderNumber           string                 `json:"order_number"`
	FulfillmentGeneration int64                  `json:"fulfillment_generation"`
	MarketCode            string                 `json:"market_code,omitempty"`
	DepotCode             string                 `json:"depot_code,omitempty"`
	BuyerType             retail_enums.BuyerType `json:"buyer_type,omitempty"`
	PickingStartedAt      *time.Time             `json:"picking_started_at,omitempty"`
	PickingCompletedAt    time.Time              `json:"picking_completed_at"`
	Lines                 []PickingFactLine      `json:"lines"`
	Pickers               []PickingFactPicker    `json:"pickers,omitempty"`
}
