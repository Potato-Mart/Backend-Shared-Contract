package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/identity/identity_enums"
	"time"
)

// UserDeviceSeenEvent is emitted when a known user authenticates or makes an
// identified request from a device/IP pair. Consumers can use it to maintain
// device inventories, IP history, and security alerts.
type UserDeviceSeenEvent struct {
	UserID    string                `json:"user_id"`
	SessionID string                `json:"session_id,omitempty"`
	Portal    identity_enums.Portal `json:"portal,omitempty"`
	DeviceKey string                `json:"device_key"`
	IPAddress string                `json:"ip_address,omitempty"`
	UserAgent string                `json:"user_agent,omitempty"`
	SeenAt    time.Time             `json:"seen_at"`
	RequestID string                `json:"request_id,omitempty"`
}
