package security

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security/security_enums"
	"time"
)

// ObjectMediaAsset is the complete stored-asset record. The bytes live in
// object storage; this record carries the storage, lifecycle, audit, and data
// protection information needed to manage the asset.
type ObjectMediaAsset struct {
	ID              string                         `json:"id"`
	Code            string                         `json:"code"`
	Filename        string                         `json:"filename"`
	Bucket          string                         `json:"bucket"`
	StoragePath     string                         `json:"storage_path"`
	URL             string                         `json:"url,omitempty"`
	Visibility      security_enums.MediaVisibility `json:"visibility"`
	MIMEType        string                         `json:"mime_type"`
	SizeBytes       int64                          `json:"size_bytes"`
	Width           int                            `json:"width,omitempty"`
	Height          int                            `json:"height,omitempty"`
	Folder          string                         `json:"folder,omitempty"`
	Status          security_enums.MediaStatus     `json:"status,omitempty"`
	References      []ObjectMediaReference         `json:"references,omitempty"`
	UploadExpiresAt *time.Time                     `json:"upload_expires_at,omitempty"`
	PurgedAt        *time.Time                     `json:"purged_at,omitempty"`

	audit.AuditFields
	DataProtectionFields
}
