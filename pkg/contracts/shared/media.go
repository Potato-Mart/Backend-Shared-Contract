package shared

import "time"

type Media struct {
	ID          string    `json:"id"`
	Filename    string    `json:"filename"`
	StoragePath string    `json:"storage_path"`
	URL         string    `json:"url"`
	MIMEType    string    `json:"mime_type"`
	SizeBytes   int64     `json:"size_bytes"`
	Width       int       `json:"width,omitempty"`
	Height      int       `json:"height,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by,omitempty"`
}
