package warehouse_enums

// PackingDiscrepancyWithdrawalMethod records how staff withdrew base units
// from a recorded packing discrepancy: Command is an explicit staff action,
// Scan is a later scan of the units.
type PackingDiscrepancyWithdrawalMethod string

const (
	PackingDiscrepancyWithdrawalMethodCommand PackingDiscrepancyWithdrawalMethod = "command"
	PackingDiscrepancyWithdrawalMethodScan    PackingDiscrepancyWithdrawalMethod = "scan"
)

// IsValid reports whether m is a known PackingDiscrepancyWithdrawalMethod.
func (m PackingDiscrepancyWithdrawalMethod) IsValid() bool {
	switch m {
	case PackingDiscrepancyWithdrawalMethodCommand, PackingDiscrepancyWithdrawalMethodScan:
		return true
	}
	return false
}

func (m PackingDiscrepancyWithdrawalMethod) String() string { return string(m) }
