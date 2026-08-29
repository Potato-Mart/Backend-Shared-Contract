package purchase_enums

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
