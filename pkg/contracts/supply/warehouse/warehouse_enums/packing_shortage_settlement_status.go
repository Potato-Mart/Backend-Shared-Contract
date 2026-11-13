package warehouse_enums

// PackingShortageSettlementStatus is the state of one shortage settlement.
// Applied means stock was settled (or there was no stock hold to settle).
// Deferred means the hold was kept and nothing changed; a count releases it.
// Reversed means every settled unit was taken back.
type PackingShortageSettlementStatus string

const (
	PackingShortageSettlementStatusApplied  PackingShortageSettlementStatus = "applied"
	PackingShortageSettlementStatusDeferred PackingShortageSettlementStatus = "deferred"
	PackingShortageSettlementStatusReversed PackingShortageSettlementStatus = "reversed"
)

// IsValid reports whether s is a known PackingShortageSettlementStatus.
func (s PackingShortageSettlementStatus) IsValid() bool {
	switch s {
	case PackingShortageSettlementStatusApplied, PackingShortageSettlementStatusDeferred,
		PackingShortageSettlementStatusReversed:
		return true
	}
	return false
}

func (s PackingShortageSettlementStatus) String() string { return string(s) }
