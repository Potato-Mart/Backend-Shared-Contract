package wallet

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
	"time"
)

// GiftCardTransaction is one stored-value balance ledger entry.
type GiftCardTransaction struct {
	ID                 string                                 `json:"id"`
	GiftCardCode       string                                 `json:"gift_card_code"`
	Delta              money.Money                            `json:"delta"`
	BalanceAfter       money.Money                            `json:"balance_after"`
	Reason             wallet_enums.GiftCardTransactionReason `json:"reason"`
	ReservationID      string                                 `json:"reservation_id,omitempty"`
	RelatedOrderNumber string                                 `json:"related_order_number,omitempty"`
	Note               string                                 `json:"note,omitempty"`
	CreatedBy          string                                 `json:"created_by,omitempty"`
	CreatedAt          time.Time                              `json:"created_at"`
}
