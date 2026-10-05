package tax

import "github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/audit"

// TaxCategory is a Pricing-owned configurable tax classification exchanged with admin consumers.
type TaxCategory struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Revision int64  `json:"revision"`
	audit.AuditFields
}
