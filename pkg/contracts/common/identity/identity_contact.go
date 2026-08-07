package identity

// IdentityLink carries shared references from a business profile to the
// canonical identity/account contracts. It contains identifiers only and no
// credential or token material.
type IdentityLink struct {
	UserID                string   `json:"user_id,omitempty"`
	AccountID             string   `json:"account_id,omitempty"`
	PrimaryAuthIdentityID string   `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs       []string `json:"auth_identity_ids,omitempty"`
}
