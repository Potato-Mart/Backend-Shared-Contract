package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order/order_enums"
)

// PricedPackageComponent freezes the accepted offer, package configuration,
// requested quantity, and commercial totals for one CASE or EACH component.
type PricedPackageComponent struct {
	AcceptedOffer         product.SellableOfferSnapshot `json:"accepted_offer"`
	RequestedPackageCount int64                         `json:"requested_package_count"`
	RequestedBaseUnits    int64                         `json:"requested_base_units"`
	PackagePrice          money.Money                   `json:"package_price"`
	TaxAmount             money.Money                   `json:"tax_amount"`
	DiscountAmount        money.Money                   `json:"discount_amount"`
	ComponentTotal        money.Money                   `json:"component_total"`
}

// LooseSubstitutionPolicySnapshot records whether unavailable sealed cases may
// be fulfilled with an exact base-unit quantity of loose EACH stock.
type LooseSubstitutionPolicySnapshot struct {
	Allowed    bool                                      `json:"allowed"`
	Source     order_enums.LooseSubstitutionPolicySource `json:"source"`
	CapturedAt time.Time                                 `json:"captured_at"`
}

// GroupOrderContext identifies an order as the consolidated fulfilment owner
// or as a participant referencing that parent fulfilment.
type GroupOrderContext struct {
	GroupOrderCode          string                     `json:"group_order_code"`
	Role                    order_enums.GroupOrderRole `json:"role"`
	ParentOrderNumber       string                     `json:"parent_order_number,omitempty"`
	ParentFulfilmentID      string                     `json:"parent_fulfilment_id,omitempty"`
	ParentAllocationLineIDs []string                   `json:"parent_allocation_line_ids,omitempty"`
}

// GroupOrderAggregateLine records one parent-owned aggregate package demand
// and its reservation and allocation references.
type GroupOrderAggregateLine struct {
	ID                   string                               `json:"id"`
	ProductSKUCode       string                               `json:"product_sku_code"`
	OfferID              string                               `json:"offer_id"`
	OfferRevision        int64                                `json:"offer_revision"`
	PackageOptionID      string                               `json:"package_option_id"`
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

// GroupOrderParticipantShare records one participant's commercial and
// collection entitlement against a parent allocation line.
type GroupOrderParticipantShare struct {
	ParticipantOrderNumber string                               `json:"participant_order_number"`
	ParticipantOrderItemID string                               `json:"participant_order_item_id"`
	ParentAllocationLineID string                               `json:"parent_allocation_line_id"`
	Sequence               int64                                `json:"sequence"`
	RequestedComposition   packaging.PackageCompositionSnapshot `json:"requested_composition"`
	FulfilledComposition   packaging.PackageCompositionSnapshot `json:"fulfilled_composition"`
	ShortageComposition    packaging.PackageCompositionSnapshot `json:"shortage_composition"`
	ReturnedComposition    packaging.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition    packaging.PackageCompositionSnapshot `json:"refunded_composition"`
	Components             []PricedPackageComponent             `json:"components"`
	DiscountAmount         money.Money                          `json:"discount_amount"`
	TaxAmount              money.Money                          `json:"tax_amount"`
	Total                  money.Money                          `json:"total"`
	RefundAmount           money.Money                          `json:"refund_amount"`
}

// GroupOrderFulfilmentPlan is the consolidated parent-owned inventory and
// participant-share snapshot for one group order.
type GroupOrderFulfilmentPlan struct {
	ID                 string                       `json:"id"`
	GroupOrderCode     string                       `json:"group_order_code"`
	ParentOrderNumber  string                       `json:"parent_order_number"`
	ParentFulfilmentID string                       `json:"parent_fulfilment_id"`
	AggregateLines     []GroupOrderAggregateLine    `json:"aggregate_lines"`
	ParticipantShares  []GroupOrderParticipantShare `json:"participant_shares"`
	Revision           int64                        `json:"revision"`
	Timezone           string                       `json:"timezone"`
	CapturedAt         time.Time                    `json:"captured_at"`
}
