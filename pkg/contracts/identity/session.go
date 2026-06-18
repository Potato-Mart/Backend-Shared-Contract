package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// LoginSession is a non-secret projection of an active login. A session is
// scoped to one portal and, for user sessions, one account/persona. The hashed
// refresh token never leaves the identity service – this type is only used for
// admin "active sessions" views and for "log out everywhere" features.
//
// APIs must validate portal, audience, and account_type consistently; this
// repository only defines the shared contract.
type LoginSession struct {
	ID                      string                   `json:"id"`
	UserID                  string                   `json:"user_id"`
	Portal                  enums.Portal             `json:"portal"`
	AccountID               string                   `json:"account_id,omitempty"`
	AccountType             enums.AccountType        `json:"account_type,omitempty"`
	Audience                string                   `json:"audience,omitempty"`
	Roles                   []string                 `json:"roles,omitempty"`
	Permissions             []string                 `json:"permissions,omitempty"`
	WholesaleOrganisationID string                   `json:"wholesale_organisation_id,omitempty"`
	MembershipID            string                   `json:"membership_id,omitempty"`
	RoleKey                 string                   `json:"role_key,omitempty"`
	DeviceID                string                   `json:"device_id,omitempty"`
	DeviceName              string                   `json:"device_name,omitempty"`
	DeviceType              enums.DeviceType         `json:"device_type,omitempty"`
	IPAddress               string                   `json:"ip_address,omitempty"`
	UserAgent               string                   `json:"user_agent,omitempty"`
	AuthMethod              enums.AuthMethod         `json:"auth_method,omitempty"`
	AuthAssuranceLevel      enums.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt           *time.Time               `json:"mfa_verified_at,omitempty"`
	RiskLevel               enums.SecurityRiskLevel  `json:"risk_level,omitempty"`
	IssuedAt                time.Time                `json:"issued_at"`
	LastSeenAt              time.Time                `json:"last_seen_at"`
	ExpiresAt               time.Time                `json:"expires_at"`
	RefreshTokenRotatedAt   *time.Time               `json:"refresh_token_rotated_at,omitempty"`
	RevokedAt               *time.Time               `json:"revoked_at,omitempty"`
	RevokedReason           string                   `json:"revoked_reason,omitempty"`
	History                 []shared.HistoryEntry    `json:"history,omitempty"`
}

// AuthTokenPair is returned on successful login or refresh.
type AuthTokenPair struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"` // always "Bearer"
	ExpiresIn    int       `json:"expires_in"` // seconds until access_token expiry
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
}
