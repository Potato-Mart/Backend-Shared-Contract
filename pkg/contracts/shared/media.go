package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/product"
)

// Media is the public projection of a stored asset (image, document,
// generated poster, etc.). The bytes live in object storage; this
// record carries everything required to render, audit, and reap.
type Media struct {
	ID          string                  `json:"id"`
	Filename    string                  `json:"filename"`
	StoragePath string                  `json:"storage_path"`
	URL         string                  `json:"url"`
	MIMEType    string                  `json:"mime_type"`
	SizeBytes   int64                   `json:"size_bytes"`
	Width       int                     `json:"width,omitempty"`
	Height      int                     `json:"height,omitempty"`
	Folder      string                  `json:"folder,omitempty"`
	Status      productenum.MediaStatus `json:"status,omitempty"`
	PurgedAt    *time.Time              `json:"purged_at,omitempty"`

	common.AuditFields
	common.DataProtectionFields
}
