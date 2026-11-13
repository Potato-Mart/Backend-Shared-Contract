package warehouse_enums

// PackingShortageSettlementReversalCause is the packing write that took
// settled units back: an explicit shortage withdrawal, a scan, or a saved
// packed quantity.
type PackingShortageSettlementReversalCause string

const (
	PackingShortageSettlementReversalCauseWithdrawal PackingShortageSettlementReversalCause = "withdrawal"
	PackingShortageSettlementReversalCauseScan       PackingShortageSettlementReversalCause = "scan"
	PackingShortageSettlementReversalCauseProgress   PackingShortageSettlementReversalCause = "progress"
)

// IsValid reports whether c is a known PackingShortageSettlementReversalCause.
func (c PackingShortageSettlementReversalCause) IsValid() bool {
	switch c {
	case PackingShortageSettlementReversalCauseWithdrawal, PackingShortageSettlementReversalCauseScan,
		PackingShortageSettlementReversalCauseProgress:
		return true
	}
	return false
}

func (c PackingShortageSettlementReversalCause) String() string { return string(c) }
