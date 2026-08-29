package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"
)

// EvidenceRecord tracks security evidence and chain-of-custody metadata.
type EvidenceRecord struct {
	ID             string            `json:"id"`
	CollectedAt    time.Time         `json:"collected_at"`
	CollectedBy    string            `json:"collected_by,omitempty"`
	Source         string            `json:"source,omitempty"`
	EvidenceType   string            `json:"evidence_type"` // e.g. "log", "screenshot", "export", "provider_report"
	Description    string            `json:"description,omitempty"`
	StorageURI     string            `json:"storage_uri,omitempty"`
	HashAlgorithm  string            `json:"hash_algorithm,omitempty"`
	HashValue      string            `json:"hash_value,omitempty"`
	ChainOfCustody []CustodyEvent    `json:"chain_of_custody,omitempty"`
	Metadata       metadata.Metadata `json:"metadata,omitempty"`

	DataProtectionFields
}
