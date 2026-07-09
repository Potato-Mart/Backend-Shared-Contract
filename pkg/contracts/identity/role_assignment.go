package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/account"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/security"
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

// EffectivePermissionSet is the resolved RBAC projection for an admitted
// account in a portal and optional business scope.
type EffectivePermissionSet struct {
	UserID                     string                          `json:"user_id"`
	AccountID                  string                          `json:"account_id"`
	AccountType                accountenum.AccountType         `json:"account_type,omitempty"`
	Portal                     accountenum.Portal              `json:"portal"`
	ScopeType                  string                          `json:"scope_type,omitempty"`
	ScopeID                    string                          `json:"scope_id,omitempty"`
	WholesaleOrganisationCode  string                          `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID       string                          `json:"organisation_access_id,omitempty"`
	Roles                      []string                        `json:"roles,omitempty"`
	Permissions                []string                        `json:"permissions,omitempty"`
	RequiresMFA                bool                            `json:"requires_mfa,omitempty"`
	RequiredAuthAssuranceLevel securityenum.AuthAssuranceLevel `json:"required_auth_assurance_level,omitempty"`
	EvaluatedAt                *time.Time                      `json:"evaluated_at,omitempty"`
}
