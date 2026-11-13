package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
)

// OrderGiftCardIssue records one gift card sold on an order. LineID is the
// cart line it came from and ItemID the OrderItem that carries its face
// value. Amount is the face value paid; BonusAmountMinor is the bonus the
// denomination policy added on issue, in the minor units of Amount's
// currency, and is omitted when there is none. The card belongs to the retail
// customer CustomerNumber.
//
// IssuanceID is Pricing's issuance reference and GiftCardCode the issued
// card's code; both are absent until the card is issued. The card's PIN is
// never recorded here. FailureReasonCode is set when Status is failed, and
// IssuedAt is when Pricing issued the card.
type OrderGiftCardIssue struct {
	LineID                    string                          `json:"line_id"`
	ItemID                    string                          `json:"item_id"`
	Amount                    money.Money                     `json:"amount"`
	BonusAmountMinor          int64                           `json:"bonus_amount_minor,omitempty"`
	CustomerNumber            string                          `json:"customer_number"`
	DenominationPolicyVersion int                             `json:"denomination_policy_version"`
	IssuanceID                string                          `json:"issuance_id,omitempty"`
	GiftCardCode              string                          `json:"gift_card_code,omitempty"`
	Status                    order_enums.GiftCardIssueStatus `json:"status"`
	FailureReasonCode         string                          `json:"failure_reason_code,omitempty"`
	IssuedAt                  *time.Time                      `json:"issued_at,omitempty"`
}
