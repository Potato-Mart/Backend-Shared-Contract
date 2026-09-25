package courier

// DeliveryCredentialRequirements is a non-sensitive, company-specific
// description of the credential and connection-test configuration required by
// a registered API implementation. Supply derives this value; operators do not
// set it.
type DeliveryCredentialRequirements struct {
	APIBaseURL                string `json:"api_base_url"`
	ConnectionAddressRequired bool   `json:"connection_address_required"`
	// AuthenticationMethods advertises only methods supported by the registered
	// implementation. Missing/empty metadata advertises no method; it does not
	// imply account/password support or successful authentication.
	AuthenticationMethods []DeliveryAuthenticationMethodRequirements `json:"authentication_methods,omitempty"`
}
