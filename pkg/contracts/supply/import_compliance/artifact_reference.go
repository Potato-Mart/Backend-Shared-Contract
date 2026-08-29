package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// ArtifactReference identifies a deterministic generated artifact stored by a
// backend-managed media service.
type ArtifactReference struct {
	ID             string                               `json:"id"`
	Kind           import_compliance_enums.ArtifactKind `json:"kind"`
	MediaCode      string                               `json:"media_code"`
	Filename       string                               `json:"filename"`
	MIMEType       string                               `json:"mime_type"`
	ChecksumSHA256 string                               `json:"checksum_sha256"`
	RevisionNumber int64                                `json:"revision_number"`
	GeneratedAt    time.Time                            `json:"generated_at"`
}
