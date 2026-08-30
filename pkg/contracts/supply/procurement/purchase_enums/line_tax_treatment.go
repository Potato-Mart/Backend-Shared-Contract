package purchase_enums

// LineTaxTreatment is the declared taxability of one supplier invoice line.
// Supplier tax registration is an entity fact and is never inferred from
// product taxability.
type LineTaxTreatment string

const (
	// LineTaxTreatmentTaxable is a taxable supply.
	LineTaxTreatmentTaxable LineTaxTreatment = "taxable"
	// LineTaxTreatmentGSTFree is a supply that carries no tax but is still
	// within the tax system.
	LineTaxTreatmentGSTFree LineTaxTreatment = "gst_free"
	// LineTaxTreatmentInputTaxed is a supply whose input tax is not
	// recoverable.
	LineTaxTreatmentInputTaxed LineTaxTreatment = "input_taxed"
	// LineTaxTreatmentOutOfScope is a supply outside the tax system.
	LineTaxTreatmentOutOfScope LineTaxTreatment = "out_of_scope"
	// LineTaxTreatmentUnknown blocks authoritative invoice or receipt
	// confirmation until an administrator resolves it.
	LineTaxTreatmentUnknown LineTaxTreatment = "unknown"
)

// IsValid reports whether t is a known LineTaxTreatment.
func (t LineTaxTreatment) IsValid() bool {
	switch t {
	case LineTaxTreatmentTaxable, LineTaxTreatmentGSTFree, LineTaxTreatmentInputTaxed,
		LineTaxTreatmentOutOfScope, LineTaxTreatmentUnknown:
		return true
	}
	return false
}

func (t LineTaxTreatment) String() string { return string(t) }
