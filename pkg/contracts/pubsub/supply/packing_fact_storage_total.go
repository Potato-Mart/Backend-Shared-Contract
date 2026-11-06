package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"

// PackingFactStorageTotal sums a stage's lines per storage type, so the
// ambient and chilled units of the shared AMBIENT_CHILLED stage stay separate.
type PackingFactStorageTotal struct {
	StorageType    classification_enums.StorageType `json:"storage_type"`
	TotalBaseUnits int64                            `json:"total_base_units"`
	LineCount      int                              `json:"line_count"`
}
