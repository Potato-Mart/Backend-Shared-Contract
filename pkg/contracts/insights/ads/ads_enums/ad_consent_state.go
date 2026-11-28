package ads_enums

// AdConsentState is one advertising consent answer carried with a stored ad
// click. Consent defaults to denied: absence of a granted state is never
// consent, and a denied state withholds the behaviour it governs.
type AdConsentState string

const (
	// AdConsentStateGranted means the visitor agreed to this use.
	AdConsentStateGranted AdConsentState = "granted"
	// AdConsentStateDenied means the visitor did not agree to this use.
	AdConsentStateDenied AdConsentState = "denied"
)

// IsValid reports whether s is a known AdConsentState.
func (s AdConsentState) IsValid() bool {
	switch s {
	case AdConsentStateGranted, AdConsentStateDenied:
		return true
	}
	return false
}

func (s AdConsentState) String() string { return string(s) }
