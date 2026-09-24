package courier

// DeliveryProviderCredentials is the standalone privileged value shape for
// courier provider credentials. Its fields are sensitive values. Serialize it
// only on an independently authorized credential write or writer-only detail
// operation; never embed it in company, connection, list, or customer models.
// Supply owns authorization, encrypted persistence, versioning, rotation, and
// log redaction. Services must not persist these values in plaintext.
type DeliveryProviderCredentials struct {
	SignInAccount string `json:"sign_in_account,omitempty"`
	Password      string `json:"password,omitempty"`
	APIBaseURL    string `json:"api_base_url,omitempty"`
	APIToken      string `json:"api_token,omitempty"`
}
