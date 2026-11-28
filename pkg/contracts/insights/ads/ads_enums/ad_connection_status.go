package ads_enums

// AdConnectionStatus is how usable one platform connection is. It never
// carries a credential: a connection reports only whether it is configured and
// working.
type AdConnectionStatus string

const (
	// AdConnectionStatusConnected is configured and working.
	AdConnectionStatusConnected AdConnectionStatus = "connected"
	// AdConnectionStatusNotConfigured has no accounts configured for this
	// environment, so ads reads have nothing to serve.
	AdConnectionStatusNotConfigured AdConnectionStatus = "not_configured"
	// AdConnectionStatusNeedsAttention is configured but degraded, for
	// example a credential about to expire or a failing sync.
	AdConnectionStatusNeedsAttention AdConnectionStatus = "needs_attention"
)

// IsValid reports whether s is a known AdConnectionStatus.
func (s AdConnectionStatus) IsValid() bool {
	switch s {
	case AdConnectionStatusConnected, AdConnectionStatusNotConfigured, AdConnectionStatusNeedsAttention:
		return true
	}
	return false
}

func (s AdConnectionStatus) String() string { return string(s) }
