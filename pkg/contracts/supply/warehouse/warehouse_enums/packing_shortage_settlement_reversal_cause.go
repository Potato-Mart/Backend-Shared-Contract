package warehouse_enums

// PackingShortageSettlementReversalCause is what took settled units back: an
// explicit shortage withdrawal, a scan, a saved packed quantity, or a newer
// refund request that was settled for the same units (superseded).
type PackingShortageSettlementReversalCause string

const (
	PackingShortageSettlementReversalCauseWithdrawal PackingShortageSettlementReversalCause = "withdrawal"
	PackingShortageSettlementReversalCauseScan       PackingShortageSettlementReversalCause = "scan"
	PackingShortageSettlementReversalCauseProgress   PackingShortageSettlementReversalCause = "progress"
	PackingShortageSettlementReversalCauseSuperseded PackingShortageSettlementReversalCause = "superseded"
)

// IsValid reports whether c is a known PackingShortageSettlementReversalCause.
func (c PackingShortageSettlementReversalCause) IsValid() bool {
	switch c {
	case PackingShortageSettlementReversalCauseWithdrawal, PackingShortageSettlementReversalCauseScan,
		PackingShortageSettlementReversalCauseProgress, PackingShortageSettlementReversalCauseSuperseded:
		return true
	}
	return false
}

func (c PackingShortageSettlementReversalCause) String() string { return string(c) }
