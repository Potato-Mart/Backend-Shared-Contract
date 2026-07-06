package warehouseenum

// StorageType classifies the physical storage conditions required by a
// product. It is used both on SKU categories and on placing areas so
// that products can only be assigned to an area with matching
// refrigeration requirements.
type StorageType string

const (
	StorageDry     StorageType = "DRY"
	StorageChilled StorageType = "CHILLED"
	StorageFrozen  StorageType = "FROZEN"
)

// IsValid reports whether s is a known StorageType.
func (s StorageType) IsValid() bool {
	switch s {
	case StorageDry, StorageChilled, StorageFrozen:
		return true
	}
	return false
}

func (s StorageType) String() string { return string(s) }
