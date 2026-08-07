package shared

// DeviceRecord captures non-secret client attributes that can be reused by
// identity device records and order source attribution.
type DeviceRecord struct {
	DeviceKey string `json:"device_key"`
	IPAddress string `json:"ip_address,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	OS        string `json:"os,omitempty"`
	Browser   string `json:"browser,omitempty"`
}
