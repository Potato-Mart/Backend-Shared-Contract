package warehouse_enums

// LocationStorageType is the storage environment of a stock location. It
// shares AMBIENT, CHILLED and FROZEN wire values with the product storage
// type; MIXED is location-only and is reserved for system-managed pools
// (the depot goods pool), which accept stock of any product storage type.
// A product can never be MIXED.
type LocationStorageType string

const (
	LocationStorageAmbient LocationStorageType = "AMBIENT"
	LocationStorageChilled LocationStorageType = "CHILLED"
	LocationStorageFrozen  LocationStorageType = "FROZEN"
	LocationStorageMixed   LocationStorageType = "MIXED"
)

// IsValid reports whether t is a known LocationStorageType.
func (t LocationStorageType) IsValid() bool {
	switch t {
	case LocationStorageAmbient, LocationStorageChilled, LocationStorageFrozen, LocationStorageMixed:
		return true
	}
	return false
}

func (t LocationStorageType) String() string { return string(t) }
