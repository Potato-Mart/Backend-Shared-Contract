package warehouse_enums

// PackageSubstitutionReason is why packing replaced a requested sealed case
// with loose units, carried in PackageSubstitutionSnapshot.reason_code.
// CaseShortageExactEach means no sealed case was available and the exact
// number of loose units came from one lot and bucket. CaseAssembledFromLoose
// means one case was made up from loose units, possibly from several lots and
// buckets, which the snapshot lists as parts; it needs no buyer consent.
// PackingCaseBreak records a sealed case opened at the packing station.
type PackageSubstitutionReason string

const (
	PackageSubstitutionReasonCaseShortageExactEach  PackageSubstitutionReason = "CASE_SHORTAGE_EXACT_EACH"
	PackageSubstitutionReasonCaseAssembledFromLoose PackageSubstitutionReason = "CASE_ASSEMBLED_FROM_LOOSE"
	PackageSubstitutionReasonPackingCaseBreak       PackageSubstitutionReason = "PACKING_CASE_BREAK"
)

// IsValid reports whether r is a known PackageSubstitutionReason.
func (r PackageSubstitutionReason) IsValid() bool {
	switch r {
	case PackageSubstitutionReasonCaseShortageExactEach, PackageSubstitutionReasonCaseAssembledFromLoose, PackageSubstitutionReasonPackingCaseBreak:
		return true
	}
	return false
}

func (r PackageSubstitutionReason) String() string { return string(r) }
