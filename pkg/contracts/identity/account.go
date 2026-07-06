package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/account"
)

// UserAccount is an account/persona record attached to one canonical user.
// AccountType controls portal admission; roles and permissions are assigned
// separately after an account has entered a portal.
type UserAccount struct {
	ID              string                    `json:"id"`
	UserID          string                    `json:"user_id"`
	AccountType     accountenum.AccountType   `json:"account_type"`
	Status          accountenum.AccountStatus `json:"status"`
	DisplayName     string                    `json:"display_name,omitempty"`
	PrimaryEmail    string                    `json:"primary_email,omitempty"`
	DefaultPortal   accountenum.Portal        `json:"default_portal,omitempty"`
	LastAccessedAt  *time.Time                `json:"last_accessed_at,omitempty"`
	SuspendedAt     *time.Time                `json:"suspended_at,omitempty"`
	SuspendedReason string                    `json:"suspended_reason,omitempty"`
	ClosedAt        *time.Time                `json:"closed_at,omitempty"`

	common.AuditFields
}

// UserAccountSummary is the compact account/persona projection returned with
// user and access-resolution responses.
type UserAccountSummary struct {
	ID             string                    `json:"id"`
	UserID         string                    `json:"user_id"`
	AccountType    accountenum.AccountType   `json:"account_type"`
	Status         accountenum.AccountStatus `json:"status"`
	DisplayName    string                    `json:"display_name,omitempty"`
	PrimaryEmail   string                    `json:"primary_email,omitempty"`
	AllowedPortals []accountenum.Portal      `json:"allowed_portals,omitempty"`
}
