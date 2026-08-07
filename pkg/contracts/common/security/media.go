package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/security/security_enums"
)

// Media is the public projection of a stored asset (image, document,
// generated poster, etc.). The bytes live in object storage; this
// record carries everything required to render, audit, and reap.
type Media struct {
	ID              string                         `json:"id"`
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
	References      []MediaReference               `json:"references,omitempty"`
	UploadExpiresAt *time.Time                     `json:"upload_expires_at,omitempty"`
	PurgedAt        *time.Time                     `json:"purged_at,omitempty"`

	audit.AuditFields
	DataProtectionFields
}

// MediaReference records an owning aggregate attachment without embedding the
// aggregate itself in the media record.
type MediaReference struct {
	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`
	Field      string `json:"field,omitempty"`
}
