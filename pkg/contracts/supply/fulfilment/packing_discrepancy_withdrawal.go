package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackingDiscrepancyWithdrawal records base units withdrawn from a recorded
// packing discrepancy, either by an explicit staff action or by a later scan.
// WithdrawnBaseUnits is positive. WithdrawnBy is required.
type PackingDiscrepancyWithdrawal struct {
	WithdrawnBaseUnits int64                                              `json:"withdrawn_base_units"`
	Method             warehouse_enums.PackingDiscrepancyWithdrawalMethod `json:"method"`
	WithdrawnBy        string                                             `json:"withdrawn_by"`
	WithdrawnAt        time.Time                                          `json:"withdrawn_at"`
}
