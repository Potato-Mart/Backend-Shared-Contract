package purchase

import common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"

// Supplier is the full supplier record. A supplier is an organisation, so it
// carries the complete organisation profile via common.OrganisationDetail
// (which embeds PartyRef for id / name / phone / email, plus registration,
// tax, addresses, branding and other organisation fields).
type Supplier struct {
	common.OrganisationDetail
	common.AuditFields
}
