package access

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security/security_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/identity/access/access_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/identity/account/account_enums"
)

// AccessTokenClaims is a framework-agnostic token-claim shape. This contract
// intentionally does not import JWT libraries or implement token validation.
//
// ScopeLevel, CountryCode, MarketCodes, and DepotCodes are the flat geographic
// scope of a workforce principal. They are absent on customer and
// service-to-service tokens. Consumers fail closed: a workforce token with no
// valid scope_level is rejected rather than treated as global.
type AccessTokenClaims struct {
	Subject                   string                            `json:"sub"`
	UserID                    string                            `json:"user_id"`
	SessionID                 string                            `json:"session_id,omitempty"`
	AuthIdentityID            string                            `json:"auth_identity_id"`
	IdentityDomain            security_enums.IdentityDomain     `json:"identity_domain"`
	AccountID                 string                            `json:"account_id,omitempty"`
	AccountType               account_enums.AccountType         `json:"account_type,omitempty"`
	Portal                    identity_enums.Portal             `json:"portal,omitempty"`
	Audience                  string                            `json:"aud,omitempty"`
	Roles                     []string                          `json:"roles,omitempty"`
	Permissions               []string                          `json:"permissions,omitempty"`
	WholesaleOrganisationCode string                            `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                            `json:"organisation_access_id,omitempty"`
	RetailCustomerNumber      string                            `json:"retail_customer_number,omitempty"`
	RoleKey                   string                            `json:"role_key,omitempty"`
	AuthAssuranceLevel        security_enums.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt             *time.Time                        `json:"mfa_verified_at,omitempty"`
	IssuedAt                  int64                             `json:"iat,omitempty"`
	ExpiresAt                 int64                             `json:"exp,omitempty"`
	ScopeLevel                access_enums.ScopeLevel           `json:"scope_level,omitempty"`
	CountryCode               geography.CountryCode             `json:"country_code,omitempty"`
	MarketCodes               []string                          `json:"market_codes,omitempty"`
	DepotCodes                []string                          `json:"depot_codes,omitempty"`
}
