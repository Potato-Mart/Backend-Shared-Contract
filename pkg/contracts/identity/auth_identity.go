package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/identity"
)

// AuthIdentity is a non-secret projection of a login identity/provider attached
// to a canonical user. It must not contain password hashes, token hashes,
// OAuth tokens, passkey private material, or IdP secrets.
type AuthIdentity struct {
	ID              string                            `json:"id"`
	UserID          string                            `json:"user_id"`
	Provider        identityenum.AuthIdentityProvider `json:"provider"`
	IdentityDomain  identityenum.IdentityDomain       `json:"identity_domain"`
	ProviderSubject string                            `json:"provider_subject,omitempty"`
	Email           string                            `json:"email,omitempty"`
	EmailVerified   bool                              `json:"email_verified,omitempty"`
	Status          identityenum.AuthIdentityStatus   `json:"status"`
	LastLoginAt     *time.Time                        `json:"last_login_at,omitempty"`
	DisabledAt      *time.Time                        `json:"disabled_at,omitempty"`
	DisabledReason  string                            `json:"disabled_reason,omitempty"`

	common.AuditFields
}

// AuthIdentitySummary is the compact admin-safe projection of an auth identity.
type AuthIdentitySummary struct {
	ID             string                            `json:"id"`
	UserID         string                            `json:"user_id"`
	Provider       identityenum.AuthIdentityProvider `json:"provider"`
	IdentityDomain identityenum.IdentityDomain       `json:"identity_domain"`
	Email          string                            `json:"email,omitempty"`
	EmailVerified  bool                              `json:"email_verified,omitempty"`
	Status         identityenum.AuthIdentityStatus   `json:"status"`
	LastLoginAt    *time.Time                        `json:"last_login_at,omitempty"`
}

// UserConnectedIdentities is an account-connection projection: every login
// provider linked to one canonical user. The backend assembles it from the
// user's AuthIdentity rows; the frontend renders it on the "connected accounts"
// screen (link Google / Apple / Line / Microsoft / Discord, etc.).
type UserConnectedIdentities struct {
	UserID     string                `json:"user_id"`
	Identities []AuthIdentitySummary `json:"identities"`
}
