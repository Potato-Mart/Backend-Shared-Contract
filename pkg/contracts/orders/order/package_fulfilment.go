package order

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
)

// PricedPackageComponent freezes the accepted offer, package configuration,
// requested quantity, and commercial totals for one CASE or EACH component.
type PricedPackageComponent struct {
	AcceptedOffer         product.SellableOfferSnapshot `json:"accepted_offer"`
	RequestedPackageCount int64                         `json:"requested_package_count"`
	RequestedBaseUnits    int64                         `json:"requested_base_units"`
	PackagePrice          common.Money                  `json:"package_price"`
	TaxAmount             common.Money                  `json:"tax_amount"`
	DiscountAmount        common.Money                  `json:"discount_amount"`
	ComponentTotal        common.Money                  `json:"component_total"`
}

// LooseSubstitutionPolicySnapshot records whether unavailable sealed cases may
// be fulfilled with an exact base-unit quantity of loose EACH stock.
type LooseSubstitutionPolicySnapshot struct {
	Allowed    bool                          `json:"allowed"`
	Source     LooseSubstitutionPolicySource `json:"source"`
	CapturedAt time.Time                     `json:"captured_at"`
}

// GroupOrderContext identifies an order as the consolidated fulfilment owner
// or as a participant referencing that parent fulfilment.
type GroupOrderContext struct {
	GroupOrderCode          string         `json:"group_order_code"`
	Role                    GroupOrderRole `json:"role"`
	ParentOrderNumber       string         `json:"parent_order_number,omitempty"`
	ParentFulfilmentID      string         `json:"parent_fulfilment_id,omitempty"`
	ParentAllocationLineIDs []string       `json:"parent_allocation_line_ids,omitempty"`
}

// GroupOrderAggregateLine records one parent-owned aggregate package demand
// and its reservation and allocation references.
type GroupOrderAggregateLine struct {
	ID                   string                            `json:"id"`
	ProductSKUCode       string                            `json:"product_sku_code"`
	OfferID              string                            `json:"offer_id"`
	OfferRevision        int64                             `json:"offer_revision"`
	PackageOptionID      string                            `json:"package_option_id"`
	RequestedComposition common.PackageCompositionSnapshot `json:"requested_composition"`
	AllocatedComposition common.PackageCompositionSnapshot `json:"allocated_composition"`
	ShortageComposition  common.PackageCompositionSnapshot `json:"shortage_composition"`
	ReturnedComposition  common.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition  common.PackageCompositionSnapshot `json:"refunded_composition"`
	Components           []PricedPackageComponent          `json:"components"`
	DiscountAmount       common.Money                      `json:"discount_amount"`
	TaxAmount            common.Money                      `json:"tax_amount"`
	Total                common.Money                      `json:"total"`
	RefundAmount         common.Money                      `json:"refund_amount"`
	ReservationID        string                            `json:"reservation_id"`
	AllocationLineID     string                            `json:"allocation_line_id"`
}

// GroupOrderParticipantShare records one participant's commercial and
// collection entitlement against a parent allocation line.
type GroupOrderParticipantShare struct {
	ParticipantOrderNumber string                            `json:"participant_order_number"`
	ParticipantOrderItemID string                            `json:"participant_order_item_id"`
	ParentAllocationLineID string                            `json:"parent_allocation_line_id"`
	Sequence               int64                             `json:"sequence"`
	RequestedComposition   common.PackageCompositionSnapshot `json:"requested_composition"`
	FulfilledComposition   common.PackageCompositionSnapshot `json:"fulfilled_composition"`
	ShortageComposition    common.PackageCompositionSnapshot `json:"shortage_composition"`
	ReturnedComposition    common.PackageCompositionSnapshot `json:"returned_composition"`
	RefundedComposition    common.PackageCompositionSnapshot `json:"refunded_composition"`
	Components             []PricedPackageComponent          `json:"components"`
	DiscountAmount         common.Money                      `json:"discount_amount"`
	TaxAmount              common.Money                      `json:"tax_amount"`
	Total                  common.Money                      `json:"total"`
	RefundAmount           common.Money                      `json:"refund_amount"`
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
