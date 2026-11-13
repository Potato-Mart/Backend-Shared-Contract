package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackingShortageSettlementReversal records settled units taken back by a
// later packing write. ReleasedBaseUnits of them had been released from the
// stock hold. WriteOffRestoredBaseUnits were put back on hand by the
// movements in WriteOffReversalMovementIDs; WriteOffKeptBaseUnits stayed
// written off for WriteOffKeptReason (for example counted_since: the location
// was counted after the write-off). RefundedButShipped marks units packed
// after their refund had completed.
type PackingShortageSettlementReversal struct {
	BaseUnits                   int64                                                  `json:"base_units"`
	Cause                       warehouse_enums.PackingShortageSettlementReversalCause `json:"cause"`
	ReleasedBaseUnits           int64                                                  `json:"released_base_units,omitempty"`
	WriteOffRestoredBaseUnits   int64                                                  `json:"write_off_restored_base_units,omitempty"`
	WriteOffKeptBaseUnits       int64                                                  `json:"write_off_kept_base_units,omitempty"`
	WriteOffKeptReason          string                                                 `json:"write_off_kept_reason,omitempty"`
	WriteOffReversalMovementIDs []string                                               `json:"write_off_reversal_movement_ids,omitempty"`
	RefundedButShipped          bool                                                   `json:"refunded_but_shipped,omitempty"`
	ReversedBy                  string                                                 `json:"reversed_by"`
	ReversedAt                  time.Time                                              `json:"reversed_at"`
}
