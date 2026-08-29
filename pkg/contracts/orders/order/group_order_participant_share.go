package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
)

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
