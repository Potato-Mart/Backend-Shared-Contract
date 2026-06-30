package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

// AuthIdentity is a non-secret projection of a login identity/provider attached
// to a canonical user. It must not contain password hashes, token hashes,
// OAuth tokens, passkey private material, or IdP secrets.
type AuthIdentity struct {
	ID              string                     `json:"id"`
	UserID          string                     `json:"user_id"`
	Provider        enums.AuthIdentityProvider `json:"provider"`
	IdentityDomain  enums.IdentityDomain       `json:"identity_domain"`
	ProviderSubject string                     `json:"provider_subject,omitempty"`
	Email           string                     `json:"email,omitempty"`
	EmailVerified   bool                       `json:"email_verified,omitempty"`
	Status          enums.AuthIdentityStatus   `json:"status"`
	LastLoginAt     *time.Time                 `json:"last_login_at,omitempty"`
	DisabledAt      *time.Time                 `json:"disabled_at,omitempty"`
	DisabledReason  string                     `json:"disabled_reason,omitempty"`

	common.AuditFields
}

// AuthIdentitySummary is the compact admin-safe projection of an auth identity.
type AuthIdentitySummary struct {
	ID             string                     `json:"id"`
	UserID         string                     `json:"user_id"`
	Provider       enums.AuthIdentityProvider `json:"provider"`
	IdentityDomain enums.IdentityDomain       `json:"identity_domain"`
	Email          string                     `json:"email,omitempty"`
	EmailVerified  bool                       `json:"email_verified,omitempty"`
	Status         enums.AuthIdentityStatus   `json:"status"`
	LastLoginAt    *time.Time                 `json:"last_login_at,omitempty"`
}
