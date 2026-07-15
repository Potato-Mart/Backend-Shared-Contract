package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/shared"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/account"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/identity"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/security"
)

// UserDevice is a non-secret projection of a browser, mobile app, or
// other client device observed for a user. The identity service owns the
// fingerprinting and trust decisions; this contract lets admin and security
// surfaces show every known device and the IP addresses seen for it.
type UserDevice struct {
	ID                 string                             `json:"id"`
	UserID             string                             `json:"user_id"`
	Portal             accountenum.Portal                 `json:"portal,omitempty"`
	DeviceName         string                             `json:"device_name,omitempty"`
	DeviceType         identityenum.DeviceType            `json:"device_type,omitempty"`
	PreferredLanguage  identityenum.UserPreferredLanguage `json:"preferred_language,omitempty"`
	IPHistory          []DeviceIP                         `json:"ip_history,omitempty"`
	Trusted            bool                               `json:"trusted"`
	RiskLevel          securityenum.SecurityRiskLevel     `json:"risk_level,omitempty"`
	TrustReason        string                             `json:"trust_reason,omitempty"`
	FirstSeenAt        time.Time                          `json:"first_seen_at"`
	LastSeenAt         time.Time                          `json:"last_seen_at"`
	LastLoginAt        *time.Time                         `json:"last_login_at,omitempty"`
	LastRiskReviewedAt *time.Time                         `json:"last_risk_reviewed_at,omitempty"`
	RevokedAt          *time.Time                         `json:"revoked_at,omitempty"`
	RevokedReason      string                             `json:"revoked_reason,omitempty"`
	History            []shared.HistoryEntry              `json:"history,omitempty"`
	shared.DeviceRecord
}

// UserDeviceSeenEvent is emitted when a known user authenticates or makes an
// identified request from a device/IP pair. Consumers can use it to maintain
// device inventories, IP history, and security alerts.
type UserDeviceSeenEvent struct {
	UserID    string             `json:"user_id"`
	SessionID string             `json:"session_id,omitempty"`
	Portal    accountenum.Portal `json:"portal,omitempty"`
	DeviceKey string             `json:"device_key,omitempty"`
	IPAddress string             `json:"ip_address,omitempty"`
	UserAgent string             `json:"user_agent,omitempty"`
	SeenAt    time.Time          `json:"seen_at"`
	RequestID string             `json:"request_id,omitempty"`
}
