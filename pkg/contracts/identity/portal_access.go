package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/common"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/account"
)

// PortalAccess records whether an account/persona may enter one front-door
// portal. It is the portal admission contract and is separate from RBAC.
type PortalAccess struct {
	ID          string                         `json:"id"`
	UserID      string                         `json:"user_id"`
	AccountID   string                         `json:"account_id"`
	AccountType accountenum.AccountType        `json:"account_type"`
	Portal      accountenum.Portal             `json:"portal"`
	Status      accountenum.PortalAccessStatus `json:"status"`
	Grant       *common.LifecycleAction        `json:"grant,omitempty"`
	Revocation  *common.LifecycleAction        `json:"revocation,omitempty"`
	ExpiresAt   *time.Time                     `json:"expires_at,omitempty"`

	common.AuditFields
}
