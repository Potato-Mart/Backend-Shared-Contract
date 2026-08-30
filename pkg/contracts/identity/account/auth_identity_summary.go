package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/identity/account/account_enums"
	"time"
)

// AuthIdentitySummary is the compact admin-safe projection of an auth identity.
type AuthIdentitySummary struct {
	ID             string                             `json:"id"`
	UserID         string                             `json:"user_id"`
	Provider       account_enums.AuthIdentityProvider `json:"provider"`
	IdentityDomain security_enums.IdentityDomain      `json:"identity_domain"`
	Email          string                             `json:"email,omitempty"`
	EmailVerified  bool                               `json:"email_verified,omitempty"`
	Status         account_enums.AuthIdentityStatus   `json:"status"`
	LastLoginAt    *time.Time                         `json:"last_login_at,omitempty"`
}
