package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
)

// SupplierAvailablePromotion is a locale-aware supplier promotion without
// product-specific qualification or pricing policy.
type SupplierAvailablePromotion struct {
	Names        []localization.LocalizedName        `json:"names"`
	Descriptions []localization.LocalizedDescription `json:"descriptions,omitempty"`
}
