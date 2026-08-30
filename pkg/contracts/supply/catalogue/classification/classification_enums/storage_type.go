package classification_enums

// StorageType classifies the physical storage conditions required by a
// catalogue product, SKU, package, or warehouse location.
type StorageType string

const (
	StorageAmbient StorageType = "AMBIENT"
	StorageChilled StorageType = "CHILLED"
	StorageFrozen  StorageType = "FROZEN"
)

// IsValid reports whether s is a supported catalogue storage type.
func (s StorageType) IsValid() bool {
	switch s {
	case StorageAmbient, StorageChilled, StorageFrozen:
		return true
	default:
		return false
	}
}

func (s StorageType) String() string { return string(s) }
