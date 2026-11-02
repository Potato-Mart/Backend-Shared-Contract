package provider

import "time"

// BankTransferEvidence is one finalized evidence object (for example a photo of
// the payer's banking app) that staff captured for a bank-transfer payment.
//
// Code is the Payments-owned media code of the stored object. The object
// itself stays private: clients obtain a short-lived download URL from
// Payments, never a public URL. SHA256 is the lowercase hex digest of the
// stored bytes, verified when the upload was finalized.
type BankTransferEvidence struct {
	Code       string    `json:"code"`
	MimeType   string    `json:"mime_type"`
	SHA256     string    `json:"sha256"`
	CapturedAt time.Time `json:"captured_at"`
	CapturedBy string    `json:"captured_by"`
}
