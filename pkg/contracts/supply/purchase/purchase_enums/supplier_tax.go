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

// InputTaxClaimStatus records whether recorded input tax may be claimed. It is
// never set to claimable without qualifying tax-invoice evidence.
type InputTaxClaimStatus string

const (
	InputTaxClaimStatusClaimable    InputTaxClaimStatus = "claimable"
	InputTaxClaimStatusNotClaimable InputTaxClaimStatus = "not_claimable"
	// InputTaxClaimStatusInsufficientEvidence is a claim blocked because
	// the qualifying tax-invoice evidence is missing or incomplete.
	InputTaxClaimStatusInsufficientEvidence InputTaxClaimStatus = "insufficient_evidence"
	InputTaxClaimStatusPendingReview        InputTaxClaimStatus = "pending_review"
)

// IsValid reports whether s is a known InputTaxClaimStatus.
func (s InputTaxClaimStatus) IsValid() bool {
	switch s {
	case InputTaxClaimStatusClaimable, InputTaxClaimStatusNotClaimable,
		InputTaxClaimStatusInsufficientEvidence, InputTaxClaimStatusPendingReview:
		return true
	}
	return false
}

func (s InputTaxClaimStatus) String() string { return string(s) }

// SupplierInvoiceStatus is the lifecycle state of a recorded supplier invoice.
type SupplierInvoiceStatus string

const (
	SupplierInvoiceStatusDraft SupplierInvoiceStatus = "draft"
	// SupplierInvoiceStatusBlocked is an invoice that cannot be confirmed
	// because a line has unknown taxability or an unknown price basis.
	SupplierInvoiceStatusBlocked   SupplierInvoiceStatus = "blocked"
	SupplierInvoiceStatusConfirmed SupplierInvoiceStatus = "confirmed"
	SupplierInvoiceStatusDisputed  SupplierInvoiceStatus = "disputed"
	SupplierInvoiceStatusCancelled SupplierInvoiceStatus = "cancelled"
)

// IsValid reports whether s is a known SupplierInvoiceStatus.
func (s SupplierInvoiceStatus) IsValid() bool {
	switch s {
	case SupplierInvoiceStatusDraft, SupplierInvoiceStatusBlocked, SupplierInvoiceStatusConfirmed,
		SupplierInvoiceStatusDisputed, SupplierInvoiceStatusCancelled:
		return true
	}
	return false
}

func (s SupplierInvoiceStatus) String() string { return string(s) }
