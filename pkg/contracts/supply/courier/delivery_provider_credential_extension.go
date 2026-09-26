package courier

// DeliveryProviderCredentialExtension is the sensitive, provider-specific
// credential value envelope for a separately authorized credential operation.
// It has the versioned values JSON shape of DeliveryProviderSettings, but the
// enclosing privileged credential model makes its values sensitive. Supply
// owns authorization, validation, encrypted persistence, explicit
// replace/remove semantics and redaction. Never embed it in DeliveryCompany or
// general read models; its values are input data, not ciphertext.
type DeliveryProviderCredentialExtension DeliveryProviderSettings
