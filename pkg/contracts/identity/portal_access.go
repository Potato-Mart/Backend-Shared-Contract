package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/account"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/security"
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

// PortalAccessDecision is the contract-only result shape for login or portal
// admission resolution.
type PortalAccessDecision struct {
	Allowed                    bool                            `json:"allowed"`
	DenyReason                 string                          `json:"deny_reason,omitempty"`
	UserID                     string                          `json:"user_id,omitempty"`
	AccountID                  string                          `json:"account_id,omitempty"`
	AccountType                accountenum.AccountType         `json:"account_type,omitempty"`
	Portal                     accountenum.Portal              `json:"portal,omitempty"`
	RequiresMFA                bool                            `json:"requires_mfa,omitempty"`
	RequiredAuthAssuranceLevel securityenum.AuthAssuranceLevel `json:"required_auth_assurance_level,omitempty"`
	Roles                      []string                        `json:"roles,omitempty"`
	Permissions                []string                        `json:"permissions,omitempty"`
}
