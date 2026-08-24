package order

// POSAttribution carries first-class in-store sale attribution (depot, event,
// register, terminal, daily session, operator, platform, form factor) on the
// order's source device.
//
// DepotCode is the trading site: depots are the only site identity in the
// platform. SessionID names the register's shared daily session, while
// OperatorUserID records who rang this particular sale inside it.
type POSAttribution struct {
	DepotCode      string `json:"depot_code,omitempty"`
	EventID        string `json:"event_id,omitempty"`
	RegisterID     string `json:"register_id,omitempty"`
	TerminalID     string `json:"terminal_id,omitempty"`
	SessionID      string `json:"session_id,omitempty"`
	OperatorUserID string `json:"operator_user_id,omitempty"`
	Platform       string `json:"platform,omitempty"`
	FormFactor     string `json:"form_factor,omitempty"`
}
