package role

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/identity/identity_enums"
	"time"
)

// RoleAssignmentRevokedEvent is emitted when a role assignment is revoked.
type RoleAssignmentRevokedEvent struct {
	RoleAssignmentID string                `json:"role_assignment_id"`
	UserID           string                `json:"user_id"`
	AccountID        string                `json:"account_id"`
	Portal           identity_enums.Portal `json:"portal"`
	RoleKey          RoleCode              `json:"role_key"`
	ScopeType        string                `json:"scope_type,omitempty"`
	ScopeID          string                `json:"scope_id,omitempty"`
	RevokedBy        string                `json:"revoked_by,omitempty"`
	RevokedAt        time.Time             `json:"revoked_at"`
	Reason           string                `json:"reason,omitempty"`
	RequestID        string                `json:"request_id,omitempty"`
}
