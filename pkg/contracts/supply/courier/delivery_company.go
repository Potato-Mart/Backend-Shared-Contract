package courier

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/geography"
)

// DeliveryCompany is the Supply-owned record for an independent API delivery
// company. Code is its immutable identity, not a registered provider identifier.
// Admins configure API endpoints, arbitrary configuration keys and capabilities
// through CRUD. Configured APIs must implement the normalized delivery protocol
// using the existing delivery request, schedule and selection contracts.
// No provider adapters, response-field mappings or stored slot catalogues apply.
// Revision identifies the effective configuration, service areas and metadata.
// Read projections must mask every sensitive Configuration value before encoding.
// Supply owns encrypted storage, privileged access and validation.
type DeliveryCompany struct {
	Code                string                              `json:"code"`
	Name                string                              `json:"name"`
	Enabled             bool                                `json:"enabled"`
	DefaultInstructions string                              `json:"default_instructions,omitempty"`
	Revision            int64                               `json:"revision"`
	CountryCodes        []geography.CountryCode             `json:"country_codes"`
	Capabilities        DeliveryCapabilities                `json:"capabilities"`
	Configuration       []DeliveryCompanyConfigurationEntry `json:"configuration"`
	ServiceAreas        []DeliveryServiceArea               `json:"service_areas"`
	// CustomMetadata is inert company-owned non-secret metadata. It must not be
	// passed to API requests. Unknown value types must survive read-modify-write.
	CustomMetadata []DeliveryCompanyCustomMetadataEntry `json:"custom_metadata,omitempty"`
	audit.AuditFields
}
