package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// LoginSession is a non-secret projection of an active staff or
// customer login. The hashed refresh token never leaves the identity
// service – this type is only used for admin "active sessions" views
// and for "log out everywhere" features.
type LoginSession struct {
	ID                    string                   `json:"id"`
	UserID                string                   `json:"user_id"`
	Portal                string                   `json:"portal"` // "control" | "store" | "partner"
	DeviceID              string                   `json:"device_id,omitempty"`
	DeviceName            string                   `json:"device_name,omitempty"`
	DeviceType            string                   `json:"device_type,omitempty"`
	IPAddress             string                   `json:"ip_address,omitempty"`
	UserAgent             string                   `json:"user_agent,omitempty"`
	AuthMethod            enums.AuthMethod         `json:"auth_method,omitempty"`
	AuthAssuranceLevel    enums.AuthAssuranceLevel `json:"auth_assurance_level,omitempty"`
	MFAVerifiedAt         *time.Time               `json:"mfa_verified_at,omitempty"`
	RiskLevel             enums.SecurityRiskLevel  `json:"risk_level,omitempty"`
	IssuedAt              time.Time                `json:"issued_at"`
	LastSeenAt            time.Time                `json:"last_seen_at"`
	ExpiresAt             time.Time                `json:"expires_at"`
	RefreshTokenRotatedAt *time.Time               `json:"refresh_token_rotated_at,omitempty"`
	RevokedAt             *time.Time               `json:"revoked_at,omitempty"`
	RevokedReason         string                   `json:"revoked_reason,omitempty"`
}

// AuthTokenPair is returned on successful login or refresh.
type AuthTokenPair struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"` // always "Bearer"
	ExpiresIn    int       `json:"expires_in"` // seconds until access_token expiry
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
}
