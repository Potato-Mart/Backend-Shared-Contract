package procurement

import (
	"time"
)

// SupplierInvoiceDocument is the immutable reference to the received document.
// ContentSHA256 makes the stored artefact tamper-evident and, with
// DuplicateKey on the invoice, protects against re-recording the same invoice.
type SupplierInvoiceDocument struct {
	Reference     string    `json:"reference"`
	ContentSHA256 string    `json:"content_sha256"`
	MediaCode     string    `json:"media_code,omitempty"`
	ReceivedAt    time.Time `json:"received_at"`
}
