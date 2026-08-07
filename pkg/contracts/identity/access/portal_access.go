package access

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/identity/account"
	"time"
)

// PortalAccess records whether an account/persona may enter one front-door
// portal. It is the portal admission contract and is separate from RBAC.
type PortalAccess struct {
	ID          string                  `json:"id"`
	UserID      string                  `json:"user_id"`
	AccountID   string                  `json:"account_id"`
	AccountType accountenum.AccountType `json:"account_type"`
	Portal      common.Portal           `json:"portal"`
	Status      PortalAccessStatus      `json:"status"`
	Grant       *common.LifecycleAction `json:"grant,omitempty"`
	Revocation  *common.LifecycleAction `json:"revocation,omitempty"`
	ExpiresAt   *time.Time              `json:"expires_at,omitempty"`

	common.AuditFields
}
