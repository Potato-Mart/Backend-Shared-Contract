package identity

import "time"

// DeviceIP records one IP address observed for a device with enough metadata
// to answer when it first and last appeared.
type DeviceIP struct {
	IPAddress   string    `json:"ip_address"`
	FirstSeenAt time.Time `json:"first_seen_at"`
	LastSeenAt  time.Time `json:"last_seen_at"`
	SeenCount   int       `json:"seen_count,omitempty"`
}
