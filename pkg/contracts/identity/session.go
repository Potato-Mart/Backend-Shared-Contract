package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/contracts/shared"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/account"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/identity"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/security"
)

// LoginSession is a non-secret projection of an active login. A session is
// scoped to one portal and, for user sessions, one account/persona. The hashed
// refresh token never leaves the identity service – this type is only used for
// admin "active sessions" views and for "log out everywhere" features.
//
// APIs must validate portal, audience, and account_type consistently; this
// repository only defines the shared contract.
type LoginSession struct {
	ID                        string                          `json:"id"`
	UserID                    string                          `json:"user_id"`
	AuthIdentityID            string                          `json:"auth_identity_id"`
	IdentityDomain            identityenum.IdentityDomain     `json:"identity_domain"`
	Portal                    accountenum.Portal              `json:"portal"`
	AccountID                 string                          `json:"account_id,omitempty"`
	AccountType               accountenum.AccountType         `json:"account_type,omitempty"`
	Audience                  string                          `json:"audience,omitempty"`
	Roles                     []string                        `json:"roles,omitempty"`
	Permissions               []string                        `json:"permissions,omitempty"`
	WholesaleOrganisationCode string                          `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                          `json:"organisation_access_id,omitempty"`
	RoleKey                   string                          `json:"role_key,omitempty"`
	DeviceKey                 string                          `json:"device_key"`
	DeviceName                string                          `json:"device_name,omitempty"`
	DeviceType                identityenum.DeviceType         `json:"device_type,omitempty"`
	IPAddress                 string                          `json:"ip_address,omitempty"`
	UserAgent                 string                          `json:"user_agent,omitempty"`
	AuthMethod                securityenum.AuthMethod         `json:"auth_method,omitempty"`
	AuthAssuranceLevel        securityenum.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt             *time.Time                      `json:"mfa_verified_at,omitempty"`
	RiskLevel                 securityenum.SecurityRiskLevel  `json:"risk_level,omitempty"`
	IssuedAt                  time.Time                       `json:"issued_at"`
	LastSeenAt                time.Time                       `json:"last_seen_at"`
	ExpiresAt                 time.Time                       `json:"expires_at"`
	RefreshTokenRotatedAt     *time.Time                      `json:"refresh_token_rotated_at,omitempty"`
	RevokedAt                 *time.Time                      `json:"revoked_at,omitempty"`
	RevokedReason             string                          `json:"revoked_reason,omitempty"`
	History                   []shared.HistoryEntry           `json:"history,omitempty"`
}

// RefreshTokenRecord is the non-secret persistence/audit projection for a
// rotating refresh token. Implementations persist only a token hash.
type RefreshTokenRecord struct {
	ID             string                      `json:"id"`
	TokenHash      string                      `json:"-"`
	UserID         string                      `json:"user_id"`
	AuthIdentityID string                      `json:"auth_identity_id"`
	IdentityDomain identityenum.IdentityDomain `json:"identity_domain"`
	AccountID      string                      `json:"account_id"`
	Portal         accountenum.Portal          `json:"portal"`
	Audience       string                      `json:"audience"`
	IssuedAt       time.Time                   `json:"issued_at"`
	ExpiresAt      time.Time                   `json:"expires_at"`
	RotatedAt      *time.Time                  `json:"rotated_at,omitempty"`
	RevokedAt      *time.Time                  `json:"revoked_at,omitempty"`
	RevokedReason  string                      `json:"revoked_reason,omitempty"`
}
