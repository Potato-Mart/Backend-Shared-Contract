package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
)

// NetGoodsValuationMovement is an immutable private posting to net-goods
// valuation, separate from CarryingCostMovement. Supply owns transactions,
// exact allocation arithmetic, stock eligibility and duplicate prevention.
type NetGoodsValuationMovement struct {
	ID        string             `json:"id"`
	SKUCode   string             `json:"sku_code"`
	DepotCode string             `json:"depot_code"`
	Currency  money.CurrencyCode `json:"currency"`

	ReferenceType   string `json:"reference_type"`
	ReferenceID     string `json:"reference_id"`
	ReferenceLineID string `json:"reference_line_id,omitempty"`
	// Sources freezes the contributing cost evidence. SourceMovementIDs
	// retains lineage through issues, transfers, returns and corrections;
	// ReversesMovementID identifies the posting explicitly reversed.
	Sources            []NetGoodsCostSourceSnapshot `json:"sources,omitempty"`
	SourceMovementIDs  []string                     `json:"source_movement_ids,omitempty"`
	ReversesMovementID string                       `json:"reverses_movement_id,omitempty"`

	// Provisional is a subset of valued, never an additional owned quantity.
	ValuedBaseUnitDelta      int64 `json:"valued_base_unit_delta"`
	ProvisionalBaseUnitDelta int64 `json:"provisional_base_unit_delta"`
	UnvaluedBaseUnitDelta    int64 `json:"unvalued_base_unit_delta"`
	// SourceCorrectionAmount is present only for a financial correction of
	// prior source valuation. Its signed amount equals the on-hand cost delta
	// plus consumed-cost variance delta plus rounding variance delta. It is
	// absent for ordinary receipts, issues and transfers: issuing stock does
	// not change an invoice source or turn ordinary consumed cost into a
	// variance. Explicit zero preserves a correction with no amount change.
	SourceCorrectionAmount *money.Money `json:"source_correction_amount,omitempty"`
	// Only OnHandCostMinorDelta contributes money to the on-hand weighted
	// average's numerator; valued quantity deltas affect its denominator.
	// Corrections affecting consumed stock retain variance evidence without
	// rewriting historical issues or customer price snapshots.
	OnHandCostMinorDelta           int64 `json:"on_hand_cost_minor_delta"`
	ConsumedCostVarianceMinorDelta int64 `json:"consumed_cost_variance_minor_delta"`
	RoundingVarianceMinorDelta     int64 `json:"rounding_variance_minor_delta"`

	BalanceValuedBaseUnitsAfter           int64 `json:"balance_valued_base_units_after"`
	BalanceProvisionalBaseUnitsAfter      int64 `json:"balance_provisional_base_units_after"`
	BalanceUnvaluedBaseUnitsAfter         int64 `json:"balance_unvalued_base_units_after"`
	BalanceNetGoodsCostMinorAfter         int64 `json:"balance_net_goods_cost_minor_after"`
	BalanceConsumedCostVarianceMinorAfter int64 `json:"balance_consumed_cost_variance_minor_after"`
	BalanceRoundingVarianceMinorAfter     int64 `json:"balance_rounding_variance_minor_after"`

	Revision      int64             `json:"revision"`
	OccurredAt    time.Time         `json:"occurred_at"`
	Actor         security.ActorRef `json:"actor"`
	RequestID     string            `json:"request_id,omitempty"`
	CorrelationID string            `json:"correlation_id,omitempty"`

	security.DataProtectionFields
}
