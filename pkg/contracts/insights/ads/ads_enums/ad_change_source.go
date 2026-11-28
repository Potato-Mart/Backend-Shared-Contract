package ads_enums

// AdChangeSource names where a recorded delivery change came from, so the
// change history can tell our own writes apart from edits made directly in
// Ads Manager or Google Ads and discovered by a later sync.
type AdChangeSource string

const (
	// AdChangeSourceAdmin is a change submitted through our own admin.
	AdChangeSourceAdmin AdChangeSource = "admin"
	// AdChangeSourcePlatform is a change made on the platform by someone
	// else and detected while syncing.
	AdChangeSourcePlatform AdChangeSource = "platform"
)

// IsValid reports whether s is a known AdChangeSource.
func (s AdChangeSource) IsValid() bool {
	switch s {
	case AdChangeSourceAdmin, AdChangeSourcePlatform:
		return true
	}
	return false
}

func (s AdChangeSource) String() string { return string(s) }
