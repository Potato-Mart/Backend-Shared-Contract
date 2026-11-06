package supply

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/supply/supply_enums"
)

// PackingFact is Supply's immutable analytical fact for one packing stage of
// one packing cycle (analytics.packing_fact on fulfilment-events). Supply
// emits one fact per stage when the order is marked packed.
//
// FactID is
// "packing:{order_number}:{fulfillment_generation}:{reopen_cycle}:{stage}".
// Reopening an order starts a new cycle and an order edit starts a new
// fulfilment generation, so a re-pack produces new facts rather than a changed
// one; a consumer counts only the latest generation and cycle of an order and
// keeps earlier ones as history. The same FactID always carries the same
// content.
//
// Times are server times. StartedAt is the first scan of a line in the stage
// during the cycle; CompletedAt is the stage's first label print in the cycle.
// Units are base units, so a case of 24 counts 24.
type PackingFact struct {
	FactID                string                         `json:"fact_id"`
	OrderNumber           string                         `json:"order_number"`
	FulfillmentGeneration int64                          `json:"fulfillment_generation"`
	ReopenCycle           int                            `json:"reopen_cycle"`
	MarketCode            string                         `json:"market_code,omitempty"`
	DepotCode             string                         `json:"depot_code,omitempty"`
	Stage                 supply_enums.PackingStageGroup `json:"stage"`
	PackedBy              string                         `json:"packed_by,omitempty"`
	StartedAt             *time.Time                     `json:"started_at,omitempty"`
	CompletedAt           *time.Time                     `json:"completed_at,omitempty"`
	PackedAt              time.Time                      `json:"packed_at"`
	PackageCount          int                            `json:"package_count"`
	Lines                 []PackingFactLine              `json:"lines"`
	StorageTotals         []PackingFactStorageTotal      `json:"storage_totals"`
}
