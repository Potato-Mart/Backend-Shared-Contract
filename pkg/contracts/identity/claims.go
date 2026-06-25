package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// AccessTokenClaims is a framework-agnostic token-claim shape. This contract
// intentionally does not import JWT libraries or implement token validation.
type AccessTokenClaims struct {
	Subject                 string                   `json:"sub"`
	UserID                  string                   `json:"user_id"`
	AccountID               string                   `json:"account_id,omitempty"`
	AccountType             enums.AccountType        `json:"account_type,omitempty"`
	Portal                  enums.Portal             `json:"portal,omitempty"`
	Audience                string                   `json:"aud,omitempty"`
	Roles                   []string                 `json:"roles,omitempty"`
	Permissions             []string                 `json:"permissions,omitempty"`
	WholesaleOrganisationID string                   `json:"wholesale_organisation_id,omitempty"`
	OrganisationAccessID    string                   `json:"organisation_access_id,omitempty"`
	RoleKey                 string                   `json:"role_key,omitempty"`
	AuthAssuranceLevel      enums.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt           *time.Time               `json:"mfa_verified_at,omitempty"`
	IssuedAt                int64                    `json:"iat,omitempty"`
	ExpiresAt               int64                    `json:"exp,omitempty"`
}
