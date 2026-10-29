package pricebook

import "time"

// OfflinePriceBookEvidence identifies the signed book the till cached. The
// digest identifies its exact signed bytes, which the issuer must retain for
// later verification. These fields are evidence, not proof of verification;
// consumers must verify the signature, the signed store binding and the
// underlying book. StoreCode refers to a field of the signed book itself.
// ExpiresAt is retained as a historical fact even when a sale occurred later.
type OfflinePriceBookEvidence struct {
	Version       string    `json:"version"`
	MarketCode    string    `json:"market_code"`
	StoreCode     string    `json:"store_code"`
	PayloadSHA256 string    `json:"payload_sha256"`
	Algorithm     string    `json:"algorithm"`
	KeyID         string    `json:"key_id"`
	Signature     string    `json:"signature_base64"`
	IssuedAt      time.Time `json:"issued_at"`
	ExpiresAt     time.Time `json:"expires_at"`
}
