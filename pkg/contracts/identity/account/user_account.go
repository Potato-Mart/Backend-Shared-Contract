package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/identity/account/account_enums"
	"time"
)

// UserAccount is an account/persona record attached to one canonical user.
// AccountType controls portal admission; roles and permissions are assigned
// separately after an account has entered a portal.
type UserAccount struct {
	ID              string                      `json:"id"`
	UserID          string                      `json:"user_id"`
	AccountType     account_enums.AccountType   `json:"account_type"`
	Status          account_enums.AccountStatus `json:"status"`
	DisplayName     string                      `json:"display_name,omitempty"`
	PrimaryEmail    string                      `json:"primary_email,omitempty"`
	DefaultPortal   identity_enums.Portal       `json:"default_portal,omitempty"`
	LastAccessedAt  *time.Time                  `json:"last_accessed_at,omitempty"`
	SuspendedAt     *time.Time                  `json:"suspended_at,omitempty"`
	SuspendedReason string                      `json:"suspended_reason,omitempty"`
	ClosedAt        *time.Time                  `json:"closed_at,omitempty"`

	audit.AuditFields
}
