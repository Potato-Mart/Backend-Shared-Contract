package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

// GiftCard is a stored-value instrument with a re-spendable balance. The
// GiftCardTransaction ledger is the source of truth for Balance; this record is
// the current projected state. It is referenced everywhere by Code (the
// business key), never by ID.
type GiftCard struct {
	ID           string                        `json:"id"`
	Code         string                        `json:"code"`
	Owner        membership.MembershipOwnerRef `json:"owner"`
	Balance      common.Money                  `json:"balance"`
	InitialValue common.Money                  `json:"initial_value"`
	Status       enums.GiftCardStatus          `json:"status"`
	IssuedAt     time.Time                     `json:"issued_at"`
	ActivatedAt  *time.Time                    `json:"activated_at,omitempty"`
	ExpiresAt    *time.Time                    `json:"expires_at,omitempty"`
	Note         string                        `json:"note,omitempty"`
	History      []shared.HistoryEntry         `json:"history,omitempty"`

	common.AuditFields
}

// GiftCardTransaction is one entry in a gift card's balance ledger. A positive
// Delta tops up (issue / top_up / refund); a negative Delta redeems.
// BalanceAfter is the running balance after this entry.
type GiftCardTransaction struct {
	ID                 string                          `json:"id"`
	GiftCardCode       string                          `json:"gift_card_code"`
	Delta              common.Money                    `json:"delta"`
	BalanceAfter       common.Money                    `json:"balance_after"`
	Reason             enums.GiftCardTransactionReason `json:"reason"`
	RelatedOrderNumber string                          `json:"related_order_number,omitempty"`
	Note               string                          `json:"note,omitempty"`
	CreatedBy          string                          `json:"created_by,omitempty"`
	CreatedAt          time.Time                       `json:"created_at"`
}
