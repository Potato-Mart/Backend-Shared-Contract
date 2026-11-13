package cart

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// CartGiftCardIssueLine is one gift card being sold on a till (POS) cart.
// Amount is the face value the customer pays. The card is issued to the
// retail customer CustomerNumber, which must be the customer bound to the
// cart. DenominationPolicyVersion is the Pricing gift-card denomination policy
// the amount was checked against, and BonusAmountMinor is the bonus that
// policy adds on issue, in the minor units of Amount's currency; it is
// omitted when there is no bonus. A gift-card line is face value, outside the
// tax scope, moves no stock and earns no points.
type CartGiftCardIssueLine struct {
	LineID                    string      `json:"line_id"`
	Amount                    money.Money `json:"amount"`
	CustomerNumber            string      `json:"customer_number"`
	DenominationPolicyVersion int         `json:"denomination_policy_version"`
	BonusAmountMinor          int64       `json:"bonus_amount_minor,omitempty"`
	AddedByUserID             string      `json:"added_by_user_id"`
	AddedAt                   time.Time   `json:"added_at"`
}
