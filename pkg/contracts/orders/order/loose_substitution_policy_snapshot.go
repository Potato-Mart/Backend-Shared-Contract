package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/orders/order/order_enums"
	"time"
)

// LooseSubstitutionPolicySnapshot records whether unavailable sealed cases may
// be fulfilled with an exact base-unit quantity of loose EACH stock.
type LooseSubstitutionPolicySnapshot struct {
	Allowed    bool                                      `json:"allowed"`
	Source     order_enums.LooseSubstitutionPolicySource `json:"source"`
	CapturedAt time.Time                                 `json:"captured_at"`
}
