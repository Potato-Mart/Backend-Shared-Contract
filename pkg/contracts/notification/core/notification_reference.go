package core

// NotificationReference links a notification to a domain fact without
// coupling the notification model to the referenced aggregate.
type NotificationReference struct {
	Kind  string `json:"kind"`
	ID    string `json:"id,omitempty"`
	Code  string `json:"code,omitempty"`
	Label string `json:"label,omitempty"`
}
