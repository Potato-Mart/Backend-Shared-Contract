package enums

// FreshnessStatus tracks the food freshness.
type FreshnessStatus string

const (
	FreshnessStatusFresh    FreshnessStatus = "FRESH"
	FreshnessStatusExpiring FreshnessStatus = "EXPIRING"
	FreshnessStatusExpired  FreshnessStatus = "EXPIRED"
)

// IsValid reports whether p is a known FreshnessStatus.
func (p FreshnessStatus) IsValid() bool {
	switch p {
	case FreshnessStatusFresh, FreshnessStatusExpiring, FreshnessStatusExpired:
		return true
	}
	return false
}

func (p FreshnessStatus) String() string { return string(p) }
