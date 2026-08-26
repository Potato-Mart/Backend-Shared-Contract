package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/localization"

// ProductContent contains the customer-facing, locale-aware product facts.
// Its image references are resolved into render-safe media only in
// SellingProduct.
type ProductContent struct {
	Name         localization.LocalizedName          `json:"name"`
	Descriptions []localization.LocalizedDescription `json:"descriptions,omitempty"`
	Localization *ProductLocalization                `json:"localization,omitempty"`
	Origin       *ProductOrigin                      `json:"origin,omitempty"`
	Images       *Images                             `json:"images,omitempty"`
}
