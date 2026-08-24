package promotion

// PromotionRelation connects qualifying products to target products. Kind is
// deliberately open so add-on purchase, BOGO, bundle, gift, discount, and
// future mechanics use the same qualifier-to-target grammar.
type PromotionRelation struct {
	ID             string          `json:"id"`
	Kind           string          `json:"kind"`
	QualifierScope PromotionScope  `json:"qualifier_scope"`
	TargetScope    PromotionScope  `json:"target_scope"`
	Terms          []PromotionTerm `json:"terms,omitempty"`
}
