package promotion

import "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/promotion/promotion_enums"

// PromotionScopeGroup is one reusable product selector and quantity range.
type PromotionScopeGroup struct {
	MatchMode          promotion_enums.PromotionMatchMode `json:"match_mode"`
	SKUCodes           []string                           `json:"sku_codes,omitempty"`
	CollectionCodes    []string                           `json:"collection_codes,omitempty"`
	CategoryTagCodes   []string                           `json:"category_tag_codes,omitempty"`
	PackageOptionCodes []string                           `json:"package_option_codes,omitempty"`
	MinimumBaseUnits   int64                              `json:"minimum_base_units,omitempty"`
	MaximumBaseUnits   *int64                             `json:"maximum_base_units,omitempty"`
}
