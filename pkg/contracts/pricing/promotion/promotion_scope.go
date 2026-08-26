package promotion

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/promotion/promotion_enums"

// PromotionScope selects canonical products. An unrestricted scope must set
// Unrestricted to true; an empty restricted scope is invalid. The outer
// MatchMode combines Groups. Within each group, values in a selector list are
// ORed, while MatchMode combines the group's non-empty selector dimensions.
type PromotionScope struct {
	Unrestricted bool                               `json:"unrestricted"`
	MatchMode    promotion_enums.PromotionMatchMode `json:"match_mode"`
	Groups       []PromotionScopeGroup              `json:"groups,omitempty"`
}
