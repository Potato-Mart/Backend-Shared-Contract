package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/temporal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// CatalogueReference pins a classification to the exact official catalogue
// version and entry that a reviewer saw.
type CatalogueReference struct {
	ID                   string                               `json:"id"`
	Jurisdiction         import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Version              string                               `json:"version"`
	EntryID              string                               `json:"entry_id,omitempty"`
	SourceURL            string                               `json:"source_url"`
	SourceChecksumSHA256 string                               `json:"source_checksum_sha256"`
	EffectiveFrom        temporal.Date                        `json:"effective_from,omitempty"`
	RetrievedAt          time.Time                            `json:"retrieved_at"`
}
