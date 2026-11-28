package ads_enums

// AdCredentialStatus reports the health of a platform credential without
// revealing it. Credentials live in the secret store; no ads record ever
// carries a token, a key or any part of one.
type AdCredentialStatus string

const (
	// AdCredentialStatusOK is valid with no expiry in sight.
	AdCredentialStatusOK AdCredentialStatus = "ok"
	// AdCredentialStatusExpiring is still valid but close enough to expiry
	// that it needs rotating.
	AdCredentialStatusExpiring AdCredentialStatus = "expiring"
	// AdCredentialStatusInvalid is rejected by the platform and must be
	// replaced before ads reads or writes can work.
	AdCredentialStatusInvalid AdCredentialStatus = "invalid"
)

// IsValid reports whether s is a known AdCredentialStatus.
func (s AdCredentialStatus) IsValid() bool {
	switch s {
	case AdCredentialStatusOK, AdCredentialStatusExpiring, AdCredentialStatusInvalid:
		return true
	}
	return false
}

func (s AdCredentialStatus) String() string { return string(s) }
