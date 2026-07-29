package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
)

// WholesaleFreightPreset defines reusable wholesale freight rules.
// System presets cannot be deleted.
type WholesaleFreightPreset struct {
	PresetKey string `json:"preset_key"`
	Label     string `json:"label"`

	// Bulk (per-box) freight rules
	BulkPerBox    common.Money  `json:"bulk_per_box"`
	BulkFreeAbove *common.Money `json:"bulk_free_above,omitempty"`
	BulkMinOrder  *common.Money `json:"bulk_min_order,omitempty"`

	// Pallet freight rules
	PalletFlat      common.Money  `json:"pallet_flat"`
	PalletFreeAbove *common.Money `json:"pallet_free_above,omitempty"`
	PalletMinOrder  *common.Money `json:"pallet_min_order,omitempty"`

	// Surcharge when no forklift is available at delivery address
	NoForkliftSurcharge common.Money `json:"no_forklift_surcharge"`

	IsSystem bool `json:"is_system"` // system presets cannot be deleted

	common.AuditFields
}
