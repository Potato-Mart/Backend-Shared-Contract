package purchase

import "github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/common"

// Supplier is the supplier snapshot carried on purchase orders.
// The contact fields come from common.PartyRef (id / name / phone / email).
type Supplier struct {
	common.PartyRef
	common.AuditFields
}
