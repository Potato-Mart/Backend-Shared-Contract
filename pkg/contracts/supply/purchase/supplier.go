package purchase

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/party"
)

// Supplier is the full supplier record. A supplier is an organisation, so it
// carries the complete organisation profile via party.OrganisationDetail
// (which embeds PartyRef for id / name / phone / email, plus registration,
// tax, addresses, branding and other organisation fields).
type Supplier struct {
	party.OrganisationDetail
	audit.AuditFields
}
