package identity

import (
	"time"

	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/account"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/identity"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/security"
)

// AccessTokenClaims is a framework-agnostic token-claim shape. This contract
// intentionally does not import JWT libraries or implement token validation.
type AccessTokenClaims struct {
	Subject                   string                          `json:"sub"`
	UserID                    string                          `json:"user_id"`
	SessionID                 string                          `json:"session_id,omitempty"`
	AuthIdentityID            string                          `json:"auth_identity_id"`
	IdentityDomain            identityenum.IdentityDomain     `json:"identity_domain"`
	AccountID                 string                          `json:"account_id,omitempty"`
	AccountType               accountenum.AccountType         `json:"account_type,omitempty"`
	Portal                    accountenum.Portal              `json:"portal,omitempty"`
	Audience                  string                          `json:"aud,omitempty"`
	Roles                     []string                        `json:"roles,omitempty"`
	Permissions               []string                        `json:"permissions,omitempty"`
	WholesaleOrganisationCode string                          `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                          `json:"organisation_access_id,omitempty"`
	RetailCustomerNumber      string                          `json:"retail_customer_number,omitempty"`
	RoleKey                   string                          `json:"role_key,omitempty"`
	AuthAssuranceLevel        securityenum.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt             *time.Time                      `json:"mfa_verified_at,omitempty"`
	IssuedAt                  int64                           `json:"iat,omitempty"`
	ExpiresAt                 int64                           `json:"exp,omitempty"`
}
