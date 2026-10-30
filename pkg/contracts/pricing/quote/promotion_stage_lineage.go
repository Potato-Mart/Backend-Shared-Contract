package quote

// PromotionStageLineage binds a retained run to its original immutable quote
// and ordinal interval. The current run's FirstUnit is rebased; RootFirstUnit
// identifies the same first native unit in the root line. Root quote identity
// and the complete predecessor chain must be independently authenticated.
// Lineage with a nonempty root quote, fingerprint and line is required for
// every clipped or rebased retained run.
type PromotionStageLineage struct {
	// These optional fields are required for negotiated manifest-based retained
	// proofs. They bind the exact root Pricing record and manifest; historical
	// v1 lineage remains byte-compatible when omitted.
	RootQuoteKey         string `json:"root_quote_key,omitempty"`
	RootQuoteRevision    string `json:"root_quote_revision,omitempty"`
	RootManifestID       string `json:"root_manifest_id,omitempty"`
	RootQuoteID          string `json:"root_quote_id"`
	RootQuoteFingerprint string `json:"root_quote_fingerprint"`
	RootLineID           string `json:"root_line_id"`
	RootFirstUnit        int64  `json:"root_first_unit"`
}
