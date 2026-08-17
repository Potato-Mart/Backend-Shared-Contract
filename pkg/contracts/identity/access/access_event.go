package access

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/customers/wholesale/wholesale_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/account/account_enums"
)

// PortalAccessGrantedEvent is emitted when an account is granted access to a
// portal.
type PortalAccessGrantedEvent struct {
	PortalAccessID string                    `json:"portal_access_id"`
	UserID         string                    `json:"user_id"`
	AccountID      string                    `json:"account_id"`
	AccountType    account_enums.AccountType `json:"account_type"`
	Portal         identity_enums.Portal     `json:"portal"`
	GrantedBy      string                    `json:"granted_by,omitempty"`
	GrantedAt      time.Time                 `json:"granted_at"`
	ExpiresAt      *time.Time                `json:"expires_at,omitempty"`
	RequestID      string                    `json:"request_id,omitempty"`
}

// PortalAccessRevokedEvent is emitted when an account loses access to a
// portal.
type PortalAccessRevokedEvent struct {
	PortalAccessID string                    `json:"portal_access_id"`
	UserID         string                    `json:"user_id"`
	AccountID      string                    `json:"account_id"`
	AccountType    account_enums.AccountType `json:"account_type"`
	Portal         identity_enums.Portal     `json:"portal"`
	RevokedBy      string                    `json:"revoked_by,omitempty"`
	RevokedAt      time.Time                 `json:"revoked_at"`
	Reason         string                    `json:"reason,omitempty"`
	RequestID      string                    `json:"request_id,omitempty"`
}

// OrganisationAccessChangedEvent is emitted when wholesale organisation access
// changes status or role.
type OrganisationAccessChangedEvent struct {
	OrganisationAccessID      string                                   `json:"organisation_access_id"`
	WholesaleOrganisationCode string                                   `json:"wholesale_organisation_code"`
	UserID                    string                                   `json:"user_id"`
	AccountID                 string                                   `json:"account_id,omitempty"`
	RoleKey                   string                                   `json:"role_key,omitempty"`
	PreviousStatus            wholesale_enums.OrganisationAccessStatus `json:"previous_status,omitempty"`
	Status                    wholesale_enums.OrganisationAccessStatus `json:"status"`
	ChangedBy                 string                                   `json:"changed_by,omitempty"`
	ChangedAt                 time.Time                                `json:"changed_at"`
	Reason                    string                                   `json:"reason,omitempty"`
	RequestID                 string                                   `json:"request_id,omitempty"`
}
