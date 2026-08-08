package access

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/security"

	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/security/security_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/identity/account/account_enums"
)

// LoginSession is a non-secret projection of an active login. A session is
// scoped to one portal and, for user sessions, one account/persona. The hashed
// refresh token never leaves the identity service – this type is only used for
// admin "active sessions" views and for "log out everywhere" features.
//
// APIs must validate portal, audience, and account_type consistently; this
// repository only defines the shared contract.
type LoginSession struct {
	ID                        string                            `json:"id"`
	UserID                    string                            `json:"user_id"`
	AuthIdentityID            string                            `json:"auth_identity_id"`
	IdentityDomain            security_enums.IdentityDomain     `json:"identity_domain"`
	Portal                    identity_enums.Portal             `json:"portal"`
	AccountID                 string                            `json:"account_id,omitempty"`
	AccountType               account_enums.AccountType         `json:"account_type,omitempty"`
	Audience                  string                            `json:"audience,omitempty"`
	Roles                     []string                          `json:"roles,omitempty"`
	Permissions               []string                          `json:"permissions,omitempty"`
	WholesaleOrganisationCode string                            `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                            `json:"organisation_access_id,omitempty"`
	RoleKey                   string                            `json:"role_key,omitempty"`
	DeviceKey                 string                            `json:"device_key"`
	DeviceName                string                            `json:"device_name,omitempty"`
	DeviceType                account_enums.DeviceType          `json:"device_type,omitempty"`
	IPAddress                 string                            `json:"ip_address,omitempty"`
	UserAgent                 string                            `json:"user_agent,omitempty"`
	AuthMethod                security_enums.AuthMethod         `json:"auth_method,omitempty"`
	AuthAssuranceLevel        security_enums.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt             *time.Time                        `json:"mfa_verified_at,omitempty"`
	RiskLevel                 security_enums.SecurityRiskLevel  `json:"risk_level,omitempty"`
	IssuedAt                  time.Time                         `json:"issued_at"`
	LastSeenAt                time.Time                         `json:"last_seen_at"`
	ExpiresAt                 time.Time                         `json:"expires_at"`
	RefreshTokenRotatedAt     *time.Time                        `json:"refresh_token_rotated_at,omitempty"`
	RevokedAt                 *time.Time                        `json:"revoked_at,omitempty"`
	RevokedReason             string                            `json:"revoked_reason,omitempty"`
	History                   []security.HistoryEntry           `json:"history,omitempty"`
}

// RefreshTokenRecord is the non-secret persistence/audit projection for a
// rotating refresh token. Implementations persist only a token hash.
type RefreshTokenRecord struct {
	ID             string                        `json:"id"`
	TokenHash      string                        `json:"-"`
	UserID         string                        `json:"user_id"`
	AuthIdentityID string                        `json:"auth_identity_id"`
	IdentityDomain security_enums.IdentityDomain `json:"identity_domain"`
	AccountID      string                        `json:"account_id"`
	Portal         identity_enums.Portal         `json:"portal"`
	Audience       string                        `json:"audience"`
	IssuedAt       time.Time                     `json:"issued_at"`
	ExpiresAt      time.Time                     `json:"expires_at"`
	RotatedAt      *time.Time                    `json:"rotated_at,omitempty"`
	RevokedAt      *time.Time                    `json:"revoked_at,omitempty"`
	RevokedReason  string                        `json:"revoked_reason,omitempty"`
}
