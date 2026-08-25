package import_compliance

// DeclarationSignatory references a managed image rather than embedding a
// base64 payload in the declaration record.
type DeclarationSignatory struct {
	Name               string `json:"name"`
	Title              string `json:"title,omitempty"`
	SignatureMediaCode string `json:"signature_media_code,omitempty"`
}
