package purchase

import "github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"

// Supplier is the supplier snapshot carried on purchase orders.
// The contact fields come from common.PartyRef (id / name / phone / email).
type Supplier struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	common.PartyRef    `bson:",inline"`
	common.AuditFields `bson:",inline"`
}

type SupplierSnapshot struct {
	SupplierID         string `json:"supplier_id,omitempty"`
	SupplierName       string `json:"supplier_name,omitempty"`
	common.PartyRef    `bson:",inline"`
	common.AuditFields `bson:",inline"`
}
