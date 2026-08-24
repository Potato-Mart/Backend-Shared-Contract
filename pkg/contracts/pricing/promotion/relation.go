package promotion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
)

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

// PromotionTerm is an open-key typed term. Exactly one of its typed value arms
// must be present; contract consumers validate that invariant at their input
// boundary. Kinds remain open through Key rather than a closed enum.
type PromotionTerm struct {
	Key              string       `json:"key"`
	StringValue      *string      `json:"string_value,omitempty"`
	IntegerValue     *int64       `json:"integer_value,omitempty"`
	BasisPointsValue *int64       `json:"basis_points_value,omitempty"`
	BooleanValue     *bool        `json:"boolean_value,omitempty"`
	MoneyValue       *money.Money `json:"money_value,omitempty"`
	TimestampValue   *time.Time   `json:"timestamp_value,omitempty"`
}

// PromotionAmount is one resolved monetary outcome keyed by an open name.
type PromotionAmount struct {
	Key    string      `json:"key"`
	Amount money.Money `json:"amount"`
}
