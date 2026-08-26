package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/device"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/identity/identity_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/account/account_enums"
	"time"
)

// UserDevice is a non-secret projection of a browser, mobile app, or
// other client device observed for a user. The identity service owns the
// fingerprinting and trust decisions; this contract lets admin and security
// surfaces show every known device and the IP addresses seen for it.
type UserDevice struct {
	ID                 string                              `json:"id"`
	UserID             string                              `json:"user_id"`
	Portal             identity_enums.Portal               `json:"portal,omitempty"`
	DeviceName         string                              `json:"device_name,omitempty"`
	DeviceType         account_enums.DeviceType            `json:"device_type,omitempty"`
	PreferredLanguage  account_enums.UserPreferredLanguage `json:"preferred_language,omitempty"`
	IPHistory          []DeviceIP                          `json:"ip_history,omitempty"`
	Trusted            bool                                `json:"trusted"`
	RiskLevel          security_enums.SecurityRiskLevel    `json:"risk_level,omitempty"`
	TrustReason        string                              `json:"trust_reason,omitempty"`
	FirstSeenAt        time.Time                           `json:"first_seen_at"`
	LastSeenAt         time.Time                           `json:"last_seen_at"`
	LastLoginAt        *time.Time                          `json:"last_login_at,omitempty"`
	LastLoginIP        string                              `json:"last_login_ip,omitempty"`
	LastRiskReviewedAt *time.Time                          `json:"last_risk_reviewed_at,omitempty"`
	RevokedAt          *time.Time                          `json:"revoked_at,omitempty"`
	RevokedReason      string                              `json:"revoked_reason,omitempty"`
	History            []security.HistoryEntry             `json:"history,omitempty"`
	device.DeviceRecord
}
