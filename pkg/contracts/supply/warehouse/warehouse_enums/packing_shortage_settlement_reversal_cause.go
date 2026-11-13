package warehouse_enums

// PackingShortageSettlementReversalCause is what took settled units back: an
// explicit shortage withdrawal, a scan, a saved packed quantity, a newer
// refund request that was settled for the same units (superseded), or the
// settlement's own station refund request being rejected or cancelled in
// Payments (refund_request_closed, from the refund.request_closed event).
type PackingShortageSettlementReversalCause string

const (
	PackingShortageSettlementReversalCauseWithdrawal PackingShortageSettlementReversalCause = "withdrawal"
	PackingShortageSettlementReversalCauseScan       PackingShortageSettlementReversalCause = "scan"
	PackingShortageSettlementReversalCauseProgress   PackingShortageSettlementReversalCause = "progress"
	PackingShortageSettlementReversalCauseSuperseded PackingShortageSettlementReversalCause = "superseded"
	// PackingShortageSettlementReversalCauseRefundRequestClosed marks a
	// reversal made because the station refund request behind the settlement
	// was rejected or cancelled.
	PackingShortageSettlementReversalCauseRefundRequestClosed PackingShortageSettlementReversalCause = "refund_request_closed"
)

// IsValid reports whether c is a known PackingShortageSettlementReversalCause.
func (c PackingShortageSettlementReversalCause) IsValid() bool {
	switch c {
	case PackingShortageSettlementReversalCauseWithdrawal, PackingShortageSettlementReversalCauseScan,
		PackingShortageSettlementReversalCauseProgress, PackingShortageSettlementReversalCauseSuperseded,
		PackingShortageSettlementReversalCauseRefundRequestClosed:
		return true
	}
	return false
}

func (c PackingShortageSettlementReversalCause) String() string { return string(c) }
