package courier

// DeliveryCredentialRequirements is a non-sensitive, company-specific
// description of the credential and connection-test configuration required by
// a registered API implementation. Supply derives this value; operators do not
// set it.
type DeliveryCredentialRequirements struct {
	APIBaseURL                string `json:"api_base_url"`
	ConnectionAddressRequired bool   `json:"connection_address_required"`
}
