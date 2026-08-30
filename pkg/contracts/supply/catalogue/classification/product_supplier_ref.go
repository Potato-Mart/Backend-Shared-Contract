package classification

// ProductSupplierRef is the code-only supplier relationship persisted with a
// product. Display names, contacts, addresses, legal identifiers, and
// operational fields are resolved from the supplier master.
type ProductSupplierRef struct {
	Code string `json:"code"`
}
