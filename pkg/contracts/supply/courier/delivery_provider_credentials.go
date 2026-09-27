package courier

// DeliveryProviderCredentials is the standalone privileged value shape for
// courier provider credentials. Its fields are sensitive values. Serialize it
// only on an independently authorized credential write or writer-only detail
// operation; never embed it in company, connection, list, or customer models.
// APIBaseURL is the sole common field; provider-specific authentication values
// belong in ProviderExtension. Supply owns authorization, encrypted persistence,
// versioning, rotation, and log redaction. Services must not persist these
// values in plaintext.
type DeliveryProviderCredentials struct {
	APIBaseURL string `json:"api_base_url,omitempty"`
	// Omission preserves the stored provider extension on partial writes; Supply
	// owns explicit replacement and removal operations.
	ProviderExtension *DeliveryProviderCredentialExtension `json:"provider_extension,omitempty"`
}
