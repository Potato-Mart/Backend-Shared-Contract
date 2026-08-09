package promotion_enums

// PromotionMatchMode controls whether a scope combines its entries using all
// requirements or any matching requirement.
type PromotionMatchMode string

const (
	PromotionMatchModeAll PromotionMatchMode = "all"
	PromotionMatchModeAny PromotionMatchMode = "any"
)

// IsValid reports whether m is a supported scope match mode.
func (m PromotionMatchMode) IsValid() bool {
	switch m {
	case PromotionMatchModeAll, PromotionMatchModeAny:
		return true
	}
	return false
}

// String returns the wire value for m.
func (m PromotionMatchMode) String() string { return string(m) }
