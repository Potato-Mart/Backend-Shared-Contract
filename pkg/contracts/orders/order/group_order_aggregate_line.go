package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"
)

// GroupOrderAggregateLine records one parent-owned aggregate package demand
// and its reservation and allocation references.
type GroupOrderAggregateLine struct {
	ID         string `json:"id"`
	SKUCode    string `json:"sku_code"`
	MarketCode string `json:"market_code"`
	// PriceSnapshot evidence for the aggregate line lives on its components;
	// the line itself carries only identity, composition, and totals.
	PackageOptionCode    string                               `json:"package_option_code"`
	RequestedComposition packaging.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition packaging.PackageCompositionSnapshot `json:"allocated_composition"`
	ShortageComposition  packaging.PackageCompositionSnapshot `json:"shortage_composition"`
	ReturnedComposition  packaging.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition  packaging.PackageCompositionSnapshot `json:"refunded_composition"`
	Components           []PricedPackageComponent             `json:"components"`
	DiscountAmount       money.Money                          `json:"discount_amount"`
	TaxAmount            money.Money                          `json:"tax_amount"`
	Total                money.Money                          `json:"total"`
	RefundAmount         money.Money                          `json:"refund_amount"`
	ReservationID        string                               `json:"reservation_id"`
	AllocationLineID     string                               `json:"allocation_line_id"`
}
