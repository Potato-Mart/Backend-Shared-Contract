package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/account"
)

// RoleAssignment grants a role key to a user in an account, portal, and
// optional business scope such as a wholesale organisation.
type RoleAssignment struct {
	ID         string                  `json:"id"`
	UserID     string                  `json:"user_id"`
	AccountID  string                  `json:"account_id"`
	Portal     accountenum.Portal      `json:"portal"`
	RoleKey    string                  `json:"role_key"`
	ScopeType  string                  `json:"scope_type,omitempty"`
	ScopeID    string                  `json:"scope_id,omitempty"`
	ExpiresAt  *time.Time              `json:"expires_at,omitempty"`
	Grant      *common.LifecycleAction `json:"grant,omitempty"`
	Revocation *common.LifecycleAction `json:"revocation,omitempty"`

	common.AuditFields
}
