package promotion

import "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/promotion/promotion_enums"

// PromotionScope selects canonical products. An unrestricted scope must set
// Unrestricted to true; an empty restricted scope is invalid. The outer
// MatchMode combines Groups. Within each group, values in a selector list are
// ORed, while MatchMode combines the group's non-empty selector dimensions.
type PromotionScope struct {
	Unrestricted bool                               `json:"unrestricted"`
	MatchMode    promotion_enums.PromotionMatchMode `json:"match_mode"`
	Groups       []PromotionScopeGroup              `json:"groups,omitempty"`
}

// PromotionScopeGroup is one reusable product selector and quantity range.
// A separate one-product group under an outer ALL mode expresses per-product
// requirements. A multi-product group expresses a combined quantity pool. A
// nil MaximumBaseUnits means the selected quantity is unlimited.
type PromotionScopeGroup struct {
	MatchMode          promotion_enums.PromotionMatchMode `json:"match_mode"`
	SKUCodes           []string                           `json:"sku_codes,omitempty"`
	CollectionCodes    []string                           `json:"collection_codes,omitempty"`
	CategoryTagCodes   []string                           `json:"category_tag_codes,omitempty"`
	PackageOptionCodes []string                           `json:"package_option_codes,omitempty"`
	MinimumBaseUnits   int64                              `json:"minimum_base_units,omitempty"`
	MaximumBaseUnits   *int64                             `json:"maximum_base_units,omitempty"`
}
