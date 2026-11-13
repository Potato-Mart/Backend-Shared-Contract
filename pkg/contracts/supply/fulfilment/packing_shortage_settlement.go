package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackingShortageSettlement records what Supply did to stock when the packing
// station raised a refund for units of a recorded shortage. The discrepancy
// keeps kind shortage; the reason is recorded here.
//
// BaseUnits is what the station sent. SettledBaseUnits is the part counted
// against the remaining shortage. The units still settled are
// SettledBaseUnits minus ReversedBaseUnits, and they never exceed the
// remaining base units of the shortage.
//
// ReleasedBaseUnits were taken off the stock hold of the line. For a no_stock
// settlement WrittenOffBaseUnits of them were also written off, as the stock
// movements in WriteOffMovementIDs. The two DepotAvailable fields are the
// signed available figure of the product in the depot before and after; they
// are absent when no stock was read.
//
// A no_stock settlement takes the order depot's own available figure to
// min(0, before+released). When the released units alone do not get it
// there, Supply writes off further free units the depot's other counted
// sellable stock still shows (or books them on the no-lot deficit); those
// units are WrittenOffBeyondReleaseBaseUnits and are not part of
// WrittenOffBaseUnits. Their movements are also listed in
// WriteOffMovementIDs. They come back only when the whole settlement is
// reversed (see PackingShortageSettlementReversal). The field is absent on
// settlements recorded before v53.36.0 and when nothing beyond the release
// was written off.
//
// A deferred settlement released nothing from the hold: DeferredReason says
// why the hold was kept. A deferred no_stock settlement may still carry
// WrittenOffBeyondReleaseBaseUnits; a deferred settlement recorded before
// v53.36.0 changed no stock. An applied settlement that released fewer units
// than it settled also carries the reason for the rest.
type PackingShortageSettlement struct {
	ID                               string                                          `json:"id"`
	Reason                           warehouse_enums.PackingShortageSettlementReason `json:"reason"`
	RefundRequestID                  string                                          `json:"refund_request_id"`
	BaseUnits                        int64                                           `json:"base_units"`
	SettledBaseUnits                 int64                                           `json:"settled_base_units"`
	Status                           warehouse_enums.PackingShortageSettlementStatus `json:"status"`
	DeferredReason                   string                                          `json:"deferred_reason,omitempty"`
	ReleasedBaseUnits                int64                                           `json:"released_base_units"`
	WrittenOffBaseUnits              int64                                           `json:"written_off_base_units"`
	WriteOffMovementIDs              []string                                        `json:"write_off_movement_ids,omitempty"`
	WrittenOffBeyondReleaseBaseUnits int64                                           `json:"written_off_beyond_release_base_units,omitempty"`
	DepotAvailableBeforeBaseUnits    *int64                                          `json:"depot_available_before_base_units,omitempty"`
	DepotAvailableAfterBaseUnits     *int64                                          `json:"depot_available_after_base_units,omitempty"`
	PackedAtSettlement               bool                                            `json:"packed_at_settlement,omitempty"`
	SettledBy                        string                                          `json:"settled_by"`
	SettledAt                        time.Time                                       `json:"settled_at"`
	ReversedBaseUnits                int64                                           `json:"reversed_base_units,omitempty"`
	ReleasedReversedBaseUnits        int64                                           `json:"released_reversed_base_units,omitempty"`
	Reversals                        []PackingShortageSettlementReversal             `json:"reversals,omitempty"`
}
