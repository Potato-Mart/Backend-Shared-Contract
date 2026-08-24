package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
)

// WholesaleFreightPreset defines reusable wholesale freight rules.
// System presets cannot be deleted.
type WholesaleFreightPreset struct {
	PresetKey string `json:"preset_key"`
	Label     string `json:"label"`

	// Bulk (per-box) freight rules
	BulkPerBox    money.Money  `json:"bulk_per_box"`
	BulkFreeAbove *money.Money `json:"bulk_free_above,omitempty"`
	BulkMinOrder  *money.Money `json:"bulk_min_order,omitempty"`

	// Pallet freight rules
	PalletFlat      money.Money  `json:"pallet_flat"`
	PalletFreeAbove *money.Money `json:"pallet_free_above,omitempty"`
	PalletMinOrder  *money.Money `json:"pallet_min_order,omitempty"`

	// Surcharge when no forklift is available at delivery address
	NoForkliftSurcharge money.Money `json:"no_forklift_surcharge"`

	IsSystem bool `json:"is_system"` // system presets cannot be deleted

	audit.AuditFields
}
