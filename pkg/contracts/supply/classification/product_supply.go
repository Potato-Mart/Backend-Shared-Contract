package classification

// ProductSupply groups customer-safe supplier and manufacturing information.
// The two sections are independently optional because manufacturing details
// may be known even when no supplier reference is available, and vice versa.
type ProductSupply struct {
	Suppliers     []ProductSupplierRef  `json:"suppliers,omitempty"`
	Manufacturing *ProductManufacturing `json:"manufacturing,omitempty"`
}
