package warehouse_enums

// StorageType classifies the physical storage conditions required by a
// product or stock location.
type StorageType string

const (
	StorageAmbient StorageType = "AMBIENT"
	StorageChilled StorageType = "CHILLED"
	StorageFrozen  StorageType = "FROZEN"
)

// IsValid reports whether s is a known StorageType.
func (s StorageType) IsValid() bool {
	switch s {
	case StorageAmbient, StorageChilled, StorageFrozen:
		return true
	}
	return false
}

func (s StorageType) String() string { return string(s) }
