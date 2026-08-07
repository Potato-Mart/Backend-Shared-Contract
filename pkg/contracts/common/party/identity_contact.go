package party

// PersonName captures the common name fields used by person-like profiles.
type PersonName struct {
	FirstName     string `json:"first_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	DisplayName   string `json:"display_name,omitempty"`
	PreferredName string `json:"preferred_name,omitempty"`
}

// ContactChannels groups common customer contact channels.
type ContactChannels struct {
	Email           string            `json:"email,omitempty"`
	Phone           string            `json:"phone,omitempty"`
	ExternalHandles map[string]string `json:"external_handles,omitempty"`
}
