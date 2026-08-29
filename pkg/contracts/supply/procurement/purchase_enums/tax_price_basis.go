package purchase_enums

// TaxPriceBasis records whether a supplier line amount already includes tax.
// Missing tax may only be calculated for an explicitly taxable line whose
// basis is known.
type TaxPriceBasis string

const (
	TaxPriceBasisInclusive TaxPriceBasis = "inclusive"
	TaxPriceBasisExclusive TaxPriceBasis = "exclusive"
	// TaxPriceBasisUnknown blocks authoritative confirmation.
	TaxPriceBasisUnknown TaxPriceBasis = "unknown"
)

// IsValid reports whether b is a known TaxPriceBasis.
func (b TaxPriceBasis) IsValid() bool {
	switch b {
	case TaxPriceBasisInclusive, TaxPriceBasisExclusive, TaxPriceBasisUnknown:
		return true
	}
	return false
}

func (b TaxPriceBasis) String() string { return string(b) }
