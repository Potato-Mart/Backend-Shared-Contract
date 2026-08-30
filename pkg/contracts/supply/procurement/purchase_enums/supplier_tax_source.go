package purchase_enums

// SupplierTaxSource records where a recorded tax amount came from, so an
// explicit supplier figure is always distinguishable from a derived one.
type SupplierTaxSource string

const (
	// SupplierTaxSourceExplicitLine is tax the supplier stated per line.
	SupplierTaxSourceExplicitLine SupplierTaxSource = "explicit_line"
	// SupplierTaxSourceInvoiceHeader is a single invoice-level tax total
	// the supplier stated without a line breakdown.
	SupplierTaxSourceInvoiceHeader SupplierTaxSource = "invoice_header"
	// SupplierTaxSourceCalculated is tax derived from an explicitly taxable
	// line with a known price basis.
	SupplierTaxSourceCalculated SupplierTaxSource = "calculated"
	// SupplierTaxSourceAbsent is no recorded tax at all.
	SupplierTaxSourceAbsent SupplierTaxSource = "absent"
)

// IsValid reports whether s is a known SupplierTaxSource.
func (s SupplierTaxSource) IsValid() bool {
	switch s {
	case SupplierTaxSourceExplicitLine, SupplierTaxSourceInvoiceHeader,
		SupplierTaxSourceCalculated, SupplierTaxSourceAbsent:
		return true
	}
	return false
}

func (s SupplierTaxSource) String() string { return string(s) }
