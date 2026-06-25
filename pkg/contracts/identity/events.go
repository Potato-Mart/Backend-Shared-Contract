package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// UserAccountCreatedEvent is emitted when an account/persona is created for a
// canonical user.
type UserAccountCreatedEvent struct {
	UserID      string            `json:"user_id"`
	AccountID   string            `json:"account_id"`
	AccountType enums.AccountType `json:"account_type"`
	CreatedBy   string            `json:"created_by,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	RequestID   string            `json:"request_id,omitempty"`
}

// UserAccountStatusChangedEvent is emitted when an account/persona lifecycle
// status changes.
type UserAccountStatusChangedEvent struct {
	UserID         string              `json:"user_id"`
	AccountID      string              `json:"account_id"`
	AccountType    enums.AccountType   `json:"account_type"`
	PreviousStatus enums.AccountStatus `json:"previous_status,omitempty"`
	Status         enums.AccountStatus `json:"status"`
	ChangedBy      string              `json:"changed_by,omitempty"`
	ChangedAt      time.Time           `json:"changed_at"`
	Reason         string              `json:"reason,omitempty"`
	RequestID      string              `json:"request_id,omitempty"`
}

// PortalAccessGrantedEvent is emitted when an account is granted access to a
// portal.
type PortalAccessGrantedEvent struct {
	PortalAccessID string            `json:"portal_access_id"`
	UserID         string            `json:"user_id"`
	AccountID      string            `json:"account_id"`
	AccountType    enums.AccountType `json:"account_type"`
	Portal         enums.Portal      `json:"portal"`
	GrantedBy      string            `json:"granted_by,omitempty"`
	GrantedAt      time.Time         `json:"granted_at"`
	ExpiresAt      *time.Time        `json:"expires_at,omitempty"`
	RequestID      string            `json:"request_id,omitempty"`
}

// PortalAccessRevokedEvent is emitted when an account loses access to a
// portal.
type PortalAccessRevokedEvent struct {
	PortalAccessID string            `json:"portal_access_id"`
	UserID         string            `json:"user_id"`
	AccountID      string            `json:"account_id"`
	AccountType    enums.AccountType `json:"account_type"`
	Portal         enums.Portal      `json:"portal"`
	RevokedBy      string            `json:"revoked_by,omitempty"`
	RevokedAt      time.Time         `json:"revoked_at"`
	Reason         string            `json:"reason,omitempty"`
	RequestID      string            `json:"request_id,omitempty"`
}

// AuthIdentityLinkedEvent is emitted when a non-secret auth identity is linked
// to a canonical user.
type AuthIdentityLinkedEvent struct {
	AuthIdentityID string                     `json:"auth_identity_id"`
	UserID         string                     `json:"user_id"`
	Provider       enums.AuthIdentityProvider `json:"provider"`
	IdentityDomain enums.IdentityDomain       `json:"identity_domain"`
	LinkedBy       string                     `json:"linked_by,omitempty"`
	LinkedAt       time.Time                  `json:"linked_at"`
	RequestID      string                     `json:"request_id,omitempty"`
}

// RoleAssignmentGrantedEvent is emitted when a role key is granted in an
// account/portal scope.
type RoleAssignmentGrantedEvent struct {
	RoleAssignmentID string       `json:"role_assignment_id"`
	UserID           string       `json:"user_id"`
	AccountID        string       `json:"account_id"`
	Portal           enums.Portal `json:"portal"`
	RoleKey          string       `json:"role_key"`
	ScopeType        string       `json:"scope_type,omitempty"`
	ScopeID          string       `json:"scope_id,omitempty"`
	GrantedBy        string       `json:"granted_by,omitempty"`
	GrantedAt        time.Time    `json:"granted_at"`
	RequestID        string       `json:"request_id,omitempty"`
}

// RoleAssignmentRevokedEvent is emitted when a role assignment is revoked.
type RoleAssignmentRevokedEvent struct {
	RoleAssignmentID string       `json:"role_assignment_id"`
	UserID           string       `json:"user_id"`
	AccountID        string       `json:"account_id"`
	Portal           enums.Portal `json:"portal"`
	RoleKey          string       `json:"role_key"`
	ScopeType        string       `json:"scope_type,omitempty"`
	ScopeID          string       `json:"scope_id,omitempty"`
	RevokedBy        string       `json:"revoked_by,omitempty"`
	RevokedAt        time.Time    `json:"revoked_at"`
	Reason           string       `json:"reason,omitempty"`
	RequestID        string       `json:"request_id,omitempty"`
}

// OrganisationAccessChangedEvent is emitted when wholesale organisation access
// changes status or role.
type OrganisationAccessChangedEvent struct {
	OrganisationAccessID    string                         `json:"organisation_access_id"`
	WholesaleOrganisationID string                         `json:"wholesale_organisation_id"`
	UserID                  string                         `json:"user_id"`
	AccountID               string                         `json:"account_id,omitempty"`
	RoleKey                 string                         `json:"role_key,omitempty"`
	PreviousStatus          enums.OrganisationAccessStatus `json:"previous_status,omitempty"`
	Status                  enums.OrganisationAccessStatus `json:"status"`
	ChangedBy               string                         `json:"changed_by,omitempty"`
	ChangedAt               time.Time                      `json:"changed_at"`
	Reason                  string                         `json:"reason,omitempty"`
	RequestID               string                         `json:"request_id,omitempty"`
}
