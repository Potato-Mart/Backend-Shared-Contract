package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackingShortageSettlementReversal records settled units taken back by a
// later packing write, or by a newer refund request settled for the same
// units (cause superseded). ReleasedBaseUnits of them had been released from the
// stock hold. WriteOffRestoredBaseUnits were put back on hand by the
// movements in WriteOffReversalMovementIDs; WriteOffKeptBaseUnits stayed
// written off for WriteOffKeptReason (for example counted_since: the location
// was counted after the write-off). RefundedButShipped marks units packed
// after their refund had completed.
//
// WriteOffBeyondReleaseRestoredBaseUnits are units of the settlement's
// WrittenOffBeyondReleaseBaseUnits put back on hand by this reversal; Supply
// restores them only when the whole settlement is reversed. They are not part
// of WriteOffRestoredBaseUnits, and their movements are also listed in
// WriteOffReversalMovementIDs. The field is absent when none were restored.
//
// Cause refund_request_closed is a station refund request that Payments
// rejected or cancelled after it was settled.
type PackingShortageSettlementReversal struct {
	BaseUnits                              int64                                                  `json:"base_units"`
	Cause                                  warehouse_enums.PackingShortageSettlementReversalCause `json:"cause"`
	ReleasedBaseUnits                      int64                                                  `json:"released_base_units,omitempty"`
	WriteOffRestoredBaseUnits              int64                                                  `json:"write_off_restored_base_units,omitempty"`
	WriteOffKeptBaseUnits                  int64                                                  `json:"write_off_kept_base_units,omitempty"`
	WriteOffKeptReason                     string                                                 `json:"write_off_kept_reason,omitempty"`
	WriteOffBeyondReleaseRestoredBaseUnits int64                                                  `json:"write_off_beyond_release_restored_base_units,omitempty"`
	WriteOffReversalMovementIDs            []string                                               `json:"write_off_reversal_movement_ids,omitempty"`
	RefundedButShipped                     bool                                                   `json:"refunded_but_shipped,omitempty"`
	ReversedBy                             string                                                 `json:"reversed_by"`
	ReversedAt                             time.Time                                              `json:"reversed_at"`
}
