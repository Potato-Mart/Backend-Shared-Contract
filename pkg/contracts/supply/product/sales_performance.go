package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/localization"
)

// SalesTotals records paid-order unit activity for one measurement period.
// Completed line refunds are represented separately and subtracted in NetUnits.
type SalesTotals struct {
	PaidOrderCount int64 `json:"paid_order_count"`
	GrossUnits     int64 `json:"gross_units"`
	RefundedUnits  int64 `json:"refunded_units"`
	NetUnits       int64 `json:"net_units"`
}

// SalesWindowStats identifies the bounded period represented by its totals.
type SalesWindowStats struct {
	WindowDays int `json:"window_days"`
	SalesTotals
}

// CategorySalesRank records one product's position within one category tag.
// Rankings use the embedded window and are computed across the complete result
// set, never from one paginated storefront response.
type CategorySalesRank struct {
	CategoryTagID   string                       `json:"category_tag_id"`
	CategoryTagSlug string                       `json:"category_tag_slug,omitempty"`
	CategoryTagName []localization.LocalizedName `json:"category_tag_name,omitempty"`
	Rank            int                          `json:"rank"`
	Population      int                          `json:"population"`
	WindowDays      int                          `json:"window_days"`
	NetUnits        int64                        `json:"net_units"`
}

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
