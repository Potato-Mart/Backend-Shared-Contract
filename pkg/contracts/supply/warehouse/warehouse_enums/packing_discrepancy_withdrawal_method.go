package warehouse_enums

// PackingDiscrepancyWithdrawalMethod records how staff changed the base units
// of a recorded packing discrepancy: Command is an explicit staff withdrawal,
// Scan is a withdrawal by a later scan of the units, and Reinstate re-records
// previously withdrawn units that never arrived as missing again.
type PackingDiscrepancyWithdrawalMethod string

const (
	PackingDiscrepancyWithdrawalMethodCommand   PackingDiscrepancyWithdrawalMethod = "command"
	PackingDiscrepancyWithdrawalMethodScan      PackingDiscrepancyWithdrawalMethod = "scan"
	PackingDiscrepancyWithdrawalMethodReinstate PackingDiscrepancyWithdrawalMethod = "reinstate"
)

// IsValid reports whether m is a known PackingDiscrepancyWithdrawalMethod.
func (m PackingDiscrepancyWithdrawalMethod) IsValid() bool {
	switch m {
	case PackingDiscrepancyWithdrawalMethodCommand, PackingDiscrepancyWithdrawalMethodScan,
		PackingDiscrepancyWithdrawalMethodReinstate:
		return true
	}
	return false
}

func (m PackingDiscrepancyWithdrawalMethod) String() string { return string(m) }
