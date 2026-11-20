package sales

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// OrderItemFact is the immutable product and merchandising snapshot used by
// sales rollups. Dimension values are canonical identifiers captured at purchase
// time, so later catalogue edits cannot rewrite historical analytics.
type OrderItemFact struct {
	ProductFactDimensions
	// ProductName is captured for order rollups only; refund facts remain
	// bound to the shared ProductFactDimensions.
	ProductName string      `json:"product_name,omitempty"`
	Gross       money.Money `json:"gross"`
	// PromotionAllocations records frozen proportions of this paid line.
	// Historical facts omit it when the producer lacks allocation evidence.
	PromotionAllocations []OrderItemPromotionAllocation `json:"promotion_allocations,omitempty"`
}
