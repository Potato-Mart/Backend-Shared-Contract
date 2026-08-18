package role

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/identity/identity_enums"
)

// RoleAssignment grants a role key to a user in an account, portal, and
// optional business scope such as a wholesale organisation.
type RoleAssignment struct {
	ID         string                 `json:"id"`
	UserID     string                 `json:"user_id"`
	AccountID  string                 `json:"account_id"`
	Portal     identity_enums.Portal  `json:"portal"`
	RoleKey    string                 `json:"role_key"`
	ScopeType  string                 `json:"scope_type,omitempty"`
	ScopeID    string                 `json:"scope_id,omitempty"`
	ExpiresAt  *time.Time             `json:"expires_at,omitempty"`
	Grant      *audit.LifecycleAction `json:"grant,omitempty"`
	Revocation *audit.LifecycleAction `json:"revocation,omitempty"`

	audit.AuditFields
}
