package purchase

import "github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"

// Supplier is the full supplier record. A supplier is a company, so it
// carries the complete organisation profile via common.CompanyDetail
// (which embeds PartyRef for id / name / phone / email, plus registration,
// tax, addresses, branding and other company fields).
type Supplier struct {
	common.CompanyDetail `bson:",inline"`
	common.AuditFields   `bson:",inline"`
}

// SupplierSnapshot is the lightweight supplier reference embedded on
// purchase orders. It shares the same identity/contact base as Supplier.
type SupplierSnapshot struct {
	common.PartyRef    `bson:",inline"`
	common.AuditFields `bson:",inline"`
}
