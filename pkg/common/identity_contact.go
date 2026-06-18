package common

// IdentityLink carries shared references from a business profile to the
// canonical identity/account contracts. It contains identifiers only and no
// credential or token material.
type IdentityLink struct {
	UserID                string   `json:"user_id,omitempty"`
	AccountID             string   `json:"account_id,omitempty"`
	PrimaryAuthIdentityID string   `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs       []string `json:"auth_identity_ids,omitempty"`
}

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
	Mobile          string            `json:"mobile,omitempty"`
	LineID          string            `json:"line_id,omitempty"`
	ExternalHandles map[string]string `json:"external_handles,omitempty"`
}
