package role

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/identity/identity_enums"
)

// RoleAssignment grants a role key to a user in an account, portal, and
// optional business scope such as a wholesale organisation.
//
// RoleKey is deliberately an open RoleCode rather than a closed enum: Portal
// selects the vocabulary. A control-portal grant carries one of the six
// workforce role_enums.UserRole keys and a wholesale-portal grant carries a
// wholesale buyer role key, and the service that owns each catalogue
// validates the code.
type RoleAssignment struct {
	ID         string                 `json:"id"`
	UserID     string                 `json:"user_id"`
	AccountID  string                 `json:"account_id"`
	Portal     identity_enums.Portal  `json:"portal"`
	RoleKey    RoleCode               `json:"role_key"`
	ScopeType  string                 `json:"scope_type,omitempty"`
	ScopeID    string                 `json:"scope_id,omitempty"`
	ExpiresAt  *time.Time             `json:"expires_at,omitempty"`
	Grant      *audit.LifecycleAction `json:"grant,omitempty"`
	Revocation *audit.LifecycleAction `json:"revocation,omitempty"`

	audit.AuditFields
}
