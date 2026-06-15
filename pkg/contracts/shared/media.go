package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
)

// MediaStatus controls visibility and lifecycle of a media row.
//
//	pending  – SAS URL issued but the upload hasn't been finalized yet.
//	active   – Upload finalized; safe to reference.
//	deleted  – Soft-deleted; the orphan reaper will purge the blob
//	           once no documents still reference the row.
type MediaStatus string

const (
	MediaStatusPending MediaStatus = "pending"
	MediaStatusActive  MediaStatus = "active"
	MediaStatusDeleted MediaStatus = "deleted"
)

// Media is the public projection of a stored asset (image, document,
// generated poster, etc.). The bytes live in object storage; this
// record carries everything required to render, audit, and reap.
//
// The pre-v3.3 fields (ID, Filename, StoragePath, URL, MIMEType,
// SizeBytes, Width, Height, CreatedAt, CreatedBy) are unchanged and
// remain wire-compatible. Folder, Status, UpdatedAt, DeletedAt and
// PurgedAt are additive optional fields introduced in v3.3.0.
type Media struct {
	ID          string      `json:"id"`
	Filename    string      `json:"filename"`
	StoragePath string      `json:"storage_path"`
	URL         string      `json:"url"`
	MIMEType    string      `json:"mime_type"`
	SizeBytes   int64       `json:"size_bytes"`
	Width       int         `json:"width,omitempty"`
	Height      int         `json:"height,omitempty"`
	Folder      string      `json:"folder,omitempty"`
	Status      MediaStatus `json:"status,omitempty"`
	PurgedAt    *time.Time  `json:"purged_at,omitempty"`

	common.AuditFields          `bson:",inline"`
	common.DataProtectionFields `bson:",inline"`
}
