package access

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/identity/access/access_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/identity/account/account_enums"
)

// PortalAccess records whether an account/persona may enter one front-door
// portal. It is the portal admission contract and is separate from RBAC.
type PortalAccess struct {
	ID          string                          `json:"id"`
	UserID      string                          `json:"user_id"`
	AccountID   string                          `json:"account_id"`
	AccountType account_enums.AccountType       `json:"account_type"`
	Portal      identity_enums.Portal           `json:"portal"`
	Status      access_enums.PortalAccessStatus `json:"status"`
	Grant       *audit.LifecycleAction          `json:"grant,omitempty"`
	Revocation  *audit.LifecycleAction          `json:"revocation,omitempty"`
	ExpiresAt   *time.Time                      `json:"expires_at,omitempty"`

	audit.AuditFields
}
