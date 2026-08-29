package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// EvidenceReference points at source material without embedding regulated or
// potentially sensitive file contents in the shared model.
type EvidenceReference struct {
	ID                   string                               `json:"id"`
	Kind                 import_compliance_enums.EvidenceKind `json:"kind"`
	MediaCode            string                               `json:"media_code,omitempty"`
	SourceURL            string                               `json:"source_url,omitempty"`
	SourceTitle          string                               `json:"source_title,omitempty"`
	SourceVersion        string                               `json:"source_version,omitempty"`
	SourceChecksumSHA256 string                               `json:"source_checksum_sha256,omitempty"`
	CapturedAt           *time.Time                           `json:"captured_at,omitempty"`
	Note                 string                               `json:"note,omitempty"`
}
