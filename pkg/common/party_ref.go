package common

// PartyRef is the shared identity-and-contact base for a person or
// organisation involved in a transaction (order customer, supplier, etc.).
// It is embedded by the structs that need these fields (Supplier,
// SupplierSnapshot, OrganisationDetail, sales.Order.Customer, ...)
// so the id/name/phone/email quartet is defined once. Each embedder adds its
// own extra fields without affecting the others.
type PartyRef struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}
