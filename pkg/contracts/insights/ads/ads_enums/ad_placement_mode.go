package ads_enums

// AdPlacementMode says whether a draft ad joins an existing platform ad set or
// brings the settings for a new campaign and ad set with it.
type AdPlacementMode string

const (
	// AdPlacementModeNew creates a new campaign and ad set from the
	// settings carried on the draft.
	AdPlacementModeNew AdPlacementMode = "new"
	// AdPlacementModeExisting joins an ad set that already exists on the
	// platform.
	AdPlacementModeExisting AdPlacementMode = "existing"
)

// IsValid reports whether m is a known AdPlacementMode.
func (m AdPlacementMode) IsValid() bool {
	switch m {
	case AdPlacementModeNew, AdPlacementModeExisting:
		return true
	}
	return false
}

func (m AdPlacementMode) String() string { return string(m) }
