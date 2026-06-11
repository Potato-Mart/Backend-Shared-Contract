package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/common"
)

// MediaUploadRequest is the body of POST /media/upload-url. The client
// describes the file it intends to upload and the server returns a
// short-lived SAS URL the client can PUT to directly.
type MediaUploadRequest struct {
	// Filename as it should appear in the media library. The server may
	// sanitize or namespace it.
	Filename string `json:"filename"`

	// MIMEType is the Content-Type the client will set on the blob PUT.
	// The server validates this against an allowlist before issuing the
	// SAS URL.
	MIMEType string `json:"mime_type"`

	// SizeBytes is the byte length the client intends to upload. The
	// server validates this against per-folder size caps; the actual
	// blob length is also validated on Finalize.
	SizeBytes int64 `json:"size_bytes"`

	// Folder selects which container/prefix to upload into – e.g.
	// "product-cover", "product-gallery", "marketing-poster", "import-invoice".
	// Each folder has its own allowed MIME types, max size, and TTL.
	Folder string `json:"folder"`

	// AttachToType / AttachToID are optional. When set, the server can
	// hint the eventual relationship for analytics ("media uploaded
	// while editing product:prod_abc"). Authoritative linking happens
	// on Finalize.
	AttachToType string `json:"attach_to_type,omitempty"`
	AttachToID   string `json:"attach_to_id,omitempty"`

	common.DataProtectionFields
}

// MediaUploadResponse is what the server returns for an upload request.
// The client must PUT the bytes to UploadURL within ExpiresAt, then
// call POST /media/finalize with MediaID to register the upload.
type MediaUploadResponse struct {
	MediaID   string    `json:"media_id"`
	UploadURL string    `json:"upload_url"`
	Method    string    `json:"method"` // always "PUT"
	Headers   []Header  `json:"headers,omitempty"`
	MaxBytes  int64     `json:"max_bytes"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Header is a single HTTP header the client must echo on its PUT.
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MediaFinalizeRequest confirms that the client successfully uploaded
// to the SAS URL. The server validates the blob exists, matches the
// claimed size, then flips the media row from PENDING to ACTIVE.
type MediaFinalizeRequest struct {
	MediaID string `json:"media_id"`
}
