package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// PortalAccess records whether an account/persona may enter one front-door
// portal. It is the portal admission contract and is separate from RBAC.
type PortalAccess struct {
	ID          string                   `json:"id"`
	UserID      string                   `json:"user_id"`
	AccountID   string                   `json:"account_id"`
	AccountType enums.AccountType        `json:"account_type"`
	Portal      enums.Portal             `json:"portal"`
	Status      enums.PortalAccessStatus `json:"status"`
	Grant       *common.LifecycleAction  `json:"grant,omitempty"`
	Revocation  *common.LifecycleAction  `json:"revocation,omitempty"`
	ExpiresAt   *time.Time               `json:"expires_at,omitempty"`

	common.AuditFields `bson:",inline"`
}

// PortalAccessDecision is the contract-only result shape for login or portal
// admission resolution.
type PortalAccessDecision struct {
	Allowed                    bool                     `json:"allowed"`
	DenyReason                 string                   `json:"deny_reason,omitempty"`
	UserID                     string                   `json:"user_id,omitempty"`
	AccountID                  string                   `json:"account_id,omitempty"`
	AccountType                enums.AccountType        `json:"account_type,omitempty"`
	Portal                     enums.Portal             `json:"portal,omitempty"`
	RequiresMFA                bool                     `json:"requires_mfa,omitempty"`
	RequiredAuthAssuranceLevel enums.AuthAssuranceLevel `json:"required_auth_assurance_level,omitempty"`
	Roles                      []string                 `json:"roles,omitempty"`
	Permissions                []string                 `json:"permissions,omitempty"`
}

// ResolvePortalAccessRequest asks an identity service to resolve whether a
// user/account can enter a portal. This contract does not implement the
// resolution logic.
type ResolvePortalAccessRequest struct {
	UserID                  string            `json:"user_id,omitempty"`
	AccountID               string            `json:"account_id,omitempty"`
	AccountType             enums.AccountType `json:"account_type,omitempty"`
	Portal                  enums.Portal      `json:"portal"`
	Audience                string            `json:"audience,omitempty"`
	AuthIdentityID          string            `json:"auth_identity_id,omitempty"`
	WholesaleOrganisationID string            `json:"wholesale_organisation_id,omitempty"`
	MembershipID            string            `json:"membership_id,omitempty"`

	Context *IdentityRequestContext `json:"context,omitempty"`
}

// GrantPortalAccessRequest grants or activates portal access for an account.
type GrantPortalAccessRequest struct {
	UserID      string                   `json:"user_id"`
	AccountID   string                   `json:"account_id"`
	AccountType enums.AccountType        `json:"account_type"`
	Portal      enums.Portal             `json:"portal"`
	Status      enums.PortalAccessStatus `json:"status,omitempty"`
	ExpiresAt   *time.Time               `json:"expires_at,omitempty"`

	Context *IdentityRequestContext `json:"context,omitempty"`
}

// RevokePortalAccessRequest revokes portal access for an account.
type RevokePortalAccessRequest struct {
	PortalAccessID string       `json:"portal_access_id,omitempty"`
	UserID         string       `json:"user_id,omitempty"`
	AccountID      string       `json:"account_id,omitempty"`
	Portal         enums.Portal `json:"portal,omitempty"`

	Context *IdentityRequestContext `json:"context,omitempty"`
}
