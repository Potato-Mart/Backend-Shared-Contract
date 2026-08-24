package promotion

import "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"

// PromotionAmount is one resolved monetary outcome keyed by an open name.
type PromotionAmount struct {
	Key    string      `json:"key"`
	Amount money.Money `json:"amount"`
}
