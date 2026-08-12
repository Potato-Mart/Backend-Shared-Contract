package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/party"
)

// Supplier is the full supplier record. A supplier is an organisation, so it
// carries the complete organisation profile via party.OrganisationDetail
// (which embeds PartyRef for id / name / phone / email, plus registration,
// tax, addresses, branding and other organisation fields).
type Supplier struct {
	party.OrganisationDetail
	audit.AuditFields
}

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
