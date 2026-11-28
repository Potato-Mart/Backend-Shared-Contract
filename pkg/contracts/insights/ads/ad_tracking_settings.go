package ads

// AdTrackingSettings is the single stored configuration for the ads module:
// how orders are attributed to ads, and what limits a delivery change is held
// to.
//
// Revision is the opaque revision a write must quote as its expected revision,
// so two people cannot overwrite each other's settings; it is carried exactly
// and never compared as a number. Both sections are optional, and an absent
// section means the stored defaults still apply rather than "no rule".
type AdTrackingSettings struct {
	Revision    string               `json:"revision,omitempty"`
	Attribution *AdAttributionPolicy `json:"attribution,omitempty"`
	Guardrails  *AdGuardrails        `json:"guardrails,omitempty"`
}
