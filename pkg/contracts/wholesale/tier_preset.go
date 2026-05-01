package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
)

// WholesaleTierPreset defines the freight rules for a wholesale shipping tier.
// Presets are admin-editable; customer_profiles.wholesale_tier references
// a tier_key from this table. The four system presets cannot be deleted.
type WholesaleTierPreset struct {
	TierKey string `json:"tier_key"`
	Label   string `json:"label"`

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

	SortOrder int       `json:"sort_order"`
	IsSystem  bool      `json:"is_system"` // system presets cannot be deleted
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
