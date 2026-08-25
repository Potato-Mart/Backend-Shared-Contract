package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"

// ProductLocalization groups secondary per-language display fields.
type ProductLocalization struct {
	OtherNames []localization.LocalizedName `json:"other_names,omitempty"`
}
