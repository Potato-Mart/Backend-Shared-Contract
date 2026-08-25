package account

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/account/account_enums"
)

// AuthIdentity is a non-secret projection of a login identity/provider attached
// to a canonical user. It must not contain password hashes, token hashes,
// OAuth tokens, passkey private material, or IdP secrets.
type AuthIdentity struct {
	ID              string                             `json:"id"`
	UserID          string                             `json:"user_id"`
	Provider        account_enums.AuthIdentityProvider `json:"provider"`
	IdentityDomain  security_enums.IdentityDomain      `json:"identity_domain"`
	ProviderSubject string                             `json:"provider_subject,omitempty"`
	Email           string                             `json:"email,omitempty"`
	EmailVerified   bool                               `json:"email_verified,omitempty"`
	Status          account_enums.AuthIdentityStatus   `json:"status"`
	LastLoginAt     *time.Time                         `json:"last_login_at,omitempty"`
	DisabledAt      *time.Time                         `json:"disabled_at,omitempty"`
	DisabledReason  string                             `json:"disabled_reason,omitempty"`

	audit.AuditFields
}
