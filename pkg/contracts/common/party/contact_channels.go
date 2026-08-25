package party

// ContactChannels groups common customer contact channels.
type ContactChannels struct {
	Email           string            `json:"email,omitempty"`
	Phone           string            `json:"phone,omitempty"`
	ExternalHandles map[string]string `json:"external_handles,omitempty"`
}
