package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
)

// CloudServiceSecurityProfile records security posture for a cloud service
// or managed platform used by a backend.
type CloudServiceSecurityProfile struct {
	ID                    string          `json:"id"`
	Provider              string          `json:"provider"` // e.g. "azure", "aws", "gcp", "adyen"
	ServiceName           string          `json:"service_name"`
	Environment           string          `json:"environment,omitempty"` // "dev", "staging", "prod"
	OwnerID               string          `json:"owner_id,omitempty"`
	Region                string          `json:"region,omitempty"`
	DataResidency         string          `json:"data_residency,omitempty"`
	EncryptionAtRest      bool            `json:"encryption_at_rest"`
	EncryptionInTransit   bool            `json:"encryption_in_transit"`
	CustomerManagedKeys   bool            `json:"customer_managed_keys,omitempty"`
	LoggingEnabled        bool            `json:"logging_enabled"`
	MonitoringEnabled     bool            `json:"monitoring_enabled"`
	BackupEnabled         bool            `json:"backup_enabled,omitempty"`
	DeletionSupported     bool            `json:"deletion_supported,omitempty"`
	DLPEnabled            bool            `json:"dlp_enabled,omitempty"`
	LastRiskReviewedAt    *time.Time      `json:"last_risk_reviewed_at,omitempty"`
	LastConfigReviewedAt  *time.Time      `json:"last_config_reviewed_at,omitempty"`
	SupplierAgreementID   string          `json:"supplier_agreement_id,omitempty"`
	ComplianceReferences  []string        `json:"compliance_references,omitempty"`
	ConfigurationBaseline string          `json:"configuration_baseline,omitempty"`
	Metadata              common.Metadata `json:"metadata,omitempty"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
}
