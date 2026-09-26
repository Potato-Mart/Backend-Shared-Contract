package courier

import "github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/metadata"

// DeliveryProviderSettings contains versioned, non-secret configuration values
// for the provider implementation selected by the parent DeliveryCompany.Code.
// Supply owns the schema, validation and execution semantics. Values retain
// their JSON types; unknown keys must survive read-modify-write and must not be
// interpreted by an implementation that does not recognize them.
type DeliveryProviderSettings struct {
	SchemaVersion int               `json:"schema_version"`
	Values        metadata.Metadata `json:"values"`
}
