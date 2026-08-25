package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"time"
)

// PromotionTerm is an open-key typed term. Services enforce the one-value-arm invariant.
type PromotionTerm struct {
	Key              string       `json:"key"`
	StringValue      *string      `json:"string_value,omitempty"`
	IntegerValue     *int64       `json:"integer_value,omitempty"`
	BasisPointsValue *int64       `json:"basis_points_value,omitempty"`
	BooleanValue     *bool        `json:"boolean_value,omitempty"`
	MoneyValue       *money.Money `json:"money_value,omitempty"`
	TimestampValue   *time.Time   `json:"timestamp_value,omitempty"`
}
