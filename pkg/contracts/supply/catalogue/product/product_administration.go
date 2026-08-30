package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
)

// ProductAdministration retains master-data history and audit information.
// Customer-facing product projections must omit this optional component.
type ProductAdministration struct {
	History []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}
