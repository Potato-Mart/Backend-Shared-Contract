package fulfilment

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// PackingDiscrepancyWithdrawal records base units withdrawn from a recorded
// packing discrepancy, either by an explicit staff action (command) or by a
// later scan (scan). A reinstate entry re-records previously withdrawn units
// that never arrived as missing again. WithdrawnBaseUnits is positive for
// every method; Method alone gives the direction. WithdrawnBy is required.
type PackingDiscrepancyWithdrawal struct {
	WithdrawnBaseUnits int64                                              `json:"withdrawn_base_units"`
	Method             warehouse_enums.PackingDiscrepancyWithdrawalMethod `json:"method"`
	WithdrawnBy        string                                             `json:"withdrawn_by"`
	WithdrawnAt        time.Time                                          `json:"withdrawn_at"`
}
