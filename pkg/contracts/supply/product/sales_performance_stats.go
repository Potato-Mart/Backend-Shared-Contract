package product

import (
	"time"
)

// SalesPerformanceStats is the computed, read-only sales-performance model.
type SalesPerformanceStats struct {
	Last7Days     SalesWindowStats    `json:"last_7_days"`
	Last30Days    SalesWindowStats    `json:"last_30_days"`
	Last90Days    SalesWindowStats    `json:"last_90_days"`
	Lifetime      SalesTotals         `json:"lifetime"`
	CategoryRanks []CategorySalesRank `json:"category_ranks,omitempty"`
	LastSoldAt    *time.Time          `json:"last_sold_at,omitempty"`
	AsOf          time.Time           `json:"as_of"`
	Timezone      string              `json:"timezone"`
}
