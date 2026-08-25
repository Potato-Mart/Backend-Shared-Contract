package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"

// SellingProductContent contains the resolved, localized product display data
// that a customer client can render without querying catalogue master records.
type SellingProductContent struct {
	Name         localization.LocalizedName          `json:"name"`
	Descriptions []localization.LocalizedDescription `json:"descriptions,omitempty"`
	OtherNames   []localization.LocalizedName        `json:"other_names,omitempty"`
	Origin       *ProductOrigin                      `json:"origin,omitempty"`
	Images       *SellingProductImages               `json:"images,omitempty"`
}
