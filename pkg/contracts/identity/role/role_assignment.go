package role

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"time"
)

// RoleAssignment grants a role key to a user in an account, portal, and
// optional business scope such as a wholesale organisation.
type RoleAssignment struct {
	ID         string                  `json:"id"`
	UserID     string                  `json:"user_id"`
	AccountID  string                  `json:"account_id"`
	Portal     common.Portal           `json:"portal"`
	RoleKey    string                  `json:"role_key"`
	ScopeType  string                  `json:"scope_type,omitempty"`
	ScopeID    string                  `json:"scope_id,omitempty"`
	ExpiresAt  *time.Time              `json:"expires_at,omitempty"`
	Grant      *common.LifecycleAction `json:"grant,omitempty"`
	Revocation *common.LifecycleAction `json:"revocation,omitempty"`

	common.AuditFields
}
