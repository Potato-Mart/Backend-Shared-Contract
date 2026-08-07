package product

// ProductSupplierRef is the customer-safe supplier identity exposed with a
// product. Supplier contacts, addresses, legal identifiers, and operational
// fields remain outside this reference.
type ProductSupplierRef struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ProductManufacturing contains customer-safe product manufacturing details.
// Its fields are optional so partially known declarations remain representable.
type ProductManufacturing struct {
	CompanyName string `json:"company_name,omitempty"`
	Location    string `json:"location,omitempty"`
}

// ProductSupply groups customer-safe supplier and manufacturing information.
// The two sections are independently optional because manufacturing details
// may be known even when no supplier reference is available, and vice versa.
type ProductSupply struct {
	Supplier      *ProductSupplierRef   `json:"supplier,omitempty"`
	Manufacturing *ProductManufacturing `json:"manufacturing,omitempty"`
}
