package ads

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdBudget is the spending allowance held on one platform object, in that ad
// account's own currency.
//
// Level says which object actually holds the budget, which is not always the
// object being read: a Meta ad set budget and a Meta campaign budget both
// appear on the rows below them. Exactly one of Daily and Lifetime is set;
// the other is absent, never a zero amount. An absent budget means the object
// holds none of its own.
type AdBudget struct {
	Level    ads_enums.AdLevel `json:"level,omitempty"`
	Daily    *money.Money      `json:"daily,omitempty"`
	Lifetime *money.Money      `json:"lifetime,omitempty"`
}
